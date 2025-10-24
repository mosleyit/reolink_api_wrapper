# Documentation Guide

This guide explains how to use and contribute to the Reolink Camera API documentation.

## 📚 Viewing the Documentation

### Online (Recommended)

Visit **https://mosleyit.github.io/reolink_api_wrapper/** to access:

1. **Swagger UI** - Interactive API explorer
   - Try out API endpoints directly in your browser
   - See real-time request/response examples
   - Generate code snippets

2. **Redoc** - Beautiful, readable documentation
   - Three-panel design for easy navigation
   - Search functionality
   - Mobile-friendly

3. **OpenAPI YAML** - Raw specification
   - Download for offline use
   - Import into tools like Postman
   - Generate clients in any language

### Locally

Run a local server to view the documentation:

```bash
# Simple Python server
cd docs
python -m http.server 8000
# Open http://localhost:8000

# Or use Node.js
npx http-server docs -p 8000
```

## 🔧 Using the OpenAPI Specification

### Generate API Clients

The OpenAPI specification can generate type-safe clients in multiple languages:

#### Python
```bash
openapi-generator-cli generate \
  -i docs/reolink-camera-api-openapi.yaml \
  -g python \
  -o ./python-client \
  --package-name reolink_api
```

#### Go
```bash
openapi-generator-cli generate \
  -i docs/reolink-camera-api-openapi.yaml \
  -g go \
  -o ./go-client \
  --package-name reolink
```

#### TypeScript/JavaScript
```bash
openapi-generator-cli generate \
  -i docs/reolink-camera-api-openapi.yaml \
  -g typescript-axios \
  -o ./ts-client
```

#### Java
```bash
openapi-generator-cli generate \
  -i docs/reolink-camera-api-openapi.yaml \
  -g java \
  -o ./java-client \
  --library okhttp-gson
```

### Import into API Tools

#### Postman
1. Open Postman
2. Click **Import**
3. Select **Link**
4. Paste: `https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml`

#### Insomnia
1. Open Insomnia
2. Click **Create** → **Import From** → **URL**
3. Paste: `https://mosleyit.github.io/reolink_api_wrapper/reolink-camera-api-openapi.yaml`

#### Bruno / Hoppscotch / Paw
Import the YAML file directly from the URL above.

## 📝 Documentation Structure

```
docs/
├── index.html                          # Landing page
├── swagger-ui.html                     # Swagger UI interface
├── redoc.html                          # Redoc interface
├── reolink-camera-api-openapi.yaml    # OpenAPI specification
├── api_guide.txt                       # Text version of PDF
├── reolink-camera-http-api-user-guide-v8.pdf  # Original PDF
└── .nojekyll                           # Disable Jekyll processing
```

## 🚀 Deployment

The documentation is automatically deployed to GitHub Pages via GitHub Actions.

### Automatic Deployment

Every push to the `main` branch triggers the deployment workflow:

1. GitHub Actions workflow runs (`.github/workflows/deploy-docs.yml`)
2. Documentation is built from the `docs/` folder
3. Deployed to https://mosleyit.github.io/reolink_api_wrapper/

### Manual Deployment

You can also trigger deployment manually:

```bash
# Via GitHub CLI
gh workflow run deploy-docs.yml

# Or via GitHub web interface
# Go to Actions → Deploy Documentation → Run workflow
```

## 🎨 Customizing the Documentation

### Updating the Landing Page

Edit `docs/index.html` to customize:
- Color scheme (currently purple gradient)
- Feature highlights
- Statistics
- Quick start examples

### Customizing Swagger UI

Edit `docs/swagger-ui.html` to change:
- Theme colors
- Default expansion settings
- Request snippet languages
- Layout options

### Customizing Redoc

Edit `docs/redoc.html` to modify:
- Color scheme
- Typography
- Sidebar appearance
- Right panel styling

## 📖 OpenAPI Specification Details

The `reolink-camera-api-openapi.yaml` file contains:

- **110+ endpoints** - All camera API operations
- **Complete schemas** - Request/response types
- **Examples** - Working examples for every endpoint
- **Error codes** - All 50+ error codes documented
- **Authentication** - Token and Basic auth methods
- **Streaming** - RTSP, RTMP, FLV protocols

### Key Sections

1. **System Commands** - Login, device info, reboot, upgrade
2. **Network Commands** - DDNS, NTP, WiFi, P2P
3. **Video Commands** - OSD, ISP, encoding, streaming
4. **PTZ Commands** - Pan, tilt, zoom, presets
5. **Alarm Commands** - Motion detection, AI detection
6. **Recording Commands** - Schedule, playback
7. **LED Commands** - IR lights, white LED

## 🔍 Searching the Documentation

### In Swagger UI
- Use the filter box at the top
- Search by endpoint name or tag

### In Redoc
- Use the search box in the left sidebar
- Searches through all endpoints and descriptions

### In the YAML File
```bash
# Search for specific commands
grep -n "cmd.*Login" docs/reolink-camera-api-openapi.yaml

# Search for error codes
grep -n "code.*-6" docs/reolink-camera-api-openapi.yaml
```

## 🤝 Contributing

To improve the documentation:

1. Edit the OpenAPI YAML file
2. Test locally using the methods above
3. Commit and push to trigger automatic deployment
4. Verify the changes at the GitHub Pages URL

### Adding New Endpoints

1. Add the endpoint to `docs/reolink-camera-api-openapi.yaml`
2. Include complete request/response schemas
3. Add working examples
4. Document all parameters
5. Test the documentation locally

### Reporting Issues

If you find errors or missing information:

1. Open an issue on GitHub
2. Include the endpoint name
3. Describe the problem
4. Provide the correct information if known

## 📊 Analytics

The documentation site includes basic analytics to track:
- Page views
- Popular endpoints
- User locations

View analytics in the GitHub repository insights.

## 🔒 Security

The documentation is public and contains no sensitive information:
- No API keys or tokens
- No real camera IPs
- Example credentials only

When using the API:
- Never commit real credentials
- Use environment variables
- Rotate tokens regularly

## 📞 Support

For help with the documentation:

- **GitHub Issues**: https://github.com/mosleyit/reolink_api_wrapper/issues
- **Discussions**: https://github.com/mosleyit/reolink_api_wrapper/discussions
- **Email**: See GitHub profile

## 📜 License

This documentation is based on the official Reolink Camera HTTP API User Guide.
Please refer to Reolink's terms of service for API usage.

