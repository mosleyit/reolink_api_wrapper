package reolink

import (
	"net/http/httptest"

	"github.com/mosleyit/reolink_api_wrapper/pkg/logger"
)

// newTestClient creates a Client for testing with the given test server.
// This is a helper function to ensure all test clients are properly initialized.
func newTestClient(server *httptest.Server) *Client {
	client := &Client{
		baseURL:    server.URL,
		httpClient: server.Client(),
		logger:     logger.NewNoOp(),
	}

	// Initialize all API structs
	client.System = &SystemAPI{client: client}
	client.Security = &SecurityAPI{client: client}
	client.Network = &NetworkAPI{client: client}
	client.Video = &VideoAPI{client: client}
	client.Encoding = &EncodingAPI{client: client}
	client.Recording = &RecordingAPI{client: client}
	client.PTZ = &PTZAPI{client: client}
	client.Alarm = &AlarmAPI{client: client}
	client.LED = &LEDAPI{client: client}
	client.AI = &AIAPI{client: client}
	client.Streaming = &StreamingAPI{client: client}

	return client
}
