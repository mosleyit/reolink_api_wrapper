package reolink

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPTZAPI_PtzCtrl(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "PtzCtrl" {
			t.Errorf("Expected cmd 'PtzCtrl', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "PtzCtrl",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test PtzCtrl
	ctx := t.Context()
	err := client.PTZ.PtzCtrl(ctx, PtzCtrlParam{
		Channel: 0,
		Op:      PTZOpLeft,
		Speed:   32,
	})
	if err != nil {
		t.Fatalf("PtzCtrl failed: %v", err)
	}
}

func TestPTZAPI_GetPtzPreset(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetPtzPreset" {
			t.Errorf("Expected cmd 'GetPtzPreset', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetPtzPreset",
			Code: 0,
			Value: json.RawMessage(`{
				"PtzPreset": [
					{
						"enable": 1,
						"id": 1,
						"name": "Front Door"
					},
					{
						"enable": 1,
						"id": 2,
						"name": "Backyard"
					},
					{
						"enable": 0,
						"id": 3,
						"name": ""
					}
				]
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzPreset
	ctx := t.Context()
	presets, err := client.PTZ.GetPtzPreset(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzPreset failed: %v", err)
	}

	// Verify response
	if len(presets) != 3 {
		t.Errorf("Expected 3 presets, got %d", len(presets))
	}
	if presets[0].ID != 1 {
		t.Errorf("Expected preset ID 1, got %d", presets[0].ID)
	}
	if presets[0].Name != "Front Door" {
		t.Errorf("Expected preset name 'Front Door', got '%s'", presets[0].Name)
	}
	if presets[0].Enable != 1 {
		t.Errorf("Expected preset enabled, got %d", presets[0].Enable)
	}
}

func TestPTZAPI_SetPtzPreset(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPtzPreset" {
			t.Errorf("Expected cmd 'SetPtzPreset', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPtzPreset",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetPtzPreset
	ctx := t.Context()
	err := client.PTZ.SetPtzPreset(ctx, PtzPreset{
		Enable: 1,
		ID:     1,
		Name:   "Test Position",
	})
	if err != nil {
		t.Fatalf("SetPtzPreset failed: %v", err)
	}
}

func TestPTZAPI_GetPtzPatrol(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send mock response
		resp := []Response{{
			Cmd:  "GetPtzPatrol",
			Code: 0,
			Value: json.RawMessage(`{
				"PtzPatrol": {
					"channel": 0,
					"enable": 1,
					"id": 1,
					"running": 0,
					"name": "Test Patrol",
					"preset": [
						{
							"id": 1,
							"dwellTime": 5,
							"speed": 32
						},
						{
							"id": 2,
							"dwellTime": 10,
							"speed": 16
						}
					]
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzPatrol
	ctx := t.Context()
	patrol, err := client.PTZ.GetPtzPatrol(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzPatrol failed: %v", err)
	}

	// Verify response
	if patrol.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", patrol.Channel)
	}
	if patrol.Enable != 1 {
		t.Errorf("Expected patrol enabled, got %d", patrol.Enable)
	}
	if patrol.Name != "Test Patrol" {
		t.Errorf("Expected patrol name 'Test Patrol', got '%s'", patrol.Name)
	}
	if len(patrol.Preset) != 2 {
		t.Errorf("Expected 2 presets in patrol, got %d", len(patrol.Preset))
	}
}

func TestPTZAPI_GetPtzGuard(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send mock response
		resp := []Response{{
			Cmd:  "GetPtzGuard",
			Code: 0,
			Value: json.RawMessage(`{
				"PtzGuard": {
					"channel": 0,
					"cmdStr": "",
					"benable": 1,
					"bexistPos": 1,
					"timeout": 60,
					"bSaveCurrentPos": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzGuard
	ctx := t.Context()
	guard, err := client.PTZ.GetPtzGuard(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzGuard failed: %v", err)
	}

	// Verify response
	if guard.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", guard.Channel)
	}
	if guard.BEnable != 1 {
		t.Errorf("Expected guard enabled, got %d", guard.BEnable)
	}
	if guard.Timeout != 60 {
		t.Errorf("Expected timeout 60, got %d", guard.Timeout)
	}
}

func TestPTZAPI_GetPtzCheckState(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetPtzCheckState" {
			t.Errorf("Expected cmd 'GetPtzCheckState', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetPtzCheckState",
			Code: 0,
			Value: json.RawMessage(`{
				"status": 0
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzCheckState
	ctx := t.Context()
	state, err := client.PTZ.GetPtzCheckState(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzCheckState failed: %v", err)
	}

	if state.Status != 0 {
		t.Errorf("Expected Status 0, got %d", state.Status)
	}
}

func TestPTZAPI_PtzCheck(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "PtzCheck" {
			t.Errorf("Expected cmd 'PtzCheck', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "PtzCheck",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test PtzCheck
	ctx := t.Context()
	err := client.PTZ.PtzCheck(ctx, 0)
	if err != nil {
		t.Fatalf("PtzCheck failed: %v", err)
	}
}

func TestPTZAPI_GetZoomFocus(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetZoomFocus" {
			t.Errorf("Expected cmd 'GetZoomFocus', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:  "GetZoomFocus",
			Code: 0,
			Value: json.RawMessage(`{
				"ZoomFocus": {
					"channel": 0,
					"zoom": {"pos": 50},
					"focus": {"pos": 30}
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetZoomFocus
	ctx := t.Context()
	zf, err := client.PTZ.GetZoomFocus(ctx, 0)
	if err != nil {
		t.Fatalf("GetZoomFocus failed: %v", err)
	}

	if zf.Zoom.Pos != 50 {
		t.Errorf("Expected Zoom.Pos 50, got %d", zf.Zoom.Pos)
	}
	if zf.Focus.Pos != 30 {
		t.Errorf("Expected Focus.Pos 30, got %d", zf.Focus.Pos)
	}
}

func TestPTZAPI_StartZoomFocus(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "StartZoomFocus" {
			t.Errorf("Expected cmd 'StartZoomFocus', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "StartZoomFocus",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test StartZoomFocus
	ctx := t.Context()
	err := client.PTZ.StartZoomFocus(ctx, 0, PTZOpZoomInc, 0)
	if err != nil {
		t.Fatalf("StartZoomFocus failed: %v", err)
	}
}

func TestPTZAPI_GetPtzTattern(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetPtzTattern" {
			t.Errorf("Expected cmd 'GetPtzTattern', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:  "GetPtzTattern",
			Code: 0,
			Value: json.RawMessage(`{
				"PtzTattern": {
					"enable": 1,
					"id": 1
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzTattern
	ctx := t.Context()
	tattern, err := client.PTZ.GetPtzTattern(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzTattern failed: %v", err)
	}

	if tattern.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", tattern.Enable)
	}
	if tattern.ID != 1 {
		t.Errorf("Expected ID 1, got %d", tattern.ID)
	}
}

func TestPTZAPI_SetPtzTattern(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPtzTattern" {
			t.Errorf("Expected cmd 'SetPtzTattern', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPtzTattern",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetPtzTattern
	ctx := t.Context()
	err := client.PTZ.SetPtzTattern(ctx, 0, PtzTattern{
		Enable: 1,
		ID:     1,
	})
	if err != nil {
		t.Fatalf("SetPtzTattern failed: %v", err)
	}
}

func TestPTZAPI_GetPtzSerial(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetPtzSerial" {
			t.Errorf("Expected cmd 'GetPtzSerial', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:  "GetPtzSerial",
			Code: 0,
			Value: json.RawMessage(`{
				"PtzSerial": {
					"channel": 0,
					"baudRate": 9600,
					"ctrlAddr": 1,
					"ctrlProtocol": "PELCO_D",
					"dataBit": "CS8",
					"flowCtrl": "none",
					"parity": "none",
					"stopBit": 1
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetPtzSerial
	ctx := t.Context()
	serial, err := client.PTZ.GetPtzSerial(ctx, 0)
	if err != nil {
		t.Fatalf("GetPtzSerial failed: %v", err)
	}

	if serial.BaudRate != 9600 {
		t.Errorf("Expected BaudRate 9600, got %d", serial.BaudRate)
	}
	if serial.CtrlProtocol != "PELCO_D" {
		t.Errorf("Expected CtrlProtocol PELCO_D, got %s", serial.CtrlProtocol)
	}
}

func TestPTZAPI_SetPtzSerial(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPtzSerial" {
			t.Errorf("Expected cmd 'SetPtzSerial', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPtzSerial",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetPtzSerial
	ctx := t.Context()
	err := client.PTZ.SetPtzSerial(ctx, PtzSerial{
		Channel:      0,
		BaudRate:     9600,
		CtrlAddr:     1,
		CtrlProtocol: "PELCO_D",
		DataBit:      "CS8",
		FlowCtrl:     "none",
		Parity:       "none",
		StopBit:      1,
	})
	if err != nil {
		t.Fatalf("SetPtzSerial failed: %v", err)
	}
}

func TestPTZAPI_GetAutoFocus(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAutoFocus" {
			t.Errorf("Expected cmd 'GetAutoFocus', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:  "GetAutoFocus",
			Code: 0,
			Value: json.RawMessage(`{
				"AutoFocus": {
					"channel": 0,
					"disable": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test GetAutoFocus
	ctx := t.Context()
	af, err := client.PTZ.GetAutoFocus(ctx, 0)
	if err != nil {
		t.Fatalf("GetAutoFocus failed: %v", err)
	}

	if af.Disable != 0 {
		t.Errorf("Expected Disable 0, got %d", af.Disable)
	}
}

func TestPTZAPI_SetAutoFocus(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAutoFocus" {
			t.Errorf("Expected cmd 'SetAutoFocus', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetAutoFocus",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetAutoFocus
	ctx := t.Context()
	err := client.PTZ.SetAutoFocus(ctx, AutoFocus{
		Channel: 0,
		Disable: 0,
	})
	if err != nil {
		t.Fatalf("SetAutoFocus failed: %v", err)
	}
}

func TestPTZAPI_SetPtzPatrol(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPtzPatrol" {
			t.Errorf("Expected cmd 'SetPtzPatrol', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPtzPatrol",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetPtzPatrol
	ctx := t.Context()
	err := client.PTZ.SetPtzPatrol(ctx, PtzPatrol{
		Channel: 0,
	})
	if err != nil {
		t.Fatalf("SetPtzPatrol failed: %v", err)
	}
}

func TestPTZAPI_SetPtzGuard(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetPtzGuard" {
			t.Errorf("Expected cmd 'SetPtzGuard', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetPtzGuard",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.PTZ = &PTZAPI{client: client}

	// Test SetPtzGuard
	ctx := t.Context()
	err := client.PTZ.SetPtzGuard(ctx, PtzGuard{
		Channel: 0,
		BEnable: 1,
		Timeout: 60,
	})
	if err != nil {
		t.Fatalf("SetPtzGuard failed: %v", err)
	}
}
