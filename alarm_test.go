package reolink

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAlarmAPI_GetMdState(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetMdState" {
			t.Errorf("Expected cmd 'GetMdState', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetMdState",
			Code: 0,
			Value: json.RawMessage(`{
				"state": 1
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test GetMdState
	ctx := context.Background()
	state, err := client.Alarm.GetMdState(ctx, 0)
	if err != nil {
		t.Fatalf("GetMdState failed: %v", err)
	}

	// Verify response
	if state != 1 {
		t.Errorf("Expected state 1 (motion detected), got %d", state)
	}
}

func TestAlarmAPI_GetMdAlarm(t *testing.T) {
	// Create a detection grid (60x33 = 1980 cells)
	gridSize := 60 * 33
	detectionGrid := strings.Repeat("1", gridSize)

	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetMdAlarm" {
			t.Errorf("Expected cmd 'GetMdAlarm', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetMdAlarm",
			Code: 0,
			Value: json.RawMessage(`{
				"MdAlarm": {
					"channel": 0,
					"scope": {
						"cols": 60,
						"rows": 33,
						"table": "` + detectionGrid + `"
					},
					"newSens": {
						"sens": [
							{
								"id": 0,
								"beginHour": 0,
								"beginMin": 0,
								"endHour": 6,
								"endMin": 0,
								"enable": 1,
								"priority": 0,
								"sensitivity": 50
							},
							{
								"id": 1,
								"beginHour": 6,
								"beginMin": 0,
								"endHour": 18,
								"endMin": 0,
								"enable": 1,
								"priority": 0,
								"sensitivity": 80
							},
							{
								"id": 2,
								"beginHour": 18,
								"beginMin": 0,
								"endHour": 24,
								"endMin": 0,
								"enable": 1,
								"priority": 0,
								"sensitivity": 50
							},
							{
								"id": 3,
								"beginHour": 0,
								"beginMin": 0,
								"endHour": 0,
								"endMin": 0,
								"enable": 0,
								"priority": 0,
								"sensitivity": 0
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

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test GetMdAlarm
	ctx := context.Background()
	mdAlarm, err := client.Alarm.GetMdAlarm(ctx, 0)
	if err != nil {
		t.Fatalf("GetMdAlarm failed: %v", err)
	}

	// Verify response
	if mdAlarm.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", mdAlarm.Channel)
	}
	if mdAlarm.Scope.Cols != 60 {
		t.Errorf("Expected 60 columns, got %d", mdAlarm.Scope.Cols)
	}
	if mdAlarm.Scope.Rows != 33 {
		t.Errorf("Expected 33 rows, got %d", mdAlarm.Scope.Rows)
	}
	if len(mdAlarm.Scope.Table) != gridSize {
		t.Errorf("Expected grid size %d, got %d", gridSize, len(mdAlarm.Scope.Table))
	}
	if len(mdAlarm.NewSens.Sens) != 4 {
		t.Errorf("Expected 4 sensitivity periods, got %d", len(mdAlarm.NewSens.Sens))
	}

	// Verify first sensitivity period
	sens := mdAlarm.NewSens.Sens[0]
	if sens.ID != 0 {
		t.Errorf("Expected sensitivity ID 0, got %d", sens.ID)
	}
	if sens.Enable != 1 {
		t.Errorf("Expected sensitivity enabled, got %d", sens.Enable)
	}
	if sens.Sensitivity != 50 {
		t.Errorf("Expected sensitivity 50, got %d", sens.Sensitivity)
	}
	if sens.BeginHour != 0 || sens.EndHour != 6 {
		t.Errorf("Expected time period 0:00-6:00, got %d:00-%d:00", sens.BeginHour, sens.EndHour)
	}
}

func TestAlarmAPI_SetMdAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetMdAlarm" {
			t.Errorf("Expected cmd 'SetMdAlarm', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetMdAlarm",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test SetMdAlarm
	ctx := context.Background()
	gridSize := 60 * 33
	config := MdAlarm{
		Channel: 0,
		Scope: MdScope{
			Cols:  60,
			Rows:  33,
			Table: strings.Repeat("1", gridSize),
		},
		NewSens: MdNewSens{
			Sens: []MdSensitivity{
				{
					ID:          0,
					BeginHour:   0,
					BeginMin:    0,
					EndHour:     24,
					EndMin:      0,
					Enable:      1,
					Priority:    0,
					Sensitivity: 80,
				},
			},
		},
	}

	err := client.Alarm.SetMdAlarm(ctx, config)
	if err != nil {
		t.Fatalf("SetMdAlarm failed: %v", err)
	}
}

func TestAlarmAPI_AudioAlarmPlay(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "AudioAlarmPlay" {
			t.Errorf("Expected cmd 'AudioAlarmPlay', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "AudioAlarmPlay",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test AudioAlarmPlay
	ctx := context.Background()
	err := client.Alarm.AudioAlarmPlay(ctx, AudioAlarmPlayParam{
		Channel:      0,
		AlarmMode:    "manul",
		ManualSwitch: 1,
		Times:        3,
	})
	if err != nil {
		t.Fatalf("AudioAlarmPlay failed: %v", err)
	}
}

func TestAlarmAPI_GetAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAlarm" {
			t.Errorf("Expected cmd 'GetAlarm', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAlarm",
			Code: 0,
			Value: json.RawMessage(`{
				"Alarm": {
					"channel": 0,
					"type": "md",
					"enable": 1,
					"scope": {
						"cols": 60,
						"rows": 33,
						"table": "` + strings.Repeat("1", 60*33) + `"
					},
					"sens": [
						{
							"beginHour": 0,
							"beginMin": 0,
							"endHour": 23,
							"endMin": 59,
							"sensitivity": 50
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
	client.Alarm = &AlarmAPI{client: client}

	// Test GetAlarm
	ctx := context.Background()
	alarm, err := client.Alarm.GetAlarm(ctx, 0, "md")
	if err != nil {
		t.Fatalf("GetAlarm failed: %v", err)
	}

	if alarm.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", alarm.Channel)
	}
	if alarm.Type != "md" {
		t.Errorf("Expected Type 'md', got '%s'", alarm.Type)
	}
	if alarm.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", alarm.Enable)
	}
}

func TestAlarmAPI_SetAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAlarm" {
			t.Errorf("Expected cmd 'SetAlarm', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetAlarm",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test SetAlarm
	ctx := context.Background()
	err := client.Alarm.SetAlarm(ctx, Alarm{
		Channel: 0,
		Type:    "md",
		Enable:  1,
	})
	if err != nil {
		t.Fatalf("SetAlarm failed: %v", err)
	}
}

func TestAlarmAPI_GetAudioAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAudioAlarm" {
			t.Errorf("Expected cmd 'GetAudioAlarm', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAudioAlarm",
			Code: 0,
			Value: json.RawMessage(`{
				"AudioAlarm": {
					"channel": 0,
					"enable": 1,
					"sensitivity": 50,
					"schedule": {
						"enable": 1,
						"table": "` + strings.Repeat("1", 168) + `"
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
	client.Alarm = &AlarmAPI{client: client}

	// Test GetAudioAlarm
	ctx := context.Background()
	audioAlarm, err := client.Alarm.GetAudioAlarm(ctx, 0)
	if err != nil {
		t.Fatalf("GetAudioAlarm failed: %v", err)
	}

	if audioAlarm.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", audioAlarm.Channel)
	}
	if audioAlarm.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", audioAlarm.Enable)
	}
	if audioAlarm.Sensitivity != 50 {
		t.Errorf("Expected Sensitivity 50, got %d", audioAlarm.Sensitivity)
	}
}

func TestAlarmAPI_SetAudioAlarm(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAudioAlarm" {
			t.Errorf("Expected cmd 'SetAudioAlarm', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetAudioAlarm",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test SetAudioAlarm
	ctx := context.Background()
	err := client.Alarm.SetAudioAlarm(ctx, AudioAlarm{
		Channel:     0,
		Enable:      1,
		Sensitivity: 50,
	})
	if err != nil {
		t.Fatalf("SetAudioAlarm failed: %v", err)
	}
}

func TestAlarmAPI_GetBuzzerAlarmV20(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetBuzzerAlarmV20" {
			t.Errorf("Expected cmd 'GetBuzzerAlarmV20', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetBuzzerAlarmV20",
			Code: 0,
			Value: json.RawMessage(`{
				"BuzzerAlarm": {
					"channel": 0,
					"enable": 1,
					"schedule": {
						"enable": 1,
						"table": {
							"MD": "` + strings.Repeat("1", 168) + `"
						}
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
	client.Alarm = &AlarmAPI{client: client}

	// Test GetBuzzerAlarmV20
	ctx := context.Background()
	buzzerAlarm, err := client.Alarm.GetBuzzerAlarmV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetBuzzerAlarmV20 failed: %v", err)
	}

	if buzzerAlarm.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", buzzerAlarm.Channel)
	}
	if buzzerAlarm.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", buzzerAlarm.Enable)
	}
}

func TestAlarmAPI_SetBuzzerAlarmV20(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetBuzzerAlarmV20" {
			t.Errorf("Expected cmd 'SetBuzzerAlarmV20', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetBuzzerAlarmV20",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test SetBuzzerAlarmV20
	ctx := context.Background()
	err := client.Alarm.SetBuzzerAlarmV20(ctx, BuzzerAlarm{
		Channel: 0,
		Enable:  1,
	})
	if err != nil {
		t.Fatalf("SetBuzzerAlarmV20 failed: %v", err)
	}
}

func TestAlarmAPI_GetAudioAlarmV20(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAudioAlarmV20" {
			t.Errorf("Expected cmd 'GetAudioAlarmV20', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAudioAlarmV20",
			Code: 0,
			Value: json.RawMessage(`{
				"AudioAlarm": {
					"channel": 0,
					"enable": 1,
					"sensitivity": 50
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test GetAudioAlarmV20
	ctx := context.Background()
	audioAlarm, err := client.Alarm.GetAudioAlarmV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetAudioAlarmV20 failed: %v", err)
	}

	// Verify response
	if audioAlarm.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", audioAlarm.Channel)
	}
	if audioAlarm.Enable != 1 {
		t.Errorf("Expected enable 1, got %d", audioAlarm.Enable)
	}
	if audioAlarm.Sensitivity != 50 {
		t.Errorf("Expected sensitivity 50, got %d", audioAlarm.Sensitivity)
	}
}

func TestAlarmAPI_SetAudioAlarmV20(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAudioAlarmV20" {
			t.Errorf("Expected cmd 'SetAudioAlarmV20', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetAudioAlarmV20",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.Alarm = &AlarmAPI{client: client}

	// Test SetAudioAlarmV20
	ctx := context.Background()
	err := client.Alarm.SetAudioAlarmV20(ctx, AudioAlarm{
		Channel:     0,
		Enable:      1,
		Sensitivity: 50,
	})
	if err != nil {
		t.Fatalf("SetAudioAlarmV20 failed: %v", err)
	}
}
