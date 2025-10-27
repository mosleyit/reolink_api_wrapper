package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mosleyit/reolink_api_wrapper/sdk/go/reolink"
)

func main() {
	// Get camera details from environment variables
	host := os.Getenv("REOLINK_HOST")
	username := os.Getenv("REOLINK_USERNAME")
	password := os.Getenv("REOLINK_PASSWORD")

	if host == "" {
		host = "192.168.1.100"
	}
	if username == "" {
		username = "admin"
	}
	if password == "" {
		log.Fatal("REOLINK_PASSWORD environment variable is required")
	}

	// Create client
	client := reolink.NewClient(host,
		reolink.WithCredentials(username, password),
		reolink.WithTimeout(30*time.Second),
	)

	// Create context
	ctx := context.Background()

	// Login
	fmt.Println("Logging in...")
	if err := client.Login(ctx); err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	defer func() {
		fmt.Println("Logging out...")
		if err := client.Logout(ctx); err != nil {
			log.Printf("Logout failed: %v", err)
		}
	}()

	fmt.Printf("Logged in successfully. Token: %s\n", client.GetToken())

	// Get device information
	fmt.Println("\nGetting device information...")
	info, err := client.System.GetDeviceInfo(ctx)
	if err != nil {
		log.Fatalf("GetDeviceInfo failed: %v", err)
	}

	fmt.Printf("Model: %s\n", info.Model)
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Serial: %s\n", info.Serial)
	fmt.Printf("Firmware: %s\n", info.FirmVer)
	fmt.Printf("Hardware: %s\n", info.HardVer)
	fmt.Printf("Channels: %d\n", info.ChannelNum)

	// Get device name
	fmt.Println("\nGetting device name...")
	name, err := client.System.GetDeviceName(ctx)
	if err != nil {
		log.Fatalf("GetDeviceName failed: %v", err)
	}
	fmt.Printf("Device Name: %s\n", name)

	// Get time configuration
	fmt.Println("\nGetting time configuration...")
	timeConfig, err := client.System.GetTime(ctx)
	if err != nil {
		log.Fatalf("GetTime failed: %v", err)
	}
	fmt.Printf("Time: %04d-%02d-%02d %02d:%02d:%02d (TZ: %d)\n",
		timeConfig.Year, timeConfig.Mon, timeConfig.Day,
		timeConfig.Hour, timeConfig.Min, timeConfig.Sec,
		timeConfig.TimeZone)

	// Get streaming URLs
	fmt.Println("\nStreaming URLs:")
	fmt.Printf("RTSP Main: %s\n", client.Streaming.GetRTSPURL(reolink.StreamMain, 1))
	fmt.Printf("RTSP Sub:  %s\n", client.Streaming.GetRTSPURL(reolink.StreamSub, 1))
	fmt.Printf("RTMP Main: %s\n", client.Streaming.GetRTMPURL(reolink.StreamMain, 0))
	fmt.Printf("FLV Main:  %s\n", client.Streaming.GetFLVURL(reolink.StreamMain, 0))

	fmt.Println("\nDone!")
}
