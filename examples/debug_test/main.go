package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	reolink "github.com/mosleyit/reolink_api_wrapper"
)

func main() {
	fmt.Println("=== Debug Login Test ===")

	// Camera configuration - use environment variables
	host := getEnv("REOLINK_HOST", "192.168.1.100")
	username := getEnv("REOLINK_USERNAME", "admin")
	password := getEnv("REOLINK_PASSWORD", "")

	if password == "" {
		log.Fatal("REOLINK_PASSWORD environment variable is required")
	}

	// Test 1: Direct HTTP request to see what the camera expects
	fmt.Println("--- Test 1: Direct HTTP Login Request ---")

	loginReq := []map[string]interface{}{
		{
			"cmd": "Login",
			"param": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": username,
					"password": password,
					"Version":  "0",
				},
			},
		},
	}

	reqBody, _ := json.MarshalIndent(loginReq, "", "  ")
	fmt.Printf("Request body:\n%s\n\n", string(reqBody))

	// Create HTTP client with TLS skip verify
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := fmt.Sprintf("https://%s/cgi-bin/api.cgi", host)
	fmt.Printf("URL: %s\n\n", url)

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response status: %d\n", resp.StatusCode)
	fmt.Printf("Response body:\n%s\n\n", string(respBody))

	// Test 2: Using SDK
	fmt.Println("--- Test 2: SDK Login ---")
	client := reolink.NewClient(host,
		reolink.WithCredentials(username, password),
		reolink.WithHTTPS(true),
		reolink.WithTimeout(30*time.Second),
		reolink.WithInsecureSkipVerify(true),
	)

	ctx := context.Background()
	if err := client.Login(ctx); err != nil {
		log.Printf("SDK Login failed: %v", err)
	} else {
		fmt.Printf("SDK Login successful! Token: %s\n", client.GetToken())
		client.Logout(ctx)
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
