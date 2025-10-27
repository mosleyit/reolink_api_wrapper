package reolink

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncodingAPI_GetEnc(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if len(req) != 1 {
			t.Fatalf("Expected 1 request, got %d", len(req))
		}

		if req[0].Cmd != "GetEnc" {
			t.Errorf("Expected cmd 'GetEnc', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetEnc",
			Code: 0,
			Value: json.RawMessage(`{
				"Enc": {
					"audio": 0,
					"channel": 0,
					"mainStream": {
						"bitRate": 4096,
						"frameRate": 20,
						"gop": 2,
						"height": 2160,
						"width": 3840,
						"profile": "High",
						"size": "3840*2160",
						"vType": "h265"
					},
					"subStream": {
						"bitRate": 256,
						"frameRate": 10,
						"gop": 1,
						"height": 360,
						"width": 640,
						"profile": "Main",
						"size": "640*360",
						"vType": "h264"
					}
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Encoding = &EncodingAPI{client: client}

	// Test GetEnc
	ctx := context.Background()
	config, err := client.Encoding.GetEnc(ctx, 0)
	if err != nil {
		t.Fatalf("GetEnc failed: %v", err)
	}

	// Verify response
	if config.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", config.Channel)
	}
	if config.Audio != 0 {
		t.Errorf("Expected audio 0, got %d", config.Audio)
	}
	if config.MainStream.VType != "h265" {
		t.Errorf("Expected main stream codec h265, got %s", config.MainStream.VType)
	}
	if config.MainStream.Width != 3840 {
		t.Errorf("Expected main stream width 3840, got %d", config.MainStream.Width)
	}
	if config.MainStream.Height != 2160 {
		t.Errorf("Expected main stream height 2160, got %d", config.MainStream.Height)
	}
	if config.SubStream.VType != "h264" {
		t.Errorf("Expected sub stream codec h264, got %s", config.SubStream.VType)
	}
}

func TestEncodingAPI_SetEnc(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetEnc" {
			t.Errorf("Expected cmd 'SetEnc', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetEnc",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Encoding = &EncodingAPI{client: client}

	// Test SetEnc
	ctx := context.Background()
	config := EncConfig{
		Channel: 0,
		Audio:   0,
		MainStream: Stream{
			VType:     "h265",
			Size:      "3840*2160",
			FrameRate: 20,
			BitRate:   4096,
			GOP:       2,
			Height:    2160,
			Width:     3840,
			Profile:   "High",
		},
		SubStream: Stream{
			VType:     "h264",
			Size:      "640*360",
			FrameRate: 10,
			BitRate:   256,
			GOP:       1,
			Height:    360,
			Width:     640,
			Profile:   "Main",
		},
	}

	err := client.Encoding.SetEnc(ctx, config)
	if err != nil {
		t.Fatalf("SetEnc failed: %v", err)
	}
}

func TestEncodingAPI_Snap(t *testing.T) {
	// Create mock server that returns a fake JPEG
	fakeJPEG := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46} // JPEG header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Verify URL contains cmd=Snap
		if r.URL.Query().Get("cmd") != "Snap" {
			t.Errorf("Expected cmd=Snap in URL, got %s", r.URL.Query().Get("cmd"))
		}

		// Send fake JPEG
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(fakeJPEG)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Encoding = &EncodingAPI{client: client}

	// Test Snap
	ctx := context.Background()
	imageData, err := client.Encoding.Snap(ctx, 0)
	if err != nil {
		t.Fatalf("Snap failed: %v", err)
	}

	// Verify response
	if len(imageData) != len(fakeJPEG) {
		t.Errorf("Expected %d bytes, got %d", len(fakeJPEG), len(imageData))
	}

	// Verify JPEG header
	if imageData[0] != 0xFF || imageData[1] != 0xD8 {
		t.Errorf("Invalid JPEG header")
	}
}

func TestEncodingAPI_Snap_Error(t *testing.T) {
	// Create mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Encoding = &EncodingAPI{client: client}

	// Test Snap with error
	ctx := context.Background()
	_, err := client.Encoding.Snap(ctx, 0)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
