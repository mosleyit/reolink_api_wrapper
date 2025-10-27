package reolink

import (
	"fmt"
)

// Error codes from the Reolink API specification
const (
	// Success
	ErrCodeSuccess = 0

	// General Errors (-1 to -49)
	ErrCodeMissingParameters      = -1
	ErrCodeUsedUpMemory           = -2
	ErrCodeCheckError             = -3
	ErrCodeParametersError        = -4
	ErrCodeMaxSessionNumber       = -5
	ErrCodeLoginRequired          = -6
	ErrCodeLoginError             = -7
	ErrCodeOperationTimeout       = -8
	ErrCodeNotSupported           = -9
	ErrCodeProtocolError          = -10
	ErrCodeFailedReadOperation    = -11
	ErrCodeFailedGetConfiguration = -12
	ErrCodeFailedSetConfiguration = -13
	ErrCodeFailedApplyMemory      = -14
	ErrCodeFailedCreateSocket     = -15
	ErrCodeFailedSendData         = -16
	ErrCodeFailedReceiveData      = -17
	ErrCodeFailedOpenFile         = -18
	ErrCodeFailedReadFile         = -19
	ErrCodeFailedWriteFile        = -20
	ErrCodeTokenError             = -21
	ErrCodeStringLengthExceeded   = -22
	ErrCodeMissingParametersAlt   = -23
	ErrCodeCommandError           = -24
	ErrCodeInternalError          = -25
	ErrCodeAbilityError           = -26
	ErrCodeInvalidUser            = -27
	ErrCodeUserAlreadyExists      = -28
	ErrCodeMaxUsersReached        = -29
	ErrCodeVersionIdentical       = -30
	ErrCodeUpgradeBusy            = -31
	ErrCodeIPConflict             = -32
	ErrCodeCloudBindEmailFirst    = -34
	ErrCodeCloudUnbindCamera      = -35
	ErrCodeCloudInfoTimeout       = -36
	ErrCodeCloudPasswordError     = -37
	ErrCodeCloudUIDError          = -38
	ErrCodeCloudUserNotExist      = -39
	ErrCodeCloudUnbindFailed      = -40
	ErrCodeCloudNotSupported      = -41
	ErrCodeCloudServerFailed      = -42
	ErrCodeCloudBindFailed        = -43
	ErrCodeCloudUnknownError      = -44
	ErrCodeCloudNeedVerifyCode    = -45
	ErrCodeDigestAuthFailed       = -46
	ErrCodeDigestNonceExpires     = -47
	ErrCodeSnapFailed             = -48
	ErrCodeChannelInvalid         = -49
	ErrCodeDeviceOffline          = -99
	ErrCodeTestFailed             = -100

	// Upgrade Errors (-101 to -105)
	ErrCodeUpgradeCheckFailed    = -101
	ErrCodeUpgradeDownloadFailed = -102
	ErrCodeUpgradeStatusFailed   = -103
	ErrCodeFrequentLogins        = -105

	// Video Recording Errors (-220 to -222)
	ErrCodeVideoDownloadError = -220
	ErrCodeVideoBusy          = -221
	ErrCodeVideoNotExist      = -222

	// Authentication Errors (-301, -310)
	ErrCodeDigestNonceError = -301
	ErrCodeAESDecryptFailed = -310

	// FTP Errors (-451 to -454)
	ErrCodeFTPLoginFailed     = -451
	ErrCodeFTPCreateDirFailed = -452
	ErrCodeFTPUploadFailed    = -453
	ErrCodeFTPConnectFailed   = -454

	// Email Errors (-480 to -485)
	ErrCodeEmailUndefined     = -480
	ErrCodeEmailConnectFailed = -481
	ErrCodeEmailAuthFailed    = -482
	ErrCodeEmailNetworkError  = -483
	ErrCodeEmailServerError   = -484
	ErrCodeEmailMemoryError   = -485

	// Login Errors (-500 to -507)
	ErrCodeIPLimitReached      = -500
	ErrCodeUserLocked          = -501
	ErrCodeUserNotOnline       = -502
	ErrCodeInvalidUsername     = -503
	ErrCodeInvalidPassword     = -504
	ErrCodeUserAlreadyLoggedIn = -505
	ErrCodeAccountLocked       = -506
	ErrCodeAccountNotActivated = -507
)

// APIError represents an error returned by the Reolink API
type APIError struct {
	Code    int    // Response code from API
	RspCode int    // Detailed error code (from error.rspCode)
	Detail  string // Error detail message
	Cmd     string // Command that caused the error
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("reolink api error: cmd=%s code=%d rspCode=%d detail=%s",
			e.Cmd, e.Code, e.RspCode, e.Detail)
	}
	return fmt.Sprintf("reolink api error: cmd=%s code=%d rspCode=%d (%s)",
		e.Cmd, e.Code, e.RspCode, errorCodeToString(e.RspCode))
}

// Is implements error comparison for errors.Is
func (e *APIError) Is(target error) bool {
	t, ok := target.(*APIError)
	if !ok {
		return false
	}
	return e.RspCode == t.RspCode
}

