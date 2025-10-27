package reolink

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSystemAPI_GetDeviceInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "GetDevInfo" {
			t.Errorf("expected GetDevInfo command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "GetDevInfo",
			Code: 0,
			Value: json.RawMessage(`{
				"DevInfo": {
					"model": "RLC-410",
					"name": "Test Camera",
					"serial": "12345678",
					"firmVer": "v3.0.0.0",
					"hardVer": "IPC_51516M5M",
					"channelNum": 1,
					"diskNum": 0,
					"wifi": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	info, err := client.System.GetDeviceInfo(ctx)
	if err != nil {
		t.Fatalf("GetDeviceInfo failed: %v", err)
	}

	if info.Model != "RLC-410" {
		t.Errorf("expected model RLC-410, got %s", info.Model)
	}

	if info.Name != "Test Camera" {
		t.Errorf("expected name Test Camera, got %s", info.Name)
	}

	if info.Serial != "12345678" {
		t.Errorf("expected serial 12345678, got %s", info.Serial)
	}
}

func TestSystemAPI_GetDeviceName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "GetDevName",
			Code:  0,
			Value: json.RawMessage(`{"DevName":{"name":"My Camera"}}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	name, err := client.System.GetDeviceName(ctx)
	if err != nil {
		t.Fatalf("GetDeviceName failed: %v", err)
	}

	if name != "My Camera" {
		t.Errorf("expected name 'My Camera', got %s", name)
	}
}

func TestSystemAPI_SetDeviceName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "SetDevName" {
			t.Errorf("expected SetDevName command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "SetDevName",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.SetDeviceName(ctx, "New Camera Name")
	if err != nil {
		t.Fatalf("SetDeviceName failed: %v", err)
	}
}

func TestSystemAPI_GetHddInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetHddInfo",
			Code: 0,
			Value: json.RawMessage(`{
				"HddInfo": [
					{
						"capacity": 1000000,
						"format": 1,
						"mount": 1,
						"size": 500000,
						"status": "ok"
					}
				]
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	hdds, err := client.System.GetHddInfo(ctx)
	if err != nil {
		t.Fatalf("GetHddInfo failed: %v", err)
	}

	if len(hdds) != 1 {
		t.Fatalf("expected 1 HDD, got %d", len(hdds))
	}

	if hdds[0].Capacity != 1000000 {
		t.Errorf("expected capacity 1000000, got %d", hdds[0].Capacity)
	}

	if hdds[0].Status != "ok" {
		t.Errorf("expected status ok, got %s", hdds[0].Status)
	}
}

func TestSystemAPI_Reboot(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "Reboot" {
			t.Errorf("expected Reboot command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "Reboot",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.Reboot(ctx)
	if err != nil {
		t.Fatalf("Reboot failed: %v", err)
	}
}

func TestSystemAPI_GetTime(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetTime",
			Code: 0,
			Value: json.RawMessage(`{
				"Time": {
					"year": 2024,
					"mon": 10,
					"day": 27,
					"hour": 14,
					"min": 30,
					"sec": 45,
					"timeFormat": "24"
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	timeConfig, err := client.System.GetTime(ctx)
	if err != nil {
		t.Fatalf("GetTime failed: %v", err)
	}

	if timeConfig.Year != 2024 {
		t.Errorf("expected year 2024, got %d", timeConfig.Year)
	}
	if timeConfig.TimeFormat != "24" {
		t.Errorf("expected timeFormat '24', got %s", timeConfig.TimeFormat)
	}
}

func TestSystemAPI_SetTime(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "SetTime" {
			t.Errorf("expected SetTime command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "SetTime",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	timeConfig := &TimeConfig{
		Year:       2024,
		Mon:        10,
		Day:        27,
		Hour:       14,
		Min:        30,
		Sec:        45,
		TimeFormat: "24",
	}
	err := client.System.SetTime(ctx, timeConfig)
	if err != nil {
		t.Fatalf("SetTime failed: %v", err)
	}
}

func TestSystemAPI_Format(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "Format" {
			t.Errorf("expected Format command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "Format",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.Format(ctx, 0)
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}
}

func TestSystemAPI_Restore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "Restore" {
			t.Errorf("expected Restore command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "Restore",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.Restore(ctx)
	if err != nil {
		t.Fatalf("Restore failed: %v", err)
	}
}

func TestSystemAPI_GetAbility(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetAbility",
			Code: 0,
			Value: json.RawMessage(`{
				"Ability": {
					"Ability": {
						"abilityChn": [
							{
								"aiTrack": {"permit": 1, "ver": 1},
								"ptzCtrl": {"permit": 1, "ver": 1}
							}
						]
					}
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	ability, err := client.System.GetAbility(ctx)
	if err != nil {
		t.Fatalf("GetAbility failed: %v", err)
	}

	if ability.AbilityInfo == nil {
		t.Error("expected AbilityInfo to be non-nil")
	}
}

func TestSystemAPI_GetAutoUpgrade(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetAutoUpgrade",
			Code: 0,
			Value: json.RawMessage(`{
				"AutoUpgrade": {
					"enable": 1
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	autoUpgrade, err := client.System.GetAutoUpgrade(ctx)
	if err != nil {
		t.Fatalf("GetAutoUpgrade failed: %v", err)
	}

	if autoUpgrade.Enable != 1 {
		t.Errorf("expected Enable 1, got %d", autoUpgrade.Enable)
	}
}

func TestSystemAPI_SetAutoUpgrade(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "SetAutoUpgrade",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.SetAutoUpgrade(ctx, true)
	if err != nil {
		t.Fatalf("SetAutoUpgrade failed: %v", err)
	}
}

func TestSystemAPI_CheckFirmware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "CheckFirmware",
			Code: 0,
			Value: json.RawMessage(`{
				"newFirmware": 1
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	check, err := client.System.CheckFirmware(ctx)
	if err != nil {
		t.Fatalf("CheckFirmware failed: %v", err)
	}

	if check.NewFirmware != 1 {
		t.Errorf("expected NewFirmware 1, got %d", check.NewFirmware)
	}
}

func TestSystemAPI_UpgradeOnline(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "UpgradeOnline",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.UpgradeOnline(ctx)
	if err != nil {
		t.Fatalf("UpgradeOnline failed: %v", err)
	}
}

func TestSystemAPI_UpgradeStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "UpgradeStatus",
			Code: 0,
			Value: json.RawMessage(`{
				"Status": {
					"Persent": 50,
					"code": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	status, err := client.System.UpgradeStatus(ctx)
	if err != nil {
		t.Fatalf("UpgradeStatus failed: %v", err)
	}

	if status.Percent != 50 {
		t.Errorf("expected Percent 50, got %d", status.Percent)
	}
	if status.Code != 0 {
		t.Errorf("expected Code 0, got %d", status.Code)
	}
}

func TestSystemAPI_UpgradePrepare(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "UpgradePrepare",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.UpgradePrepare(ctx, false, "firmware.pak")
	if err != nil {
		t.Fatalf("UpgradePrepare failed: %v", err)
	}
}

func TestSystemAPI_GetSysCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "GetSysCfg" {
			t.Errorf("expected GetSysCfg command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "GetSysCfg",
			Code: 0,
			Value: json.RawMessage(`{
				"SysCfg": {
					"LockTime": 300,
					"allowedTimes": 5,
					"loginLock": 1
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	cfg, err := client.System.GetSysCfg(ctx)
	if err != nil {
		t.Fatalf("GetSysCfg failed: %v", err)
	}

	if cfg.LockTime != 300 {
		t.Errorf("expected LockTime 300, got %d", cfg.LockTime)
	}
	if cfg.AllowedTimes != 5 {
		t.Errorf("expected AllowedTimes 5, got %d", cfg.AllowedTimes)
	}
	if cfg.LoginLock != 1 {
		t.Errorf("expected LoginLock 1, got %d", cfg.LoginLock)
	}
}

func TestSystemAPI_SetSysCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "SetSysCfg" {
			t.Errorf("expected SetSysCfg command, got %v", req)
		}

		resp := []Response{{
			Cmd:   "SetSysCfg",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.SetSysCfg(ctx, SysCfg{
		LockTime:     300,
		AllowedTimes: 5,
		LoginLock:    1,
	})
	if err != nil {
		t.Fatalf("SetSysCfg failed: %v", err)
	}
}

func TestSystemAPI_Upgrade(t *testing.T) {
	client := NewClient("192.168.1.100")

	ctx := context.Background()
	firmware := []byte("fake firmware data")

	// This should return an error indicating it's not implemented
	err := client.System.Upgrade(ctx, firmware)
	if err == nil {
		t.Fatal("Expected Upgrade to return error (not implemented)")
	}

	// Verify the error message indicates to use alternative methods
	expectedMsg := "Upgrade endpoint not yet implemented"
	if !contains(err.Error(), expectedMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedMsg, err.Error())
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestSystemAPI_GetAutoMaint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetAutoMaint",
			Code: 0,
			Value: json.RawMessage(`{
				"AutoMaint": {
					"enable": 1,
					"hour": 2,
					"min": 30
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	autoMaint, err := client.System.GetAutoMaint(ctx)
	if err != nil {
		t.Fatalf("GetAutoMaint failed: %v", err)
	}

	if autoMaint.Enable != 1 {
		t.Errorf("expected Enable 1, got %d", autoMaint.Enable)
	}
	if autoMaint.Hour != 2 {
		t.Errorf("expected Hour 2, got %d", autoMaint.Hour)
	}
	if autoMaint.Min != 30 {
		t.Errorf("expected Min 30, got %d", autoMaint.Min)
	}
}

func TestSystemAPI_SetAutoMaint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "SetAutoMaint" {
			t.Errorf("expected SetAutoMaint command, got %v", req)
		}

		resp := []Response{{
			Cmd:   "SetAutoMaint",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.System.SetAutoMaint(ctx, AutoMaint{
		Enable: 1,
		Hour:   2,
		Min:    30,
	})
	if err != nil {
		t.Fatalf("SetAutoMaint failed: %v", err)
	}
}

func TestSystemAPI_GetChannelStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "Getchannelstatus",
			Code: 0,
			Value: json.RawMessage(`{
				"status": [
					{
						"channel": 0,
						"name": "Camera 1",
						"online": 1
					},
					{
						"channel": 1,
						"name": "Camera 2",
						"online": 0
					}
				]
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	channelStatus, err := client.System.GetChannelStatus(ctx)
	if err != nil {
		t.Fatalf("GetChannelStatus failed: %v", err)
	}

	if len(channelStatus.Status) != 2 {
		t.Fatalf("expected 2 channels, got %d", len(channelStatus.Status))
	}

	if channelStatus.Status[0].Channel != 0 {
		t.Errorf("expected channel 0, got %d", channelStatus.Status[0].Channel)
	}
	if channelStatus.Status[0].Name != "Camera 1" {
		t.Errorf("expected name 'Camera 1', got %s", channelStatus.Status[0].Name)
	}
	if channelStatus.Status[0].Online != 1 {
		t.Errorf("expected online 1, got %d", channelStatus.Status[0].Online)
	}
}
