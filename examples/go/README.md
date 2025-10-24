# Go Examples for Reolink Camera API

This directory contains Go examples for using the Reolink Camera API.

## Prerequisites

- Go 1.18+

## Option 1: Using OpenAPI Generator (Recommended)

Generate a Go client from the OpenAPI specification:

```bash
# Install OpenAPI Generator (requires Java)
# Or use Docker:
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
  -i /local/docs/reolink-camera-api-openapi.yaml \
  -g go \
  -o /local/examples/go/generated-client \
  --package-name reolink

# Use the generated client
cd generated-client
go mod init github.com/mosleyit/reolink_api_wrapper/examples/go/generated-client
go mod tidy
```

### Using the Generated Client

```go
package main

import (
    "context"
    "fmt"
    "github.com/mosleyit/reolink_api_wrapper/examples/go/generated-client"
)

func main() {
    cfg := reolink.NewConfiguration()
    cfg.Host = "192.168.1.100"
    cfg.Scheme = "http"
    
    client := reolink.NewAPIClient(cfg)
    ctx := context.Background()
    
    // Login
    loginReq := []map[string]interface{}{
        {
            "cmd": "Login",
            "param": map[string]interface{}{
                "User": map[string]string{
                    "userName": "admin",
                    "password": "your_password",
                },
            },
        },
    }
    
    resp, _, err := client.DefaultApi.LoginPost(ctx).Body(loginReq).Execute()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Response: %+v\n", resp)
}
```

## Option 2: Manual Implementation

See `basic_example.go` for a simple implementation.

### Basic Usage

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type ReolinkCamera struct {
    Host     string
    Username string
    Password string
    Token    string
    BaseURL  string
    Client   *http.Client
}

type APIRequest struct {
    Cmd   string                 `json:"cmd"`
    Param map[string]interface{} `json:"param,omitempty"`
    Token string                 `json:"token,omitempty"`
}

type APIResponse struct {
    Cmd   string                 `json:"cmd"`
    Code  int                    `json:"code"`
    Value map[string]interface{} `json:"value,omitempty"`
}

func NewReolinkCamera(host, username, password string) *ReolinkCamera {
    return &ReolinkCamera{
        Host:     host,
        Username: username,
        Password: password,
        BaseURL:  fmt.Sprintf("http://%s/cgi-bin/api.cgi", host),
        Client:   &http.Client{},
    }
}

