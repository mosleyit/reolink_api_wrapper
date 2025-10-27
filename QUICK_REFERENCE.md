# Quick Reference: Full Canonical Restructure

## TL;DR

**What:** Restructure Go SDK from `sdk/go/reolink/` to repository root with canonical layout  
**Why:** Industry-standard structure, better organization, simpler imports  
**Impact:** Breaking change - requires v2.0.0  
**User Impact:** Only import path changes, API stays the same  
**Time:** 12-18 hours of work  

---

## Key Changes at a Glance

### 1. Module Path
```diff
- module github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink
+ module github.com/mosleyit/reolink_api_wrapper
```

### 2. Installation
```diff
- go get github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink
+ go get github.com/mosleyit/reolink_api_wrapper@v2
```

### 3. Import Statement
```diff
- import "github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
+ import "github.com/mosleyit/reolink_api_wrapper"
```

### 4. Usage (NO CHANGE)
```go
// Exactly the same!
client := reolink.NewClient("192.168.1.100",
    reolink.WithCredentials("admin", "password"))
```

---

## File Movements

### Core Files → Root
```
sdk/go/reolink/client.go       → client.go
sdk/go/reolink/config.go       → config.go
sdk/go/reolink/errors.go       → errors.go
sdk/go/reolink/go.mod          → go.mod (update module path)
sdk/go/reolink/Makefile        → Makefile
sdk/go/reolink/.gitignore      → .gitignore
```

### API Modules → api/
```
sdk/go/reolink/system.go       → api/system/service.go
sdk/go/reolink/security.go     → api/security/service.go
sdk/go/reolink/network.go      → api/network/service.go
sdk/go/reolink/video.go        → api/video/service.go
sdk/go/reolink/encoding.go     → api/encoding/service.go
sdk/go/reolink/recording.go    → api/recording/service.go
sdk/go/reolink/ptz.go          → api/ptz/service.go
sdk/go/reolink/alarm.go        → api/alarm/service.go
sdk/go/reolink/led.go          → api/led/service.go
sdk/go/reolink/ai.go           → api/ai/service.go
sdk/go/reolink/streaming.go    → api/streaming/service.go
```

### Shared Types → api/common/
```
sdk/go/reolink/models.go       → api/common/types.go
```

### Utilities → pkg/
```
sdk/go/reolink/logger.go       → pkg/logger/logger.go
```

### Internal → internal/
```
sdk/go/reolink/testing.go      → internal/testing/helpers.go
(extract from client.go)       → internal/httpclient/
```

### Examples → examples/
```
sdk/go/reolink/examples/       → examples/ (update imports)
```

### Tests → Stay with source
```
sdk/go/reolink/*_test.go       → Same location as source files
```

---

## New Files to Create

```
LICENSE                        # MIT or Apache 2.0
CHANGELOG.md                   # Version history starting with v2.0.0
version.go                     # SDK version constant
doc.go                         # Package documentation
```

---

## Package Name Changes

| File Location | Old Package | New Package |
|--------------|-------------|-------------|
| Root (client.go, etc.) | `reolink` | `reolink` ✅ |
| api/system/ | `reolink` | `system` ⚠️ |
| api/security/ | `reolink` | `security` ⚠️ |
| api/network/ | `reolink` | `network` ⚠️ |
| api/video/ | `reolink` | `video` ⚠️ |
| api/encoding/ | `reolink` | `encoding` ⚠️ |
| api/recording/ | `reolink` | `recording` ⚠️ |
| api/ptz/ | `reolink` | `ptz` ⚠️ |
| api/alarm/ | `reolink` | `alarm` ⚠️ |
| api/led/ | `reolink` | `led` ⚠️ |
| api/ai/ | `reolink` | `ai` ⚠️ |
| api/streaming/ | `reolink` | `streaming` ⚠️ |
| api/common/ | N/A | `common` 🆕 |
| pkg/logger/ | `reolink` | `logger` ⚠️ |
| internal/httpclient/ | N/A | `httpclient` 🆕 |
| internal/testing/ | `reolink` | `testing` ⚠️ |

---

## Migration Checklist

### Preparation
- [ ] Create backup branch
- [ ] Run baseline tests
- [ ] Create directory structure

### New Files
- [ ] Create LICENSE
- [ ] Create CHANGELOG.md
- [ ] Create version.go
- [ ] Create doc.go

### Extract Packages
- [ ] Create pkg/logger/
- [ ] Create internal/testing/
- [ ] Create internal/httpclient/
- [ ] Create api/common/

### Migrate API Modules
- [ ] api/system/
- [ ] api/security/
- [ ] api/network/
- [ ] api/video/
- [ ] api/encoding/
- [ ] api/recording/
- [ ] api/ptz/
- [ ] api/alarm/
- [ ] api/led/
- [ ] api/ai/
- [ ] api/streaming/

### Update Root
- [ ] Move client.go to root
- [ ] Move config.go to root
- [ ] Move errors.go to root
- [ ] Move tests to root
- [ ] Update go.mod
- [ ] Update Makefile
- [ ] Move .gitignore

### Update Examples
- [ ] Update basic/main.go imports
- [ ] Update debug_test/main.go imports
- [ ] Update hardware_test/main.go imports
- [ ] Update examples/README.md

