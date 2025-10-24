# JavaScript/TypeScript Examples for Reolink Camera API

This directory contains JavaScript and TypeScript examples for using the Reolink Camera API.

## Prerequisites

- Node.js 14+
- npm or yarn

## Option 1: Using OpenAPI Generator (Recommended)

Generate a TypeScript/JavaScript client from the OpenAPI specification:

```bash
# Install OpenAPI Generator
npm install -g @openapitools/openapi-generator-cli

# Generate TypeScript client with Axios
openapi-generator-cli generate \
  -i ../../docs/reolink-camera-api-openapi.yaml \
  -g typescript-axios \
  -o ./generated-client

# Install dependencies
cd generated-client
npm install
```

### Using the Generated Client (TypeScript)

```typescript
import { Configuration, DefaultApi } from './generated-client';

const config = new Configuration({
  basePath: 'http://192.168.1.100/cgi-bin'
});

const api = new DefaultApi(config);

async function main() {
  try {
    // Login
    const loginResponse = await api.loginPost([{
      cmd: 'Login',
      param: {
        User: {
          userName: 'admin',
          password: 'your_password'
        }
      }
    }]);
    
    const token = loginResponse.data[0].value.Token.name;
    console.log(`Token: ${token}`);
    
  } catch (error) {
    console.error('Error:', error);
  }
}

main();
```

## Option 2: Manual Implementation

### Installation

```bash
npm install axios
# or
yarn add axios
```

### Basic Usage (JavaScript)

```javascript
const axios = require('axios');

class ReolinkCamera {
  constructor(host, username, password) {
    this.host = host;
    this.username = username;
    this.password = password;
    this.token = null;
    this.baseURL = `http://${host}/cgi-bin/api.cgi`;
  }

  async sendRequest(requests) {
    try {
      const response = await axios.post(this.baseURL, requests);
      return response.data;
    } catch (error) {
      throw new Error(`API request failed: ${error.message}`);
    }
  }

  async login() {
    const requests = [{
      cmd: 'Login',
      param: {
        User: {
          userName: this.username,
          password: this.password
        }
      }
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`Login failed with code: ${responses[0].code}`);
    }

    this.token = responses[0].value.Token.name;
    return this.token;
  }

  async logout() {
    const requests = [{
      cmd: 'Logout',
      token: this.token
    }];

    return await this.sendRequest(requests);
  }

  async getDeviceInfo() {
    const requests = [{
      cmd: 'GetDevInfo',
      token: this.token
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`Request failed with code: ${responses[0].code}`);
    }

    return responses[0].value;
  }

  async getHDDInfo() {
    const requests = [{
      cmd: 'GetHddInfo',
      token: this.token
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`Request failed with code: ${responses[0].code}`);
    }

    return responses[0].value;
  }

  async ptzControl(operation, speed = 32) {
    const requests = [{
      cmd: 'PtzCtrl',
      param: {
        channel: 0,
        op: operation,
        speed: speed
      },
      token: this.token
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`PTZ control failed with code: ${responses[0].code}`);
    }

    return responses[0];
  }

  async getSnapshot(channel = 0) {
    const url = `http://${this.host}/cgi-bin/api.cgi?cmd=Snap&channel=${channel}&token=${this.token}`;
    
    const response = await axios.get(url, {
      responseType: 'arraybuffer'
    });

    return response.data;
  }

  async startRecording(channel = 0) {
    const requests = [{
      cmd: 'StartRecord',
      param: {
        channel: channel
      },
      token: this.token
    }];

    return await this.sendRequest(requests);
  }

  async stopRecording(channel = 0) {
    const requests = [{
      cmd: 'StopRecord',
      param: {
        channel: channel
      },
      token: this.token
    }];

    return await this.sendRequest(requests);
  }
}

// Example usage
async function main() {
  const camera = new ReolinkCamera('192.168.1.100', 'admin', 'your_password');

  try {
    // Login
    const token = await camera.login();
    console.log(`Logged in with token: ${token}`);

    // Get device info
    const info = await camera.getDeviceInfo();
    console.log('Device Info:', info);

    // Get HDD info
    const hdd = await camera.getHDDInfo();
    console.log('HDD Info:', hdd);

    // Get snapshot
    const snapshot = await camera.getSnapshot();
    const fs = require('fs');
    fs.writeFileSync('snapshot.jpg', snapshot);
    console.log('Snapshot saved to snapshot.jpg');

    // Logout
    await camera.logout();
    console.log('Logged out');

  } catch (error) {
    console.error('Error:', error.message);
  }
}

