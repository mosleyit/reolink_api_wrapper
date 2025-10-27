# Reolink Camera API Documentation

Complete OpenAPI 3.0.3 specification for the Reolink Camera HTTP API (Version 8).

## Overview

This repository contains comprehensive API documentation for Reolink IP cameras, converted from the official PDF documentation into a structured OpenAPI specification format. It includes ready-to-use examples in multiple programming languages.

## Repository Structure

```
reolink_api_wrapper/
├── docs/                          # Documentation files
│   ├── reolink-camera-api-openapi.yaml    # Complete OpenAPI 3.0.3 spec (8,898 lines)
│   ├── api_guide.txt                       # Text conversion of PDF (20,717 lines)
│   └── reolink-camera-http-api-user-guide-v8.pdf  # Original PDF
├── sdk/                           # SDK implementations
│   └── go/reolink/                # Production-ready Go SDK
│       ├── README.md              # SDK documentation
│       └── examples/              # Working Go examples
└── README.md                      # This file
```

## Quick Start

### Option 1: Use the Go SDK (Recommended)

A production-ready Go SDK is available:

```bash
go get github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink
```

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
)

func main() {
    client := reolink.NewClient("192.168.1.100",
        reolink.WithCredentials("admin", "password"))

    ctx := context.Background()
    if err := client.Login(ctx); err != nil {
        log.Fatal(err)
    }
    defer client.Logout(ctx)

    info, err := client.System.GetDeviceInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Camera: %s (Firmware: %s)\n", info.Model, info.FirmVer)
}
```

See [sdk/go/reolink/README.md](sdk/go/reolink/README.md) for complete documentation.

### Option 2: Generate a Client from OpenAPI Spec

Use OpenAPI Generator to create a client in your language:

```bash
# Python
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g python -o ./python-client

# TypeScript/Axios
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g typescript-axios -o ./ts-client

# Java
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g java -o ./java-client
```

### Option 3: Manual Implementation

Use the OpenAPI specification to implement your own client, or see the [Go SDK examples](sdk/go/reolink/examples/) for reference implementations.

## Features

### Complete API Coverage

- **System Commands**: Login, Logout, GetDevInfo, Reboot, Upgrade, etc.
- **Security Commands**: User management, password changes
- **Network Commands**: DDNS, NTP, Network ports, WiFi, P2P, UPnP
- **Video Input Commands**: OSD, ISP, Mask, Crop, Stitch
- **Encoding Commands**: Stream configuration, resolution, bitrate
- **Record Commands**: Recording schedules, playback
- **PTZ Commands**: Pan/Tilt/Zoom control, presets, patrols, patterns
- **Alarm Commands**: Motion detection, AI detection (people, vehicle, pets)
- **LED Commands**: IR lights, white LED, power LED
- **AI Commands**: AI detection configuration and state

### OpenAPI Specification Highlights

- **100% Self-Contained**: No need to reference the original PDF
- **Production-Ready**: Generate API clients in any language
- **IDE-Friendly**: Full autocomplete support with detailed schemas
- **Comprehensive Examples**: Every command includes working examples
- **Version Support**: Both standard and V20 enhanced commands
- **Model-Specific Notes**: Special features for different camera models

### Authentication Methods

- **Token-based**: Long session (3600 second lease time)
- **Basic Authentication**: Short session with credentials in URL

### Supported Protocols

- **HTTP API**: POST requests to `/cgi-bin/api.cgi`
- **RTSP**: Real-Time Streaming Protocol
- **RTMP**: Real-Time Messaging Protocol
- **FLV**: Flash Video streaming

### Video Codecs

- H.264
- H.265

## 📚 Interactive Documentation

### 🌐 View Online (Recommended)

**Live Documentation:** https://mosleyit.github.io/reolink_api_wrapper/

The documentation is hosted on GitHub Pages with multiple viewing options:

1. **[Swagger UI](https://mosleyit.github.io/reolink_api_wrapper/swagger-ui.html)** - Interactive API explorer with "Try it out" functionality
   - Test API endpoints directly from your browser
   - See request/response examples
   - Generate code snippets in multiple languages

2. **[Redoc](https://mosleyit.github.io/reolink_api_wrapper/redoc.html)** - Beautiful three-panel documentation
   - Clean, responsive design
   - Perfect for reading and understanding the API
   - Search functionality

3. **[Go SDK Documentation](https://mosleyit.github.io/reolink_api_wrapper/godoc.html)** - Complete Go SDK API reference
   - Static HTML documentation
   - Also available on [pkg.go.dev](https://pkg.go.dev/github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink)
   - Full type definitions and examples

4. **[Download OpenAPI YAML](https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml)** - Raw specification file
   - Use with your own tools
   - Generate clients in any language
   - Import into Postman, Insomnia, etc.

### 💻 View Locally

You can also run the documentation locally:

#### Option 1: Simple HTTP Server

```bash
# Python 3
cd docs
python -m http.server 8000

