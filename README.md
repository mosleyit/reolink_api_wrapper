# Reolink Camera API - Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/mosleyit/reolink_api_wrapper.svg)](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/mosleyit/reolink_api_wrapper)](https://goreportcard.com/report/github.com/mosleyit/reolink_api_wrapper)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Production-ready Go SDK for the Reolink Camera HTTP API with 100% API coverage.

## Features

- ✅ **100% API Coverage** - All 130 endpoints across 11 modules
- ✅ **Type-Safe** - Comprehensive Go types for all API requests and responses
- ✅ **Well-Tested** - 269 unit tests with 60% coverage
- ✅ **Hardware-Validated** - Tested on real Reolink cameras
- ✅ **Context-Aware** - Full context.Context support for timeouts and cancellation
- ✅ **Production-Ready** - Used in production environments
- ✅ **Comprehensive Documentation** - Complete API documentation and examples

## Installation

```bash
go get github.com/mosleyit/reolink_api_wrapper
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mosleyit/reolink_api_wrapper"
)

func main() {
    // Create client
    client := reolink.NewClient("192.168.1.100",
        reolink.WithCredentials("admin", "password"))

    // Authenticate
    ctx := context.Background()
    if err := client.Login(ctx); err != nil {
        log.Fatal(err)
    }
    defer client.Logout(ctx)

    // Get device information
    info, err := client.System.GetDeviceInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Camera: %s (Firmware: %s)\n", info.Model, info.FirmVer)
}
```

## API Modules

The SDK is organized into 11 domain-specific modules:

| Module | Description | Endpoints |
|--------|-------------|-----------|
| **System** | Device info, time, maintenance, reboot, firmware | 15 |
| **Security** | Users, authentication, certificates, sessions | 12 |
| **Network** | Network config, WiFi, DDNS, NTP, Email, FTP, Push | 10 |
| **Video** | OSD, image settings, ISP, privacy mask, crop | 13 |
| **Encoding** | Stream configuration, snapshots, video encoding | 6 |
| **Recording** | Recording config, search, download, playback | 10 |
| **PTZ** | Pan/Tilt/Zoom control, presets, patrols, guard | 18 |
| **Alarm** | Motion detection, AI alarms, audio alarms, buzzer | 24 |
| **LED** | IR lights, white LED, power LED control | 6 |
| **AI** | AI detection, auto-tracking, auto-focus | 13 |
| **Streaming** | RTSP, RTMP, FLV URL helpers | 3 |

## Examples

See the [examples/](examples/) directory for complete working examples:

- **[basic](examples/basic/)** - Simple example showing authentication and device info
- **[debug_test](examples/debug_test/)** - Debug tool for testing API calls
- **[hardware_test](examples/hardware_test/)** - Comprehensive hardware validation suite

## Documentation

- **[API Documentation](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper)** - Complete Go package documentation
- **[OpenAPI Spec](docs/reolink-camera-api-openapi.yaml)** - Full API specification (8,898 lines)
- **[Online Docs](https://mosleyit.github.io/reolink_api_wrapper/)** - Interactive API documentation
- **[CHANGELOG](CHANGELOG.md)** - Version history and migration guide

## Repository Structure

```
reolink_api_wrapper/
├── *.go                           # SDK source files (root package)
├── *_test.go                      # Unit tests
├── api/                           # API-specific packages
│   └── common/                    # Shared types and utilities
├── pkg/                           # Public packages
│   └── logger/                    # Logger interface and implementations
├── examples/                      # Ready-to-run examples
│   ├── basic/                     # Simple usage example
│   ├── debug_test/                # Debug tool
│   └── hardware_test/             # Hardware validation
├── docs/                          # Documentation files
│   ├── reolink-camera-api-openapi.yaml    # OpenAPI 3.0.3 spec
│   └── ...                        # Additional documentation
├── LICENSE                        # MIT License
├── CHANGELOG.md                   # Version history
└── README.md                      # This file
```

## Configuration Options

The client supports various configuration options:

```go
client := reolink.NewClient("192.168.1.100",
    reolink.WithCredentials("admin", "password"),
    reolink.WithHTTPS(true),
    reolink.WithTimeout(30*time.Second),
    reolink.WithLogger(myLogger),
)
```

## Development

### Quick Start with Make

This project includes a comprehensive Makefile for common development tasks:

```bash
# Show all available commands
make help

# Run tests
make test

# Build all examples
make build

# Run linter
make lint

# Format code
make fmt

# Run all checks (format, lint, test) and build
make all
```

### Common Make Targets

| Target | Description |
|--------|-------------|
| `make test` | Run all tests |
| `make test-coverage` | Run tests with coverage report |
| `make build` | Build all example binaries |
| `make lint` | Run linter |
| `make fmt` | Format all code |
| `make clean` | Remove built binaries and coverage files |
| `make verify` | Run all verification checks (CI-friendly) |
| `make install-tools` | Install development tools |

Run `make help` to see all 37 available targets.

### Manual Testing

You can also run tests directly with Go:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...
```

## Hardware Compatibility

The SDK has been tested and validated on real Reolink hardware:

- Reolink IPC cameras (various models)
- Reolink NVR systems
- Firmware versions 8.x and later

## OpenAPI Specification

The complete API specification is available in [`docs/reolink-camera-api-openapi.yaml`](docs/reolink-camera-api-openapi.yaml):

- **110+ API endpoints** fully documented
- **Complete request/response schemas** with types and constraints
- **Working examples** for every command
- **Error codes** (-1 to -507) categorized
- **Streaming protocol details** (RTSP, RTMP, FLV)

### View Online

**Live Documentation:** <https://mosleyit.github.io/reolink_api_wrapper/>

- **[Swagger UI](https://mosleyit.github.io/reolink_api_wrapper/swagger-ui.html)** - Interactive API explorer
- **[Redoc](https://mosleyit.github.io/reolink_api_wrapper/redoc.html)** - Beautiful documentation

## Generate Clients in Other Languages

Use OpenAPI Generator to create clients in your language:

```bash
# Python
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g python -o ./python-client

# TypeScript/Axios
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g typescript-axios -o ./ts-client

# Java
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g java -o ./java-client
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

Use the OpenAPI specification to implement your own client, or see the [examples/](examples/) for reference implementations.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Based on the official Reolink Camera HTTP API User Guide (Version 8)
- OpenAPI specification: [`docs/reolink-camera-api-openapi.yaml`](docs/reolink-camera-api-openapi.yaml)

## Support

- **Issues**: [GitHub Issues](https://github.com/mosleyit/reolink_api_wrapper/issues)
- **Documentation**: [pkg.go.dev](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper)
- **API Spec**: [OpenAPI Documentation](https://mosleyit.github.io/reolink_api_wrapper/)

## Version History

See [CHANGELOG.md](CHANGELOG.md) for version history and migration guides.

## Related Projects

- **OpenAPI Specification**: Complete API documentation in [`docs/`](docs/)
