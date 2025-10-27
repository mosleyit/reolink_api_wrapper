package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// LEDAPI provides access to LED and light control API endpoints
type LEDAPI struct {
	client *Client
}

// LED state constants
const (
	LEDStateAuto = "Auto"
	LEDStateOn   = "On"
	LEDStateOff  = "Off"
)

// IrLights represents IR lights configuration
type IrLights struct {
	State string `json:"state"` // Auto, On, Off
}

// IrLightsValue wraps IrLights for API response
type IrLightsValue struct {
	IrLights IrLights `json:"IrLights"`
}

// IrLightsParam represents parameters for SetIrLights
type IrLightsParam struct {
	IrLights struct {
		Channel int    `json:"channel"` // Channel number
		State   string `json:"state"`   // Auto, On, Off
	} `json:"IrLights"`
}

// PowerLed represents power LED configuration
type PowerLed struct {
	State string `json:"state"` // Auto, Off
}

// PowerLedValue wraps PowerLed for API response
type PowerLedValue struct {
	PowerLed PowerLed `json:"PowerLed"`
}

// PowerLedParam represents parameters for SetPowerLed
type PowerLedParam struct {
	PowerLed struct {
		Channel int    `json:"channel"` // Channel number
		State   string `json:"state"`   // Auto, Off
	} `json:"PowerLed"`
}

// WhiteLedSchedule represents lighting schedule
type WhiteLedSchedule struct {
	StartHour int `json:"StartHour"` // Start hour (0-23)
	StartMin  int `json:"StartMin"`  // Start minute (0-59)
	EndHour   int `json:"EndHour"`   // End hour (0-23)
	EndMin    int `json:"EndMin"`    // End minute (0-59)
}

// WhiteLedAiDetect represents AI detection types for white LED
type WhiteLedAiDetect struct {
	People  int `json:"people"`  // 0=disabled, 1=enabled
	Vehicle int `json:"vehicle"` // 0=disabled, 1=enabled
	DogCat  int `json:"dog_cat"` // 0=disabled, 1=enabled
	Face    int `json:"face"`    // 0=disabled, 1=enabled
}

// WhiteLed represents white LED configuration
type WhiteLed struct {
	Channel          int              `json:"channel"`          // Channel number
	State            int              `json:"state"`            // 0=off, 1=on
	Mode             int              `json:"mode"`             // 0=always on, 1=alarm trigger, 2=auto with AI
	Bright           int              `json:"bright"`           // Brightness (0-100)
	LightingSchedule WhiteLedSchedule `json:"LightingSchedule"` // Schedule for mode 2
	WlAiDetectType   WhiteLedAiDetect `json:"wlAiDetectType"`   // AI detection types
}

// WhiteLedValue wraps WhiteLed for API response
type WhiteLedValue struct {
	WhiteLed WhiteLed `json:"WhiteLed"`
}

// WhiteLedParam represents parameters for SetWhiteLed
type WhiteLedParam struct {
	WhiteLed WhiteLed `json:"WhiteLed"`
}

