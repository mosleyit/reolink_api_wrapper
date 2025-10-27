package reolink

import (
	"errors"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	err := NewAPIError("GetDevInfo", 0, ErrCodeLoginRequired, "please login first")

	expected := "reolink api error: cmd=GetDevInfo code=0 rspCode=-6 detail=please login first"
	if err.Error() != expected {
		t.Errorf("expected error message %q, got %q", expected, err.Error())
	}
}

func TestAPIError_ErrorWithoutDetail(t *testing.T) {
	err := NewAPIError("GetDevInfo", 0, ErrCodeLoginRequired, "")

	if err.Error() == "" {
		t.Error("expected non-empty error message")
	}

	// Should include the error code description
	if err.Error() != "reolink api error: cmd=GetDevInfo code=0 rspCode=-6 (login required)" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}

func TestAPIError_Is(t *testing.T) {
	err1 := NewAPIError("GetDevInfo", 0, ErrCodeLoginRequired, "please login first")
	err2 := NewAPIError("GetTime", 0, ErrCodeLoginRequired, "login required")
	err3 := NewAPIError("GetDevInfo", 0, ErrCodeParametersError, "param error")

	if !errors.Is(err1, err2) {
		t.Error("expected errors with same RspCode to match")
	}

	if errors.Is(err1, err3) {
		t.Error("expected errors with different RspCode to not match")
	}
}

func TestErrorCodeToString(t *testing.T) {
	tests := []struct {
		code     int
		expected string
	}{
		{ErrCodeSuccess, "success"},
		{ErrCodeLoginRequired, "login required"},
		{ErrCodeParametersError, "parameters error"},
		{ErrCodeNotSupported, "not supported"},
		{ErrCodeTokenError, "token error"},
		{ErrCodeInvalidUser, "invalid user"},
		{ErrCodeDeviceOffline, "device offline"},
		{ErrCodeUpgradeCheckFailed, "upgrade checking firmware failed"},
		{ErrCodeVideoNotExist, "the video file does not exist"},
		{ErrCodeFTPConnectFailed, "cannot connect FTP server"},
		{ErrCodeEmailAuthFailed, "email auth user failed"},
		{ErrCodeInvalidPassword, "invalid password"},
		{-9999, "unknown error code: -9999"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := errorCodeToString(tt.code)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestResponse_ToAPIError(t *testing.T) {
	// Test response with error detail
	resp := Response{
		Cmd:  "GetDevInfo",
		Code: 0,
		Error: &ErrorDetail{
			RspCode: ErrCodeLoginRequired,
			Detail:  "please login first",
		},
	}

	apiErr := resp.ToAPIError()
	if apiErr == nil {
		t.Fatal("expected APIError, got nil")
	}

	if apiErr.RspCode != ErrCodeLoginRequired {
		t.Errorf("expected RspCode %d, got %d", ErrCodeLoginRequired, apiErr.RspCode)
	}

	if apiErr.Detail != "please login first" {
		t.Errorf("expected detail 'please login first', got %q", apiErr.Detail)
	}

	// Test response with non-zero code but no error detail
	resp2 := Response{
		Cmd:  "GetDevInfo",
		Code: -1,
	}

	apiErr2 := resp2.ToAPIError()
	if apiErr2 == nil {
		t.Fatal("expected APIError, got nil")
	}

	if apiErr2.Code != -1 {
		t.Errorf("expected Code -1, got %d", apiErr2.Code)
	}

	// Test successful response
	resp3 := Response{
		Cmd:  "GetDevInfo",
		Code: 0,
	}

	apiErr3 := resp3.ToAPIError()
	if apiErr3 != nil {
		t.Errorf("expected nil for successful response, got %v", apiErr3)
	}
}
