package common

import (
	"encoding/json"
)

// Request represents a single API request command
type Request struct {
	Cmd    string      `json:"cmd"`              // Command name
	Action int         `json:"action,omitempty"` // 0: get value only, 1: get initial, range, and value
	Param  interface{} `json:"param,omitempty"`  // Command-specific parameters
	Token  string      `json:"token,omitempty"`  // Authentication token (added by client)
}

// Response represents a single API response
type Response struct {
	Cmd     string          `json:"cmd"`               // Command name
	Code    int             `json:"code"`              // Response code (0 = success)
	Value   json.RawMessage `json:"value,omitempty"`   // Response data (present when code = 0)
	Error   *ErrorDetail    `json:"error,omitempty"`   // Error details (present when error occurs)
	Initial json.RawMessage `json:"initial,omitempty"` // Initial/default values (when action = 1)
	Range   json.RawMessage `json:"range,omitempty"`   // Valid ranges/options (when action = 1)
}

// ErrorDetail represents detailed error information in a response
type ErrorDetail struct {
	RspCode int    `json:"rspCode"` // Detailed error code
	Detail  string `json:"detail"`  // Error detail message
}

// Channel represents a camera channel
type Channel struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Online int    `json:"online"` // 1 = online, 0 = offline
	Status string `json:"status"`
}

// Schedule represents a time schedule configuration
type Schedule struct {
	Enable int        `json:"enable"`
	Table  [][]string `json:"table"` // 7x48 array representing week schedule
}

// StreamType represents video stream type
type StreamType string

const (
	StreamMain StreamType = "main" // Main stream (high quality)
	StreamSub  StreamType = "sub"  // Sub stream (low quality)
	StreamExt  StreamType = "ext"  // External stream
)

// StreamConfig represents video stream configuration
type StreamConfig struct {
	Channel    int    `json:"channel"`
	MainStream Stream `json:"mainStream"`
	SubStream  Stream `json:"subStream"`
}

// Stream represents individual stream settings
type Stream struct {
	VType     string `json:"vType"`     // Video codec: "h264" or "h265"
	Size      string `json:"size"`      // Resolution: "2560*1440", "1920*1080", etc.
	FrameRate int    `json:"frameRate"` // Frames per second
	BitRate   int    `json:"bitRate"`   // Bitrate in kbps
	GOP       int    `json:"gop"`       // Group of pictures
	Height    int    `json:"height"`    // Video height in pixels
	Width     int    `json:"width"`     // Video width in pixels
	Profile   string `json:"profile"`   // H.264/H.265 profile (Base, Main, High)
}

