package reolink

import (
	"testing"
)

func TestStreamingAPI_GetRTSPURL(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("admin", "password"))

	tests := []struct {
		name       string
		channel    int
		streamType StreamType
		expected   string
	}{
		{
			name:       "Main stream channel 1",
			channel:    1,
			streamType: StreamMain,
			expected:   "rtsp://admin:password@192.168.1.100:554/Preview_01_main",
		},
		{
			name:       "Sub stream channel 1",
			channel:    1,
			streamType: StreamSub,
			expected:   "rtsp://admin:password@192.168.1.100:554/Preview_01_sub",
		},
		{
			name:       "Main stream channel 2",
			channel:    2,
			streamType: StreamMain,
			expected:   "rtsp://admin:password@192.168.1.100:554/Preview_02_main",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := client.Streaming.GetRTSPURL(tt.streamType, tt.channel)
			if url != tt.expected {
				t.Errorf("expected URL '%s', got '%s'", tt.expected, url)
			}
		})
	}
}

func TestStreamingAPI_GetRTMPURL(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("admin", "password"))

	tests := []struct {
		name       string
		channel    int
		streamType StreamType
		expected   string
	}{
		{
			name:       "Main stream channel 0",
			channel:    0,
			streamType: StreamMain,
			expected:   "rtmp://192.168.1.100/bcs/channel0_main.bcs?channel=0&stream=0&user=admin&password=password",
		},
		{
			name:       "Sub stream channel 0",
			channel:    0,
			streamType: StreamSub,
			expected:   "rtmp://192.168.1.100/bcs/channel0_sub.bcs?channel=0&stream=1&user=admin&password=password",
		},
		{
			name:       "Main stream channel 1",
			channel:    1,
			streamType: StreamMain,
			expected:   "rtmp://192.168.1.100/bcs/channel1_main.bcs?channel=1&stream=0&user=admin&password=password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := client.Streaming.GetRTMPURL(tt.streamType, tt.channel)
			if url != tt.expected {
				t.Errorf("expected URL '%s', got '%s'", tt.expected, url)
			}
		})
	}
}

func TestStreamingAPI_GetFLVURL(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("admin", "password"), WithHTTPS(true))

	tests := []struct {
		name       string
		channel    int
		streamType StreamType
		expected   string
	}{
		{
			name:       "Main stream channel 0",
			channel:    0,
			streamType: StreamMain,
			expected:   "https://192.168.1.100/flv?port=1935&app=bcs&stream=channel0_main.bcs&user=admin&password=password",
		},
		{
			name:       "Sub stream channel 0",
			channel:    0,
			streamType: StreamSub,
			expected:   "https://192.168.1.100/flv?port=1935&app=bcs&stream=channel0_sub.bcs&user=admin&password=password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := client.Streaming.GetFLVURL(tt.streamType, tt.channel)
			if url != tt.expected {
				t.Errorf("expected URL '%s', got '%s'", tt.expected, url)
			}
		})
	}
}
