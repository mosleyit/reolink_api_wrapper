# Code Transformation Examples

This document shows concrete examples of how code will change during the restructure.

---

## Example 1: System API Module

### Before: `sdk/go/reolink/system.go`

```go
package reolink

import (
    "context"
    "encoding/json"
    "fmt"
)

// SystemAPI provides access to system-related API endpoints
type SystemAPI struct {
    client *Client
}

// GetDeviceInfo retrieves device information
func (s *SystemAPI) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {
    s.client.logger.Debug("getting device info")

    req := []Request{{
        Cmd:    "GetDevInfo",
        Action: 0,
    }}

    var resp []Response
    if err := s.client.do(ctx, req, &resp); err != nil {
        s.client.logger.Error("failed to get device info: %v", err)
        return nil, fmt.Errorf("GetDevInfo request failed: %w", err)
    }

    if len(resp) == 0 {
        return nil, fmt.Errorf("empty response")
    }

    if apiErr := resp[0].ToAPIError(); apiErr != nil {
        return nil, apiErr
    }

    var value DeviceInfoValue
    if err := json.Unmarshal(resp[0].Value, &value); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &value.DevInfo, nil
}
```

### After: `api/system/service.go`

```go
package system

import (
    "context"
    "encoding/json"
    "fmt"
    
    "github.com/mosleyit/reolink_api_wrapper/api/common"
    "github.com/mosleyit/reolink_api_wrapper/internal/httpclient"
    "github.com/mosleyit/reolink_api_wrapper/pkg/logger"
)

// Service provides access to system-related API endpoints
type Service struct {
    client *httpclient.Client
    logger logger.Logger
}

// NewService creates a new system API service
func NewService(client *httpclient.Client, log logger.Logger) *Service {
    return &Service{
        client: client,
        logger: log,
    }
}

// GetDeviceInfo retrieves device information
func (s *Service) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {
    s.logger.Debug("getting device info")

    req := []common.Request{{
        Cmd:    "GetDevInfo",
        Action: 0,
    }}

    var resp []common.Response
    if err := s.client.Do(ctx, req, &resp); err != nil {
        s.logger.Error("failed to get device info: %v", err)
        return nil, fmt.Errorf("GetDevInfo request failed: %w", err)
    }

    if len(resp) == 0 {
        return nil, fmt.Errorf("empty response")
    }

    if apiErr := resp[0].ToAPIError(); apiErr != nil {
        return nil, apiErr
    }

    var value DeviceInfoValue
    if err := json.Unmarshal(resp[0].Value, &value); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &value.DevInfo, nil
}
```

**Key Changes:**
- Package name: `reolink` → `system`
- Struct name: `SystemAPI` → `Service`
- Added imports for `common`, `httpclient`, `logger`
- Client type: `*Client` → `*httpclient.Client`
- Logger type: `Logger` → `logger.Logger`
- Request/Response: `Request` → `common.Request`
- Added `NewService()` constructor

---

## Example 2: System Types

### Before: `sdk/go/reolink/models.go` (excerpt)

```go
package reolink

// DeviceInfo represents device information from GetDevInfo
type DeviceInfo struct {
    B485         int    `json:"B485"`
    IOInputNum   int    `json:"IOInputNum"`
    IOOutputNum  int    `json:"IOOutputNum"`
    AudioNum     int    `json:"audioNum"`
    BuildDay     string `json:"buildDay"`
    CfgVer       string `json:"cfgVer"`
    ChannelNum   int    `json:"channelNum"`
    Detail       string `json:"detail"`
    DiskNum      int    `json:"diskNum"`
    ExactType    string `json:"exactType"`
    FirmVer      string `json:"firmVer"`
    FrameworkVer int    `json:"frameworkVer"`
    HardVer      string `json:"hardVer"`
    Model        string `json:"model"`
    Name         string `json:"name"`
    PakSuffix    string `json:"pakSuffix"`
    Serial       string `json:"serial"`
    Type         string `json:"type"`
    Wifi         int    `json:"wifi"`
}

// DeviceInfoValue wraps DeviceInfo for API response
type DeviceInfoValue struct {
    DevInfo DeviceInfo `json:"DevInfo"`
}
```

### After: `api/system/types.go`

```go
package system

// DeviceInfo represents device information from GetDevInfo
type DeviceInfo struct {
    B485         int    `json:"B485"`
    IOInputNum   int    `json:"IOInputNum"`
    IOOutputNum  int    `json:"IOOutputNum"`
    AudioNum     int    `json:"audioNum"`
    BuildDay     string `json:"buildDay"`
    CfgVer       string `json:"cfgVer"`
    ChannelNum   int    `json:"channelNum"`
    Detail       string `json:"detail"`
    DiskNum      int    `json:"diskNum"`
    ExactType    string `json:"exactType"`
    FirmVer      string `json:"firmVer"`
    FrameworkVer int    `json:"frameworkVer"`
    HardVer      string `json:"hardVer"`
    Model        string `json:"model"`
    Name         string `json:"name"`
    PakSuffix    string `json:"pakSuffix"`
    Serial       string `json:"serial"`
    Type         string `json:"type"`
    Wifi         int    `json:"wifi"`
}

// DeviceInfoValue wraps DeviceInfo for API response
type DeviceInfoValue struct {
    DevInfo DeviceInfo `json:"DevInfo"`
}
```

