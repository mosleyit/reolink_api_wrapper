// Package reolink provides a Go client for the Reolink Camera HTTP API.
//
// The SDK supports all Reolink camera API endpoints including system management,
// PTZ control, motion detection, AI features, and video streaming.
//
// # Installation
//
//	go get github.com/mosleyit/reolink_api_wrapper@v2
//
// # Quick Start
//
// Basic usage example:
//
//	package main
//
//	import (
//	    "context"
//	    "fmt"
//	    "log"
//	    "github.com/mosleyit/reolink_api_wrapper"
//	)
//
//	func main() {
//	    // Create client
//	    client := reolink.NewClient("192.168.1.100",
//	        reolink.WithCredentials("admin", "password"))
//
//	    // Authenticate
//	    ctx := context.Background()
//	    if err := client.Login(ctx); err != nil {
//	        log.Fatal(err)
//	    }
//	    defer client.Logout(ctx)
//
//	    // Get device information
//	    info, err := client.System.GetDeviceInfo(ctx)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Camera: %s, Firmware: %s\n", info.Model, info.FirmVer)
//	}
//
// # API Modules
//
// The SDK is organized into domain-specific modules:
//
//   - System: Device info, time, maintenance, reboot, firmware updates
//   - Security: Users, authentication, certificates, online sessions
//   - Network: Network config, WiFi, DDNS, NTP, Email, FTP, Push notifications
//   - Video: OSD, image settings, ISP, privacy mask, crop
//   - Encoding: Stream configuration, snapshots, video encoding
//   - Recording: Recording config, search, download, playback
//   - PTZ: Pan/Tilt/Zoom control, presets, patrols, guard positions
//   - Alarm: Motion detection, AI alarms, audio alarms, buzzer
//   - LED: IR lights, white LED, power LED control
//   - AI: AI detection, auto-tracking, auto-focus
//   - Streaming: RTSP, RTMP, FLV URL helpers
//
// # Configuration Options
//
// The client supports various configuration options:
//
//	client := reolink.NewClient("192.168.1.100",
//	    reolink.WithCredentials("admin", "password"),
//	    reolink.WithHTTPS(true),
//	    reolink.WithTimeout(30*time.Second),
//	    reolink.WithLogger(myLogger),
//	)
//
// # Authentication
//
// The SDK handles authentication automatically:
//
//	// Login to obtain token
//	if err := client.Login(ctx); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Token is automatically included in subsequent requests
//	info, err := client.System.GetDeviceInfo(ctx)
//
//	// Logout when done
//	defer client.Logout(ctx)
//
// # Error Handling
//
// The SDK provides comprehensive error handling:
//
//	info, err := client.System.GetDeviceInfo(ctx)
//	if err != nil {
//	    // Check if it's an API error
//	    if apiErr, ok := err.(*reolink.APIError); ok {
//	        fmt.Printf("API Error: %s (code: %d)\n", apiErr.Message, apiErr.Code)
//	    } else {
//	        fmt.Printf("Network Error: %v\n", err)
//	    }
//	    return
//	}
//
// # Streaming
//
// Get streaming URLs for RTSP, RTMP, or FLV:
//
//	// RTSP URL for main stream
//	rtspURL := client.Streaming.GetRTSPURL(reolink.StreamMain, 1)
//	fmt.Printf("RTSP: %s\n", rtspURL)
//
//	// RTMP URL for sub stream
//	rtmpURL := client.Streaming.GetRTMPURL(reolink.StreamSub, 0)
//	fmt.Printf("RTMP: %s\n", rtmpURL)
//
// # PTZ Control
//
// Control Pan/Tilt/Zoom operations:
//
//	// Move camera right
//	err := client.PTZ.PtzCtrl(ctx, "Right", 32, 0, 1)
//
//	// Go to preset position
//	err := client.PTZ.SetPtzPreset(ctx, "call", 1, "Preset1")
//
// # Motion Detection
//
// Configure and monitor motion detection:
//
//	// Get motion detection state
//	state, err := client.Alarm.GetMdState(ctx, 0)
//	if state.State == 1 {
//	    fmt.Println("Motion detected!")
//	}
//
//	// Configure motion detection
//	alarm := reolink.MdAlarm{
//	    Enable:    1,
//	    Sensitivity: 50,
//	}
//	err := client.Alarm.SetMdAlarm(ctx, alarm)
//
// # Recording
//
// Search and download recordings:
//
//	// Search for recordings
//	results, err := client.Recording.Search(ctx, searchParams)
//
//	// Download recording
//	data, err := client.Recording.Download(ctx, fileName, 0)
//
// # Context Support
//
// All API calls support context for timeout and cancellation:
//
//	// With timeout
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	info, err := client.System.GetDeviceInfo(ctx)
//
// # Logging
//
// Enable logging for debugging:
//
//	logger := reolink.NewStdLogger(os.Stderr)
//	client := reolink.NewClient("192.168.1.100",
//	    reolink.WithCredentials("admin", "password"),
//	    reolink.WithLogger(logger),
//	)
//
// # Hardware Compatibility
//
// The SDK has been tested and validated on real Reolink hardware:
//
//   - Reolink IPC cameras (various models)
//   - Reolink NVR systems
//   - Firmware versions 8.x and later
//
// # API Coverage
//
// The SDK provides 100% coverage of the Reolink HTTP API:
//
//   - 130 API endpoints implemented
//   - 11 API modules
//   - 269 unit tests
//   - 60.5% test coverage
//   - Hardware validated
//
// # Documentation
//
// Complete API documentation is available at:
//
//   - OpenAPI Spec: https://github.com/mosleyit/reolink_api_wrapper/blob/main/docs/reolink-camera-api-openapi.yaml
//   - Online Docs: https://mosleyit.github.io/reolink_api_wrapper/
//   - Examples: https://github.com/mosleyit/reolink_api_wrapper/tree/main/examples
//
// # Version
//
// Current SDK version: 2.0.0
//
// # License
//
// MIT License - see LICENSE file for details
package reolink
