package reolink

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Option is a functional option for configuring the Client
type Option func(*Client)

// WithCredentials sets the username and password for authentication
func WithCredentials(username, password string) Option {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

// WithHTTPS enables or disables HTTPS
func WithHTTPS(useHTTPS bool) Option {
	return func(c *Client) {
		c.useHTTPS = useHTTPS
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithInsecureSkipVerify sets whether to skip TLS certificate verification
func WithInsecureSkipVerify(skip bool) Option {
	return func(c *Client) {
		if transport, ok := c.httpClient.Transport.(*http.Transport); ok {
			if transport.TLSClientConfig == nil {
				transport.TLSClientConfig = &tls.Config{}
			}
			transport.TLSClientConfig.InsecureSkipVerify = skip
		}
	}
}

// WithTLSConfig sets a custom TLS configuration
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *Client) {
		if transport, ok := c.httpClient.Transport.(*http.Transport); ok {
			transport.TLSClientConfig = tlsConfig
		}
	}
}

// WithToken sets an existing authentication token
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// WithLogger sets a custom logger for the client
func WithLogger(logger Logger) Option {
	return func(c *Client) {
		if logger != nil {
			c.logger = logger
		}
	}
}
