package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// VideoAPI provides access to video input and encoding API endpoints
type VideoAPI struct {
	client *Client
}

// Osd represents On-Screen Display configuration
type Osd struct {
	Channel    int        `json:"channel"`    // Channel number
	BgColor    int        `json:"bgcolor"`    // Background color (0=transparent, 1=black)
	OsdChannel OsdChannel `json:"osdChannel"` // Camera name display settings
	OsdTime    OsdTime    `json:"osdTime"`    // Timestamp display settings
	Watermark  int        `json:"watermark"`  // Watermark enable (0=off, 1=on)
}

// OsdChannel represents camera name display settings
type OsdChannel struct {
	Enable int    `json:"enable"` // 0=disabled, 1=enabled
	Name   string `json:"name"`   // Camera name
	Pos    string `json:"pos"`    // Position: "Upper Left", "Upper Right", "Lower Left", "Lower Right", "Top Center", "Bottom Center"
}

// OsdTime represents timestamp display settings
type OsdTime struct {
	Enable int    `json:"enable"` // 0=disabled, 1=enabled
	Pos    string `json:"pos"`    // Position: "Upper Left", "Upper Right", "Lower Left", "Lower Right", "Top Center", "Bottom Center"
}

// OsdValue represents the response value for GetOsd
type OsdValue struct {
	Osd Osd `json:"Osd"`
}

// Image represents image quality settings
type Image struct {
	Channel    int `json:"channel"`    // Channel number
	Bright     int `json:"bright"`     // Brightness (0-255, default 128)
	Contrast   int `json:"contrast"`   // Contrast (0-255, default 128)
	Saturation int `json:"saturation"` // Saturation (0-255, default 128)
	Hue        int `json:"hue"`        // Hue (0-255, default 128)
	Sharpen    int `json:"sharpen"`    // Sharpness (0-255, default 128)
}

// ImageValue represents the response value for GetImage
type ImageValue struct {
	Image Image `json:"Image"`
}

// IspGain represents gain range settings
type IspGain struct {
	Min int `json:"min"` // Minimum gain value
	Max int `json:"max"` // Maximum gain value
}

// Isp represents Image Signal Processor settings
type Isp struct {
	Channel     int     `json:"channel"`     // Channel number
	AntiFlicker string  `json:"antiFlicker"` // "Outdoor", "50Hz", "60Hz"
	Exposure    string  `json:"exposure"`    // "Auto", "Manual"
	Gain        IspGain `json:"gain"`        // Gain range (min/max)
	DayNight    string  `json:"dayNight"`    // "Auto", "Color", "Black&White"
	BackLight   string  `json:"backLight"`   // "Off", "BackLightControl", "DynamicRangeControl", "Off"
	Blc         int     `json:"blc"`         // Backlight compensation (0-255)
	Drc         int     `json:"drc"`         // Dynamic range control (0-255)
	Rotation    int     `json:"rotation"`    // Rotation angle (0, 90, 180, 270)
	Mirroring   int     `json:"mirroring"`   // Mirror (0=off, 1=on)
	Nr3d        int     `json:"nr3d"`        // 3D noise reduction (0-100)
}

// IspValue represents the response value for GetIsp
type IspValue struct {
	Isp Isp `json:"Isp"`
}

// Mask represents privacy mask configuration
type Mask struct {
	Channel int        `json:"channel"` // Channel number
	Enable  int        `json:"enable"`  // 0=disabled, 1=enabled
	Area    []MaskArea `json:"area"`    // Privacy mask areas (up to 4)
}

// MaskArea represents a single privacy mask area
type MaskArea struct {
	Screen MaskScreen `json:"screen"` // Screen dimensions
	X      int        `json:"x"`      // X coordinate
	Y      int        `json:"y"`      // Y coordinate
	Width  int        `json:"width"`  // Width
	Height int        `json:"height"` // Height
}

// MaskScreen represents screen dimensions for mask area
type MaskScreen struct {
	Height int `json:"height"` // Screen height
	Width  int `json:"width"`  // Screen width
}

