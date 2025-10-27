package reolink

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestErrorPaths_EmptyResponse tests empty response handling across APIs
func TestErrorPaths_EmptyResponse(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"GetDeviceInfo", func(c *Client, ctx context.Context) error {
			_, err := c.System.GetDeviceInfo(ctx)
			return err
		}},
		{"GetUsers", func(c *Client, ctx context.Context) error {
			_, err := c.Security.GetUsers(ctx)
			return err
		}},
		{"GetNetPort", func(c *Client, ctx context.Context) error {
			_, err := c.Network.GetNetPort(ctx)
			return err
		}},
		{"GetOsd", func(c *Client, ctx context.Context) error {
			_, err := c.Video.GetOsd(ctx, 0)
			return err
		}},
		{"GetEnc", func(c *Client, ctx context.Context) error {
			_, err := c.Encoding.GetEnc(ctx, 0)
			return err
		}},
		{"GetRec", func(c *Client, ctx context.Context) error {
			_, err := c.Recording.GetRec(ctx, 0)
			return err
		}},
		{"GetPtzPreset", func(c *Client, ctx context.Context) error {
			_, err := c.PTZ.GetPtzPreset(ctx, 0)
			return err
		}},
		{"GetMdState", func(c *Client, ctx context.Context) error {
			_, err := c.Alarm.GetMdState(ctx, 0)
			return err
		}},
		{"GetIrLights", func(c *Client, ctx context.Context) error {
			_, err := c.LED.GetIrLights(ctx)
			return err
		}},
		{"GetAiCfg", func(c *Client, ctx context.Context) error {
			_, err := c.AI.GetAiCfg(ctx, 0)
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Return empty array
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]Response{})
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for empty response, got nil")
			}
			if !errors.Is(err, errors.New("empty response")) && err.Error() != "empty response" {
				// Check if error message contains "empty response"
				if err.Error() != "" && err.Error() != "empty response" {
					// Error message is not empty and not "empty response"
					t.Logf("Got error: %v", err)
				}
			}
		})
	}
}

// TestErrorPaths_APIError tests API error handling across APIs
func TestErrorPaths_APIError(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"GetDeviceInfo", func(c *Client, ctx context.Context) error {
			_, err := c.System.GetDeviceInfo(ctx)
			return err
		}},
		{"GetUsers", func(c *Client, ctx context.Context) error {
			_, err := c.Security.GetUsers(ctx)
			return err
		}},
		{"GetNetPort", func(c *Client, ctx context.Context) error {
			_, err := c.Network.GetNetPort(ctx)
			return err
		}},
		{"GetOsd", func(c *Client, ctx context.Context) error {
			_, err := c.Video.GetOsd(ctx, 0)
			return err
		}},
		{"GetEnc", func(c *Client, ctx context.Context) error {
			_, err := c.Encoding.GetEnc(ctx, 0)
			return err
		}},
		{"GetRec", func(c *Client, ctx context.Context) error {
			_, err := c.Recording.GetRec(ctx, 0)
			return err
		}},
		{"GetPtzPreset", func(c *Client, ctx context.Context) error {
			_, err := c.PTZ.GetPtzPreset(ctx, 0)
			return err
		}},
		{"GetMdState", func(c *Client, ctx context.Context) error {
			_, err := c.Alarm.GetMdState(ctx, 0)
			return err
		}},
		{"GetIrLights", func(c *Client, ctx context.Context) error {
			_, err := c.LED.GetIrLights(ctx)
			return err
		}},
		{"GetAiCfg", func(c *Client, ctx context.Context) error {
			_, err := c.AI.GetAiCfg(ctx, 0)
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := []Response{{
					Cmd:  tt.name,
					Code: 0,
					Error: &ErrorDetail{
						RspCode: ErrCodeLoginRequired,
						Detail:  "login required",
					},
				}}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for API error response, got nil")
			}

			var apiErr *APIError
			if errors.As(err, &apiErr) {
				if apiErr.RspCode != ErrCodeLoginRequired {
					t.Errorf("expected RspCode %d, got %d", ErrCodeLoginRequired, apiErr.RspCode)
				}
			}
		})
	}
}

// TestErrorPaths_InvalidJSON tests invalid JSON handling
func TestErrorPaths_InvalidJSON(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"GetDeviceInfo", func(c *Client, ctx context.Context) error {
			_, err := c.System.GetDeviceInfo(ctx)
			return err
		}},
		{"GetUsers", func(c *Client, ctx context.Context) error {
			_, err := c.Security.GetUsers(ctx)
			return err
		}},
		{"GetNetPort", func(c *Client, ctx context.Context) error {
			_, err := c.Network.GetNetPort(ctx)
			return err
		}},
		{"GetOsd", func(c *Client, ctx context.Context) error {
			_, err := c.Video.GetOsd(ctx, 0)
			return err
		}},
		{"GetEnc", func(c *Client, ctx context.Context) error {
			_, err := c.Encoding.GetEnc(ctx, 0)
			return err
		}},
		{"GetRec", func(c *Client, ctx context.Context) error {
			_, err := c.Recording.GetRec(ctx, 0)
			return err
		}},
		{"GetPtzPreset", func(c *Client, ctx context.Context) error {
			_, err := c.PTZ.GetPtzPreset(ctx, 0)
			return err
		}},
		{"GetMdState", func(c *Client, ctx context.Context) error {
			_, err := c.Alarm.GetMdState(ctx, 0)
			return err
		}},
		{"GetIrLights", func(c *Client, ctx context.Context) error {
			_, err := c.LED.GetIrLights(ctx)
			return err
		}},
		{"GetAiCfg", func(c *Client, ctx context.Context) error {
			_, err := c.AI.GetAiCfg(ctx, 0)
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := []Response{{
					Cmd:   tt.name,
					Code:  0,
					Value: json.RawMessage(`{"invalid": json}`), // Invalid JSON
				}}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for invalid JSON, got nil")
			}
		})
	}
}

