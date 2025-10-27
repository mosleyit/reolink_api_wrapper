package reolink

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// RecordingAPI provides methods for recording configuration and playback
type RecordingAPI struct {
	client *Client
}

// Rec represents recording configuration
type Rec struct {
	Channel   int         `json:"channel"`
	Overwrite int         `json:"overwrite"`         // 0=stop when full, 1=overwrite oldest
	PostRec   string      `json:"postRec"`           // Post-recording duration: "30 Seconds", "1 Minute", etc.
	PreRec    int         `json:"preRec"`            // Pre-recording: 0=off, 1=on
	SaveDay   int         `json:"saveDay,omitempty"` // Days to keep recordings (v2.0 only)
	Schedule  RecSchedule `json:"schedule"`
}

// RecSchedule represents recording schedule configuration
type RecSchedule struct {
	Enable  int         `json:"enable"`            // 0=disabled, 1=enabled
	Channel int         `json:"channel,omitempty"` // Channel number (v2.0 only)
	Table   interface{} `json:"table"`             // string for v1, RecScheduleTable for v2.0
}

// RecScheduleTable represents v2.0 schedule with multiple alarm types
type RecScheduleTable struct {
	MD        string `json:"MD,omitempty"`         // Motion detection schedule (168 chars)
	TIMING    string `json:"TIMING,omitempty"`     // Timing schedule (168 chars)
	AIPeople  string `json:"AI_PEOPLE,omitempty"`  // AI people detection schedule (168 chars)
	AIVehicle string `json:"AI_VEHICLE,omitempty"` // AI vehicle detection schedule (168 chars)
	AIDogCat  string `json:"AI_DOG_CAT,omitempty"` // AI dog/cat detection schedule (168 chars)
}

// RecValue represents the response value for GetRec/GetRecV20
type RecValue struct {
	Rec Rec `json:"Rec"`
}

// SearchParam represents parameters for searching recordings
type SearchParam struct {
	Search SearchCriteria `json:"Search"`
}

// SearchCriteria represents search criteria for recordings
type SearchCriteria struct {
	Channel    int       `json:"channel"`
	OnlyStatus int       `json:"onlyStatus"` // 0=main stream, 1=sub stream
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	StreamType string    `json:"streamType,omitempty"` // "main" or "sub"
}

