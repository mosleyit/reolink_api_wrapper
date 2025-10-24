# Reolink Camera API Documentation

Complete OpenAPI 3.0.3 specification for the Reolink Camera HTTP API (Version 8).

## Overview

This repository contains comprehensive API documentation for Reolink IP cameras, converted from the official PDF documentation into a structured OpenAPI specification format.

## Files

- **reolink-camera-api-openapi.yaml** - Complete OpenAPI 3.0.3 specification (8,898 lines)
  - 110+ API endpoints fully documented
  - Complete request/response schemas
  - All parameters with types, constraints, and descriptions
  - Working examples for every command
  - Error codes (-1 to -507) categorized
  - V20 enhanced commands with schedule tables
  - Streaming protocol details (RTSP, RTMP, FLV)

- **api_guide.txt** - Text conversion of the original PDF (20,717 lines)
  - Source material for the OpenAPI specification
  - Complete API reference with examples

- **reolink-camera-http-api-user-guide-v8.pdf** - Original PDF documentation
  - Official Reolink Camera HTTP API User Guide Version 8
  - Dated 2023-4

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

## Usage

### Using the OpenAPI Specification

The OpenAPI YAML file can be used to:

1. **Generate API clients** in multiple languages (Python, JavaScript, Go, Java, etc.)
2. **Generate interactive documentation** using tools like Swagger UI or Redoc
3. **Validate API requests/responses** in your application
4. **Enable IDE autocomplete** when working with the API

### Example Tools

```bash
# Generate a Python client
openapi-generator-cli generate -i reolink-camera-api-openapi.yaml -g python -o ./python-client

# Generate a JavaScript/TypeScript client
openapi-generator-cli generate -i reolink-camera-api-openapi.yaml -g typescript-axios -o ./ts-client

# Serve interactive documentation
npx @redocly/cli preview-docs reolink-camera-api-openapi.yaml
```

## API Basics

### Base URL
```
http://<camera-ip>/cgi-bin/api.cgi
```

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
    }
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

See the OpenAPI spec for complete error code documentation.

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

## Statistics

- **Total Endpoints**: 110+
- **Total Lines (OpenAPI)**: 8,898
- **Error Codes Documented**: 50+
- **Command Categories**: 10+
- **Completeness**: 100%

## Version

- **API Version**: 8
- **Documentation Date**: 2023-4
- **OpenAPI Version**: 3.0.3

## License

This documentation is based on the official Reolink Camera HTTP API User Guide. Please refer to Reolink's terms of service for API usage.

## Contributing

Contributions are welcome! If you find any issues or have improvements, please open an issue or pull request.

## Acknowledgments

- Original documentation by Reolink
- Converted to OpenAPI 3.0.3 specification for developer convenience

