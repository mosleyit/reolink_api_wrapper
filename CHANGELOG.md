# Changelog

All notable changes to the Reolink Camera API Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-10-27

### Initial Release

Production-ready Go SDK for the Reolink Camera HTTP API with 100% API coverage.

### Features

- ✅ **100% API Coverage** - All 130 endpoints across 11 modules
  - System API (15 endpoints)
  - Security API (12 endpoints)
  - Network API (10 endpoints)
  - Video API (13 endpoints)
  - Encoding API (6 endpoints)
  - Recording API (10 endpoints)
  - PTZ API (18 endpoints)
  - Alarm API (24 endpoints)
  - LED API (6 endpoints)
  - AI API (13 endpoints)
  - Streaming helpers (3 endpoints)
- ✅ **Type-Safe** - Comprehensive Go types for all API requests and responses
- ✅ **Well-Tested** - 269 unit tests with 60% coverage
- ✅ **Hardware-Validated** - Tested on real Reolink cameras
- ✅ **Context-Aware** - Full context.Context support for timeouts and cancellation
- ✅ **Production-Ready** - Used in production environments
- ✅ **Comprehensive Documentation** - Complete API documentation and examples

### Includes

- LICENSE file (MIT License)
- CHANGELOG.md for version tracking
- `version.go` with SDK version information
- `doc.go` with comprehensive package documentation
- Comprehensive error handling with 130+ error codes
- Functional options pattern for configuration
- Automatic authentication and token management
- HTTPS/TLS support with self-signed certificate support
- Hardware validation on real Reolink devices
- Examples for common use cases (basic, debug, hardware test)
- Makefile with 37 targets for development automation
