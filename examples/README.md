# Reolink Go SDK - Examples

This directory contains example programs demonstrating how to use the Reolink Go SDK.

## Prerequisites

1. **Go 1.19 or later** installed
2. **A Reolink camera** on your network
3. **Camera credentials** (username and password)

## Available Examples

### 1. Basic Example (`basic/`)

**Purpose**: Demonstrates basic SDK usage - login, get device info, logout.

**What it does:**
- Connects to a Reolink camera
- Authenticates and obtains a token
- Retrieves device information
- Gets device name and time configuration
- Properly logs out

**Run it:**
```bash
cd basic
export REOLINK_HOST=192.168.1.100
export REOLINK_USERNAME=admin
export REOLINK_PASSWORD=yourpassword
go run main.go
```

**Expected output:**
```
Logging in...
Logged in successfully. Token: abc123...
Device Information:
  Model: RLC-810A
  Firmware: v3.0.0.136_20121100
  Hardware: IPC_51516M5M
  ...
```

### 2. Hardware Test (`hardware_test/`)

**Purpose**: Comprehensive integration testing against real Reolink hardware.

**What it does:**
- Tests all major API modules
- Validates System, Security, Network, Video, PTZ, Alarm, LED, AI APIs
- Reports success/failure for each endpoint
- Useful for validating SDK against your specific camera model

**Run it:**
```bash
cd hardware_test
export REOLINK_HOST=192.168.1.100
export REOLINK_USERNAME=admin
export REOLINK_PASSWORD=yourpassword
go run main.go
```

**Use cases:**
- Validate SDK compatibility with your camera model
- Test which features your camera supports
- Debug API issues
- Verify firmware compatibility

### 3. Debug Test (`debug_test/`)

**Purpose**: Low-level debugging tool for troubleshooting API issues.

**What it does:**
- Makes direct HTTP requests to the camera
- Compares SDK behavior with raw HTTP requests
- Useful for debugging authentication or API issues

**Run it:**
```bash
cd debug_test
# Edit main.go to set your camera IP and credentials
go run main.go
```

**When to use:**
- Troubleshooting login issues
- Debugging API request/response format
- Comparing SDK behavior with direct HTTP calls
- Investigating camera-specific quirks

## Configuration

All examples support environment variables for configuration:

```bash
# Required
export REOLINK_HOST=192.168.1.100        # Camera IP address
export REOLINK_USERNAME=admin             # Camera username
export REOLINK_PASSWORD=yourpassword      # Camera password

# Optional
export REOLINK_HTTPS=true                 # Use HTTPS (default: false)
export REOLINK_SKIP_VERIFY=true           # Skip TLS verification (default: false)
```

Alternatively, you can hardcode values in the example files (not recommended for production).

## Building Examples

Each example can be built into a standalone binary:

```bash
# Build basic example
cd basic
go build -o basic main.go
./basic

# Build hardware test
cd hardware_test
go build -o hardware_test main.go
./hardware_test
```

## Common Use Cases

### Get Device Information

See `basic/main.go` for a complete example:

```go
info, err := client.System.GetDeviceInfo(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Model: %s, Firmware: %s\n", info.Model, info.FirmVer)
```

### Control PTZ Camera

```go
// Move camera right
err := client.PTZ.StartPTZ(ctx, 0, "Right", 32)
if err != nil {
    log.Fatal(err)
}

// Stop movement after 2 seconds
time.Sleep(2 * time.Second)
err = client.PTZ.StopPTZ(ctx, 0)
```

### Get Stream URLs

```go
// Get RTSP URL for main stream
rtspURL := client.Streaming.GetRTSPURL(reolink.StreamMain, 0)
fmt.Printf("RTSP URL: %s\n", rtspURL)
// Output: rtsp://admin:password@192.168.1.100:554/Preview_01_main

// Get RTMP URL for sub stream
rtmpURL := client.Streaming.GetRTMPURL(reolink.StreamSub, 0)
fmt.Printf("RTMP URL: %s\n", rtmpURL)
```

### Configure Motion Detection

```go
// Get current motion detection config
md, err := client.Alarm.GetMd(ctx, 0)
if err != nil {
    log.Fatal(err)
}

// Enable motion detection
md.Enable = 1
md.Sensitivity = 50

// Update configuration
err = client.Alarm.SetMd(ctx, 0, md)
if err != nil {
    log.Fatal(err)
}
```

### Manage Users

```go
// Get all users
users, err := client.Security.GetUsers(ctx)
if err != nil {
    log.Fatal(err)
}

// Add a new user
err = client.Security.AddUser(ctx, reolink.User{
    UserName: "guest",
    Password: "guestpass",
    Level:    "guest",
})
```

## Troubleshooting

### Connection Issues

**Problem**: `connection refused` or `timeout`

**Solutions:**
- Verify camera IP address is correct
- Ensure camera is on the same network
- Check firewall settings
- Try pinging the camera: `ping 192.168.1.100`

### Authentication Issues

**Problem**: `login failed` or `invalid credentials`

**Solutions:**
- Verify username and password are correct
- Check if camera account is locked (too many failed attempts)
- Try logging in via web interface first
- Use `debug_test` example to see raw API responses

### HTTPS/TLS Issues

**Problem**: `x509: certificate signed by unknown authority`

**Solutions:**
- Use `WithInsecureSkipVerify(true)` for self-signed certificates
- Or install the camera's certificate in your system trust store
- For production, use proper TLS certificates

```go
client := reolink.NewClient(host,
    reolink.WithCredentials(username, password),
    reolink.WithHTTPS(true),
    reolink.WithInsecureSkipVerify(true), // Only for testing!
)
```

### Feature Not Supported

**Problem**: API returns error or unexpected response

**Solutions:**
- Check if your camera model supports the feature
- Verify firmware version compatibility
- Run `hardware_test` to see which APIs your camera supports

## Adding Your Own Examples

To create a new example:

1. Create a new directory: `mkdir my_example`
2. Create `main.go` with your code
3. Add module dependency:
   ```go
   import "github.com/mosleyit/reolink_api_wrapper"
   ```
4. Build and run:
   ```bash
   cd my_example
   go mod init my_example
   go mod edit -replace github.com/mosleyit/reolink_api_wrapper=../..
   go mod tidy
   go run main.go
   ```

## Additional Resources

- **[Main README](../README.md)** - SDK overview and features
- **[pkg.go.dev](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper)** - Complete API reference
- **[GitHub Pages Docs](https://mosleyit.github.io/reolink_api_wrapper/)** - Documentation hub
- **[OpenAPI Spec](../docs/reolink-camera-api-openapi.yaml)** - Complete API specification

## Support

If you encounter issues:

1. Check the [Troubleshooting](#troubleshooting) section above
2. Run the `debug_test` example to see raw API behavior
3. Open an issue: https://github.com/mosleyit/reolink_api_wrapper/issues

## License

These examples are provided as-is for demonstration purposes. Use at your own risk.

