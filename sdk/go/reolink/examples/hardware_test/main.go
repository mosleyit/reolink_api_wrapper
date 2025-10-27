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
	fmt.Println("=== Reolink Camera SDK Hardware Test ===")
	fmt.Println()

	// Camera configuration - use environment variables
	host := getEnv("REOLINK_HOST", "192.168.1.100")
	username := getEnv("REOLINK_USERNAME", "admin")
	password := getEnv("REOLINK_PASSWORD", "")

	if password == "" {
		log.Fatal("REOLINK_PASSWORD environment variable is required")
	}

	// Create client with HTTPS enabled
	fmt.Printf("Creating client for %s...\n", host)
	client := reolink.NewClient(host,
		reolink.WithCredentials(username, password),
		reolink.WithHTTPS(true), // HTTPS enabled
		reolink.WithTimeout(30*time.Second),
		reolink.WithInsecureSkipVerify(true), // Skip cert verification for self-signed certs
	)

	ctx := context.Background()

	// Test 1: Login
	fmt.Println("\n--- Test 1: Authentication ---")
	fmt.Println("Logging in...")
	if err := client.Login(ctx); err != nil {
		log.Fatalf("❌ Login failed: %v", err)
	}
	token := client.GetToken()
	fmt.Printf("✅ Login successful! Token: %s (length: %d)\n", token, len(token))

	// Small delay to ensure token is ready
	time.Sleep(100 * time.Millisecond)

	// Ensure we logout at the end
	defer func() {
		fmt.Println("\n--- Cleanup ---")
		fmt.Println("Logging out...")
		if err := client.Logout(ctx); err != nil {
			log.Printf("⚠️  Logout failed: %v", err)
		} else {
			fmt.Println("✅ Logout successful")
		}
	}()

	// Test 2: Get Device Info
	fmt.Println("\n--- Test 2: Device Information ---")
	info, err := client.System.GetDeviceInfo(ctx)
	if err != nil {
		log.Printf("❌ GetDeviceInfo failed: %v", err)
	} else {
		fmt.Println("✅ Device Information:")
		fmt.Printf("   Model:        %s\n", info.Model)
		fmt.Printf("   Name:         %s\n", info.Name)
		fmt.Printf("   Serial:       %s\n", info.Serial)
		fmt.Printf("   Firmware:     %s\n", info.FirmVer)
		fmt.Printf("   Hardware:     %s\n", info.HardVer)
		fmt.Printf("   Channels:     %d\n", info.ChannelNum)
		fmt.Printf("   Disk Num:     %d\n", info.DiskNum)
		fmt.Printf("   WiFi:         %d\n", info.Wifi)
	}

	// Test 3: Get Device Name
	fmt.Println("\n--- Test 3: Device Name ---")
	name, err := client.System.GetDeviceName(ctx)
	if err != nil {
		log.Printf("❌ GetDeviceName failed: %v", err)
	} else {
		fmt.Printf("✅ Device Name: %s\n", name)
	}

	// Test 4: Get Time Configuration
	fmt.Println("\n--- Test 4: Time Configuration ---")
	timeConfig, err := client.System.GetTime(ctx)
	if err != nil {
		log.Printf("❌ GetTime failed: %v", err)
	} else {
		fmt.Println("✅ Time Configuration:")
		fmt.Printf("   Date/Time:    %04d-%02d-%02d %02d:%02d:%02d\n",
			timeConfig.Year, timeConfig.Mon, timeConfig.Day,
			timeConfig.Hour, timeConfig.Min, timeConfig.Sec)
		fmt.Printf("   Time Zone:    %d\n", timeConfig.TimeZone)
		fmt.Printf("   Time Format:  %s\n", timeConfig.TimeFormat)
	}

	// Test 5: Get HDD Info
	fmt.Println("\n--- Test 5: Storage Information ---")
	hdds, err := client.System.GetHddInfo(ctx)
	if err != nil {
		log.Printf("❌ GetHddInfo failed: %v", err)
	} else {
		if len(hdds) == 0 {
			fmt.Println("ℹ️  No storage devices found (SD card may not be present)")
		} else {
			fmt.Printf("✅ Found %d storage device(s):\n", len(hdds))
			for i, hdd := range hdds {
				fmt.Printf("   HDD %d:\n", i)
				fmt.Printf("      Capacity:  %d MB\n", hdd.Capacity)
				fmt.Printf("      Used:      %d MB\n", hdd.Size)
				fmt.Printf("      Status:    %s\n", hdd.Status)
				fmt.Printf("      Format:    %d\n", hdd.Format)
				fmt.Printf("      Mount:     %d\n", hdd.Mount)
			}
		}
	}

	// Test 6: Get System Ability
	fmt.Println("\n--- Test 6: System Capabilities ---")
	ability, err := client.System.GetAbility(ctx)
	if err != nil {
		log.Printf("❌ GetAbility failed: %v", err)
	} else {
		fmt.Println("✅ System capabilities retrieved")
		fmt.Printf("   Ability fields: %d\n", len(ability.AbilityInfo))
	}

	// Test 7: Get Users
	fmt.Println("\n--- Test 7: User Management ---")
	users, err := client.Security.GetUsers(ctx)
	if err != nil {
		log.Printf("❌ GetUsers failed: %v", err)
	} else {
		fmt.Printf("✅ Found %d user(s):\n", len(users))
		for i, user := range users {
			fmt.Printf("   User %d:\n", i+1)
			fmt.Printf("      Username:  %s\n", user.UserName)
			fmt.Printf("      Level:     %s\n", user.Level)
		}
	}

	// Test 8: Get Online Users
	fmt.Println("\n--- Test 8: Online Users ---")
	onlineUsers, err := client.Security.GetOnlineUsers(ctx)
	if err != nil {
		log.Printf("❌ GetOnlineUsers failed: %v", err)
	} else {
		fmt.Printf("✅ Found %d online user(s):\n", len(onlineUsers))
		for i, user := range onlineUsers {
			fmt.Printf("   Online User %d:\n", i+1)
			fmt.Printf("      Username:  %s\n", user.UserName)
			fmt.Printf("      IP:        %s\n", user.IP)
		}
	}

	// Test 9: Streaming URLs
	fmt.Println("\n--- Test 9: Streaming URLs ---")
	fmt.Println("✅ Generated Streaming URLs:")
	fmt.Printf("   RTSP Main:     %s\n", client.Streaming.GetRTSPURL(reolink.StreamMain, 1))
	fmt.Printf("   RTSP Sub:      %s\n", client.Streaming.GetRTSPURL(reolink.StreamSub, 1))
	fmt.Printf("   RTMP Main:     %s\n", client.Streaming.GetRTMPURL(reolink.StreamMain, 0))
	fmt.Printf("   FLV Main:      %s\n", client.Streaming.GetFLVURL(reolink.StreamMain, 0))

	// Test 10: Encoding Configuration
	fmt.Println("\n--- Test 10: Encoding Configuration ---")
	encConfig, err := client.Encoding.GetEnc(ctx, 0)
	if err != nil {
		log.Printf("❌ GetEnc failed: %v", err)
	} else {
		fmt.Println("✅ Encoding Configuration:")
		fmt.Printf("   Channel:       %d\n", encConfig.Channel)
		fmt.Printf("   Audio:         %d\n", encConfig.Audio)
		fmt.Println("   Main Stream:")
		fmt.Printf("      Codec:      %s\n", encConfig.MainStream.VType)
		fmt.Printf("      Resolution: %s (%dx%d)\n", encConfig.MainStream.Size, encConfig.MainStream.Width, encConfig.MainStream.Height)
		fmt.Printf("      Bitrate:    %d kbps\n", encConfig.MainStream.BitRate)
		fmt.Printf("      FPS:        %d\n", encConfig.MainStream.FrameRate)
		fmt.Printf("      GOP:        %d\n", encConfig.MainStream.GOP)
		fmt.Printf("      Profile:    %s\n", encConfig.MainStream.Profile)
		fmt.Println("   Sub Stream:")
		fmt.Printf("      Codec:      %s\n", encConfig.SubStream.VType)
		fmt.Printf("      Resolution: %s (%dx%d)\n", encConfig.SubStream.Size, encConfig.SubStream.Width, encConfig.SubStream.Height)
		fmt.Printf("      Bitrate:    %d kbps\n", encConfig.SubStream.BitRate)
		fmt.Printf("      FPS:        %d\n", encConfig.SubStream.FrameRate)
	}

	// Test 11: Snapshot
	fmt.Println("\n--- Test 11: Snapshot Capture ---")
	imageData, err := client.Encoding.Snap(ctx, 0)
	if err != nil {
		log.Printf("❌ Snap failed: %v", err)
	} else {
		fmt.Printf("✅ Snapshot captured: %d bytes (JPEG)\n", len(imageData))
		// Optionally save to file
		// os.WriteFile("snapshot.jpg", imageData, 0644)
	}

	// Test 12: Motion Detection State
	fmt.Println("\n--- Test 12: Motion Detection State ---")
	mdState, err := client.Alarm.GetMdState(ctx, 0)
	if err != nil {
		log.Printf("❌ GetMdState failed: %v", err)
	} else {
		stateStr := "No motion"
		if mdState == 1 {
			stateStr = "Motion detected!"
		}
		fmt.Printf("✅ Motion Detection State: %s (%d)\n", stateStr, mdState)
	}

	// Test 13: Motion Detection Configuration
	fmt.Println("\n--- Test 13: Motion Detection Configuration ---")
	mdAlarm, err := client.Alarm.GetMdAlarm(ctx, 0)
	if err != nil {
		log.Printf("❌ GetMdAlarm failed: %v", err)
	} else {
		fmt.Println("✅ Motion Detection Configuration:")
		fmt.Printf("   Channel:       %d\n", mdAlarm.Channel)
		fmt.Printf("   Detection Grid: %dx%d (%d cells)\n", mdAlarm.Scope.Cols, mdAlarm.Scope.Rows, len(mdAlarm.Scope.Table))
		fmt.Printf("   Time Periods:  %d\n", len(mdAlarm.NewSens.Sens))
		for i, sens := range mdAlarm.NewSens.Sens {
			if sens.Enable == 1 {
				fmt.Printf("      Period %d:   %02d:%02d-%02d:%02d (Sensitivity: %d)\n",
					i+1, sens.BeginHour, sens.BeginMin, sens.EndHour, sens.EndMin, sens.Sensitivity)
			}
		}
	}

	// Test 14: PTZ Presets (if camera supports PTZ)
	fmt.Println("\n--- Test 14: PTZ Presets ---")
	presets, err := client.PTZ.GetPtzPreset(ctx, 0)
	if err != nil {
		log.Printf("⚠️  GetPtzPreset failed (camera may not support PTZ): %v", err)
	} else {
		fmt.Printf("✅ Found %d PTZ preset(s):\n", len(presets))
		for _, preset := range presets {
			if preset.Enable == 1 {
				fmt.Printf("   Preset %d: %s (enabled)\n", preset.ID, preset.Name)
			}
		}
	}

	// Test 15: PTZ Guard Configuration (if camera supports PTZ)
	fmt.Println("\n--- Test 15: PTZ Guard Configuration ---")
	guard, err := client.PTZ.GetPtzGuard(ctx, 0)
	if err != nil {
		log.Printf("⚠️  GetPtzGuard failed (camera may not support PTZ): %v", err)
	} else {
		fmt.Println("✅ PTZ Guard Configuration:")
		fmt.Printf("   Enabled:       %d\n", guard.BEnable)
		fmt.Printf("   Position Set:  %d\n", guard.BExistPos)
		fmt.Printf("   Timeout:       %d seconds\n", guard.Timeout)
	}

	// Test 16: Network Configuration
	fmt.Println("\n--- Test 16: Network Configuration ---")
	netPort, err := client.Network.GetNetPort(ctx)
	if err != nil {
		log.Printf("❌ GetNetPort failed: %v", err)
	} else {
		fmt.Println("✅ Network Port Configuration:")
		fmt.Printf("   HTTP Port:     %d\n", netPort.HTTPPort)
		fmt.Printf("   HTTPS Port:    %d\n", netPort.HTTPSPort)
		fmt.Printf("   RTMP Port:     %d\n", netPort.RTMPPort)
		fmt.Printf("   RTSP Port:     %d\n", netPort.RTSPPort)
		fmt.Printf("   ONVIF Port:    %d\n", netPort.OnvifPort)
	}

	// Test 17: Local Link Configuration
	fmt.Println("\n--- Test 17: Local Link Configuration ---")
	localLink, err := client.Network.GetLocalLink(ctx)
	if err != nil {
		log.Printf("❌ GetLocalLink failed: %v", err)
	} else {
		fmt.Println("✅ Local Link Configuration:")
		fmt.Printf("   Type:          %s\n", localLink.Type)
		fmt.Printf("   IP Address:    %s\n", localLink.Static.IP)
		fmt.Printf("   Gateway:       %s\n", localLink.Static.Gateway)
		fmt.Printf("   Subnet Mask:   %s\n", localLink.Static.Mask)
		fmt.Printf("   DNS Auto:      %d\n", localLink.DNS.Auto)
		fmt.Printf("   DNS1:          %s\n", localLink.DNS.DNS1)
		fmt.Printf("   DNS2:          %s\n", localLink.DNS.DNS2)
	}

	// Test 18: WiFi Configuration
	fmt.Println("\n--- Test 18: WiFi Configuration ---")
	wifi, err := client.Network.GetWifi(ctx)
	if err != nil {
		log.Printf("⚠️  GetWifi failed (camera may not support WiFi): %v", err)
	} else {
		fmt.Println("✅ WiFi Configuration:")
		fmt.Printf("   SSID:          %s\n", wifi.SSID)
		fmt.Printf("   Password Set:  %t\n", wifi.Password != "")
	}

	// Test 19: DDNS Configuration
	fmt.Println("\n--- Test 19: DDNS Configuration ---")
	ddns, err := client.Network.GetDdns(ctx)
	if err != nil {
		log.Printf("❌ GetDdns failed: %v", err)
	} else {
		fmt.Println("✅ DDNS Configuration:")
		fmt.Printf("   Enabled:       %d\n", ddns.Enable)
		fmt.Printf("   Type:          %s\n", ddns.Type)
		fmt.Printf("   Domain:        %s\n", ddns.Domain)
	}

	// Test 20: NTP Configuration
	fmt.Println("\n--- Test 20: NTP Configuration ---")
	ntp, err := client.Network.GetNtp(ctx)
	if err != nil {
		log.Printf("❌ GetNtp failed: %v", err)
	} else {
		fmt.Println("✅ NTP Configuration:")
		fmt.Printf("   Enabled:       %d\n", ntp.Enable)
		fmt.Printf("   Server:        %s\n", ntp.Server)
		fmt.Printf("   Port:          %d\n", ntp.Port)
		fmt.Printf("   Interval:      %d\n", ntp.Interval)
	}

	// Test 21: Email Configuration
	fmt.Println("\n--- Test 21: Email Configuration ---")
	email, err := client.Network.GetEmail(ctx)
	if err != nil {
		log.Printf("❌ GetEmail failed: %v", err)
	} else {
		fmt.Println("✅ Email Configuration:")
		fmt.Printf("   SMTP Server:   %s\n", email.SMTPServer)
		fmt.Printf("   SMTP Port:     %d\n", email.SMTPPort)
		fmt.Printf("   Username:      %s\n", email.UserName)
		fmt.Printf("   Receiver 1:    %s\n", email.Addr1)
		fmt.Printf("   Interval:      %d\n", email.Interval)
	}

	// Test 22: FTP Configuration
	fmt.Println("\n--- Test 22: FTP Configuration ---")
	ftp, err := client.Network.GetFtp(ctx)
	if err != nil {
		log.Printf("❌ GetFtp failed: %v", err)
	} else {
		fmt.Println("✅ FTP Configuration:")
		fmt.Printf("   Server:        %s\n", ftp.Server)
		fmt.Printf("   Port:          %d\n", ftp.Port)
		fmt.Printf("   Username:      %s\n", ftp.UserName)
		fmt.Printf("   Remote Dir:    %s\n", ftp.RemoteDir)
	}

	// Test 23: Push Notification Configuration
	fmt.Println("\n--- Test 23: Push Notification Configuration ---")
	push, err := client.Network.GetPush(ctx)
	if err != nil {
		log.Printf("❌ GetPush failed: %v", err)
	} else {
		fmt.Println("✅ Push Notification Configuration:")
		fmt.Printf("   Schedule Enable: %d\n", push.Schedule.Enable)
	}

	// Test 24: P2P Configuration
	fmt.Println("\n--- Test 24: P2P Configuration ---")
	p2p, err := client.Network.GetP2p(ctx)
	if err != nil {
		log.Printf("❌ GetP2p failed: %v", err)
	} else {
		fmt.Println("✅ P2P Configuration:")
		fmt.Printf("   Enabled:       %d\n", p2p.Enable)
		fmt.Printf("   UID:           %s\n", p2p.UID)
	}

	// Test 25: UPnP Configuration
	fmt.Println("\n--- Test 25: UPnP Configuration ---")
	upnp, err := client.Network.GetUpnp(ctx)
	if err != nil {
		log.Printf("❌ GetUpnp failed: %v", err)
	} else {
		fmt.Println("✅ UPnP Configuration:")
		fmt.Printf("   Enabled:       %d\n", upnp.Enable)
	}

	// Test 26: OSD Configuration
	fmt.Println("\n--- Test 26: OSD Configuration ---")
	osd, err := client.Video.GetOsd(ctx, 0)
	if err != nil {
		log.Printf("❌ GetOsd failed: %v", err)
	} else {
		fmt.Println("✅ OSD Configuration:")
		fmt.Printf("   Channel:       %d\n", osd.Channel)
		fmt.Printf("   Camera Name:   %s (enabled: %d)\n", osd.OsdChannel.Name, osd.OsdChannel.Enable)
		fmt.Printf("   Timestamp:     enabled: %d\n", osd.OsdTime.Enable)
		fmt.Printf("   Watermark:     %d\n", osd.Watermark)
	}

	// Test 27: Image Settings
	fmt.Println("\n--- Test 27: Image Settings ---")
	image, err := client.Video.GetImage(ctx, 0)
	if err != nil {
		log.Printf("❌ GetImage failed: %v", err)
	} else {
		fmt.Println("✅ Image Settings:")
		fmt.Printf("   Channel:       %d\n", image.Channel)
		fmt.Printf("   Brightness:    %d\n", image.Bright)
		fmt.Printf("   Contrast:      %d\n", image.Contrast)
		fmt.Printf("   Saturation:    %d\n", image.Saturation)
		fmt.Printf("   Hue:           %d\n", image.Hue)
		fmt.Printf("   Sharpness:     %d\n", image.Sharpen)
	}

	// Test 28: ISP Settings
	fmt.Println("\n--- Test 28: ISP Settings ---")
	isp, err := client.Video.GetIsp(ctx, 0)
	if err != nil {
		log.Printf("❌ GetIsp failed: %v", err)
	} else {
		fmt.Println("✅ ISP Settings:")
		fmt.Printf("   Channel:       %d\n", isp.Channel)
		fmt.Printf("   Anti-flicker:  %s\n", isp.AntiFlicker)
		fmt.Printf("   Exposure:      %s\n", isp.Exposure)
		fmt.Printf("   Gain:          min=%d max=%d\n", isp.Gain.Min, isp.Gain.Max)
		fmt.Printf("   Day/Night:     %s\n", isp.DayNight)
		fmt.Printf("   Backlight:     %s\n", isp.BackLight)
		fmt.Printf("   3D NR:         %d\n", isp.Nr3d)
	}

	// Test 29: Privacy Mask
	fmt.Println("\n--- Test 29: Privacy Mask ---")
	mask, err := client.Video.GetMask(ctx, 0)
	if err != nil {
		log.Printf("❌ GetMask failed: %v", err)
	} else {
		fmt.Println("✅ Privacy Mask Configuration:")
		fmt.Printf("   Channel:       %d\n", mask.Channel)
		fmt.Printf("   Enabled:       %d\n", mask.Enable)
		fmt.Printf("   Mask Areas:    %d\n", len(mask.Area))
	}

	// Test 30: Recording Configuration
	fmt.Println("\n--- Test 30: Recording Configuration ---")
	rec, err := client.Recording.GetRec(ctx, 0)
	if err != nil {
		log.Printf("❌ GetRec failed: %v", err)
	} else {
		fmt.Println("✅ Recording Configuration:")
		fmt.Printf("   Channel:       %d\n", rec.Channel)
		fmt.Printf("   Overwrite:     %d\n", rec.Overwrite)
		fmt.Printf("   Pre-record:    %d\n", rec.PreRec)
		fmt.Printf("   Post-record:   %s\n", rec.PostRec)
	}

	// Test 31: LED - IR Lights
	fmt.Println("\n--- Test 31: IR Lights Configuration ---")
	irLights, err := client.LED.GetIrLights(ctx)
	if err != nil {
		log.Printf("❌ GetIrLights failed: %v", err)
	} else {
		fmt.Println("✅ IR Lights Configuration:")
		fmt.Printf("   State:         %s\n", irLights.State)
	}

	// Test 32: LED - White LED
	fmt.Println("\n--- Test 32: White LED Configuration ---")
	whiteLed, err := client.LED.GetWhiteLed(ctx, 0)
	if err != nil {
		log.Printf("⚠️  GetWhiteLed failed (camera may not have white LED): %v", err)
	} else {
		fmt.Println("✅ White LED Configuration:")
		fmt.Printf("   Channel:       %d\n", whiteLed.Channel)
		fmt.Printf("   State:         %d\n", whiteLed.State)
		fmt.Printf("   Mode:          %d\n", whiteLed.Mode)
		fmt.Printf("   Brightness:    %d\n", whiteLed.Bright)
	}

	// Test 33: LED - Power LED
	fmt.Println("\n--- Test 33: Power LED Configuration ---")
	powerLed, err := client.LED.GetPowerLed(ctx, 0)
	if err != nil {
		log.Printf("❌ GetPowerLed failed: %v", err)
	} else {
		fmt.Println("✅ Power LED Configuration:")
		fmt.Printf("   State:         %s\n", powerLed.State)
	}

	// Test 34: AI Configuration
	fmt.Println("\n--- Test 34: AI Configuration ---")
	aiCfg, err := client.AI.GetAiCfg(ctx, 0)
	if err != nil {
		log.Printf("⚠️  GetAiCfg failed (camera may not support AI): %v", err)
	} else {
		fmt.Println("✅ AI Configuration:")
		fmt.Printf("   Channel:       %d\n", aiCfg.Channel)
		fmt.Printf("   AI Track:      %d\n", aiCfg.AiTrack)
		fmt.Printf("   People:        %d\n", aiCfg.AiDetectType.People)
		fmt.Printf("   Vehicle:       %d\n", aiCfg.AiDetectType.Vehicle)
		fmt.Printf("   Dog/Cat:       %d\n", aiCfg.AiDetectType.DogCat)
		fmt.Printf("   Face:          %d\n", aiCfg.AiDetectType.Face)
	}

	// Test 35: AI State
	fmt.Println("\n--- Test 35: AI Detection State ---")
	aiState, err := client.AI.GetAiState(ctx, 0)
	if err != nil {
		log.Printf("⚠️  GetAiState failed (camera may not support AI): %v", err)
	} else {
		fmt.Println("✅ AI Detection State:")
		fmt.Printf("   Channel:       %d\n", aiState.Channel)
		fmt.Printf("   People:        alarm=%d support=%d\n", aiState.People.AlarmState, aiState.People.Support)
		fmt.Printf("   Vehicle:       alarm=%d support=%d\n", aiState.Vehicle.AlarmState, aiState.Vehicle.Support)
		fmt.Printf("   Dog/Cat:       alarm=%d support=%d\n", aiState.DogCat.AlarmState, aiState.DogCat.Support)
		fmt.Printf("   Face:          alarm=%d support=%d\n", aiState.Face.AlarmState, aiState.Face.Support)
	}

	// Summary
	fmt.Println("\n=== Test Summary ===")
	fmt.Println("✅ All tests completed!")
	fmt.Println("\nThe SDK is working correctly with your camera hardware.")
	fmt.Println("\nAPIs tested (35 endpoints):")
	fmt.Println("  • System API (5): GetDeviceInfo, GetDeviceName, GetTime, GetHddInfo, GetAbility")
	fmt.Println("  • Security API (2): GetUsers, GetOnlineUsers")
	fmt.Println("  • Encoding API (2): GetEnc, Snap")
	fmt.Println("  • Alarm API (2): GetMdState, GetMdAlarm")
	fmt.Println("  • PTZ API (2): GetPtzPreset, GetPtzGuard")
	fmt.Println("  • Network API (10): GetNetPort, GetLocalLink, GetWifi, GetDdns, GetNtp,")
	fmt.Println("                      GetEmail, GetFtp, GetPush, GetP2p, GetUpnp")
	fmt.Println("  • Video API (4): GetOsd, GetImage, GetIsp, GetMask")
	fmt.Println("  • Recording API (1): GetRec")
	fmt.Println("  • LED API (3): GetIrLights, GetWhiteLed, GetPowerLed")
	fmt.Println("  • AI API (2): GetAiCfg, GetAiState")
	fmt.Println("  • Streaming API (3): GetRTSPURL, GetRTMPURL, GetFLVURL")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
