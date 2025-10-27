package reolink

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAIAPI_GetAiCfg(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAiCfg" {
			t.Errorf("Expected cmd 'GetAiCfg', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAiCfg",
			Code: 0,
			Value: json.RawMessage(`{
				"channel": 0,
				"aiTrack": 1,
				"AiDetectType": {
					"people": 1,
					"vehicle": 1,
					"dog_cat": 1,
					"face": 0
				},
				"trackType": {
					"people": 1,
					"vehicle": 0,
					"dog_cat": 0,
					"face": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.AI = &AIAPI{client: client}

	// Test GetAiCfg
	ctx := t.Context()
	cfg, err := client.AI.GetAiCfg(ctx, 0)
	if err != nil {
		t.Fatalf("GetAiCfg failed: %v", err)
	}

	// Verify response
	if cfg.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", cfg.Channel)
	}
	if cfg.AiTrack != 1 {
		t.Errorf("Expected aiTrack 1, got %d", cfg.AiTrack)
	}
	if cfg.AiDetectType.People != 1 {
		t.Errorf("Expected people detection enabled, got %d", cfg.AiDetectType.People)
	}
	if cfg.AiDetectType.Vehicle != 1 {
		t.Errorf("Expected vehicle detection enabled, got %d", cfg.AiDetectType.Vehicle)
	}
	if cfg.AiDetectType.DogCat != 1 {
		t.Errorf("Expected dog_cat detection enabled, got %d", cfg.AiDetectType.DogCat)
	}
	if cfg.AiDetectType.Face != 0 {
		t.Errorf("Expected face detection disabled, got %d", cfg.AiDetectType.Face)
	}
	if cfg.TrackType.People != 1 {
		t.Errorf("Expected people tracking enabled, got %d", cfg.TrackType.People)
	}
}

func TestAIAPI_SetAiCfg(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "SetAiCfg" {
			t.Errorf("Expected cmd 'SetAiCfg', got '%s'", req[0].Cmd)
		}

		// Send success response
		resp := []Response{{
			Cmd:   "SetAiCfg",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.AI = &AIAPI{client: client}

	// Test SetAiCfg
	ctx := t.Context()
	config := AiCfg{
		Channel: 0,
		AiTrack: 1,
		AiDetectType: AiDetectType{
			People:  1,
			Vehicle: 1,
			DogCat:  1,
			Face:    0,
		},
		TrackType: AiTrackType{
			People:  1,
			Vehicle: 0,
			DogCat:  0,
			Face:    0,
		},
	}

	err := client.AI.SetAiCfg(ctx, config)
	if err != nil {
		t.Fatalf("SetAiCfg failed: %v", err)
	}
}

func TestAIAPI_GetAiState(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req[0].Cmd != "GetAiState" {
			t.Errorf("Expected cmd 'GetAiState', got '%s'", req[0].Cmd)
		}

		// Send mock response
		resp := []Response{{
			Cmd:  "GetAiState",
			Code: 0,
			Value: json.RawMessage(`{
				"channel": 0,
				"people": {
					"alarm_state": 1,
					"support": 1
				},
				"vehicle": {
					"alarm_state": 0,
					"support": 1
				},
				"dog_cat": {
					"alarm_state": 0,
					"support": 1
				},
				"face": {
					"alarm_state": 0,
					"support": 0
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := newTestClient(server)
	client.AI = &AIAPI{client: client}

	// Test GetAiState
	ctx := t.Context()
	state, err := client.AI.GetAiState(ctx, 0)
	if err != nil {
		t.Fatalf("GetAiState failed: %v", err)
	}

	// Verify response
	if state.Channel != 0 {
		t.Errorf("Expected channel 0, got %d", state.Channel)
	}
	if state.People.AlarmState != 1 {
		t.Errorf("Expected people alarm state 1, got %d", state.People.AlarmState)
	}
	if state.People.Support != 1 {
		t.Errorf("Expected people support 1, got %d", state.People.Support)
	}
	if state.Vehicle.AlarmState != 0 {
		t.Errorf("Expected vehicle alarm state 0, got %d", state.Vehicle.AlarmState)
	}
	if state.Vehicle.Support != 1 {
		t.Errorf("Expected vehicle support 1, got %d", state.Vehicle.Support)
	}
	if state.DogCat.AlarmState != 0 {
		t.Errorf("Expected dog_cat alarm state 0, got %d", state.DogCat.AlarmState)
	}
	if state.DogCat.Support != 1 {
		t.Errorf("Expected dog_cat support 1, got %d", state.DogCat.Support)
	}
	if state.Face.AlarmState != 0 {
		t.Errorf("Expected face alarm state 0, got %d", state.Face.AlarmState)
	}
	if state.Face.Support != 0 {
		t.Errorf("Expected face support 0, got %d", state.Face.Support)
	}
}
