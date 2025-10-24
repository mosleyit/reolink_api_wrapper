# Python Examples for Reolink Camera API

This directory contains Python examples for using the Reolink Camera API.

## Prerequisites

- Python 3.7+
- `requests` library

## Installation

```bash
pip install requests
```

## Option 1: Using the OpenAPI Generator (Recommended)

Generate a Python client from the OpenAPI specification:

```bash
# Install OpenAPI Generator
pip install openapi-generator-cli

# Generate Python client
openapi-generator-cli generate \
  -i ../../docs/reolink-camera-api-openapi.yaml \
  -g python \
  -o ./generated-client \
  --package-name reolink_api

# Install the generated client
cd generated-client
pip install .
```

### Using the Generated Client

```python
from reolink_api import ApiClient, Configuration
from reolink_api.api import default_api

# Configure the client
config = Configuration(
    host="http://192.168.1.100/cgi-bin/api.cgi"
)

# Create API client
with ApiClient(config) as api_client:
    api = default_api.DefaultApi(api_client)
    
    # Login
    login_response = api.login_post([{
        "cmd": "Login",
        "param": {
            "User": {
                "userName": "admin",
                "password": "your_password"
            }
        }
    }])
    
    token = login_response[0]['value']['Token']['name']
    print(f"Token: {token}")
```

## Option 2: Manual Implementation

See `basic_example.py` for a simple implementation using the `requests` library.

### Basic Usage

```python
import requests
import json

class ReolinkCamera:
    def __init__(self, host, username, password):
        self.host = host
        self.username = username
        self.password = password
        self.token = None
        self.base_url = f"http://{host}/cgi-bin/api.cgi"
    
    def login(self):
        """Login and get token"""
        payload = [{
            "cmd": "Login",
            "param": {
                "User": {
                    "userName": self.username,
                    "password": self.password
                }
            }
        }]
        
        response = requests.post(self.base_url, json=payload)
        result = response.json()
        
        if result[0]['code'] == 0:
            self.token = result[0]['value']['Token']['name']
            return self.token
        else:
            raise Exception(f"Login failed: {result[0]['code']}")
    
    def logout(self):
        """Logout and invalidate token"""
        payload = [{
            "cmd": "Logout",
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()
    
    def get_device_info(self):
        """Get device information"""
        payload = [{
            "cmd": "GetDevInfo",
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()
    
    def get_hdd_info(self):
        """Get HDD/SD card information"""
        payload = [{
            "cmd": "GetHddInfo",
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()
    
    def ptz_control(self, operation, speed=32):
        """
        Control PTZ (Pan-Tilt-Zoom)
        
        Operations:
        - Up, Down, Left, Right
        - ZoomInc, ZoomDec (Zoom In/Out)
        - FocusInc, FocusDec (Focus In/Out)
        - Auto, Stop
        """
        payload = [{
            "cmd": "PtzCtrl",
            "param": {
                "channel": 0,
                "op": operation,
                "speed": speed
            },
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()
    
    def get_snapshot(self, channel=0):
        """Get a snapshot image"""
        url = f"http://{self.host}/cgi-bin/api.cgi?cmd=Snap&channel={channel}&token={self.token}"
        response = requests.get(url)
        return response.content
    
    def start_recording(self, channel=0):
        """Start manual recording"""
        payload = [{
            "cmd": "StartRecord",
            "param": {
                "channel": channel
            },
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()
    
    def stop_recording(self, channel=0):
        """Stop manual recording"""
        payload = [{
            "cmd": "StopRecord",
            "param": {
                "channel": channel
            },
            "token": self.token
        }]
        
        response = requests.post(self.base_url, json=payload)
        return response.json()

# Example usage
if __name__ == "__main__":
    camera = ReolinkCamera("192.168.1.100", "admin", "your_password")
    
    try:
        # Login
        token = camera.login()
        print(f"Logged in with token: {token}")
        
        # Get device info
        info = camera.get_device_info()
        print(f"Device: {info}")
        
        # Get HDD info
        hdd = camera.get_hdd_info()
        print(f"Storage: {hdd}")
        
        # Get snapshot
        snapshot = camera.get_snapshot()
        with open("snapshot.jpg", "wb") as f:
            f.write(snapshot)
        print("Snapshot saved to snapshot.jpg")
        
        # Logout
        camera.logout()
        print("Logged out")
        
    except Exception as e:
        print(f"Error: {e}")
```

## Advanced Examples

### Streaming Video (RTSP)

```python
import cv2

# RTSP URL format
rtsp_url = f"rtsp://{username}:{password}@{host}:554/h264Preview_01_main"

# Open video stream
cap = cv2.VideoCapture(rtsp_url)

while True:
    ret, frame = cap.read()
    if not ret:
        break
    
    cv2.imshow('Reolink Camera', frame)
    
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

cap.release()
cv2.destroyAllWindows()
```

### Motion Detection Events

```python
def setup_motion_detection(camera, sensitivity=50):
    """Configure motion detection"""
    payload = [{
        "cmd": "SetMdAlarm",
        "param": {
            "Alarm": {
                "channel": 0,
                "enable": 1,
                "sensitivity": sensitivity,
                "scope": "1" * 4800  # Full screen detection (80x60 grid)
            }
        },
        "token": camera.token
    }]
    
    response = requests.post(camera.base_url, json=payload)
    return response.json()
```

### AI Detection Configuration

```python
def setup_ai_detection(camera, detect_people=True, detect_vehicle=True):
    """Configure AI detection"""
    ai_types = []
    if detect_people:
        ai_types.append("people")
    if detect_vehicle:
        ai_types.append("vehicle")
    
    payload = [{
        "cmd": "SetAiCfg",
        "param": {
            "AiDetection": {
                "channel": 0,
                "ai_type": ai_types
            }
        },
        "token": camera.token
    }]
    
    response = requests.post(camera.base_url, json=payload)
    return response.json()
```

## Error Handling

```python
def handle_api_response(response):
    """Handle API response and errors"""
    result = response.json()
    
    if result[0]['code'] == 0:
        return result[0]['value']
    else:
        error_codes = {
            -1: "Unknown error",
            -2: "Invalid parameter",
            -3: "Operation failed",
            -6: "Invalid username or password",
            -8: "Invalid token"
        }
        
        code = result[0]['code']
        message = error_codes.get(code, f"Error code: {code}")
        raise Exception(message)
```

## See Also

- `basic_example.py` - Simple example with basic operations
- `advanced_example.py` - Advanced features (PTZ, AI, streaming)
- `async_example.py` - Async/await implementation using aiohttp