// SearchResult represents a search result
type SearchResult struct {
	Channel   int       `json:"channel"`
	FileName  string    `json:"fileName"`
	FileSize  int64     `json:"fileSize"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Type      string    `json:"type"` // "MD", "TIMING", "AI_PEOPLE", etc.
}

// SearchValue represents the response value for Search
type SearchValue struct {
	SearchResult []SearchResult `json:"SearchResult"`
}

// GetRec gets recording configuration (v1.0)
func (r *RecordingAPI) GetRec(ctx context.Context, channel int) (*Rec, error) {
	r.client.logger.Debug("getting recording configuration: channel=%d", channel)

	req := []Request{{
		Cmd:    "GetRec",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to get recording configuration: %v", err)
		return nil, fmt.Errorf("GetRec request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetRec")
		r.client.logger.Error("failed to get recording configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		r.client.logger.Error("failed to get recording configuration: %v", err)
		return nil, err
	}

	var value RecValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		r.client.logger.Error("failed to parse recording configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetRec response: %w", err)
	}

	return &value.Rec, nil
}

// SetRec sets recording configuration (v1.0)
func (r *RecordingAPI) SetRec(ctx context.Context, rec Rec) error {
	r.client.logger.Info("setting recording configuration: channel=%d", rec.Channel)

	req := []Request{{
		Cmd: "SetRec",
		Param: map[string]interface{}{
			"Rec": rec,
		},
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to set recording configuration: %v", err)
		return fmt.Errorf("SetRec request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetRec")
		r.client.logger.Error("failed to set recording configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		r.client.logger.Error("failed to set recording configuration: %v", apiErr)
		return apiErr
	}

	r.client.logger.Info("successfully set recording configuration")
	return nil
}

// GetRecV20 gets recording configuration (v2.0 with enhanced features)
func (r *RecordingAPI) GetRecV20(ctx context.Context, channel int) (*Rec, error) {
	r.client.logger.Debug("getting recording configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd:    "GetRecV20",
		Action: 0,
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to get recording configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetRecV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetRecV20")
		r.client.logger.Error("failed to get recording configuration (v2.0): %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		r.client.logger.Error("failed to get recording configuration (v2.0): %v", err)
		return nil, err
	}

	var value RecValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		r.client.logger.Error("failed to parse recording configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse GetRecV20 response: %w", err)
	}

	return &value.Rec, nil
}

// SetRecV20 sets recording configuration (v2.0 with enhanced features)
func (r *RecordingAPI) SetRecV20(ctx context.Context, rec Rec) error {
	r.client.logger.Info("setting recording configuration (v2.0): channel=%d", rec.Channel)

	req := []Request{{
		Cmd: "SetRecV20",
		Param: map[string]interface{}{
			"Rec": rec,
		},
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to set recording configuration (v2.0): %v", err)
		return fmt.Errorf("SetRecV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetRecV20")
		r.client.logger.Error("failed to set recording configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		r.client.logger.Error("failed to set recording configuration (v2.0): %v", apiErr)
		return apiErr
	}

	r.client.logger.Info("successfully set recording configuration (v2.0)")
	return nil
}

// Search searches for recordings by time range
func (r *RecordingAPI) Search(ctx context.Context, channel int, startTime, endTime time.Time, streamType string) ([]SearchResult, error) {
	r.client.logger.Info("searching recordings: channel=%d start=%s end=%s stream=%s",
		channel, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), streamType)

	onlyStatus := 0
	if streamType == "sub" {
		onlyStatus = 1
	}

	req := []Request{{
		Cmd:    "Search",
		Action: 0,
		Param: SearchParam{
			Search: SearchCriteria{
				Channel:    channel,
				OnlyStatus: onlyStatus,
				StartTime:  startTime,
				EndTime:    endTime,
				StreamType: streamType,
			},
		},
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to search recordings: %v", err)
		return nil, fmt.Errorf("Search request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from Search")
		r.client.logger.Error("failed to search recordings: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		r.client.logger.Error("failed to search recordings: %v", err)
		return nil, err
	}

	var value SearchValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		r.client.logger.Error("failed to parse search recordings response: %v", err)
		return nil, fmt.Errorf("failed to parse Search response: %w", err)
	}

	r.client.logger.Info("successfully searched recordings: found=%d", len(value.SearchResult))
	return value.SearchResult, nil
}

// Download downloads a recording file
// Returns the URL to download the file via GET request
func (r *RecordingAPI) Download(source, output string) string {
	r.client.logger.Info("generating download URL: source=%s", source)

	url := fmt.Sprintf("%s?cmd=Download&source=%s&output=%s&token=%s",
		r.client.baseURL, source, output, r.client.token)

	r.client.logger.Debug("generated download URL")
	return url
}

// Playback returns the URL for streaming playback of a recording
func (r *RecordingAPI) Playback(source, output string) string {
	r.client.logger.Info("generating playback URL: source=%s", source)

	url := fmt.Sprintf("%s?cmd=Playback&source=%s&output=%s&token=%s",
		r.client.baseURL, source, output, r.client.token)

	r.client.logger.Debug("generated playback URL")
	return url
}

// NvrDownload downloads a recording from NVR
// This is a placeholder - actual implementation depends on NVR-specific parameters
func (r *RecordingAPI) NvrDownload(ctx context.Context, params map[string]interface{}) error {
	r.client.logger.Info("downloading recording from NVR")

	req := []Request{{
		Cmd:   "NvrDownload",
		Param: params,
	}}

	var resp []Response
	if err := r.client.do(ctx, req, &resp); err != nil {
		r.client.logger.Error("failed to download recording from NVR: %v", err)
		return fmt.Errorf("NvrDownload request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from NvrDownload")
		r.client.logger.Error("failed to download recording from NVR: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		r.client.logger.Error("failed to download recording from NVR: %v", apiErr)
		return apiErr
	}

	r.client.logger.Debug("successfully initiated NVR download")
	return nil
}