main();
```

### TypeScript Implementation

```typescript
import axios, { AxiosInstance } from 'axios';

interface APIRequest {
  cmd: string;
  param?: Record<string, any>;
  token?: string;
  action?: number;
}

interface APIResponse {
  cmd: string;
  code: number;
  value?: Record<string, any>;
}

class ReolinkCamera {
  private host: string;
  private username: string;
  private password: string;
  private token: string | null = null;
  private baseURL: string;
  private client: AxiosInstance;

  constructor(host: string, username: string, password: string) {
    this.host = host;
    this.username = username;
    this.password = password;
    this.baseURL = `http://${host}/cgi-bin/api.cgi`;
    this.client = axios.create({
      baseURL: this.baseURL,
      timeout: 10000
    });
  }

  private async sendRequest(requests: APIRequest[]): Promise<APIResponse[]> {
    try {
      const response = await this.client.post<APIResponse[]>('', requests);
      return response.data;
    } catch (error) {
      throw new Error(`API request failed: ${error}`);
    }
  }

  async login(): Promise<string> {
    const requests: APIRequest[] = [{
      cmd: 'Login',
      param: {
        User: {
          userName: this.username,
          password: this.password
        }
      }
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`Login failed with code: ${responses[0].code}`);
    }

    this.token = responses[0].value!.Token.name;
    return this.token;
  }

  async logout(): Promise<void> {
    if (!this.token) {
      throw new Error('Not logged in');
    }

    const requests: APIRequest[] = [{
      cmd: 'Logout',
      token: this.token
    }];

    await this.sendRequest(requests);
    this.token = null;
  }

  async getDeviceInfo(): Promise<Record<string, any>> {
    if (!this.token) {
      throw new Error('Not logged in');
    }

    const requests: APIRequest[] = [{
      cmd: 'GetDevInfo',
      token: this.token
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`Request failed with code: ${responses[0].code}`);
    }

    return responses[0].value!;
  }

  async ptzControl(operation: string, speed: number = 32): Promise<void> {
    if (!this.token) {
      throw new Error('Not logged in');
    }

    const requests: APIRequest[] = [{
      cmd: 'PtzCtrl',
      param: {
        channel: 0,
        op: operation,
        speed: speed
      },
      token: this.token
    }];

    const responses = await this.sendRequest(requests);
    
    if (responses[0].code !== 0) {
      throw new Error(`PTZ control failed with code: ${responses[0].code}`);
    }
  }
}

export default ReolinkCamera;
```

## Advanced Examples

### Browser Usage (Fetch API)

```javascript
class ReolinkCameraBrowser {
  constructor(host, username, password) {
    this.host = host;
    this.username = username;
    this.password = password;
    this.token = null;
    this.baseURL = `http://${host}/cgi-bin/api.cgi`;
  }

  async sendRequest(requests) {
    const response = await fetch(this.baseURL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requests)
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  }

  // ... rest of the methods similar to above
}
```

### React Hook Example

```typescript
import { useState, useEffect } from 'react';
import ReolinkCamera from './ReolinkCamera';

function useCameraSnapshot(host: string, username: string, password: string) {
  const [snapshot, setSnapshot] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const captureSnapshot = async () => {
    setLoading(true);
    setError(null);

    try {
      const camera = new ReolinkCamera(host, username, password);
      await camera.login();
      
      const snapshotData = await camera.getSnapshot();
      const blob = new Blob([snapshotData], { type: 'image/jpeg' });
      const url = URL.createObjectURL(blob);
      
      setSnapshot(url);
      
      await camera.logout();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  return { snapshot, loading, error, captureSnapshot };
}

export default useCameraSnapshot;
```

## See Also

- `basic_example.js` - Simple Node.js example
- `typescript_example.ts` - TypeScript implementation
- `browser_example.html` - Browser-based example
- `react_example.tsx` - React component example