# Then open http://localhost:8000 in your browser
```

#### Option 2: Using Redoc CLI

```bash
npx @redocly/cli preview-docs docs/reolink-camera-api-openapi.yaml
```

#### Option 3: Using Swagger UI (Docker)

```bash
docker run -p 8080:8080 -e SWAGGER_JSON=/docs/reolink-camera-api-openapi.yaml \
  -v $(pwd)/docs:/docs swaggerapi/swagger-ui

# Then open http://localhost:8080
```

### 📖 OpenAPI Specification

The [`docs/reolink-camera-api-openapi.yaml`](docs/reolink-camera-api-openapi.yaml) file contains:

- **110+ API endpoints** fully documented
- **Complete request/response schemas** with types and constraints
- **All parameters** documented with descriptions
- **Working examples** for every command
- **Error codes** (-1 to -507) categorized
- **V20 enhanced commands** with schedule tables
- **Streaming protocol details** (RTSP, RTMP, FLV)

## Go SDK

A production-ready Go SDK is available in [sdk/go/reolink/](sdk/go/reolink/). See the [SDK README](sdk/go/reolink/README.md) for:
- Installation instructions
- Complete API documentation
- Usage examples
- Testing guide

The SDK includes working examples in [sdk/go/reolink/examples/](sdk/go/reolink/examples/) demonstrating common use cases.

## API Basics

### Base URL
```
http://<camera-ip>/cgi-bin/api.cgi
```

### Authentication

Two methods are supported:

1. **Token-based (Recommended)**: Login once, use token for subsequent requests (3600 second lease)
2. **Basic Auth**: Include credentials in each request URL

### Request Format
All requests are POST with JSON body:
```json
[
  {
    "cmd": "CommandName",
    "param": {
      "Parameter": {
        "field": "value"
      }
    },
    "token": "your-token-here"
  }
]
```

### Response Format
```json
[
  {
    "cmd": "CommandName",
    "code": 0,
    "value": {
      "Data": {
        "field": "value"
      }
    }
  }
]
```

### Common Error Codes

- `0`: Success
- `-1`: Unknown error
- `-2`: Invalid parameter
- `-3`: Operation failed
- `-6`: Invalid username or password
- `-8`: Invalid token
- `-16`: Need to login first

See the [OpenAPI spec](docs/reolink-camera-api-openapi.yaml) for complete error code documentation (50+ codes).

## Key Features

### Schedule Tables
V20 commands support detailed schedule tables with alarm types:
- **MD**: Motion detection
- **TIMING**: Scheduled recording
- **AI_PEOPLE**: AI people detection
- **AI_VEHICLE**: AI vehicle detection
- **AI_DOG_CAT**: AI pet detection

Each schedule is a 168-character string (7 days × 24 hours).

### Motion Detection Scope
Configurable detection zones using grid tables:
- 80×60 grid (4800 characters)
- 80×45 grid (3600 characters)

### PTZ Operations
19 different PTZ operation types including:
- Up, Down, Left, Right
- Zoom In/Out, Focus In/Out
- Auto, Stop, Patrol, Preset
- And more...

## 🔧 Using with Popular Tools

### Postman

1. Open Postman
2. Click **Import**
3. Select **Link** and paste: `https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml`
4. Postman will create a collection with all endpoints

### Insomnia

1. Open Insomnia
2. Click **Create** → **Import From** → **URL**
3. Paste: `https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml`
4. All endpoints will be imported

### VS Code (REST Client Extension)

1. Install the REST Client extension
2. Create a `.http` file
3. Use the OpenAPI spec as reference for endpoints

### Paw / RapidAPI

Import the OpenAPI YAML file directly into these tools for a complete API client.

## 📊 Statistics

- **Total Endpoints**: 110+
- **Total Lines (OpenAPI)**: 8,898
- **Error Codes Documented**: 50+
- **Command Categories**: 10+
- **API Coverage**: 100%
- **Go SDK**: Production-ready with 60%+ test coverage

## 📌 Version

- **API Version**: 8
- **Documentation Date**: 2023-4
- **OpenAPI Version**: 3.0.3
- **Documentation URL**: https://mosleyit.github.io/reolink_api_wrapper/

## License

This documentation is based on the official Reolink Camera HTTP API User Guide. Please refer to Reolink's terms of service for API usage.

## Contributing

Contributions are welcome! If you find any issues or have improvements, please open an issue or pull request.

## Acknowledgments

- Original documentation by Reolink
- Converted to OpenAPI 3.0.3 specification for developer convenience

