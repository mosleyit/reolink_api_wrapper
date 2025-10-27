// Package reolink provides a Go client for the Reolink Camera HTTP API.
//
// The client supports all API endpoints including system management,
// PTZ control, motion detection, AI features, and video streaming.
//
// Basic usage:
//
//	client := reolink.NewClient("192.168.1.100",
//	    reolink.WithCredentials("admin", "password"))
//
//	ctx := context.Background()
//	if err := client.Login(ctx); err != nil {
//	    log.Fatal(err)
//	}
//	defer client.Logout(ctx)
//
//	info, err := client.System.GetDeviceInfo(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Camera: %s\n", info.Model)
//
// For more examples, see the examples/ directory.
package reolink

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Client represents a Reolink camera API client
type Client struct {
	host       string
	baseURL    string
	httpClient *http.Client
	username   string
	password   string
	token      string
	tokenMu    sync.RWMutex
	useHTTPS   bool
	logger     Logger

	// API modules
	System    *SystemAPI
	Security  *SecurityAPI
	Network   *NetworkAPI
	Video     *VideoAPI
	Encoding  *EncodingAPI
	Recording *RecordingAPI
	PTZ       *PTZAPI
	Alarm     *AlarmAPI
	LED       *LEDAPI
	AI        *AIAPI
	Streaming *StreamingAPI
}

// NewClient creates a new Reolink API client
func NewClient(host string, opts ...Option) *Client {
	c := &Client{
		host:     host,
		useHTTPS: false,
		logger:   &NoOpLogger{}, // Default to no-op logger
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // Reolink cameras often use self-signed certs
				},
			},
		},
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	// Set base URL
	c.updateBaseURL()

	// Initialize API modules
	c.System = &SystemAPI{client: c}
	c.Security = &SecurityAPI{client: c}
	c.Network = &NetworkAPI{client: c}
	c.Video = &VideoAPI{client: c}
	c.Encoding = &EncodingAPI{client: c}
	c.Recording = &RecordingAPI{client: c}
	c.PTZ = &PTZAPI{client: c}
	c.Alarm = &AlarmAPI{client: c}
	c.LED = &LEDAPI{client: c}
	c.AI = &AIAPI{client: c}
	c.Streaming = &StreamingAPI{client: c}

	return c
}

// updateBaseURL updates the base URL based on current settings
func (c *Client) updateBaseURL() {
	scheme := "http"
	if c.useHTTPS {
		scheme = "https"
	}
	c.baseURL = fmt.Sprintf("%s://%s/cgi-bin/api.cgi", scheme, c.host)
}

// do executes an API request
func (c *Client) do(ctx context.Context, requests []Request, response interface{}) error {
	// Add token to requests if available
	c.tokenMu.RLock()
	token := c.token
	c.tokenMu.RUnlock()

	if token != "" {
		for i := range requests {
			requests[i].Token = token
		}
	}

	// Marshal request
	reqBody, err := json.Marshal(requests)
	if err != nil {
		c.logger.Error("failed to marshal request: %v", err)
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Build URL with cmd parameter
	url := c.baseURL
	if len(requests) > 0 {
		url = fmt.Sprintf("%s?cmd=%s", c.baseURL, requests[0].Cmd)
		if token != "" {
			url = fmt.Sprintf("%s&token=%s", url, token)
		}
		c.logger.Debug("API request: cmd=%s", requests[0].Cmd)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(reqBody))
	if err != nil {
		c.logger.Error("failed to create request: %v", err)
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Execute request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		c.logger.Error("failed to execute request: %v", err)
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer httpResp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		c.logger.Error("failed to read response: %v", err)
		return fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("API response: status=%d, body_len=%d", httpResp.StatusCode, len(respBody))

	// Check HTTP status
	if httpResp.StatusCode != http.StatusOK {
		c.logger.Warn("unexpected status code: %d", httpResp.StatusCode)
		return fmt.Errorf("unexpected status code: %d, body: %s", httpResp.StatusCode, string(respBody))
	}

	// Unmarshal response
	if err := json.Unmarshal(respBody, response); err != nil {
		c.logger.Error("failed to unmarshal response: %v", err)
		return fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(respBody))
	}

	return nil
}

// Login authenticates with the camera and obtains a token
func (c *Client) Login(ctx context.Context) error {
	if c.username == "" || c.password == "" {
		return fmt.Errorf("username and password are required")
	}

	c.logger.Info("logging in to camera at %s", c.host)

	req := []Request{{
		Cmd: "Login",
		Param: LoginParam{
			User: LoginUser{
				UserName: c.username,
				Password: c.password,
				Version:  "0", // No encryption
			},
		},
	}}

	var resp []Response
	if err := c.do(ctx, req, &resp); err != nil {
		c.logger.Error("login failed: %v", err)
		return fmt.Errorf("login request failed: %w", err)
	}

	if len(resp) == 0 {
		return fmt.Errorf("empty response")
	}

	// Check for errors
	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		c.logger.Error("login failed with API error: %v", apiErr)
		return apiErr
	}

	// Parse login response
	var loginValue LoginValue
	if err := json.Unmarshal(resp[0].Value, &loginValue); err != nil {
		return fmt.Errorf("failed to parse login response: %w", err)
	}

	// Store token
	c.tokenMu.Lock()
	c.token = loginValue.Token.Name
	c.tokenMu.Unlock()

	c.logger.Info("successfully logged in, token lease time: %d seconds", loginValue.Token.LeaseTime)

	return nil
}

// Logout invalidates the current token
func (c *Client) Logout(ctx context.Context) error {
	c.logger.Info("logging out from camera at %s", c.host)

	req := []Request{{
		Cmd: "Logout",
	}}

	var resp []Response
	if err := c.do(ctx, req, &resp); err != nil {
		c.logger.Error("logout failed: %v", err)
		return fmt.Errorf("logout request failed: %w", err)
	}

	if len(resp) == 0 {
		return fmt.Errorf("empty response")
	}

	// Check for errors
	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		c.logger.Error("logout failed with API error: %v", apiErr)
		return apiErr
	}

	// Clear token
	c.tokenMu.Lock()
	c.token = ""
	c.tokenMu.Unlock()

	c.logger.Info("successfully logged out")

	return nil
}

// GetToken returns the current authentication token
func (c *Client) GetToken() string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.token
}

// SetToken sets the authentication token manually
func (c *Client) SetToken(token string) {
	c.tokenMu.Lock()
	c.token = token
	c.tokenMu.Unlock()
}

// IsAuthenticated returns true if the client has a valid token
func (c *Client) IsAuthenticated() bool {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.token != ""
}

// Host returns the camera host address
func (c *Client) Host() string {
	return c.host
}

// BaseURL returns the base API URL
func (c *Client) BaseURL() string {
	return c.baseURL
}