func (r *ReolinkCamera) sendRequest(requests []APIRequest) ([]APIResponse, error) {
    jsonData, err := json.Marshal(requests)
    if err != nil {
        return nil, err
    }
    
    resp, err := r.Client.Post(r.BaseURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    var responses []APIResponse
    err = json.Unmarshal(body, &responses)
    if err != nil {
        return nil, err
    }
    
    return responses, nil
}

func (r *ReolinkCamera) Login() error {
    requests := []APIRequest{
        {
            Cmd: "Login",
            Param: map[string]interface{}{
                "User": map[string]string{
                    "userName": r.Username,
                    "password": r.Password,
                },
            },
        },
    }
    
    responses, err := r.sendRequest(requests)
    if err != nil {
        return err
    }
    
    if responses[0].Code != 0 {
        return fmt.Errorf("login failed with code: %d", responses[0].Code)
    }
    
    token := responses[0].Value["Token"].(map[string]interface{})
    r.Token = token["name"].(string)
    
    return nil
}

func (r *ReolinkCamera) Logout() error {
    requests := []APIRequest{
        {
            Cmd:   "Logout",
            Token: r.Token,
        },
    }
    
    _, err := r.sendRequest(requests)
    return err
}

func (r *ReolinkCamera) GetDeviceInfo() (map[string]interface{}, error) {
    requests := []APIRequest{
        {
            Cmd:   "GetDevInfo",
            Token: r.Token,
        },
    }
    
    responses, err := r.sendRequest(requests)
    if err != nil {
        return nil, err
    }
    
    if responses[0].Code != 0 {
        return nil, fmt.Errorf("request failed with code: %d", responses[0].Code)
    }
    
    return responses[0].Value, nil
}

func (r *ReolinkCamera) GetHDDInfo() (map[string]interface{}, error) {
    requests := []APIRequest{
        {
            Cmd:   "GetHddInfo",
            Token: r.Token,
        },
    }
    
    responses, err := r.sendRequest(requests)
    if err != nil {
        return nil, err
    }
    
    if responses[0].Code != 0 {
        return nil, fmt.Errorf("request failed with code: %d", responses[0].Code)
    }
    
    return responses[0].Value, nil
}

func (r *ReolinkCamera) PTZControl(operation string, speed int) error {
    requests := []APIRequest{
        {
            Cmd: "PtzCtrl",
            Param: map[string]interface{}{
                "channel": 0,
                "op":      operation,
                "speed":   speed,
            },
            Token: r.Token,
        },
    }
    
    responses, err := r.sendRequest(requests)
    if err != nil {
        return err
    }
    
    if responses[0].Code != 0 {
        return fmt.Errorf("PTZ control failed with code: %d", responses[0].Code)
    }
    
    return nil
}

func (r *ReolinkCamera) GetSnapshot(channel int) ([]byte, error) {
    url := fmt.Sprintf("http://%s/cgi-bin/api.cgi?cmd=Snap&channel=%d&token=%s",
        r.Host, channel, r.Token)
    
    resp, err := r.Client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    return io.ReadAll(resp.Body)
}

func main() {
    camera := NewReolinkCamera("192.168.1.100", "admin", "your_password")
    
    // Login
    err := camera.Login()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Logged in with token: %s\n", camera.Token)
    
    // Get device info
    info, err := camera.GetDeviceInfo()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Device Info: %+v\n", info)
    
    // Get HDD info
    hdd, err := camera.GetHDDInfo()
    if err != nil {
        panic(err)
    }
    fmt.Printf("HDD Info: %+v\n", hdd)
    
    // Get snapshot
    snapshot, err := camera.GetSnapshot(0)
    if err != nil {
        panic(err)
    }
    
    // Save snapshot
    err = os.WriteFile("snapshot.jpg", snapshot, 0644)
    if err != nil {
        panic(err)
    }
    fmt.Println("Snapshot saved to snapshot.jpg")
    
    // Logout
    err = camera.Logout()
    if err != nil {
        panic(err)
    }
    fmt.Println("Logged out")
}
```

## Advanced Examples

### Concurrent Operations

```go
func (r *ReolinkCamera) GetMultipleInfo() error {
    var wg sync.WaitGroup
    errChan := make(chan error, 3)
    
    wg.Add(3)
    
    go func() {
        defer wg.Done()
        _, err := r.GetDeviceInfo()
        if err != nil {
            errChan <- err
        }
    }()
    
    go func() {
        defer wg.Done()
        _, err := r.GetHDDInfo()
        if err != nil {
            errChan <- err
        }
    }()
    
    go func() {
        defer wg.Done()
        _, err := r.GetSnapshot(0)
        if err != nil {
            errChan <- err
        }
    }()
    
    wg.Wait()
    close(errChan)
    
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

### RTSP Streaming

```go
import (
    "github.com/deepch/vdk/format/rtsp"
)

func streamRTSP(host, username, password string) error {
    rtspURL := fmt.Sprintf("rtsp://%s:%s@%s:554/h264Preview_01_main",
        username, password, host)
    
    session, err := rtsp.Dial(rtspURL)
    if err != nil {
        return err
    }
    defer session.Close()
    
    // Process stream...
    return nil
}
```

## See Also

- `basic_example.go` - Simple example with basic operations
- `advanced_example.go` - Advanced features (PTZ, AI, streaming)
- `concurrent_example.go` - Concurrent operations example