**Key Changes:**
- Package name: `reolink` → `system`
- Types stay the same (no breaking changes)

---

## Example 3: Common Types

### Before: `sdk/go/reolink/models.go` (excerpt)

```go
package reolink

import "encoding/json"

// Request represents a single API request command
type Request struct {
    Cmd    string      `json:"cmd"`
    Action int         `json:"action,omitempty"`
    Param  interface{} `json:"param,omitempty"`
    Token  string      `json:"token,omitempty"`
}

// Response represents a single API response
type Response struct {
    Cmd     string          `json:"cmd"`
    Code    int             `json:"code"`
    Value   json.RawMessage `json:"value,omitempty"`
    Error   *ErrorDetail    `json:"error,omitempty"`
    Initial json.RawMessage `json:"initial,omitempty"`
    Range   json.RawMessage `json:"range,omitempty"`
}

// ErrorDetail represents detailed error information
type ErrorDetail struct {
    RspCode int    `json:"rspCode"`
    Detail  string `json:"detail"`
}
```

### After: `api/common/types.go`

```go
package common

import "encoding/json"

// Request represents a single API request command
type Request struct {
    Cmd    string      `json:"cmd"`
    Action int         `json:"action,omitempty"`
    Param  interface{} `json:"param,omitempty"`
    Token  string      `json:"token,omitempty"`
}

// Response represents a single API response
type Response struct {
    Cmd     string          `json:"cmd"`
    Code    int             `json:"code"`
    Value   json.RawMessage `json:"value,omitempty"`
    Error   *ErrorDetail    `json:"error,omitempty"`
    Initial json.RawMessage `json:"initial,omitempty"`
    Range   json.RawMessage `json:"range,omitempty"`
}

// ErrorDetail represents detailed error information
type ErrorDetail struct {
    RspCode int    `json:"rspCode"`
    Detail  string `json:"detail"`
}

// ToAPIError converts a Response to an APIError if it contains an error
func (r *Response) ToAPIError() error {
    // Implementation moved from models.go
    // Will need to import the error type from root package
    return nil // placeholder
}
```

**Key Changes:**
- Package name: `reolink` → `common`
- Types stay the same
- Shared across all API modules

---

## Example 4: Main Client

### Before: `sdk/go/reolink/client.go`

```go
package reolink

import (
    "bytes"
    "context"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"
)

// Client represents a Reolink camera API client
type Client struct {
    host       string
    baseURL    string
    httpClient *http.Client
    username   string
    password   string
    token      string
    tokenMu    sync.RWMutex
    useHTTPS   bool
    logger     Logger

    // API modules
    System    *SystemAPI
    Security  *SecurityAPI
    Network   *NetworkAPI
    Video     *VideoAPI
    Encoding  *EncodingAPI
    Recording *RecordingAPI
    PTZ       *PTZAPI
    Alarm     *AlarmAPI
    LED       *LEDAPI
    AI        *AIAPI
    Streaming *StreamingAPI
}

// NewClient creates a new Reolink API client
func NewClient(host string, opts ...Option) *Client {
    c := &Client{
        host:     host,
        useHTTPS: false,
        logger:   &NoOpLogger{},
        httpClient: &http.Client{
            Timeout: 30 * time.Second,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: true,
                },
            },
        },
    }

    for _, opt := range opts {
        opt(c)
    }

    c.updateBaseURL()

    // Initialize API modules
    c.System = &SystemAPI{client: c}
    c.Security = &SecurityAPI{client: c}
    // ... etc

    return c
}
```

### After: `client.go`

