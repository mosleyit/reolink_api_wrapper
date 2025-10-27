package reolink

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("192.168.1.100")

	if client.host != "192.168.1.100" {
		t.Errorf("expected host 192.168.1.100, got %s", client.host)
	}

	if client.baseURL != "http://192.168.1.100/cgi-bin/api.cgi" {
		t.Errorf("expected baseURL http://192.168.1.100/cgi-bin/api.cgi, got %s", client.baseURL)
	}

	if client.System == nil {
		t.Error("System API not initialized")
	}

	if client.Security == nil {
		t.Error("Security API not initialized")
	}
}

func TestNewClientWithOptions(t *testing.T) {
	client := NewClient("192.168.1.100",
		WithCredentials("admin", "password"),
		WithHTTPS(true),
		WithTimeout(10*time.Second),
	)

	if client.username != "admin" {
		t.Errorf("expected username admin, got %s", client.username)
	}

	if client.password != "password" {
		t.Errorf("expected password password, got %s", client.password)
	}

	if !client.useHTTPS {
		t.Error("expected HTTPS to be enabled")
	}

	if client.baseURL != "https://192.168.1.100/cgi-bin/api.cgi" {
		t.Errorf("expected baseURL https://192.168.1.100/cgi-bin/api.cgi, got %s", client.baseURL)
	}

	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("expected timeout 10s, got %v", client.httpClient.Timeout)
	}
}

func TestLogin(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Parse request
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request: %v", err)
		}

		if len(req) != 1 || req[0].Cmd != "Login" {
			t.Errorf("expected Login command, got %v", req)
		}

		// Return mock response
		resp := []Response{{
			Cmd:   "Login",
			Code:  0,
			Value: json.RawMessage(`{"Token":{"name":"test-token","leaseTime":3600}}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client pointing to mock server
	client := NewClient(server.URL[7:], // Remove "http://"
		WithCredentials("admin", "password"))
	client.baseURL = server.URL // Override baseURL for testing

	// Test login
	ctx := context.Background()
	err := client.Login(ctx)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if client.GetToken() != "test-token" {
		t.Errorf("expected token test-token, got %s", client.GetToken())
	}

	if !client.IsAuthenticated() {
		t.Error("expected client to be authenticated")
	}
}

func TestClient_GetToken(t *testing.T) {
	client := NewClient("192.168.1.100")

	// Initially no token
	if client.GetToken() != "" {
		t.Errorf("expected empty token, got '%s'", client.GetToken())
	}

	// Set a token
	client.token = "test_token_456"

	// Should return the token
	if client.GetToken() != "test_token_456" {
		t.Errorf("expected token 'test_token_456', got '%s'", client.GetToken())
	}
}

func TestClient_IsAuthenticated(t *testing.T) {
	client := NewClient("192.168.1.100")

	// Initially not authenticated
	if client.IsAuthenticated() {
		t.Error("expected client to not be authenticated initially")
	}

	// Set a token
	client.token = "test_token_123"

	// Now should be authenticated
	if !client.IsAuthenticated() {
		t.Error("expected client to be authenticated after setting token")
	}
}

func TestLoginError(t *testing.T) {
	// Create mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "Login",
			Code: 0,
			Error: &ErrorDetail{
				RspCode: ErrCodeLoginError,
				Detail:  "invalid credentials",
			},
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:],
		WithCredentials("admin", "wrong"))
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.Login(ctx)
	if err == nil {
		t.Fatal("expected login to fail")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}

	if apiErr.RspCode != ErrCodeLoginError {
		t.Errorf("expected error code %d, got %d", ErrCodeLoginError, apiErr.RspCode)
	}
}

func TestLogout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "Logout",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.SetToken("test-token")

	ctx := context.Background()
	err := client.Logout(ctx)
	if err != nil {
		t.Fatalf("Logout failed: %v", err)
	}

	if client.GetToken() != "" {
		t.Errorf("expected token to be cleared, got %s", client.GetToken())
	}

	if client.IsAuthenticated() {
		t.Error("expected client to not be authenticated")
	}
}

func TestTokenManagement(t *testing.T) {
	client := NewClient("192.168.1.100")

	if client.IsAuthenticated() {
		t.Error("expected new client to not be authenticated")
	}

	client.SetToken("test-token")

	if !client.IsAuthenticated() {
		t.Error("expected client to be authenticated after setting token")
	}

	if client.GetToken() != "test-token" {
		t.Errorf("expected token test-token, got %s", client.GetToken())
	}
}

func TestClientHost(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{"IP address", "192.168.1.100"},
		{"Hostname", "camera.local"},
		{"FQDN", "camera.example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.host)
			if client.Host() != tt.host {
				t.Errorf("expected host %s, got %s", tt.host, client.Host())
			}
		})
	}
}

func TestClientBaseURL(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		useHTTPS bool
		expected string
	}{
		{"HTTP", "192.168.1.100", false, "http://192.168.1.100/cgi-bin/api.cgi"},
		{"HTTPS", "192.168.1.100", true, "https://192.168.1.100/cgi-bin/api.cgi"},
		{"HTTP with hostname", "camera.local", false, "http://camera.local/cgi-bin/api.cgi"},
		{"HTTPS with hostname", "camera.local", true, "https://camera.local/cgi-bin/api.cgi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.host, WithHTTPS(tt.useHTTPS))
			if client.BaseURL() != tt.expected {
				t.Errorf("expected baseURL %s, got %s", tt.expected, client.BaseURL())
			}
		})
	}
}
