package reolink

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/mosleyit/reolink_api_wrapper/pkg/logger"
)

func TestWithCredentials(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("testuser", "testpass"))

	if client.username != "testuser" {
		t.Errorf("expected username 'testuser', got '%s'", client.username)
	}

	if client.password != "testpass" {
		t.Errorf("expected password 'testpass', got '%s'", client.password)
	}
}

func TestWithHTTPS(t *testing.T) {
	tests := []struct {
		name     string
		useHTTPS bool
		expected string
	}{
		{
			name:     "HTTPS enabled",
			useHTTPS: true,
			expected: "https://192.168.1.100/cgi-bin/api.cgi",
		},
		{
			name:     "HTTPS disabled",
			useHTTPS: false,
			expected: "http://192.168.1.100/cgi-bin/api.cgi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient("192.168.1.100", WithHTTPS(tt.useHTTPS))

			if client.useHTTPS != tt.useHTTPS {
				t.Errorf("expected useHTTPS %v, got %v", tt.useHTTPS, client.useHTTPS)
			}

			if client.baseURL != tt.expected {
				t.Errorf("expected baseURL '%s', got '%s'", tt.expected, client.baseURL)
			}
		})
	}
}

func TestWithTimeout(t *testing.T) {
	timeout := 60 * time.Second
	client := NewClient("192.168.1.100", WithTimeout(timeout))

	if client.httpClient.Timeout != timeout {
		t.Errorf("expected timeout %v, got %v", timeout, client.httpClient.Timeout)
	}
}

func TestWithInsecureSkipVerify(t *testing.T) {
	client := NewClient("192.168.1.100", WithInsecureSkipVerify(true))

	// The client should have been created successfully
	if client == nil {
		t.Fatal("expected client to be non-nil")
		return
	}

	// The client should be usable (InsecureSkipVerify doesn't force HTTPS)
	if client.httpClient == nil {
		t.Error("expected httpClient to be non-nil")
	}
}

func TestWithLogger(t *testing.T) {
	log := logger.NewStdLogger(nil)
	client := NewClient("192.168.1.100", WithLogger(log))

	if client.logger != log {
		t.Error("expected custom logger to be set")
	}
}

func TestWithLoggerNil(t *testing.T) {
	client := NewClient("192.168.1.100", WithLogger(nil))

	// Should keep the default NoOpLogger
	// We can't check the type directly since it's in another package,
	// but we can verify it's not nil
	if client.logger == nil {
		t.Error("expected logger to be non-nil when nil logger is provided")
	}
}

func TestMultipleOptions(t *testing.T) {
	client := NewClient("192.168.1.100",
		WithCredentials("admin", "password"),
		WithHTTPS(true),
		WithTimeout(30*time.Second),
		WithInsecureSkipVerify(true),
	)

	if client.username != "admin" {
		t.Errorf("expected username 'admin', got '%s'", client.username)
	}

	if client.password != "password" {
		t.Errorf("expected password 'password', got '%s'", client.password)
	}

	if !client.useHTTPS {
		t.Error("expected HTTPS to be enabled")
	}

	if client.httpClient.Timeout != 30*time.Second {
		t.Errorf("expected timeout 30s, got %v", client.httpClient.Timeout)
	}

	if client.baseURL != "https://192.168.1.100/cgi-bin/api.cgi" {
		t.Errorf("expected baseURL 'https://192.168.1.100/cgi-bin/api.cgi', got '%s'", client.baseURL)
	}
}

func TestStreamTypeConstants(t *testing.T) {
	if StreamMain != "main" {
		t.Errorf("expected StreamMain to be 'main', got '%s'", StreamMain)
	}

	if StreamSub != "sub" {
		t.Errorf("expected StreamSub to be 'sub', got '%s'", StreamSub)
	}

	if StreamExt != "ext" {
		t.Errorf("expected StreamExt to be 'ext', got '%s'", StreamExt)
	}
}

func TestWithHTTPClient(t *testing.T) {
	customClient := &http.Client{
		Timeout: 90 * time.Second,
	}
	client := NewClient("192.168.1.100", WithHTTPClient(customClient))

	if client.httpClient != customClient {
		t.Error("expected custom HTTP client to be set")
	}

	if client.httpClient.Timeout != 90*time.Second {
		t.Errorf("expected timeout 90s, got %v", client.httpClient.Timeout)
	}
}

func TestWithTLSConfig(t *testing.T) {
	customTLSConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,
	}
	client := NewClient("192.168.1.100", WithTLSConfig(customTLSConfig))

	if transport, ok := client.httpClient.Transport.(*http.Transport); ok {
		if transport.TLSClientConfig != customTLSConfig {
			t.Error("expected custom TLS config to be set")
		}
		if transport.TLSClientConfig.MinVersion != tls.VersionTLS12 {
			t.Errorf("expected MinVersion TLS 1.2, got %v", transport.TLSClientConfig.MinVersion)
		}
		if transport.TLSClientConfig.MaxVersion != tls.VersionTLS13 {
			t.Errorf("expected MaxVersion TLS 1.3, got %v", transport.TLSClientConfig.MaxVersion)
		}
	} else {
		t.Error("expected http.Transport")
	}
}

func TestWithToken(t *testing.T) {
	token := "test-token-12345"
	client := NewClient("192.168.1.100", WithToken(token))

	if client.token != token {
		t.Errorf("expected token '%s', got '%s'", token, client.token)
	}
}
