# Changelog

All notable changes to the Reolink Camera API Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-10-27

### Breaking Changes

- **Restructured SDK to canonical Go layout**
  - Moved SDK from `sdk/go/reolink/` to repository root
  - Changed module path from `github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink` to `github.com/mosleyit/reolink_api_wrapper`
  - Organized API modules into separate packages under `api/` directory
  - Moved internal implementation to `internal/` directory
  - Moved shared utilities to `pkg/` directory

### Migration Guide

To upgrade from v1.x to v2.0:

1. Update your `go.mod`:
   ```bash
   go get github.com/mosleyit/reolink_api_wrapper@v2
   ```

2. Update your import statement:
   ```go
   // Before (v1.x)
   import "github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
   
   // After (v2.0)
   import "github.com/mosleyit/reolink_api_wrapper"
   ```

3. No other code changes needed! The public API remains 100% compatible.

### Added

- LICENSE file (MIT License)
- CHANGELOG.md for version tracking
- `version.go` with SDK version information
- `doc.go` with comprehensive package documentation
- Improved package organization with clear separation of concerns
- Better internal structure for maintainability

### Changed

- Module path (breaking change - see above)
- Internal package organization (no impact on public API)
- File structure follows canonical Go SDK layout

### Maintained

- ✅ 100% API coverage (130 endpoints across 11 modules)
- ✅ All 269 unit tests passing
- ✅ 60.5% test coverage
- ✅ Complete backward compatibility for public API
- ✅ All functionality unchanged

## [1.0.0] - 2024

### Added

- Initial release of Reolink Camera API Go SDK
- Complete implementation of all Reolink HTTP API endpoints
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
- Comprehensive error handling with 130+ error codes
- Context-aware operations
- Functional options pattern for configuration
- Automatic authentication and token management
- HTTPS/TLS support
- Comprehensive test suite (157 unit tests)
- Hardware validation on real Reolink devices
- Examples for common use cases

---

## Version History

- **v2.0.0** (2025-10-27) - Canonical restructure, breaking import path change
- **v1.0.0** (2024) - Initial release with full API coverage

