package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// AlarmAPI provides access to alarm and motion detection API endpoints
type AlarmAPI struct {
	client *Client
}

// MdScope represents motion detection scope/area
type MdScope struct {
	Cols  int    `json:"cols"`  // Number of columns in grid (typically 80)
	Rows  int    `json:"rows"`  // Number of rows in grid (typically 60 or 45)
	Table string `json:"table"` // Grid string of 1s and 0s (length = cols × rows)
}

// MdSensitivity represents time-based sensitivity settings
type MdSensitivity struct {
	ID          int `json:"id"`          // Time period ID (0-3)
	BeginHour   int `json:"beginHour"`   // Start hour (0-23)
	BeginMin    int `json:"beginMin"`    // Start minute (0-59)
	EndHour     int `json:"endHour"`     // End hour (0-23)
	EndMin      int `json:"endMin"`      // End minute (0-59)
	Enable      int `json:"enable"`      // 0=disabled, 1=enabled
	Priority    int `json:"priority"`    // Priority level
	Sensitivity int `json:"sensitivity"` // Sensitivity (0-100, higher = more sensitive)
}

// MdNewSens wraps sensitivity array
type MdNewSens struct {
	Sens []MdSensitivity `json:"sens"` // Array of up to 4 time periods
}

// MdAlarm represents motion detection alarm configuration
type MdAlarm struct {
	Channel int       `json:"channel"` // Channel number
	Scope   MdScope   `json:"scope"`   // Detection area
	NewSens MdNewSens `json:"newSens"` // Time-based sensitivity
}

// MdAlarmValue wraps MdAlarm for API response
type MdAlarmValue struct {
	MdAlarm MdAlarm `json:"MdAlarm"`
}

// MdAlarmParam represents parameters for SetMdAlarm
type MdAlarmParam struct {
	MdAlarm MdAlarm `json:"MdAlarm"`
}

// MdStateValue represents motion detection state
type MdStateValue struct {
	State int `json:"state"` // 0=no motion, 1=motion detected
}

// AudioAlarmPlayParam represents parameters for AudioAlarmPlay
type AudioAlarmPlayParam struct {
	Channel      int    `json:"channel"`       // Channel number
	AlarmMode    string `json:"alarm_mode"`    // Alarm mode
	ManualSwitch int    `json:"manual_switch"` // 0=off, 1=on
	Times        int    `json:"times"`         // Number of times to play
}

// GetMdState gets current motion detection state
func (a *AlarmAPI) GetMdState(ctx context.Context, channel int) (int, error) {
	a.client.logger.Debug("getting motion detection state: channel=%d", channel)

	req := []Request{{
		Cmd: "GetMdState",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get motion detection state: %v", err)
		return 0, fmt.Errorf("GetMdState request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get motion detection state: %v", err)
		return 0, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get motion detection state: %v", apiErr)
		return 0, apiErr
	}

	var value MdStateValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse motion detection state response: %v", err)
		return 0, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved motion detection state: state=%d", value.State)
	return value.State, nil
}

// GetMdAlarm gets motion detection alarm configuration
func (a *AlarmAPI) GetMdAlarm(ctx context.Context, channel int) (*MdAlarm, error) {
	a.client.logger.Debug("getting motion detection alarm configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetMdAlarm",
		Action: 1, // Get initial, range, and value
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get motion detection alarm configuration: %v", err)
		return nil, fmt.Errorf("GetMdAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get motion detection alarm configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get motion detection alarm configuration: %v", apiErr)
		return nil, apiErr
	}

	var value MdAlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse motion detection alarm configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved motion detection alarm configuration: channel=%d",
		value.MdAlarm.Channel)
	return &value.MdAlarm, nil
}

