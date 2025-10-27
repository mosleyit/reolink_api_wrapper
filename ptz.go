package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// PTZAPI provides access to Pan-Tilt-Zoom control endpoints
type PTZAPI struct {
	client *Client
}

// PTZ operation constants
const (
	PTZOpStop        = "Stop"
	PTZOpLeft        = "Left"
	PTZOpRight       = "Right"
	PTZOpUp          = "Up"
	PTZOpDown        = "Down"
	PTZOpLeftUp      = "LeftUp"
	PTZOpLeftDown    = "LeftDown"
	PTZOpRightUp     = "RightUp"
	PTZOpRightDown   = "RightDown"
	PTZOpZoomInc     = "ZoomInc"
	PTZOpZoomDec     = "ZoomDec"
	PTZOpFocusInc    = "FocusInc"
	PTZOpFocusDec    = "FocusDec"
	PTZOpIrisInc     = "IrisInc"
	PTZOpIrisDec     = "IrisDec"
	PTZOpAuto        = "Auto"
	PTZOpToPos       = "ToPos"
	PTZOpStartPatrol = "StartPatrol"
	PTZOpStopPatrol  = "StopPatrol"
)

// PtzPreset represents a PTZ preset position
type PtzPreset struct {
	Enable int    `json:"enable"` // 0=disabled, 1=enabled
	ID     int    `json:"id"`     // Preset ID (1-64)
	Name   string `json:"name"`   // Preset name
}

// PtzPresetValue wraps preset array for API response
type PtzPresetValue struct {
	PtzPreset []PtzPreset `json:"PtzPreset"`
}

// PtzPresetParam represents parameters for SetPtzPreset
type PtzPresetParam struct {
	PtzPreset PtzPreset `json:"PtzPreset"`
}

// PtzPatrolPreset represents a preset in a patrol
type PtzPatrolPreset struct {
	ID        int `json:"id"`        // Preset ID (1-64)
	DwellTime int `json:"dwellTime"` // Dwell time at preset (seconds)
	Speed     int `json:"speed"`     // Movement speed
}

// PtzPatrol represents a PTZ patrol/tour configuration
type PtzPatrol struct {
	Channel int               `json:"channel"` // Channel number
	Enable  int               `json:"enable"`  // 0=disabled, 1=enabled
	ID      int               `json:"id"`      // Patrol ID
	Running int               `json:"running"` // 0=stopped, 1=running
	Name    string            `json:"name"`    // Patrol name
	Preset  []PtzPatrolPreset `json:"preset"`  // List of presets (max 16)
}

// PtzPatrolValue wraps patrol for API response
type PtzPatrolValue struct {
	PtzPatrol PtzPatrol `json:"PtzPatrol"`
}

// PtzPatrolParam represents parameters for SetPtzPatrol
type PtzPatrolParam struct {
	PtzPatrol PtzPatrol `json:"PtzPatrol"`
}

// PtzGuard represents PTZ guard/home position configuration
type PtzGuard struct {
	Channel         int    `json:"channel"`         // Channel number
	CmdStr          string `json:"cmdStr"`          // Command string
	BEnable         int    `json:"benable"`         // 0=disabled, 1=enabled
	BExistPos       int    `json:"bexistPos"`       // Whether guard position exists
	Timeout         int    `json:"timeout"`         // Timeout in seconds (typically 60)
	BSaveCurrentPos int    `json:"bSaveCurrentPos"` // 1=save current position as guard
}

// PtzGuardValue wraps guard for API response
type PtzGuardValue struct {
	PtzGuard PtzGuard `json:"PtzGuard"`
}

// PtzGuardParam represents parameters for SetPtzGuard
type PtzGuardParam struct {
	PtzGuard PtzGuard `json:"PtzGuard"`
}

// PtzCtrlParam represents parameters for PtzCtrl command
type PtzCtrlParam struct {
	Channel int    `json:"channel"`         // Channel number
	Op      string `json:"op"`              // Operation (use PTZOp* constants)
	Speed   int    `json:"speed,omitempty"` // Speed (1-64, optional)
	ID      int    `json:"id,omitempty"`    // Preset/Patrol ID (optional)
}

