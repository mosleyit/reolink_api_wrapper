package reolink

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EncodingAPI provides access to encoding/video stream configuration endpoints
type EncodingAPI struct {
	client *Client
}

// EncConfig represents encoding configuration
type EncConfig struct {
	Audio      int    `json:"audio"`      // 0=disabled, 1=enabled
	Channel    int    `json:"channel"`    // Channel number
	MainStream Stream `json:"mainStream"` // Main stream configuration
	SubStream  Stream `json:"subStream"`  // Sub stream configuration
}

// EncValue wraps EncConfig for API response
type EncValue struct {
	Enc EncConfig `json:"Enc"`
}

// EncParam represents parameters for SetEnc
type EncParam struct {
	Enc EncConfig `json:"Enc"`
}

// GetEnc gets encoding configuration for a channel
func (e *EncodingAPI) GetEnc(ctx context.Context, channel int) (*EncConfig, error) {
	e.client.logger.Debug("getting encoding configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetEnc",
		Action: 0, // Get value only
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := e.client.do(ctx, req, &resp); err != nil {
		e.client.logger.Error("failed to get encoding configuration: %v", err)
		return nil, fmt.Errorf("GetEnc request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		e.client.logger.Error("failed to get encoding configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		e.client.logger.Error("failed to get encoding configuration: %v", apiErr)
		return nil, apiErr
	}

	var value EncValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		e.client.logger.Error("failed to parse encoding configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.Enc, nil
}

// SetEnc sets encoding configuration for a channel
func (e *EncodingAPI) SetEnc(ctx context.Context, config EncConfig) error {
	e.client.logger.Info("setting encoding configuration: channel=%d main_res=%dx%d bitrate=%d",
		config.Channel, config.MainStream.Width, config.MainStream.Height, config.MainStream.BitRate)

	req := []Request{{
		Cmd: "SetEnc",
		Param: EncParam{
			Enc: config,
		},
	}}

	var resp []Response
	if err := e.client.do(ctx, req, &resp); err != nil {
		e.client.logger.Error("failed to set encoding configuration: %v", err)
		return fmt.Errorf("SetEnc request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		e.client.logger.Error("failed to set encoding configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		e.client.logger.Error("failed to set encoding configuration: %v", apiErr)
		return apiErr
	}

	e.client.logger.Info("successfully set encoding configuration")
	return nil
}

// Snap captures a snapshot image from the specified channel
// Returns the image data as a byte slice
func (e *EncodingAPI) Snap(ctx context.Context, channel int) ([]byte, error) {
	e.client.logger.Debug("capturing snapshot: channel=%d", channel)

	// Build URL with query parameters
	url := fmt.Sprintf("%s?cmd=Snap&channel=%d&rs=snapshot", e.client.baseURL, channel)

	// Add token if available
	e.client.tokenMu.RLock()
	token := e.client.token
	e.client.tokenMu.RUnlock()

	if token != "" {
		url = fmt.Sprintf("%s&token=%s", url, token)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		e.client.logger.Error("failed to create snapshot request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute request
	httpResp, err := e.client.httpClient.Do(httpReq)
	if err != nil {
		e.client.logger.Error("snapshot request failed: %v", err)
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer httpResp.Body.Close()

	// Check status code
	if httpResp.StatusCode != http.StatusOK {
		err := fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
		e.client.logger.Error("snapshot request failed: %v", err)
		return nil, err
	}

	// Check content type
	contentType := httpResp.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/jpg" {
		err := fmt.Errorf("unexpected content type: %s", contentType)
		e.client.logger.Error("snapshot request failed: %v", err)
		return nil, err
	}

	// Read image data
	imageData, err := io.ReadAll(httpResp.Body)
	if err != nil {
		e.client.logger.Error("failed to read snapshot image data: %v", err)
		return nil, fmt.Errorf("failed to read image data: %w", err)
	}

	e.client.logger.Info("successfully captured snapshot: size=%d bytes", len(imageData))
	return imageData, nil
}