// SetMdAlarm sets motion detection alarm configuration
func (a *AlarmAPI) SetMdAlarm(ctx context.Context, config MdAlarm) error {
	a.client.logger.Info("setting motion detection alarm configuration: channel=%d",
		config.Channel)

	req := []Request{{
		Cmd: "SetMdAlarm",
		Param: MdAlarmParam{
			MdAlarm: config,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set motion detection alarm configuration: %v", err)
		return fmt.Errorf("SetMdAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set motion detection alarm configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set motion detection alarm configuration: %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set motion detection alarm configuration")
	return nil
}

// AudioAlarmPlay plays audio alarm sound
func (a *AlarmAPI) AudioAlarmPlay(ctx context.Context, param AudioAlarmPlayParam) error {
	a.client.logger.Info("playing audio alarm: channel=%d", param.Channel)

	req := []Request{{
		Cmd:   "AudioAlarmPlay",
		Param: param,
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to play audio alarm: %v", err)
		return fmt.Errorf("AudioAlarmPlay request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to play audio alarm: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to play audio alarm: %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully played audio alarm")
	return nil
}

// Alarm represents general alarm configuration
type Alarm struct {
	Channel int      `json:"channel"` // Channel number
	Type    string   `json:"type"`    // Alarm type (e.g., "md" for motion detection)
	Enable  int      `json:"enable"`  // 0=disabled, 1=enabled
	Scope   MdScope  `json:"scope"`   // Detection area
	Sens    []MdSens `json:"sens"`    // Time-based sensitivity settings
}

// MdSens represents simplified sensitivity settings for GetAlarm
type MdSens struct {
	BeginHour   int `json:"beginHour"`   // Start hour (0-23)
	BeginMin    int `json:"beginMin"`    // Start minute (0-59)
	EndHour     int `json:"endHour"`     // End hour (0-23)
	EndMin      int `json:"endMin"`      // End minute (0-59)
	Sensitivity int `json:"sensitivity"` // Sensitivity (0-100)
}

// AlarmValue wraps Alarm for API response
type AlarmValue struct {
	Alarm Alarm `json:"Alarm"`
}

// AudioAlarm represents audio detection alarm configuration
type AudioAlarm struct {
	Channel     int                `json:"channel"`     // Channel number
	Enable      int                `json:"enable"`      // 0=disabled, 1=enabled
	Sensitivity int                `json:"sensitivity"` // Audio sensitivity (0-100)
	Schedule    AudioAlarmSchedule `json:"schedule"`    // Schedule configuration
}

// AudioAlarmSchedule represents audio alarm schedule
type AudioAlarmSchedule struct {
	Enable int         `json:"enable"` // 0=disabled, 1=enabled
	Table  interface{} `json:"table"`  // string for v1, map for v2.0
}

// AudioAlarmValue wraps AudioAlarm for API response
type AudioAlarmValue struct {
	AudioAlarm AudioAlarm `json:"AudioAlarm"`
}

// BuzzerAlarm represents buzzer alarm configuration
type BuzzerAlarm struct {
	Channel  int                 `json:"channel"`  // Channel number
	Enable   int                 `json:"enable"`   // 0=disabled, 1=enabled
	Schedule BuzzerAlarmSchedule `json:"schedule"` // Schedule configuration
}

// BuzzerAlarmSchedule represents buzzer alarm schedule
type BuzzerAlarmSchedule struct {
	Enable int         `json:"enable"` // 0=disabled, 1=enabled
	Table  interface{} `json:"table"`  // string for v1, map for v2.0
}

// BuzzerAlarmValue wraps BuzzerAlarm for API response
type BuzzerAlarmValue struct {
	BuzzerAlarm BuzzerAlarm `json:"BuzzerAlarm"`
}

// GetAlarm gets general alarm configuration
func (a *AlarmAPI) GetAlarm(ctx context.Context, channel int, alarmType string) (*Alarm, error) {
	a.client.logger.Debug("getting alarm configuration: channel=%d type=%s", channel, alarmType)

	req := []Request{{
		Cmd:    "GetAlarm",
		Action: 1,
		Param: map[string]interface{}{
			"Alarm": map[string]interface{}{
				"channel": channel,
				"type":    alarmType,
			},
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get alarm configuration: %v", err)
		return nil, fmt.Errorf("GetAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get alarm configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get alarm configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse alarm configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved alarm configuration: type=%s enable=%d", value.Alarm.Type, value.Alarm.Enable)
	return &value.Alarm, nil
}

// SetAlarm sets general alarm configuration
func (a *AlarmAPI) SetAlarm(ctx context.Context, alarm Alarm) error {
	a.client.logger.Info("setting alarm configuration: channel=%d type=%s enable=%d", alarm.Channel, alarm.Type, alarm.Enable)

	req := []Request{{
		Cmd: "SetAlarm",
		Param: map[string]interface{}{
			"Alarm": alarm,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set alarm configuration: %v", err)
		return fmt.Errorf("SetAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set alarm configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set alarm configuration: %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set alarm configuration")
	return nil
}

// GetAudioAlarm gets audio detection alarm configuration
func (a *AlarmAPI) GetAudioAlarm(ctx context.Context, channel int) (*AudioAlarm, error) {
	a.client.logger.Debug("getting audio alarm configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetAudioAlarm",
		Action: 1,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get audio alarm configuration: %v", err)
		return nil, fmt.Errorf("GetAudioAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get audio alarm configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get audio alarm configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AudioAlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse audio alarm configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved audio alarm configuration: enable=%d sensitivity=%d",
		value.AudioAlarm.Enable, value.AudioAlarm.Sensitivity)
	return &value.AudioAlarm, nil
}

// SetAudioAlarm sets audio detection alarm configuration
func (a *AlarmAPI) SetAudioAlarm(ctx context.Context, audioAlarm AudioAlarm) error {
	a.client.logger.Info("setting audio alarm configuration: channel=%d enable=%d sensitivity=%d",
		audioAlarm.Channel, audioAlarm.Enable, audioAlarm.Sensitivity)

	req := []Request{{
		Cmd: "SetAudioAlarm",
		Param: map[string]interface{}{
			"Audio": audioAlarm,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set audio alarm configuration: %v", err)
		return fmt.Errorf("SetAudioAlarm request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set audio alarm configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set audio alarm configuration: %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set audio alarm configuration")
	return nil
}

// GetAudioAlarmV20 gets audio detection alarm configuration (v2.0)
func (a *AlarmAPI) GetAudioAlarmV20(ctx context.Context, channel int) (*AudioAlarm, error) {
	a.client.logger.Debug("getting audio alarm configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd:    "GetAudioAlarmV20",
		Action: 1,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get audio alarm configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetAudioAlarmV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get audio alarm configuration (v2.0): %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get audio alarm configuration (v2.0): %v", apiErr)
		return nil, apiErr
	}

	var value AudioAlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse audio alarm configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved audio alarm configuration (v2.0): enable=%d sensitivity=%d",
		value.AudioAlarm.Enable, value.AudioAlarm.Sensitivity)
	return &value.AudioAlarm, nil
}

// SetAudioAlarmV20 sets audio detection alarm configuration (v2.0)
func (a *AlarmAPI) SetAudioAlarmV20(ctx context.Context, audioAlarm AudioAlarm) error {
	a.client.logger.Info("setting audio alarm configuration (v2.0): channel=%d enable=%d sensitivity=%d",
		audioAlarm.Channel, audioAlarm.Enable, audioAlarm.Sensitivity)

	req := []Request{{
		Cmd: "SetAudioAlarmV20",
		Param: map[string]interface{}{
			"Audio": audioAlarm,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set audio alarm configuration (v2.0): %v", err)
		return fmt.Errorf("SetAudioAlarmV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set audio alarm configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set audio alarm configuration (v2.0): %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set audio alarm configuration (v2.0)")
	return nil
}

// GetBuzzerAlarmV20 gets buzzer alarm configuration (v2.0)
func (a *AlarmAPI) GetBuzzerAlarmV20(ctx context.Context, channel int) (*BuzzerAlarm, error) {
	a.client.logger.Debug("getting buzzer alarm configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd:    "GetBuzzerAlarmV20",
		Action: 1,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to get buzzer alarm configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetBuzzerAlarmV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to get buzzer alarm configuration (v2.0): %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to get buzzer alarm configuration (v2.0): %v", apiErr)
		return nil, apiErr
	}

	var value BuzzerAlarmValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		a.client.logger.Error("failed to parse buzzer alarm configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	a.client.logger.Info("successfully retrieved buzzer alarm configuration (v2.0): enable=%d",
		value.BuzzerAlarm.Enable)
	return &value.BuzzerAlarm, nil
}

// SetBuzzerAlarmV20 sets buzzer alarm configuration (v2.0)
func (a *AlarmAPI) SetBuzzerAlarmV20(ctx context.Context, buzzerAlarm BuzzerAlarm) error {
	a.client.logger.Info("setting buzzer alarm configuration (v2.0): channel=%d enable=%d",
		buzzerAlarm.Channel, buzzerAlarm.Enable)

	req := []Request{{
		Cmd: "SetBuzzerAlarmV20",
		Param: map[string]interface{}{
			"BuzzerAlarm": buzzerAlarm,
		},
	}}

	var resp []Response
	if err := a.client.do(ctx, req, &resp); err != nil {
		a.client.logger.Error("failed to set buzzer alarm configuration (v2.0): %v", err)
		return fmt.Errorf("SetBuzzerAlarmV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		a.client.logger.Error("failed to set buzzer alarm configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		a.client.logger.Error("failed to set buzzer alarm configuration (v2.0): %v", apiErr)
		return apiErr
	}

	a.client.logger.Info("successfully set buzzer alarm configuration (v2.0)")
	return nil
}
