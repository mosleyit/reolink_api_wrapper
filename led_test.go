package reolink

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLEDAPI_GetIrLights(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetIrLights" {
			t.Errorf("Expected cmd 'GetIrLights', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetIrLights",
			Code: 0,
			Value: json.RawMessage(`{
				"IrLights": {
					"state": "Auto"
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test GetIrLights
	ctx := t.Context()
	irLights, err := client.LED.GetIrLights(ctx)
	if err != nil {
		t.Fatalf("GetIrLights failed: %v", err)
	}

	// Verify response
	if irLights.State != "Auto" {
		t.Errorf("Expected state 'Auto', got '%s'", irLights.State)
	}
}

func TestLEDAPI_SetIrLights(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetIrLights" {
			t.Errorf("Expected cmd 'SetIrLights', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetIrLights",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test SetIrLights
	ctx := t.Context()
	err := client.LED.SetIrLights(ctx, 0, LEDStateOn)
	if err != nil {
		t.Fatalf("SetIrLights failed: %v", err)
	}
}

func TestLEDAPI_GetPowerLed(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetPowerLed" {
			t.Errorf("Expected cmd 'GetPowerLed', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetPowerLed",
			Code: 0,
			Value: json.RawMessage(`{
				"PowerLed": {
					"state": "Auto"
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test GetPowerLed
	ctx := t.Context()
	powerLed, err := client.LED.GetPowerLed(ctx, 0)
	if err != nil {
		t.Fatalf("GetPowerLed failed: %v", err)
	}

	// Verify response
	if powerLed.State != "Auto" {
		t.Errorf("Expected state 'Auto', got '%s'", powerLed.State)
	}
}

func TestLEDAPI_SetPowerLed(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPowerLed" {
			t.Errorf("Expected cmd 'SetPowerLed', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPowerLed",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test SetPowerLed
	ctx := t.Context()
	err := client.LED.SetPowerLed(ctx, 0, LEDStateOff)
	if err != nil {
		t.Fatalf("SetPowerLed failed: %v", err)
	}
}

func TestLEDAPI_GetWhiteLed(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetWhiteLed" {
			t.Errorf("Expected cmd 'GetWhiteLed', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetWhiteLed",
			Code: 0,
			Value: json.RawMessage(`{
				"WhiteLed": {
					"channel": 0,
					"state": 1,
					"mode": 1,
					"bright": 80,
					"LightingSchedule": {
						"StartHour": 18,
						"StartMin": 0,
						"EndHour": 6,
						"EndMin": 0
					},
					"wlAiDetectType": {
						"people": 1,
						"vehicle": 1,
						"dog_cat": 0,
						"face": 0
					}
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test GetWhiteLed
	ctx := t.Context()
	whiteLed, err := client.LED.GetWhiteLed(ctx, 0)
	if err != nil {
		t.Fatalf("GetWhiteLed failed: %v", err)
	}

	// Verify response
	if whiteLed.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", whiteLed.Channel)
	}
	if whiteLed.State != 1 {
		t.Errorf("Expected state 1, got %d", whiteLed.State)
	}
	if whiteLed.Mode != 1 {
		t.Errorf("Expected mode 1, got %d", whiteLed.Mode)
	}
	if whiteLed.Bright != 80 {
		t.Errorf("Expected brightness 80, got %d", whiteLed.Bright)
	}
	if whiteLed.WlAiDetectType.People != 1 {
		t.Errorf("Expected people detection enabled, got %d", whiteLed.WlAiDetectType.People)
	}
}

func TestLEDAPI_SetWhiteLed(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetWhiteLed" {
			t.Errorf("Expected cmd 'SetWhiteLed', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetWhiteLed",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test SetWhiteLed
	ctx := t.Context()
	config := WhiteLed{
		Channel: 0,
		State:   1,
		Mode:    1,
		Bright:  80,
		LightingSchedule: WhiteLedSchedule{
			StartHour: 18,
			StartMin:  0,
			EndHour:   6,
			EndMin:    0,
		},
		WlAiDetectType: WhiteLedAiDetect{
			People:  1,
			Vehicle: 1,
			DogCat:  0,
			Face:    0,
		},
	}

	err := client.LED.SetWhiteLed(ctx, config)
	if err != nil {
		t.Fatalf("SetWhiteLed failed: %v", err)
	}
}

func TestLEDAPI_GetAiAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAiAlarm" {
			t.Errorf("Expected cmd 'GetAiAlarm', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAiAlarm",
			Code: 0,
			Value: json.RawMessage(`{
				"AiAlarm": {
					"channel": 0,
					"ai_type": "people",
					"sensitivity": 10,
					"stay_time": 0,
					"width": 80,
					"height": 60,
					"min_target_height": 0.0,
					"max_target_height": 1.0,
					"min_target_width": 0.0,
					"max_target_width": 1.0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test GetAiAlarm
	ctx := t.Context()
	aiAlarm, err := client.LED.GetAiAlarm(ctx, 0, "people")
	if err != nil {
		t.Fatalf("GetAiAlarm failed: %v", err)
	}

	if aiAlarm.AiType != "people" {
		t.Errorf("Expected ai_type 'people', got '%s'", aiAlarm.AiType)
	}

	if aiAlarm.Sensitivity != 10 {
		t.Errorf("Expected sensitivity 10, got %d", aiAlarm.Sensitivity)
	}
}

func TestLEDAPI_SetAiAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAiAlarm" {
			t.Errorf("Expected cmd 'SetAiAlarm', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "SetAiAlarm",
			Code: 0,
			Value: json.RawMessage(`{
				"rspCode": 200
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test SetAiAlarm
	ctx := t.Context()
	alarm := AiAlarm{
		AiType:          "people",
		Sensitivity:     10,
		StayTime:        0,
		Width:           80,
		Height:          60,
		MinTargetHeight: 0.0,
		MaxTargetHeight: 1.0,
		MinTargetWidth:  0.0,
		MaxTargetWidth:  1.0,
	}

	err := client.LED.SetAiAlarm(ctx, 0, alarm)
	if err != nil {
		t.Fatalf("SetAiAlarm failed: %v", err)
	}
}

func TestLEDAPI_SetAlarmArea(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAlarmArea" {
			t.Errorf("Expected cmd 'SetAlarmArea', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "SetAlarmArea",
			Code: 0,
			Value: json.RawMessage(`{
				"rspCode": 200
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.LED = &LEDAPI{client: client}

	// Test SetAlarmArea
	ctx := t.Context()
	params := map[string]interface{}{
		"channel": 0,
		"area":    "test_area",
	}

	err := client.LED.SetAlarmArea(ctx, params)
	if err != nil {
		t.Fatalf("SetAlarmArea failed: %v", err)
	}
}
