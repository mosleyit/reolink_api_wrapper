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
├── examples/                      # Language-specific examples
│   ├── python/                    # Python examples and guides
│   ├── go/                        # Go examples and guides
│   ├── javascript/                # JavaScript/TypeScript examples
│   └── java/                      # Java examples and guides
└── README.md                      # This file
```

## Quick Start

### 1. Choose Your Language

- **[Python](examples/python/)** - Simple and powerful, great for scripting
- **[Go](examples/go/)** - High performance, excellent for services
- **[JavaScript/TypeScript](examples/javascript/)** - Perfect for web apps and Node.js
- **[Java](examples/java/)** - Enterprise-ready, Spring Boot compatible

### 2. Generate a Client (Recommended)

Use OpenAPI Generator to create a type-safe client in your language:

```bash
# Python
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g python -o ./python-client

# Go
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g go -o ./go-client

# TypeScript/Axios
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g typescript-axios -o ./ts-client

# Java
openapi-generator-cli generate -i docs/reolink-camera-api-openapi.yaml -g java -o ./java-client
```

### 3. Or Use Manual Implementation

Each language directory contains ready-to-use example code. See the README in each directory for details.

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

The documentation is hosted on GitHub Pages with three viewing options:

1. **[Swagger UI](https://mosleyit.github.io/reolink_api_wrapper/swagger-ui.html)** - Interactive API explorer with "Try it out" functionality
   - Test API endpoints directly from your browser
   - See request/response examples
   - Generate code snippets in multiple languages

2. **[Redoc](https://mosleyit.github.io/reolink_api_wrapper/redoc.html)** - Beautiful three-panel documentation
   - Clean, responsive design
   - Perfect for reading and understanding the API
   - Search functionality

3. **[Download OpenAPI YAML](https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml)** - Raw specification file
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

## Language-Specific Examples

### Python

```python
from reolink_camera import ReolinkCamera

camera = ReolinkCamera("192.168.1.100", "admin", "password")
camera.login()

# Get device info
info = camera.get_device_info()
print(info)

# Get snapshot
snapshot = camera.get_snapshot()
with open("snapshot.jpg", "wb") as f:
    f.write(snapshot)

camera.logout()
```

See [examples/python/](examples/python/) for complete examples.

### Go

```go
package main

import "github.com/mosleyit/reolink_api_wrapper/examples/go"

func main() {
    camera := NewReolinkCamera("192.168.1.100", "admin", "password")
    camera.Login()

    info, _ := camera.GetDeviceInfo()
    fmt.Printf("Device: %+v\n", info)

    camera.Logout()
}
```

See [examples/go/](examples/go/) for complete examples.

### JavaScript/TypeScript

```javascript
const ReolinkCamera = require('./reolink-camera');

const camera = new ReolinkCamera('192.168.1.100', 'admin', 'password');

async function main() {
  await camera.login();

  const info = await camera.getDeviceInfo();
  console.log(info);

  await camera.logout();
}

main();
```

See [examples/javascript/](examples/javascript/) for complete examples.

### Java

```java
ReolinkCamera camera = new ReolinkCamera("192.168.1.100", "admin", "password");

camera.login();

Map<String, Object> info = camera.getDeviceInfo();
System.out.println(info);

camera.logout();
```

See [examples/java/](examples/java/) for complete examples.

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
- **Completeness**: 100%
- **Languages Supported**: Python, Go, JavaScript/TypeScript, Java

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