// GetIrLights gets IR lights configuration
func (l *LEDAPI) GetIrLights(ctx context.Context) (*IrLights, error) {
	l.client.logger.Debug("getting IR lights configuration")

	req := []Request{{
		Cmd:    "GetIrLights",
		Action: 1, // Get initial, range, and value
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to get IR lights configuration: %v", err)
		return nil, fmt.Errorf("GetIrLights request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to get IR lights configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to get IR lights configuration: %v", apiErr)
		return nil, apiErr
	}

	var value IrLightsValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		l.client.logger.Error("failed to parse IR lights configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	l.client.logger.Info("successfully retrieved IR lights configuration: state=%s", value.IrLights.State)
	return &value.IrLights, nil
}

// SetIrLights sets IR lights configuration
func (l *LEDAPI) SetIrLights(ctx context.Context, channel int, state string) error {
	l.client.logger.Info("setting IR lights configuration: channel=%d state=%s", channel, state)

	var param IrLightsParam
	param.IrLights.Channel = channel
	param.IrLights.State = state

	req := []Request{{
		Cmd:   "SetIrLights",
		Param: param,
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to set IR lights configuration: %v", err)
		return fmt.Errorf("SetIrLights request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to set IR lights configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to set IR lights configuration: %v", apiErr)
		return apiErr
	}

	l.client.logger.Info("successfully set IR lights configuration")
	return nil
}

// GetPowerLed gets power LED configuration
func (l *LEDAPI) GetPowerLed(ctx context.Context, channel int) (*PowerLed, error) {
	l.client.logger.Debug("getting power LED configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetPowerLed",
		Action: 1,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to get power LED configuration: %v", err)
		return nil, fmt.Errorf("GetPowerLed request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to get power LED configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to get power LED configuration: %v", apiErr)
		return nil, apiErr
	}

	var value PowerLedValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		l.client.logger.Error("failed to parse power LED configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	l.client.logger.Info("successfully retrieved power LED configuration: state=%s", value.PowerLed.State)
	return &value.PowerLed, nil
}

// SetPowerLed sets power LED configuration
func (l *LEDAPI) SetPowerLed(ctx context.Context, channel int, state string) error {
	l.client.logger.Info("setting power LED configuration: channel=%d state=%s", channel, state)

	var param PowerLedParam
	param.PowerLed.Channel = channel
	param.PowerLed.State = state

	req := []Request{{
		Cmd:   "SetPowerLed",
		Param: param,
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to set power LED configuration: %v", err)
		return fmt.Errorf("SetPowerLed request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to set power LED configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to set power LED configuration: %v", apiErr)
		return apiErr
	}

	l.client.logger.Info("successfully set power LED configuration")
	return nil
}

// GetWhiteLed gets white LED configuration
func (l *LEDAPI) GetWhiteLed(ctx context.Context, channel int) (*WhiteLed, error) {
	l.client.logger.Debug("getting white LED configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetWhiteLed",
		Action: 1,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to get white LED configuration: %v", err)
		return nil, fmt.Errorf("GetWhiteLed request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to get white LED configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to get white LED configuration: %v", apiErr)
		return nil, apiErr
	}

	var value WhiteLedValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		l.client.logger.Error("failed to parse white LED configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	l.client.logger.Info("successfully retrieved white LED configuration: state=%d mode=%d bright=%d",
		value.WhiteLed.State, value.WhiteLed.Mode, value.WhiteLed.Bright)
	return &value.WhiteLed, nil
}

// SetWhiteLed sets white LED configuration
func (l *LEDAPI) SetWhiteLed(ctx context.Context, config WhiteLed) error {
	l.client.logger.Info("setting white LED configuration: channel=%d state=%d mode=%d bright=%d",
		config.Channel, config.State, config.Mode, config.Bright)

	req := []Request{{
		Cmd: "SetWhiteLed",
		Param: WhiteLedParam{
			WhiteLed: config,
		},
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to set white LED configuration: %v", err)
		return fmt.Errorf("SetWhiteLed request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to set white LED configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to set white LED configuration: %v", apiErr)
		return apiErr
	}

	l.client.logger.Info("successfully set white LED configuration")
	return nil
}

// AiAlarm represents AI-based alarm configuration
type AiAlarm struct {
	Channel         int     `json:"channel"`
	AiType          string  `json:"ai_type"`           // people, vehicle, dog_cat, face
	Sensitivity     int     `json:"sensitivity"`       // Sensitivity level
	StayTime        int     `json:"stay_time"`         // Stay time in seconds
	Width           int     `json:"width"`             // Detection area width
	Height          int     `json:"height"`            // Detection area height
	Scope           *Scope  `json:"scope,omitempty"`   // Detection scope/area
	MinTargetHeight float64 `json:"min_target_height"` // Minimum target height (0.0-1.0)
	MaxTargetHeight float64 `json:"max_target_height"` // Maximum target height (0.0-1.0)
	MinTargetWidth  float64 `json:"min_target_width"`  // Minimum target width (0.0-1.0)
	MaxTargetWidth  float64 `json:"max_target_width"`  // Maximum target width (0.0-1.0)
}

// Scope represents detection scope/area
type Scope struct {
	Area string `json:"area"` // Detection grid (width × height characters)
}

// AiAlarmValue wraps AiAlarm for API response
type AiAlarmValue struct {
	AiAlarm AiAlarm `json:"AiAlarm"`
}

// AiAlarmParam represents parameters for SetAiAlarm
type AiAlarmParam struct {
	Channel int     `json:"channel"`
	AiAlarm AiAlarm `json:"AiAlarm"`
}

// GetAiAlarm gets AI-based alarm configuration
func (l *LEDAPI) GetAiAlarm(ctx context.Context, channel int, aiType string) (*AiAlarm, error) {
	l.client.logger.Debug("getting AI alarm configuration: channel=%d aiType=%s", channel, aiType)

	req := []Request{{
		Cmd:    "GetAiAlarm",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to get AI alarm configuration: %v", err)
		return nil, fmt.Errorf("GetAiAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to get AI alarm configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to get AI alarm configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AiAlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		l.client.logger.Error("failed to parse AI alarm configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetAiAlarm response: %w", err)
	}

	l.client.logger.Info("successfully retrieved AI alarm configuration: aiType=%s sensitivity=%d",
		value.AiAlarm.AiType, value.AiAlarm.Sensitivity)
	return &value.AiAlarm, nil
}

// SetAiAlarm sets AI-based alarm configuration
func (l *LEDAPI) SetAiAlarm(ctx context.Context, channel int, alarm AiAlarm) error {
	l.client.logger.Info("setting AI alarm configuration: channel=%d aiType=%s sensitivity=%d",
		channel, alarm.AiType, alarm.Sensitivity)

	req := []Request{{
		Cmd: "SetAiAlarm",
		Param: AiAlarmParam{
			Channel: channel,
			AiAlarm: alarm,
		},
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to set AI alarm configuration: %v", err)
		return fmt.Errorf("SetAiAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to set AI alarm configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to set AI alarm configuration: %v", apiErr)
		return apiErr
	}

	l.client.logger.Info("successfully set AI alarm configuration")
	return nil
}

// SetAlarmArea sets alarm detection area/zone
func (l *LEDAPI) SetAlarmArea(ctx context.Context, params map[string]interface{}) error {
	l.client.logger.Info("setting alarm detection area")

	req := []Request{{
		Cmd:   "SetAlarmArea",
		Param: params,
	}}

	var resp []Response
	if err := l.client.do(ctx, req, &resp); err != nil {
		l.client.logger.Error("failed to set alarm detection area: %v", err)
		return fmt.Errorf("SetAlarmArea request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		l.client.logger.Error("failed to set alarm detection area: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		l.client.logger.Error("failed to set alarm detection area: %v", apiErr)
		return apiErr
	}

	l.client.logger.Info("successfully set alarm detection area")
	return nil
}