```go
package reolink

import (
    "context"
    "time"
    
    "github.com/mosleyit/reolink_api_wrapper/api/system"
    "github.com/mosleyit/reolink_api_wrapper/api/security"
    "github.com/mosleyit/reolink_api_wrapper/api/network"
    "github.com/mosleyit/reolink_api_wrapper/api/video"
    "github.com/mosleyit/reolink_api_wrapper/api/encoding"
    "github.com/mosleyit/reolink_api_wrapper/api/recording"
    "github.com/mosleyit/reolink_api_wrapper/api/ptz"
    "github.com/mosleyit/reolink_api_wrapper/api/alarm"
    "github.com/mosleyit/reolink_api_wrapper/api/led"
    "github.com/mosleyit/reolink_api_wrapper/api/ai"
    "github.com/mosleyit/reolink_api_wrapper/api/streaming"
    "github.com/mosleyit/reolink_api_wrapper/internal/httpclient"
    "github.com/mosleyit/reolink_api_wrapper/pkg/logger"
)

// Client represents a Reolink camera API client
type Client struct {
    httpClient *httpclient.Client
    logger     logger.Logger

    // API modules
    System    *system.Service
    Security  *security.Service
    Network   *network.Service
    Video     *video.Service
    Encoding  *encoding.Service
    Recording *recording.Service
    PTZ       *ptz.Service
    Alarm     *alarm.Service
    LED       *led.Service
    AI        *ai.Service
    Streaming *streaming.Service
}

// NewClient creates a new Reolink API client
func NewClient(host string, opts ...Option) *Client {
    // Create HTTP client
    httpClient := httpclient.New(host, opts...)
    
    // Create logger
    log := logger.NewNoOp()
    
    c := &Client{
        httpClient: httpClient,
        logger:     log,
    }

    // Apply options
    for _, opt := range opts {
        opt(c)
    }

    // Initialize API modules
    c.System = system.NewService(httpClient, c.logger)
    c.Security = security.NewService(httpClient, c.logger)
    c.Network = network.NewService(httpClient, c.logger)
    c.Video = video.NewService(httpClient, c.logger)
    c.Encoding = encoding.NewService(httpClient, c.logger)
    c.Recording = recording.NewService(httpClient, c.logger)
    c.PTZ = ptz.NewService(httpClient, c.logger)
    c.Alarm = alarm.NewService(httpClient, c.logger)
    c.LED = led.NewService(httpClient, c.logger)
    c.AI = ai.NewService(httpClient, c.logger)
    c.Streaming = streaming.NewService(httpClient, c.logger)

    return c
}

// Login authenticates with the camera
func (c *Client) Login(ctx context.Context) error {
    return c.httpClient.Login(ctx)
}

// Logout invalidates the current token
func (c *Client) Logout(ctx context.Context) error {
    return c.httpClient.Logout(ctx)
}
```

**Key Changes:**
- HTTP logic moved to `internal/httpclient`
- API module types: `*SystemAPI` → `*system.Service`
- Imports from new package structure
- Cleaner separation of concerns

---

## Example 5: Example Usage

### Before: `sdk/go/reolink/examples/basic/main.go`

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
    fmt.Printf("Model: %s\n", info.Model)
}
```

### After: `examples/basic/main.go`

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/mosleyit/reolink_api_wrapper"
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
    fmt.Printf("Model: %s\n", info.Model)
}
```

**Key Changes:**
- Import path only: `.../sdk/go/reolink` → `.../reolink_api_wrapper`
- **Everything else stays exactly the same!**

---

## Example 6: Logger Package

### Before: `sdk/go/reolink/logger.go`

```go
package reolink

import (
    "fmt"
    "io"
    "log"
    "os"
)

// Logger is the interface for logging
type Logger interface {
    Debug(msg string, args ...interface{})
    Info(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Error(msg string, args ...interface{})
}

// NoOpLogger does nothing
type NoOpLogger struct{}

func (l *NoOpLogger) Debug(msg string, args ...interface{}) {}
func (l *NoOpLogger) Info(msg string, args ...interface{})  {}
func (l *NoOpLogger) Warn(msg string, args ...interface{})  {}
func (l *NoOpLogger) Error(msg string, args ...interface{}) {}
```

### After: `pkg/logger/logger.go`

```go
package logger

import (
    "fmt"
    "io"
    "log"
    "os"
)

// Logger is the interface for logging
type Logger interface {
    Debug(msg string, args ...interface{})
    Info(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Error(msg string, args ...interface{})
}

// NoOpLogger does nothing
type NoOpLogger struct{}

// NewNoOp creates a new no-op logger
func NewNoOp() Logger {
    return &NoOpLogger{}
}

func (l *NoOpLogger) Debug(msg string, args ...interface{}) {}
func (l *NoOpLogger) Info(msg string, args ...interface{})  {}
func (l *NoOpLogger) Warn(msg string, args ...interface{})  {}
func (l *NoOpLogger) Error(msg string, args ...interface{}) {}
```

**Key Changes:**
- Package name: `reolink` → `logger`
- Added `NewNoOp()` constructor
- Can be imported by any package

---

## Summary of Transformations

| Aspect | Before | After |
|--------|--------|-------|
| **Module Path** | `.../sdk/go/reolink` | `.../reolink_api_wrapper` |
| **Root Package** | `package reolink` | `package reolink` ✅ |
| **API Packages** | `package reolink` | `package system`, etc. |
| **Struct Names** | `SystemAPI` | `system.Service` |
| **Common Types** | In `models.go` | In `api/common/types.go` |
| **Logger** | In `logger.go` | In `pkg/logger/logger.go` |
| **HTTP Client** | In `client.go` | In `internal/httpclient/` |
| **Imports** | Single package | Multiple packages |
| **Public API** | `client.System.GetDeviceInfo()` | `client.System.GetDeviceInfo()` ✅ |

**Key Insight:** The public API remains identical. Users only update imports!