// TestErrorPaths_NetworkError tests network error handling
func TestErrorPaths_NetworkError(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"GetDeviceInfo", func(c *Client, ctx context.Context) error {
			_, err := c.System.GetDeviceInfo(ctx)
			return err
		}},
		{"SetDeviceName", func(c *Client, ctx context.Context) error {
			return c.System.SetDeviceName(ctx, "test")
		}},
		{"GetUsers", func(c *Client, ctx context.Context) error {
			_, err := c.Security.GetUsers(ctx)
			return err
		}},
		{"AddUser", func(c *Client, ctx context.Context) error {
			return c.Security.AddUser(ctx, User{UserName: "test"})
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a server that immediately closes connections
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Close connection immediately
				hj, ok := w.(http.Hijacker)
				if ok {
					conn, _, _ := hj.Hijack()
					conn.Close()
				}
			}))
			server.Close() // Close server to cause network errors

			client := NewClient("invalid-host:9999")
			client.baseURL = "http://invalid-host:9999"
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected network error, got nil")
			}
		})
	}
}

// TestErrorPaths_SetOperations tests error paths for Set operations
func TestErrorPaths_SetOperations(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"SetDeviceName", func(c *Client, ctx context.Context) error {
			return c.System.SetDeviceName(ctx, "test")
		}},
		{"SetTime", func(c *Client, ctx context.Context) error {
			return c.System.SetTime(ctx, &TimeConfig{})
		}},
		{"AddUser", func(c *Client, ctx context.Context) error {
			return c.Security.AddUser(ctx, User{UserName: "test"})
		}},
		{"ModifyUser", func(c *Client, ctx context.Context) error {
			return c.Security.ModifyUser(ctx, User{UserName: "test"})
		}},
		{"SetOsd", func(c *Client, ctx context.Context) error {
			return c.Video.SetOsd(ctx, Osd{Channel: 0})
		}},
		{"SetEnc", func(c *Client, ctx context.Context) error {
			return c.Encoding.SetEnc(ctx, EncConfig{Channel: 0})
		}},
		{"SetRec", func(c *Client, ctx context.Context) error {
			return c.Recording.SetRec(ctx, Rec{Channel: 0})
		}},
		{"SetMdAlarm", func(c *Client, ctx context.Context) error {
			return c.Alarm.SetMdAlarm(ctx, MdAlarm{Channel: 0})
		}},
		{"SetIrLights", func(c *Client, ctx context.Context) error {
			return c.LED.SetIrLights(ctx, 0, "Auto")
		}},
		{"SetAiCfg", func(c *Client, ctx context.Context) error {
			return c.AI.SetAiCfg(ctx, AiCfg{Channel: 0})
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_EmptyResponse", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]Response{})
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for empty response, got nil")
			}
		})

		t.Run(tt.name+"_APIError", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := []Response{{
					Cmd:  tt.name,
					Code: 0,
					Error: &ErrorDetail{
						RspCode: ErrCodeParametersError,
						Detail:  "invalid parameters",
					},
				}}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for API error response, got nil")
			}

			var apiErr *APIError
			if errors.As(err, &apiErr) {
				if apiErr.RspCode != ErrCodeParametersError {
					t.Errorf("expected RspCode %d, got %d", ErrCodeParametersError, apiErr.RspCode)
				}
			}
		})
	}
}

// TestErrorPaths_DestructiveOperations tests error paths for destructive operations
func TestErrorPaths_DestructiveOperations(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Client, context.Context) error
	}{
		{"Format", func(c *Client, ctx context.Context) error {
			return c.System.Format(ctx, 0)
		}},
		{"Reboot", func(c *Client, ctx context.Context) error {
			return c.System.Reboot(ctx)
		}},
		{"Restore", func(c *Client, ctx context.Context) error {
			return c.System.Restore(ctx)
		}},
		{"DeleteUser", func(c *Client, ctx context.Context) error {
			return c.Security.DeleteUser(ctx, "testuser")
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_APIError", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := []Response{{
					Cmd:  tt.name,
					Code: 0,
					Error: &ErrorDetail{
						RspCode: ErrCodeOperationTimeout,
						Detail:  "operation timeout",
					},
				}}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}))
			defer server.Close()

			client := newTestClient(server)
			ctx := context.Background()

			err := tt.fn(client, ctx)
			if err == nil {
				t.Error("expected error for API error response, got nil")
			}
		})
	}
}
