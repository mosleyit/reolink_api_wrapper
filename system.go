package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// SystemAPI provides access to system-related API endpoints
type SystemAPI struct {
	client *Client
}

// GetDeviceInfo retrieves device information including model, firmware version,
// hardware version, and capabilities.
//
// This is typically one of the first calls made after authentication to determine
// what features the camera supports.
//
// Example:
//
//	info, err := client.System.GetDeviceInfo(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Model: %s, Firmware: %s\n", info.Model, info.FirmVer)
//
// Returns an error if the request fails or if the camera returns an error code.
func (s *SystemAPI) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {
	s.client.logger.Debug("getting device info")

	req := []Request{{
		Cmd:    "GetDevInfo",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get device info: %v", err)
		return nil, fmt.Errorf("GetDevInfo request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get device info: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get device info: %v", apiErr)
		return nil, apiErr
	}

	var value DeviceInfoValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse device info response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	s.client.logger.Info("successfully retrieved device info: model=%s firmware=%s", value.DevInfo.Model, value.DevInfo.FirmVer)
	return &value.DevInfo, nil
}

// GetDeviceName retrieves the device name
func (s *SystemAPI) GetDeviceName(ctx context.Context) (string, error) {
	s.client.logger.Debug("getting device name")

	req := []Request{{
		Cmd:    "GetDevName",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get device name: %v", err)
		return "", fmt.Errorf("GetDevName request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get device name: %v", err)
		return "", err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get device name: %v", apiErr)
		return "", apiErr
	}

	var value DeviceNameValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse device name response: %v", err)
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return value.DevName.Name, nil
}

// SetDeviceName sets the device name
func (s *SystemAPI) SetDeviceName(ctx context.Context, name string) error {
	s.client.logger.Info("setting device name to: %s", name)

	req := []Request{{
		Cmd: "SetDevName",
		Param: DeviceNameParam{
			DevName: DeviceName{
				Name: name,
			},
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to set device name: %v", err)
		return fmt.Errorf("SetDevName request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to set device name: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to set device name: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully set device name")
	return nil
}

// GetTime retrieves the current time configuration
func (s *SystemAPI) GetTime(ctx context.Context) (*TimeConfig, error) {
	s.client.logger.Debug("getting time configuration")

	req := []Request{{
		Cmd:    "GetTime",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get time configuration: %v", err)
		return nil, fmt.Errorf("GetTime request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get time configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get time configuration: %v", apiErr)
		return nil, apiErr
	}

	var value TimeValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse time configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.Time, nil
}

// SetTime sets the time configuration
func (s *SystemAPI) SetTime(ctx context.Context, timeConfig *TimeConfig) error {
	s.client.logger.Info("setting time configuration")

	req := []Request{{
		Cmd: "SetTime",
		Param: TimeParam{
			Time: *timeConfig,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to set time configuration: %v", err)
		return fmt.Errorf("SetTime request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to set time configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to set time configuration: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully set time configuration")
	return nil
}

// GetHddInfo retrieves hard disk information
func (s *SystemAPI) GetHddInfo(ctx context.Context) ([]HddInfo, error) {
	s.client.logger.Debug("getting HDD info")

	req := []Request{{
		Cmd:    "GetHddInfo",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get HDD info: %v", err)
		return nil, fmt.Errorf("GetHddInfo request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get HDD info: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get HDD info: %v", apiErr)
		return nil, apiErr
	}

	var value HddInfoValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse HDD info response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(value.HddInfo) > 0 {
		s.client.logger.Info("successfully retrieved HDD info: count=%d", len(value.HddInfo))
	}
	return value.HddInfo, nil
}

// Format formats a storage device
func (s *SystemAPI) Format(ctx context.Context, hddID int) error {
	s.client.logger.Warn("formatting disk (destructive operation): hdd_id=%d", hddID)

	req := []Request{{
		Cmd: "Format",
		Param: FormatParam{
			Hdd: struct {
				ID int `json:"id"`
			}{
				ID: hddID,
			},
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to format disk: %v", err)
		return fmt.Errorf("Format request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to format disk: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to format disk: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully formatted disk")
	return nil
}

// Reboot reboots the device
func (s *SystemAPI) Reboot(ctx context.Context) error {
	s.client.logger.Warn("rebooting device (system restart)")

	req := []Request{{
		Cmd: "Reboot",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to reboot device: %v", err)
		return fmt.Errorf("Reboot request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to reboot device: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to reboot device: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully initiated device reboot")
	return nil
}

// Restore restores factory default settings
func (s *SystemAPI) Restore(ctx context.Context) error {
	s.client.logger.Warn("restoring factory defaults (destructive operation)")

	req := []Request{{
		Cmd: "Restore",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to restore factory defaults: %v", err)
		return fmt.Errorf("Restore request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to restore factory defaults: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to restore factory defaults: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully initiated factory restore")
	return nil
}

// GetAbility retrieves system capabilities
func (s *SystemAPI) GetAbility(ctx context.Context) (*Ability, error) {
	s.client.logger.Debug("getting system capabilities")

	req := []Request{{
		Cmd:    "GetAbility",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get system capabilities: %v", err)
		return nil, fmt.Errorf("GetAbility request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get system capabilities: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get system capabilities: %v", apiErr)
		return nil, apiErr
	}

	var value AbilityValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse system capabilities response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	s.client.logger.Info("successfully retrieved system capabilities")
	return &value.Ability, nil
}

// GetAutoMaint gets automatic maintenance configuration
func (s *SystemAPI) GetAutoMaint(ctx context.Context) (*AutoMaint, error) {
	s.client.logger.Debug("getting automatic maintenance configuration")

	req := []Request{{
		Cmd:    "GetAutoMaint",
		Action: 0, // Get value only
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get automatic maintenance configuration: %v", err)
		return nil, fmt.Errorf("GetAutoMaint request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get automatic maintenance configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get automatic maintenance configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AutoMaintValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse automatic maintenance configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.AutoMaint, nil
}

// SetAutoMaint sets automatic maintenance configuration
func (s *SystemAPI) SetAutoMaint(ctx context.Context, config AutoMaint) error {
	s.client.logger.Info("setting automatic maintenance configuration")

	req := []Request{{
		Cmd: "SetAutoMaint",
		Param: AutoMaintParam{
			AutoMaint: config,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to set automatic maintenance configuration: %v", err)
		return fmt.Errorf("SetAutoMaint request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to set automatic maintenance configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to set automatic maintenance configuration: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully set automatic maintenance configuration")
	return nil
}

// GetChannelStatus gets status of all channels (for NVR)
func (s *SystemAPI) GetChannelStatus(ctx context.Context) (*ChannelStatusValue, error) {
	s.client.logger.Debug("getting channel status")

	req := []Request{{
		Cmd: "Getchannelstatus",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get channel status: %v", err)
		return nil, fmt.Errorf("GetChannelStatus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get channel status: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get channel status: %v", apiErr)
		return nil, apiErr
	}

	var value ChannelStatusValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse channel status response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value, nil
}

// AutoUpgrade represents automatic upgrade configuration
type AutoUpgrade struct {
	Enable int `json:"enable"` // 0=disabled, 1=enabled
}

// AutoUpgradeValue wraps AutoUpgrade for API response
type AutoUpgradeValue struct {
	AutoUpgrade AutoUpgrade `json:"AutoUpgrade"`
}

// FirmwareCheck represents firmware check result
type FirmwareCheck struct {
	NewFirmware int `json:"newFirmware"` // 0=no new firmware, 1=new firmware available
}

// UpgradeStatusInfo represents upgrade status information
type UpgradeStatusInfo struct {
	Percent int `json:"Persent"` // Note: API uses "Persent" (typo in API)
	Code    int `json:"code"`    // Status code
}

// UpgradeStatusValue wraps UpgradeStatusInfo for API response
type UpgradeStatusValue struct {
	Status UpgradeStatusInfo `json:"Status"`
}

// GetAutoUpgrade gets automatic upgrade configuration
func (s *SystemAPI) GetAutoUpgrade(ctx context.Context) (*AutoUpgrade, error) {
	s.client.logger.Debug("getting automatic upgrade configuration")

	req := []Request{{
		Cmd: "GetAutoUpgrade",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get automatic upgrade configuration: %v", err)
		return nil, fmt.Errorf("GetAutoUpgrade request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get automatic upgrade configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get automatic upgrade configuration: %v", apiErr)
		return nil, apiErr
	}

	var value AutoUpgradeValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse automatic upgrade configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.AutoUpgrade, nil
}

// SetAutoUpgrade sets automatic upgrade configuration
func (s *SystemAPI) SetAutoUpgrade(ctx context.Context, enable bool) error {
	s.client.logger.Info("setting automatic upgrade: enabled=%v", enable)

	enableInt := 0
	if enable {
		enableInt = 1
	}

	req := []Request{{
		Cmd: "SetAutoUpgrade",
		Param: map[string]interface{}{
			"AutoUpgrade": map[string]interface{}{
				"enable": enableInt,
			},
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to set automatic upgrade: %v", err)
		return fmt.Errorf("SetAutoUpgrade request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to set automatic upgrade: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to set automatic upgrade: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully set automatic upgrade configuration")
	return nil
}

// CheckFirmware checks for new firmware online
func (s *SystemAPI) CheckFirmware(ctx context.Context) (*FirmwareCheck, error) {
	s.client.logger.Info("checking for firmware updates")

	req := []Request{{
		Cmd: "CheckFirmware",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to check firmware: %v", err)
		return nil, fmt.Errorf("CheckFirmware request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to check firmware: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to check firmware: %v", apiErr)
		return nil, apiErr
	}

	var value FirmwareCheck
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse firmware check response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if value.NewFirmware == 1 {
		s.client.logger.Info("firmware check complete: new firmware available")
	} else {
		s.client.logger.Info("firmware check complete: no new firmware available")
	}
	return &value, nil
}

// UpgradeOnline starts online firmware upgrade
func (s *SystemAPI) UpgradeOnline(ctx context.Context) error {
	s.client.logger.Warn("starting online firmware upgrade (system change)")

	req := []Request{{
		Cmd: "UpgradeOnline",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to start online firmware upgrade: %v", err)
		return fmt.Errorf("UpgradeOnline request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to start online firmware upgrade: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to start online firmware upgrade: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully started online firmware upgrade")
	return nil
}

// UpgradeStatus gets current firmware upgrade status
func (s *SystemAPI) UpgradeStatus(ctx context.Context) (*UpgradeStatusInfo, error) {
	s.client.logger.Debug("getting firmware upgrade status")

	req := []Request{{
		Cmd: "UpgradeStatus",
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get firmware upgrade status: %v", err)
		return nil, fmt.Errorf("UpgradeStatus request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get firmware upgrade status: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get firmware upgrade status: %v", apiErr)
		return nil, apiErr
	}

	var value UpgradeStatusValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse firmware upgrade status response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.Status, nil
}

// UpgradePrepare prepares device for firmware upgrade
// restoreCfg: whether to restore configuration (false=keep config, true=reset to defaults)
// fileName: firmware file name (e.g., "firmware.pak")
func (s *SystemAPI) UpgradePrepare(ctx context.Context, restoreCfg bool, fileName string) error {
	s.client.logger.Info("preparing firmware upgrade: file=%s restore_cfg=%v", fileName, restoreCfg)

	restoreCfgInt := 0
	if restoreCfg {
		restoreCfgInt = 1
	}

	req := []Request{{
		Cmd:    "UpgradePrepare",
		Action: 1,
		Param: map[string]interface{}{
			"restoreCfg": restoreCfgInt,
			"fileName":   fileName,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to prepare firmware upgrade: %v", err)
		return fmt.Errorf("UpgradePrepare request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to prepare firmware upgrade: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to prepare firmware upgrade: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully prepared firmware upgrade")
	return nil
}

// SysCfg represents system configuration (login lock settings)
type SysCfg struct {
	LockTime     int `json:"LockTime"`     // Login lock time in seconds (0-300)
	AllowedTimes int `json:"allowedTimes"` // Maximum login attempts (0-5)
	LoginLock    int `json:"loginLock"`    // 0=disabled, 1=enabled
}

// SysCfgValue wraps SysCfg for API response
type SysCfgValue struct {
	SysCfg SysCfg `json:"SysCfg"`
}

// GetSysCfg gets system configuration (login lock settings)
func (s *SystemAPI) GetSysCfg(ctx context.Context) (*SysCfg, error) {
	s.client.logger.Debug("getting system configuration")

	req := []Request{{
		Cmd:    "GetSysCfg",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get system configuration: %v", err)
		return nil, fmt.Errorf("GetSysCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get system configuration: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get system configuration: %v", apiErr)
		return nil, apiErr
	}

	var value SysCfgValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse system configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.SysCfg, nil
}

// SetSysCfg sets system configuration (login lock settings)
func (s *SystemAPI) SetSysCfg(ctx context.Context, cfg SysCfg) error {
	s.client.logger.Info("setting system configuration")

	req := []Request{{
		Cmd:    "SetSysCfg",
		Action: 0,
		Param: map[string]interface{}{
			"SysCfg": cfg,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to set system configuration: %v", err)
		return fmt.Errorf("SetSysCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to set system configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to set system configuration: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully set system configuration")
	return nil
}

// Upgrade uploads and installs firmware upgrade
// Note: This command can only carry up to 40K packets at a time.
// It needs to be called several times to complete the device update for larger firmware files.
// The firmware parameter should be the raw firmware file bytes (.pak file)
func (s *SystemAPI) Upgrade(ctx context.Context, firmware []byte) error {
	s.client.logger.Warn("Upgrade endpoint not yet implemented (stub)")
	// This is a complex multipart/form-data upload that requires special handling
	// For now, we return an error indicating this is not yet implemented
	// Users should use UpgradePrepare + UpgradeOnline + UpgradeStatus instead
	return fmt.Errorf("Upgrade endpoint not yet implemented - use UpgradePrepare/UpgradeOnline/UpgradeStatus for firmware upgrades")
}