// MaskValue represents the response value for GetMask
type MaskValue struct {
	Mask Mask `json:"Mask"`
}

// GetOsd gets On-Screen Display configuration
func (v *VideoAPI) GetOsd(ctx context.Context, channel int) (*Osd, error) {
	v.client.logger.Debug("getting OSD configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetOsd",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get OSD configuration: %v", err)
		return nil, fmt.Errorf("GetOsd request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetOsd")
		v.client.logger.Error("failed to get OSD configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get OSD configuration: %v", err)
		return nil, err
	}

	var value OsdValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse OSD configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetOsd response: %w", err)
	}

	return &value.Osd, nil
}

// SetOsd sets On-Screen Display configuration
func (v *VideoAPI) SetOsd(ctx context.Context, osd Osd) error {
	v.client.logger.Info("setting OSD configuration: channel=%d", osd.Channel)

	req := []Request{{
		Cmd: "SetOsd",
		Param: map[string]interface{}{
			"Osd": osd,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set OSD configuration: %v", err)
		return fmt.Errorf("SetOsd request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetOsd")
		v.client.logger.Error("failed to set OSD configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set OSD configuration: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set OSD configuration")
	return nil
}

// GetImage gets image quality settings
func (v *VideoAPI) GetImage(ctx context.Context, channel int) (*Image, error) {
	v.client.logger.Debug("getting image settings: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetImage",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get image settings: %v", err)
		return nil, fmt.Errorf("GetImage request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetImage")
		v.client.logger.Error("failed to get image settings: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get image settings: %v", err)
		return nil, err
	}

	var value ImageValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse image settings response: %v", err)
		return nil, fmt.Errorf("failed to parse GetImage response: %w", err)
	}

	return &value.Image, nil
}

// SetImage sets image quality settings
func (v *VideoAPI) SetImage(ctx context.Context, image Image) error {
	v.client.logger.Info("setting image settings: channel=%d", image.Channel)

	req := []Request{{
		Cmd: "SetImage",
		Param: map[string]interface{}{
			"Image": image,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set image settings: %v", err)
		return fmt.Errorf("SetImage request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetImage")
		v.client.logger.Error("failed to set image settings: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set image settings: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set image settings")
	return nil
}

// GetIsp gets Image Signal Processor settings
func (v *VideoAPI) GetIsp(ctx context.Context, channel int) (*Isp, error) {
	v.client.logger.Debug("getting ISP settings: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetIsp",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get ISP settings: %v", err)
		return nil, fmt.Errorf("GetIsp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetIsp")
		v.client.logger.Error("failed to get ISP settings: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get ISP settings: %v", err)
		return nil, err
	}

	var value IspValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse ISP settings response: %v", err)
		return nil, fmt.Errorf("failed to parse GetIsp response: %w", err)
	}

	return &value.Isp, nil
}

// SetIsp sets Image Signal Processor settings
func (v *VideoAPI) SetIsp(ctx context.Context, isp Isp) error {
	v.client.logger.Info("setting ISP settings: channel=%d", isp.Channel)

	req := []Request{{
		Cmd: "SetIsp",
		Param: map[string]interface{}{
			"Isp": isp,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set ISP settings: %v", err)
		return fmt.Errorf("SetIsp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetIsp")
		v.client.logger.Error("failed to set ISP settings: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set ISP settings: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set ISP settings")
	return nil
}

// GetMask gets privacy mask configuration
func (v *VideoAPI) GetMask(ctx context.Context, channel int) (*Mask, error) {
	v.client.logger.Debug("getting privacy mask configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetMask",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get privacy mask configuration: %v", err)
		return nil, fmt.Errorf("GetMask request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetMask")
		v.client.logger.Error("failed to get privacy mask configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get privacy mask configuration: %v", err)
		return nil, err
	}

	var value MaskValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse privacy mask configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetMask response: %w", err)
	}

	return &value.Mask, nil
}

// SetMask sets privacy mask configuration
func (v *VideoAPI) SetMask(ctx context.Context, mask Mask) error {
	v.client.logger.Info("setting privacy mask configuration: channel=%d", mask.Channel)

	req := []Request{{
		Cmd: "SetMask",
		Param: map[string]interface{}{
			"Mask": mask,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set privacy mask configuration: %v", err)
		return fmt.Errorf("SetMask request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetMask")
		v.client.logger.Error("failed to set privacy mask configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set privacy mask configuration: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set privacy mask configuration")
	return nil
}

// Crop represents video crop/zoom configuration
type Crop struct {
	Channel      int `json:"channel"`      // Channel number
	ScreenWidth  int `json:"screenWidth"`  // Screen width
	ScreenHeight int `json:"screenHeight"` // Screen height
	CropWidth    int `json:"cropWidth"`    // Width of crop area
	CropHeight   int `json:"cropHeight"`   // Height of crop area
	TopLeftX     int `json:"topLeftX"`     // Distance from left boundary
	TopLeftY     int `json:"topLeftY"`     // Distance from top boundary
}

// CropValue represents the response value for GetCrop
type CropValue struct {
	Crop Crop `json:"Crop"`
}

// Stitch represents video stitching configuration for multi-lens cameras
type Stitch struct {
	Distance    float64 `json:"distance"`    // Distance between images (2.0-20.0)
	StitchXMove int     `json:"stitchXMove"` // Adjust pixels horizontally (-100 to 100)
	StitchYMove int     `json:"stitchYMove"` // Adjust pixels vertically (-100 to 100)
}

// StitchValue represents the response value for GetStitch
type StitchValue struct {
	Stitch Stitch `json:"stitch"`
}

// GetCrop gets video crop/zoom configuration
func (v *VideoAPI) GetCrop(ctx context.Context, channel int) (*Crop, error) {
	v.client.logger.Debug("getting crop configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetCrop",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get crop configuration: %v", err)
		return nil, fmt.Errorf("GetCrop request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetCrop")
		v.client.logger.Error("failed to get crop configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get crop configuration: %v", err)
		return nil, err
	}

	var value CropValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse crop configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetCrop response: %w", err)
	}

	return &value.Crop, nil
}

// SetCrop sets video crop/zoom configuration
func (v *VideoAPI) SetCrop(ctx context.Context, crop Crop) error {
	v.client.logger.Info("setting crop configuration: channel=%d", crop.Channel)

	req := []Request{{
		Cmd: "SetCrop",
		Param: map[string]interface{}{
			"Crop": crop,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set crop configuration: %v", err)
		return fmt.Errorf("SetCrop request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetCrop")
		v.client.logger.Error("failed to set crop configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set crop configuration: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set crop configuration")
	return nil
}

// GetStitch gets video stitching configuration (for multi-lens cameras)
func (v *VideoAPI) GetStitch(ctx context.Context) (*Stitch, error) {
	v.client.logger.Debug("getting stitch configuration")

	req := []Request{{
		Cmd:    "GetStitch",
		Action: 1,
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to get stitch configuration: %v", err)
		return nil, fmt.Errorf("GetStitch request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetStitch")
		v.client.logger.Error("failed to get stitch configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		v.client.logger.Error("failed to get stitch configuration: %v", err)
		return nil, err
	}

	var value StitchValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		v.client.logger.Error("failed to parse stitch configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetStitch response: %w", err)
	}

	return &value.Stitch, nil
}

// SetStitch sets video stitching configuration (for multi-lens cameras)
func (v *VideoAPI) SetStitch(ctx context.Context, stitch Stitch) error {
	v.client.logger.Info("setting stitch configuration")

	req := []Request{{
		Cmd: "SetStitch",
		Param: map[string]interface{}{
			"stitch": stitch,
		},
	}}

	var resp []Response
	if err := v.client.do(ctx, req, &resp); err != nil {
		v.client.logger.Error("failed to set stitch configuration: %v", err)
		return fmt.Errorf("SetStitch request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetStitch")
		v.client.logger.Error("failed to set stitch configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		v.client.logger.Error("failed to set stitch configuration: %v", apiErr)
		return apiErr
	}

	v.client.logger.Info("successfully set stitch configuration")
	return nil
}
