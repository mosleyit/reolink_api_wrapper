package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// AIAPI provides access to AI detection and tracking API endpoints
type AIAPI struct {
	client *Client
}

// AiDetectType represents AI detection type configuration
type AiDetectType struct {
	People  int `json:"people"`  // 0=disabled, 1=enabled
	Vehicle int `json:"vehicle"` // 0=disabled, 1=enabled
	DogCat  int `json:"dog_cat"` // 0=disabled, 1=enabled
	Face    int `json:"face"`    // 0=disabled, 1=enabled
}

// AiTrackType represents AI tracking type configuration
type AiTrackType struct {
	People  int `json:"people"`  // 0=disabled, 1=enabled
	Vehicle int `json:"vehicle"` // 0=disabled, 1=enabled
	DogCat  int `json:"dog_cat"` // 0=disabled, 1=enabled
	Face    int `json:"face"`    // 0=disabled, 1=enabled
}

// AiCfg represents AI configuration
type AiCfg struct {
	Channel      int          `json:"channel"`      // Channel number
	AiTrack      int          `json:"aiTrack"`      // AI tracking switch (0=off, 1=on)
	AiDetectType AiDetectType `json:"AiDetectType"` // AI detection types
	TrackType    AiTrackType  `json:"trackType"`    // AI tracking types
}

// AiDetectState represents AI detection state for a specific type
type AiDetectState struct {
	AlarmState int `json:"alarm_state"` // 0=no alarm, 1=alarm detected
	Support    int `json:"support"`     // 0=not supported, 1=supported
}

// AiState represents AI alarm state
type AiState struct {
	Channel int           `json:"channel"` // Channel number
	People  AiDetectState `json:"people"`  // People detection state
	Vehicle AiDetectState `json:"vehicle"` // Vehicle detection state
	DogCat  AiDetectState `json:"dog_cat"` // Dog/cat detection state
	Face    AiDetectState `json:"face"`    // Face detection state
}

// GetAiCfg gets AI configuration
func (a *AIAPI) GetAiCfg(ctx context.Context, channel int) (*AiCfg, error) {
	a.client.logger.Debug("getting AI configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetAiCfg",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get AI configuration: %v", err)
		return nil, fmt.Errorf("GetAiCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get AI configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get AI configuration: %v", apiErr)
		return nil, apiErr
	}

	var cfg AiCfg
	if err := json.Unmarshal(resp[0].Value, &cfg); err != nil {
		a.client.logger.Error("failed to parse AI configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &cfg, nil
}

// SetAiCfg sets AI configuration
func (a *AIAPI) SetAiCfg(ctx context.Context, config AiCfg) error {
	a.client.logger.Info("setting AI configuration: channel=%d people=%d vehicle=%d dog_cat=%d face=%d",
		config.Channel, config.AiDetectType.People, config.AiDetectType.Vehicle,
		config.AiDetectType.DogCat, config.AiDetectType.Face)

	req := []Request{{
		Cmd:    "SetAiCfg",
		Action: 0,
		Param:  config,
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set AI configuration: %v", err)
		return fmt.Errorf("SetAiCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set AI configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set AI configuration: %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set AI configuration")
	return nil
}

// GetAiState gets AI alarm state
func (a *AIAPI) GetAiState(ctx context.Context, channel int) (*AiState, error) {
	a.client.logger.Debug("getting AI state: channel=%d", channel)

	req := []Request{{
		Cmd: "GetAiState",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get AI state: %v", err)
		return nil, fmt.Errorf("GetAiState request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get AI state: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get AI state: %v", apiErr)
		return nil, apiErr
	}

	var state AiState
	if err := json.Unmarshal(resp[0].Value, &state); err != nil {
		a.client.logger.Error("failed to parse AI state response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved AI state: people=%d vehicle=%d dog_cat=%d face=%d",
		state.People.AlarmState, state.Vehicle.AlarmState, state.DogCat.AlarmState, state.Face.AlarmState)
	return &state, nil
}
