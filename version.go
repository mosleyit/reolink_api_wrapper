package reolink

// Version is the current SDK version following semantic versioning
const Version = "1.0.0"

// UserAgent returns the user agent string for HTTP requests
func UserAgent() string {
	return "reolink-go-sdk/" + Version
}