// errorCodeToString returns a human-readable description of an error code
func errorCodeToString(code int) string {
	switch code {
	case ErrCodeSuccess:
		return "success"
	case ErrCodeMissingParameters, ErrCodeMissingParametersAlt:
		return "missing parameters"
	case ErrCodeUsedUpMemory:
		return "used up memory"
	case ErrCodeCheckError:
		return "check error"
	case ErrCodeParametersError:
		return "parameters error"
	case ErrCodeMaxSessionNumber:
		return "reached the max session number"
	case ErrCodeLoginRequired:
		return "login required"
	case ErrCodeLoginError:
		return "login error"
	case ErrCodeOperationTimeout:
		return "operation timeout"
	case ErrCodeNotSupported:
		return "not supported"
	case ErrCodeProtocolError:
		return "protocol error"
	case ErrCodeFailedReadOperation:
		return "failed to read operation"
	case ErrCodeFailedGetConfiguration:
		return "failed to get configuration"
	case ErrCodeFailedSetConfiguration:
		return "failed to set configuration"
	case ErrCodeFailedApplyMemory:
		return "failed to apply for memory"
	case ErrCodeFailedCreateSocket:
		return "failed to create socket"
	case ErrCodeFailedSendData:
		return "failed to send data"
	case ErrCodeFailedReceiveData:
		return "failed to receive data"
	case ErrCodeFailedOpenFile:
		return "failed to open file"
	case ErrCodeFailedReadFile:
		return "failed to read file"
	case ErrCodeFailedWriteFile:
		return "failed to write file"
	case ErrCodeTokenError:
		return "token error"
	case ErrCodeStringLengthExceeded:
		return "the length of the string exceeds the limitation"
	case ErrCodeCommandError:
		return "command error"
	case ErrCodeInternalError:
		return "internal error"
	case ErrCodeAbilityError:
		return "ability error"
	case ErrCodeInvalidUser:
		return "invalid user"
	case ErrCodeUserAlreadyExists:
		return "user already exists"
	case ErrCodeMaxUsersReached:
		return "reached the maximum number of users"
	case ErrCodeVersionIdentical:
		return "the version is identical to the current one"
	case ErrCodeUpgradeBusy:
		return "ensure only one user can upgrade (busy)"
	case ErrCodeIPConflict:
		return "modify IP conflicted with used IP"
	case ErrCodeCloudBindEmailFirst:
		return "cloud login need bind email first"
	case ErrCodeCloudUnbindCamera:
		return "cloud login unbind camera"
	case ErrCodeCloudInfoTimeout:
		return "cloud login get information out of time"
	case ErrCodeCloudPasswordError:
		return "cloud login password error"
	case ErrCodeCloudUIDError:
		return "cloud bind camera uid error"
	case ErrCodeCloudUserNotExist:
		return "cloud login user doesn't exist"
	case ErrCodeCloudUnbindFailed:
		return "cloud unbind camera failed"
	case ErrCodeCloudNotSupported:
		return "the device doesn't support cloud"
	case ErrCodeCloudServerFailed:
		return "cloud login server failed"
	case ErrCodeCloudBindFailed:
		return "cloud bind camera failed"
	case ErrCodeCloudUnknownError:
		return "cloud unknown error"
	case ErrCodeCloudNeedVerifyCode:
		return "cloud bind camera need verify code"
	case ErrCodeDigestAuthFailed:
		return "digest authentication failed"
	case ErrCodeDigestNonceExpires:
		return "digest authentication nonce expires"
	case ErrCodeSnapFailed:
		return "snap a picture failed"
	case ErrCodeChannelInvalid:
		return "channel is invalid"
	case ErrCodeDeviceOffline:
		return "device offline"
	case ErrCodeTestFailed:
		return "test email/ftp/wifi failed"
	case ErrCodeUpgradeCheckFailed:
		return "upgrade checking firmware failed"
	case ErrCodeUpgradeDownloadFailed:
		return "upgrade download online failed"
	case ErrCodeUpgradeStatusFailed:
		return "upgrade get upgrade status failed"
	case ErrCodeFrequentLogins:
		return "frequent logins, please try again later"
	case ErrCodeVideoDownloadError:
		return "error downloading video file"
	case ErrCodeVideoBusy:
		return "busy video recording task"
	case ErrCodeVideoNotExist:
		return "the video file does not exist"
	case ErrCodeDigestNonceError:
		return "digest authentication nonce error"
	case ErrCodeAESDecryptFailed:
		return "AES decryption failure"
	case ErrCodeFTPLoginFailed:
		return "FTP test login failed"
	case ErrCodeFTPCreateDirFailed:
		return "create FTP directory failed"
	case ErrCodeFTPUploadFailed:
		return "upload FTP file failed"
	case ErrCodeFTPConnectFailed:
		return "cannot connect FTP server"
	case ErrCodeEmailUndefined:
		return "email undefined error"
	case ErrCodeEmailConnectFailed:
		return "cannot connect email server"
	case ErrCodeEmailAuthFailed:
		return "email auth user failed"
	case ErrCodeEmailNetworkError:
		return "email network error"
	case ErrCodeEmailServerError:
		return "something wrong with email server"
	case ErrCodeEmailMemoryError:
		return "something wrong with memory"
	case ErrCodeIPLimitReached:
		return "the number of IP addresses reaches the upper limit"
	case ErrCodeUserLocked:
		return "user locked"
	case ErrCodeUserNotOnline:
		return "user not online"
	case ErrCodeInvalidUsername:
		return "invalid username"
	case ErrCodeInvalidPassword:
		return "invalid password"
	case ErrCodeUserAlreadyLoggedIn:
		return "user already logged in"
	case ErrCodeAccountLocked:
		return "account locked"
	case ErrCodeAccountNotActivated:
		return "account not activated"
	default:
		return fmt.Sprintf("unknown error code: %d", code)
	}
}

// NewAPIError creates a new APIError
func NewAPIError(cmd string, code, rspCode int, detail string) *APIError {
	return &APIError{
		Cmd:     cmd,
		Code:    code,
		RspCode: rspCode,
		Detail:  detail,
	}
}
