package reolink

import (
	"fmt"
)

// StreamingAPI provides helpers for generating streaming URLs
type StreamingAPI struct {
	client *Client
}

// GetRTSPURL generates an RTSP URL for the specified stream type and channel
//
// The channel parameter is 0-based (e.g., 0, 1, 2) and will be converted to
// the 1-based channel numbers used in RTSP URLs (e.g., 01, 02, 03).
//
// Example:
//
//	url := client.Streaming.GetRTSPURL(reolink.StreamMain, 0)
//	// rtsp://admin:password@192.168.1.100:554/Preview_01_main
func (s *StreamingAPI) GetRTSPURL(streamType StreamType, channel int) string {
	s.client.logger.Debug("generating RTSP URL: stream=%s channel=%d", streamType, channel)

	scheme := "rtsp"
	port := 554

	// Format channel with leading zero (01, 02, etc.)
	// RTSP uses 1-based channel numbers, so add 1 to the 0-based channel parameter
	channelStr := fmt.Sprintf("%02d", channel+1)

	// Build URL with credentials
	var url string
	if s.client.username != "" && s.client.password != "" {
		url = fmt.Sprintf("%s://%s:%s@%s:%d/Preview_%s_%s",
			scheme, s.client.username, s.client.password,
			s.client.host, port, channelStr, streamType)
	} else {
		url = fmt.Sprintf("%s://%s:%d/Preview_%s_%s",
			scheme, s.client.host, port, channelStr, streamType)
	}

	s.client.logger.Debug("generated RTSP URL")
	return url
}

// GetRTMPURL generates an RTMP URL for the specified stream type and channel
//
// Channel IDs start from 0 for RTMP URLs (e.g., 0, 1, 2)
// Only H.264 encoding is supported for RTMP
//
// Example:
//
//	url := client.Streaming.GetRTMPURL(reolink.StreamMain, 0)
//	// rtmp://192.168.1.100/bcs/channel0_main.bcs?channel=0&stream=0&user=admin&password=password
func (s *StreamingAPI) GetRTMPURL(streamType StreamType, channelID int) string {
	s.client.logger.Debug("generating RTMP URL: stream=%s channel=%d", streamType, channelID)

	stream := 0
	if streamType == StreamSub {
		stream = 1
	}

	url := fmt.Sprintf("rtmp://%s/bcs/channel%d_%s.bcs?channel=%d&stream=%d&user=%s&password=%s",
		s.client.host, channelID, streamType, channelID, stream,
		s.client.username, s.client.password)

	s.client.logger.Debug("generated RTMP URL")
	return url
}

// GetFLVURL generates an FLV URL for the specified stream type and channel
//
// Channel IDs start from 0 for FLV URLs (e.g., 0, 1, 2)
// Only H.264 encoding is supported for FLV
//
// Example:
//
//	url := client.Streaming.GetFLVURL(reolink.StreamMain, 0)
//	// https://192.168.1.100/flv?port=1935&app=bcs&stream=channel0_main.bcs&user=admin&password=password
func (s *StreamingAPI) GetFLVURL(streamType StreamType, channelID int) string {
	s.client.logger.Debug("generating FLV URL: stream=%s channel=%d", streamType, channelID)

	scheme := "http"
	if s.client.useHTTPS {
		scheme = "https"
	}

	url := fmt.Sprintf("%s://%s/flv?port=1935&app=bcs&stream=channel%d_%s.bcs&user=%s&password=%s",
		scheme, s.client.host, channelID, streamType,
		s.client.username, s.client.password)

	s.client.logger.Debug("generated FLV URL")
	return url
}