// PtzCtrl controls PTZ movement
func (p *PTZAPI) PtzCtrl(ctx context.Context, param PtzCtrlParam) error {
	p.client.logger.Info("controlling PTZ: channel=%d op=%s speed=%d",
		param.Channel, param.Op, param.Speed)

	req := []Request{{
		Cmd:   "PtzCtrl",
		Param: param,
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to control PTZ: %v", err)
		return fmt.Errorf("PtzCtrl request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to control PTZ: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to control PTZ: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully controlled PTZ")
	return nil
}

// GetPtzPreset gets PTZ preset positions
func (p *PTZAPI) GetPtzPreset(ctx context.Context, channel int) ([]PtzPreset, error) {
	p.client.logger.Debug("getting PTZ presets: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetPtzPreset",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ presets: %v", err)
		return nil, fmt.Errorf("GetPtzPreset request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ presets: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ presets: %v", apiErr)
		return nil, apiErr
	}

	var value PtzPresetValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse PTZ presets response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ presets: count=%d", len(value.PtzPreset))
	return value.PtzPreset, nil
}

// SetPtzPreset sets or calls a PTZ preset position
func (p *PTZAPI) SetPtzPreset(ctx context.Context, preset PtzPreset) error {
	p.client.logger.Info("setting PTZ preset: id=%d name=%s", preset.ID, preset.Name)

	req := []Request{{
		Cmd: "SetPtzPreset",
		Param: PtzPresetParam{
			PtzPreset: preset,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set PTZ preset: %v", err)
		return fmt.Errorf("SetPtzPreset request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set PTZ preset: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set PTZ preset: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set PTZ preset")
	return nil
}

// GetPtzPatrol gets PTZ patrol/tour configuration
func (p *PTZAPI) GetPtzPatrol(ctx context.Context, channel int) (*PtzPatrol, error) {
	p.client.logger.Debug("getting PTZ patrol configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetPtzPatrol",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ patrol configuration: %v", err)
		return nil, fmt.Errorf("GetPtzPatrol request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ patrol configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ patrol configuration: %v", apiErr)
		return nil, apiErr
	}

	var value PtzPatrolValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse PTZ patrol configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ patrol configuration")
	return &value.PtzPatrol, nil
}

// SetPtzPatrol sets PTZ patrol/tour configuration
func (p *PTZAPI) SetPtzPatrol(ctx context.Context, patrol PtzPatrol) error {
	p.client.logger.Info("setting PTZ patrol configuration: channel=%d", patrol.Channel)

	req := []Request{{
		Cmd: "SetPtzPatrol",
		Param: PtzPatrolParam{
			PtzPatrol: patrol,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set PTZ patrol configuration: %v", err)
		return fmt.Errorf("SetPtzPatrol request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set PTZ patrol configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set PTZ patrol configuration: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set PTZ patrol configuration")
	return nil
}

// GetPtzGuard gets PTZ guard/home position configuration
func (p *PTZAPI) GetPtzGuard(ctx context.Context, channel int) (*PtzGuard, error) {
	p.client.logger.Debug("getting PTZ guard configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetPtzGuard",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ guard configuration: %v", err)
		return nil, fmt.Errorf("GetPtzGuard request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ guard configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ guard configuration: %v", apiErr)
		return nil, apiErr
	}

	var value PtzGuardValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse PTZ guard configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ guard configuration: enable=%d timeout=%d",
		value.PtzGuard.BEnable, value.PtzGuard.Timeout)
	return &value.PtzGuard, nil
}

// SetPtzGuard sets PTZ guard/home position configuration
func (p *PTZAPI) SetPtzGuard(ctx context.Context, guard PtzGuard) error {
	p.client.logger.Info("setting PTZ guard configuration: channel=%d enable=%d timeout=%d",
		guard.Channel, guard.BEnable, guard.Timeout)

	req := []Request{{
		Cmd: "SetPtzGuard",
		Param: PtzGuardParam{
			PtzGuard: guard,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set PTZ guard configuration: %v", err)
		return fmt.Errorf("SetPtzGuard request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set PTZ guard configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set PTZ guard configuration: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set PTZ guard configuration")
	return nil
}

// PtzCheckState represents PTZ calibration check state
type PtzCheckState struct {
	Status int `json:"status"` // Check state status (0=idle, 1=checking)
}

// GetPtzCheckState gets PTZ calibration check state
func (p *PTZAPI) GetPtzCheckState(ctx context.Context, channel int) (*PtzCheckState, error) {
	p.client.logger.Debug("getting PTZ check state: channel=%d", channel)

	req := []Request{{
		Cmd: "GetPtzCheckState",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ check state: %v", err)
		return nil, fmt.Errorf("GetPtzCheckState request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ check state: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ check state: %v", apiErr)
		return nil, apiErr
	}

	var state PtzCheckState
	if err := json.Unmarshal(resp[0].Value, &state); err != nil {
		p.client.logger.Error("failed to parse PTZ check state response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ check state: status=%d", state.Status)
	return &state, nil
}

// PtzCheck performs PTZ calibration check
func (p *PTZAPI) PtzCheck(ctx context.Context, channel int) error {
	p.client.logger.Info("performing PTZ calibration check: channel=%d", channel)

	req := []Request{{
		Cmd: "PtzCheck",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to perform PTZ calibration check: %v", err)
		return fmt.Errorf("PtzCheck request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to perform PTZ calibration check: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to perform PTZ calibration check: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully performed PTZ calibration check")
	return nil
}

// ZoomFocus represents zoom and focus position
type ZoomFocus struct {
	Channel int `json:"channel"` // Channel number
	Zoom    struct {
		Pos int `json:"pos"` // Zoom position
	} `json:"zoom"`
	Focus struct {
		Pos int `json:"pos"` // Focus position
	} `json:"focus"`
}

// ZoomFocusValue wraps ZoomFocus for API response
type ZoomFocusValue struct {
	ZoomFocus ZoomFocus `json:"ZoomFocus"`
}

// GetZoomFocus gets current zoom and focus position
func (p *PTZAPI) GetZoomFocus(ctx context.Context, channel int) (*ZoomFocus, error) {
	p.client.logger.Debug("getting zoom/focus position: channel=%d", channel)

	req := []Request{{
		Cmd: "GetZoomFocus",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get zoom/focus position: %v", err)
		return nil, fmt.Errorf("GetZoomFocus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get zoom/focus position: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get zoom/focus position: %v", apiErr)
		return nil, apiErr
	}

	var value ZoomFocusValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse zoom/focus position response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved zoom/focus position: zoom=%d focus=%d",
		value.ZoomFocus.Zoom.Pos, value.ZoomFocus.Focus.Pos)
	return &value.ZoomFocus, nil
}

// StartZoomFocus starts zoom or focus operation
// op: ZoomInc, ZoomDec, FocusInc, FocusDec
// pos: target position (optional, set to 0 if not used)
func (p *PTZAPI) StartZoomFocus(ctx context.Context, channel int, op string, pos int) error {
	p.client.logger.Info("starting zoom/focus operation: channel=%d op=%s pos=%d", channel, op, pos)

	req := []Request{{
		Cmd: "StartZoomFocus",
		Param: map[string]interface{}{
			"ZoomFocus": map[string]interface{}{
				"channel": channel,
				"op":      op,
				"pos":     pos,
			},
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to start zoom/focus operation: %v", err)
		return fmt.Errorf("StartZoomFocus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to start zoom/focus operation: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to start zoom/focus operation: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully started zoom/focus operation")
	return nil
}

// PtzTattern represents PTZ pattern/track configuration
// Note: API uses "Tattern" (typo) instead of "Pattern"
type PtzTattern struct {
	Enable int `json:"enable"` // 0=disabled, 1=enabled
	ID     int `json:"id"`     // Track ID (1-6)
}

// PtzTatternValue wraps PtzTattern for API response
type PtzTatternValue struct {
	PtzTattern PtzTattern `json:"PtzTattern"`
}

// GetPtzTattern gets PTZ pattern/track configuration
func (p *PTZAPI) GetPtzTattern(ctx context.Context, channel int) (*PtzTattern, error) {
	p.client.logger.Debug("getting PTZ pattern configuration: channel=%d", channel)

	req := []Request{{
		Cmd: "GetPtzTattern",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ pattern configuration: %v", err)
		return nil, fmt.Errorf("GetPtzTattern request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ pattern configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ pattern configuration: %v", apiErr)
		return nil, apiErr
	}

	var value PtzTatternValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse PTZ pattern configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ pattern configuration: enable=%d id=%d",
		value.PtzTattern.Enable, value.PtzTattern.ID)
	return &value.PtzTattern, nil
}

// SetPtzTattern sets PTZ pattern/track configuration
func (p *PTZAPI) SetPtzTattern(ctx context.Context, channel int, tattern PtzTattern) error {
	p.client.logger.Info("setting PTZ pattern configuration: channel=%d enable=%d id=%d",
		channel, tattern.Enable, tattern.ID)

	req := []Request{{
		Cmd: "SetPtzTattern",
		Param: map[string]interface{}{
			"PtzTattern": tattern,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set PTZ pattern configuration: %v", err)
		return fmt.Errorf("SetPtzTattern request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set PTZ pattern configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set PTZ pattern configuration: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set PTZ pattern configuration")
	return nil
}

// PtzSerial represents PTZ serial port configuration
type PtzSerial struct {
	Channel      int    `json:"channel"`      // Channel number
	BaudRate     int    `json:"baudRate"`     // Baud rate (1200, 2400, 4800, 9600)
	CtrlAddr     int    `json:"ctrlAddr"`     // Control address (1-64)
	CtrlProtocol string `json:"ctrlProtocol"` // Control protocol (PELCO_D, PELCO_P)
	DataBit      string `json:"dataBit"`      // Data bits (CS5, CS6, CS7, CS8)
	FlowCtrl     string `json:"flowCtrl"`     // Flow control (none, hard, xon, xoff)
	Parity       string `json:"parity"`       // Parity (none, odd, even)
	StopBit      int    `json:"stopBit"`      // Stop bits (1, 2)
}

// PtzSerialValue wraps PtzSerial for API response
type PtzSerialValue struct {
	PtzSerial PtzSerial `json:"PtzSerial"`
}

// GetPtzSerial gets PTZ serial port configuration
func (p *PTZAPI) GetPtzSerial(ctx context.Context, channel int) (*PtzSerial, error) {
	p.client.logger.Debug("getting PTZ serial configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetPtzSerial",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get PTZ serial configuration: %v", err)
		return nil, fmt.Errorf("GetPtzSerial request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get PTZ serial configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get PTZ serial configuration: %v", apiErr)
		return nil, apiErr
	}

	var value PtzSerialValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse PTZ serial configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved PTZ serial configuration: protocol=%s baudRate=%d",
		value.PtzSerial.CtrlProtocol, value.PtzSerial.BaudRate)
	return &value.PtzSerial, nil
}

// SetPtzSerial sets PTZ serial port configuration
func (p *PTZAPI) SetPtzSerial(ctx context.Context, serial PtzSerial) error {
	p.client.logger.Info("setting PTZ serial configuration: channel=%d protocol=%s baudRate=%d",
		serial.Channel, serial.CtrlProtocol, serial.BaudRate)

	req := []Request{{
		Cmd: "SetPtzSerial",
		Param: map[string]interface{}{
			"PtzSerial": serial,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set PTZ serial configuration: %v", err)
		return fmt.Errorf("SetPtzSerial request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set PTZ serial configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set PTZ serial configuration: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set PTZ serial configuration")
	return nil
}

// AutoFocus represents auto focus configuration
type AutoFocus struct {
	Channel int `json:"channel"` // Channel number
	Disable int `json:"disable"` // 0=enable autofocus, 1=forbid autofocus
}

// AutoFocusValue wraps AutoFocus for API response
type AutoFocusValue struct {
	AutoFocus AutoFocus `json:"AutoFocus"`
}

// GetAutoFocus gets auto focus configuration
func (p *PTZAPI) GetAutoFocus(ctx context.Context, channel int) (*AutoFocus, error) {
	p.client.logger.Debug("getting auto focus configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetAutoFocus",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to get auto focus configuration: %v", err)
		return nil, fmt.Errorf("GetAutoFocus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to get auto focus configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to get auto focus configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AutoFocusValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		p.client.logger.Error("failed to parse auto focus configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	p.client.logger.Info("successfully retrieved auto focus configuration: disable=%d", value.AutoFocus.Disable)
	return &value.AutoFocus, nil
}

// SetAutoFocus sets auto focus configuration
func (p *PTZAPI) SetAutoFocus(ctx context.Context, autoFocus AutoFocus) error {
	p.client.logger.Info("setting auto focus configuration: channel=%d disable=%d",
		autoFocus.Channel, autoFocus.Disable)

	req := []Request{{
		Cmd:    "SetAutoFocus",
		Action: 0,
		Param: map[string]interface{}{
			"AutoFocus": autoFocus,
		},
	}}

	var resp []Response
	if err := p.client.do(ctx, req, &resp); err != nil {
		p.client.logger.Error("failed to set auto focus configuration: %v", err)
		return fmt.Errorf("SetAutoFocus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		p.client.logger.Error("failed to set auto focus configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		p.client.logger.Error("failed to set auto focus configuration: %v", apiErr)
		return apiErr
	}

	p.client.logger.Info("successfully set auto focus configuration")
	return nil
}