### Testing
- [ ] Run all unit tests
- [ ] Build all examples
- [ ] Run integration tests
- [ ] Verify go mod tidy
- [ ] Test coverage check

### Documentation
- [ ] Update main README.md
- [ ] Create migration guide
- [ ] Update code examples
- [ ] Update badges

### Cleanup
- [ ] Remove sdk/ directory
- [ ] Final test run
- [ ] Tag v2.0.0

---

## Testing Commands

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run with coverage
go test -cover ./...

# Build all packages
go build ./...

# Build examples
cd examples/basic && go build
cd examples/debug_test && go build
cd examples/hardware_test && go build

# Verify module
go mod verify
go mod tidy

# Vet code
go vet ./...
```

---

## Risk Mitigation

### High Risk
1. **Import path changes** - All users must update
   - Mitigation: Clear migration guide, maintain v1 branch

2. **Package reorganization** - Complex refactor
   - Mitigation: Migrate one module at a time, test after each

3. **Breaking changes** - Requires major version bump
   - Mitigation: Keep public API identical, only internal changes

### Medium Risk
1. **Test failures** - Tests may break during migration
   - Mitigation: Run tests after each phase

2. **Example breakage** - Examples need import updates
   - Mitigation: Update and test each example

### Low Risk
1. **Documentation** - Docs need updates
   - Mitigation: Update as final step

---

## Rollback Plan

If something goes wrong:

1. **Keep v1.x branch active**
   ```bash
   git checkout -b v1-maintenance
   git tag v1.x.x
   ```

2. **Work on v2 in separate branch**
   ```bash
   git checkout -b v2-restructure
   ```

3. **Can revert if needed**
   ```bash
   git checkout main
   git reset --hard v1.x.x
   ```

---

## Success Metrics

✅ All 157 unit tests pass  
✅ Test coverage ≥ 60%  
✅ All 3 examples build and run  
✅ `go mod verify` succeeds  
✅ `go vet ./...` passes  
✅ Public API unchanged  
✅ Documentation updated  
✅ LICENSE added  
✅ CHANGELOG.md created  

---

## Timeline

| Phase | Duration |
|-------|----------|
| Preparation | 30 min |
| Extract packages | 2-3 hours |
| Migrate API modules | 4-6 hours |
| Update root & examples | 2-3 hours |
| Testing | 2-3 hours |
| Documentation | 1-2 hours |
| Cleanup | 30 min |
| **Total** | **12-18 hours** |

---

## Questions to Answer Before Starting

1. **License**: MIT or Apache 2.0?
2. **Version**: v2.0.0 or v1.0.0?
3. **Backward compat**: Keep v1.x branch?
4. **HTTP refactor**: Full extraction or minimal?
5. **Type split**: All in common/ or split by domain?

---

## Post-Migration

1. **Tag release**
   ```bash
   git tag v2.0.0
   git push origin v2.0.0
   ```

2. **Update pkg.go.dev**
   - Auto-updates on tag push
   - Verify docs render

3. **Announce**
   - GitHub release notes
   - Migration guide in README
   - Update badges

4. **Monitor**
   - Watch for issues
   - Be ready to patch
   - Support v1.x if needed

---

## Benefits of Canonical Structure

✅ **Simpler imports** - Shorter, cleaner path  
✅ **Better organization** - Clear separation of concerns  
✅ **Standard layout** - Familiar to Go developers  
✅ **Easier navigation** - Logical file structure  
✅ **Professional** - Industry best practices  
✅ **Maintainable** - Easier to extend and modify  
✅ **Discoverable** - Clear package boundaries  

---

## User Migration Guide

For users upgrading from v1 to v2:

### Step 1: Update go.mod
```bash
go get github.com/mosleyit/reolink_api_wrapper@v2
```

### Step 2: Update imports
```diff
- import "github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
+ import "github.com/mosleyit/reolink_api_wrapper"
```

### Step 3: No code changes needed!
```go
// Everything else stays the same
client := reolink.NewClient("192.168.1.100",
    reolink.WithCredentials("admin", "password"))
```

### Step 4: Test
```bash
go test ./...
```

**That's it!** The API is 100% compatible.

---

## Comparison: Before vs After

### Before (v1.x)
```
✗ Nested 3 levels deep (sdk/go/reolink/)
✗ Long import path
✗ Flat file structure (36 files in one dir)
✗ No LICENSE
✗ No CHANGELOG
✗ Non-standard layout
✓ Works perfectly
✓ 100% API coverage
✓ Well tested
```

### After (v2.0)
```
✓ At repository root
✓ Short import path
✓ Organized by domain
✓ LICENSE included
✓ CHANGELOG.md
✓ Canonical Go layout
✓ Works perfectly
✓ 100% API coverage
✓ Well tested
```

---

## Final Recommendation

**Proceed with full canonical restructure as v2.0.0**

**Pros:**
- Industry-standard structure
- Better long-term maintainability
- Simpler import path
- Professional appearance
- Easier for contributors

**Cons:**
- Breaking change (import path)
- Significant refactoring effort
- Users must update imports

**Verdict:** The benefits far outweigh the one-time migration cost. The public API remains identical, so users only need to update their import statements. This is the right time to make this change before the SDK gains wider adoption.

