# Reolink Camera API - Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink.svg)](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink)
[![Go Report Card](https://goreportcard.com/badge/github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink)](https://goreportcard.com/report/github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink)

Production-ready Go SDK for the Reolink Camera HTTP API with **100% API coverage** (130 endpoints).

## Status

✅ **Production Ready** - All API endpoints implemented and tested.

**Implementation Status:**
- ✅ **100% API Coverage** (130/130 endpoints)
- ✅ **60.5% Test Coverage** (157 unit tests)
- ✅ **Hardware Tested** on real Reolink cameras
- ✅ **Type-safe** with comprehensive error handling
- ✅ **Context-aware** operations
- ✅ **Production-ready** authentication and token management

**API Modules:**
- ✅ System API (15/15 endpoints) - Device info, time, maintenance, reboot
- ✅ Security API (12/12 endpoints) - Users, authentication, certificates
- ✅ Network API (10/10 endpoints) - Network, WiFi, DDNS, NTP, Email, FTP, Push
- ✅ Video API (13/13 endpoints) - OSD, image settings, ISP, mask, crop
- ✅ Encoding API (6/6 endpoints) - Stream configuration, snapshots
- ✅ Recording API (10/10 endpoints) - Recording config, search, download
- ✅ PTZ API (18/18 endpoints) - Pan/Tilt/Zoom, presets, patrols, guard
- ✅ Alarm API (24/24 endpoints) - Motion, AI, audio, buzzer, push
- ✅ LED API (6/6 endpoints) - IR lights, white LED, power LED
- ✅ AI API (13/13 endpoints) - AI detection, auto-tracking, auto-focus
- ✅ Streaming Helpers (3/3 endpoints) - RTSP, RTMP, FLV URL generation

## Features

- ✅ **Complete API Coverage** - All 130 Reolink HTTP API endpoints
- ✅ **Type-safe** client with comprehensive error handling
- ✅ **Context-aware** operations with timeout support
- ✅ **Functional options** pattern for flexible configuration
- ✅ **Automatic authentication** and token management
- ✅ **HTTPS/TLS support** with custom certificate handling
- ✅ **Streaming helpers** for RTSP, RTMP, and FLV URLs
- ✅ **130+ error codes** with descriptive messages
- ✅ **Comprehensive tests** - 157 unit tests, hardware validated
- ✅ **Production-ready** - Used with real Reolink IPC and NVR devices

## Installation

```bash
go get github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
)

func main() {
    // Create client
    client := reolink.NewClient("192.168.1.100",
        reolink.WithCredentials("admin", "password"),
        reolink.WithHTTPS(false),
    )

    ctx := context.Background()

    // Login
    if err := client.Login(ctx); err != nil {
        log.Fatal(err)
    }
    defer client.Logout(ctx)

    // Get device information
    info, err := client.System.GetDeviceInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Camera Model: %s\n", info.Model)
    fmt.Printf("Firmware: %s\n", info.FirmVer)
    fmt.Printf("Hardware: %s\n", info.HardVer)
}
```

## Documentation

- 💡 **[Examples](examples/)** - Complete usage examples for all major features
- 📖 **[OpenAPI Spec](../../../docs/reolink-camera-api-openapi.yaml)** - Complete API specification

## API Modules

The SDK is organized into functional modules with complete API coverage:

```go
client.System      // 15 endpoints: Device info, time, maintenance, reboot, upgrade
client.Security    // 12 endpoints: Users, authentication, certificates, encryption
client.Network     // 10 endpoints: Network, WiFi, DDNS, NTP, Email, FTP, Push, P2P, UPnP
client.Video       // 13 endpoints: OSD, image settings, ISP, mask, crop, zoom focus
client.Encoding    // 6 endpoints:  Stream configuration, snapshots
client.Recording   // 10 endpoints: Recording config, schedule, search, download
client.PTZ         // 18 endpoints: Pan/Tilt/Zoom, presets, patrols, guard, calibration
client.Alarm       // 24 endpoints: Motion, AI, audio, buzzer, push notifications
client.LED         // 6 endpoints:  IR lights, white LED, power LED
client.AI          // 13 endpoints: AI detection, auto-tracking, auto-focus
client.Streaming   // 3 helpers:    RTSP, RTMP, FLV URL generation
```

**Total: 130 endpoints fully implemented and tested.**

## Examples

### Get Device Information

```go
info, err := client.System.GetDeviceInfo(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Model: %s, Firmware: %s\n", info.Model, info.FirmVer)
```

### Get Device Name

```go
name, err := client.System.GetDeviceName(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Device Name: %s\n", name)
```

### Set Device Name

```go
err := client.System.SetDeviceName(ctx, "Front Door Camera")
if err != nil {
    log.Fatal(err)
}
```

### Get Time Configuration

```go
timeConfig, err := client.System.GetTime(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Time: %04d-%02d-%02d %02d:%02d:%02d\n",
    timeConfig.Year, timeConfig.Mon, timeConfig.Day,
    timeConfig.Hour, timeConfig.Min, timeConfig.Sec)
```

### Get RTSP Stream URL

```go
url := client.Streaming.GetRTSPURL(reolink.StreamMain, 1)
fmt.Printf("RTSP URL: %s\n", url)
// Output: rtsp://admin:password@192.168.1.100:554/Preview_01_main
```

### User Management

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
if err != nil {
    log.Fatal(err)
}
```

## Configuration Options

The SDK uses the functional options pattern for flexible configuration:

```go
client := reolink.NewClient("192.168.1.100",
    // Authentication (required)
    reolink.WithCredentials("admin", "password"),

    // HTTPS (default: false)
    reolink.WithHTTPS(true),

    // Custom timeout (default: 30s)
    reolink.WithTimeout(60 * time.Second),

    // Custom HTTP client (for advanced use cases)
    reolink.WithHTTPClient(customClient),

    // Custom TLS configuration
    reolink.WithTLSConfig(&tls.Config{
        InsecureSkipVerify: false,
        MinVersion:         tls.VersionTLS12,
    }),

    // Skip TLS verification (for self-signed certs - not recommended for production)
    reolink.WithInsecureSkipVerify(true),

    // Set existing token (skip login)
    reolink.WithToken("existing-token"),

    // Custom logger (default: NoOpLogger)
    reolink.WithLogger(customLogger),
)
```

### Security Best Practices

⚠️ **Important Security Considerations:**

1. **Use HTTPS in Production**: Always use `WithHTTPS(true)` for production deployments
2. **Avoid InsecureSkipVerify**: Only use `WithInsecureSkipVerify(true)` for testing with self-signed certificates
3. **Secure Credentials**: Never hardcode credentials - use environment variables or secure vaults
4. **Token Management**: Tokens are managed automatically, but you can retrieve them with `client.GetToken()`
5. **Network Security**: Ensure cameras are on a secure network, not exposed to the internet

**Recommended Production Configuration:**

```go
client := reolink.NewClient(os.Getenv("REOLINK_HOST"),
    reolink.WithCredentials(
        os.Getenv("REOLINK_USERNAME"),
        os.Getenv("REOLINK_PASSWORD"),
    ),
    reolink.WithHTTPS(true),
    reolink.WithTLSConfig(&tls.Config{
        MinVersion: tls.VersionTLS12,
    }),
    reolink.WithTimeout(30 * time.Second),
)
```

## Error Handling

The SDK provides typed errors for better error handling:

```go
info, err := client.System.GetDeviceInfo(ctx)
if err != nil {
    var apiErr *reolink.APIError
    if errors.As(err, &apiErr) {
        switch apiErr.RspCode {
        case reolink.ErrCodeLoginRequired:
            // Re-authenticate
            client.Login(ctx)
        case reolink.ErrCodeParametersError:
            // Fix parameters
            log.Printf("Invalid parameters: %s", apiErr.Detail)
        default:
            log.Printf("API error: %s (rspCode: %d)", apiErr.Error(), apiErr.RspCode)
        }
    } else {
        // Network or other error
        log.Fatal(err)
    }
}
```

## Testing

### Unit Tests

The SDK includes 157 comprehensive unit tests with 60.5% code coverage:

```bash
# Run all tests
go test ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run with verbose output
go test -v ./...

# Run specific test
go test -v -run TestSystemAPI_GetDeviceInfo ./...
```

### Integration Tests

Integration tests validate the SDK against real Reolink hardware.

**Hardware tested:**
- Reolink RLC-810A (IPC)
- Reolink RLN8-410 (NVR)

To run integration tests with your camera:

```bash
export REOLINK_HOST=192.168.1.100
export REOLINK_USERNAME=admin
export REOLINK_PASSWORD=yourpassword

# Run hardware tests
cd examples/hardware_test
go run main.go
```

See [examples/](examples/) for more integration test examples.

## Development

The SDK uses standard Go development practices:

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run linter (if installed)
golangci-lint run ./...

# Format code
go fmt ./...
```

## API Reference

Complete API documentation is available at:
- **OpenAPI Spec**: [docs/reolink-camera-api-openapi.yaml](../../../docs/reolink-camera-api-openapi.yaml)
- **Online Docs**: https://mosleyit.github.io/reolink_api_wrapper/

## Contributing

Contributions are welcome! This SDK has 100% API coverage, but there's always room for improvement:

**Areas for contribution:**
- Additional integration tests with different camera models
- Performance optimizations
- Additional examples and use cases
- Documentation improvements
- Bug fixes and error handling improvements

**Contribution process:**
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for new features (maintain >60% coverage)
4. Update documentation
5. Submit a pull request

**Development setup:**
```bash
# Clone the repository
git clone https://github.com/mosleyit/reolink_api_wrapper.git
cd reolink_api_wrapper/sdk/go/reolink

# Run tests
go test ./...

# Run linter (if installed)
golangci-lint run ./...
```

## Roadmap

- [x] Complete API implementation (130/130 endpoints)
- [x] Comprehensive unit tests (157 tests, 60.5% coverage)
- [x] Hardware validation on real cameras
- [x] Production-ready error handling
- [ ] Increase test coverage to 80%+
- [ ] WebSocket support for real-time events
- [ ] Batch operations for NVR multi-channel management
- [ ] Performance benchmarks and optimizations

## License

This SDK is based on the official Reolink Camera HTTP API. Please refer to Reolink's terms of service for API usage.

**Note:** This is an unofficial SDK and is not affiliated with or endorsed by Reolink.

## Support

- **Issues**: https://github.com/mosleyit/reolink_api_wrapper/issues
- **Discussions**: https://github.com/mosleyit/reolink_api_wrapper/discussions
- **Documentation**: See [docs/](../../../docs/) directory
- **API Reference**: [OpenAPI Specification](../../../docs/reolink-camera-api-openapi.yaml)

## Acknowledgments

- Based on the official **Reolink Camera HTTP API User Guide (Version 8)**
- OpenAPI specification: [reolink-camera-api-openapi.yaml](../../../docs/reolink-camera-api-openapi.yaml)
- Tested on real Reolink hardware (RLC-810A, RLN8-410)

## Related Projects

- **Python SDK**: [sdk/python/](../../python/) (if available)
- **JavaScript SDK**: [sdk/javascript/](../../javascript/) (if available)
- **Java SDK**: [sdk/java/](../../java/) (if available)

---

**Made with ❤️ for the Reolink community**

