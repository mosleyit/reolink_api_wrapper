package reolink

import (
	"encoding/json"
)

// Request represents a single API request command
type Request struct {
	Cmd    string      `json:"cmd"`              // Command name
	Action int         `json:"action,omitempty"` // 0: get value only, 1: get initial, range, and value
	Param  interface{} `json:"param,omitempty"`  // Command-specific parameters
	Token  string      `json:"token,omitempty"`  // Authentication token (added by client)
}

// Response represents a single API response
type Response struct {
	Cmd     string          `json:"cmd"`               // Command name
	Code    int             `json:"code"`              // Response code (0 = success)
	Value   json.RawMessage `json:"value,omitempty"`   // Response data (present when code = 0)
	Error   *ErrorDetail    `json:"error,omitempty"`   // Error details (present when error occurs)
	Initial json.RawMessage `json:"initial,omitempty"` // Initial/default values (when action = 1)
	Range   json.RawMessage `json:"range,omitempty"`   // Valid ranges/options (when action = 1)
}

// ErrorDetail represents detailed error information in a response
type ErrorDetail struct {
	RspCode int    `json:"rspCode"` // Detailed error code
	Detail  string `json:"detail"`  // Error detail message
}

// ToAPIError converts a Response to an APIError if it contains an error
func (r *Response) ToAPIError() *APIError {
	if r.Error != nil {
		return NewAPIError(r.Cmd, r.Code, r.Error.RspCode, r.Error.Detail)
	}
	if r.Code != 0 {
		return NewAPIError(r.Cmd, r.Code, r.Code, "")
	}
	return nil
}

// LoginParam represents the parameters for the Login command
type LoginParam struct {
	User LoginUser `json:"User"`
}

// LoginUser represents user credentials for login
type LoginUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Version  string `json:"Version,omitempty"` // "0": no encryption, "1": private encryption
}

// LoginValue represents the response value from a Login command
type LoginValue struct {
	Token TokenInfo `json:"Token"`
}

// TokenInfo represents authentication token information
type TokenInfo struct {
	Name      string `json:"name"`      // Token value
	LeaseTime int    `json:"leaseTime"` // Token validity duration in seconds
}

// DeviceInfo represents device information from GetDevInfo
type DeviceInfo struct {
	B485         int    `json:"B485"`
	IOInputNum   int    `json:"IOInputNum"`
	IOOutputNum  int    `json:"IOOutputNum"`
	AudioNum     int    `json:"audioNum"`
	BuildDay     string `json:"buildDay"`
	CfgVer       string `json:"cfgVer"`
	ChannelNum   int    `json:"channelNum"`
	Detail       string `json:"detail"`
	DiskNum      int    `json:"diskNum"`
	ExactType    string `json:"exactType"`
	FirmVer      string `json:"firmVer"`
	FrameworkVer int    `json:"frameworkVer"`
	HardVer      string `json:"hardVer"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	PakSuffix    string `json:"pakSuffix"`
	Serial       string `json:"serial"`
	Type         string `json:"type"`
	Wifi         int    `json:"wifi"`
}

// DeviceInfoValue wraps DeviceInfo for API response
type DeviceInfoValue struct {
	DevInfo DeviceInfo `json:"DevInfo"`
}

// DeviceName represents device name from GetDevName
type DeviceName struct {
	Name string `json:"name"`
}

// DeviceNameValue wraps DeviceName for API response
type DeviceNameValue struct {
	DevName DeviceName `json:"DevName"`
}

// DeviceNameParam represents parameters for SetDevName
type DeviceNameParam struct {
	DevName DeviceName `json:"DevName"`
}

// TimeConfig represents time configuration
type TimeConfig struct {
	Year       int    `json:"year"`
	Mon        int    `json:"mon"`
	Day        int    `json:"day"`
	Hour       int    `json:"hour"`
	Min        int    `json:"min"`
	Sec        int    `json:"sec"`
	TimeZone   int    `json:"timeZone"`
	TimeFormat string `json:"timeFormat,omitempty"` // "DD/MM/YYYY" or "MM/DD/YYYY" or "YYYY/MM/DD"
}

// TimeValue wraps TimeConfig for API response
type TimeValue struct {
	Time TimeConfig `json:"Time"`
}

// TimeParam represents parameters for SetTime
type TimeParam struct {
	Time TimeConfig `json:"Time"`
}

// DstConfig represents daylight saving time configuration
type DstConfig struct {
	Enable    int `json:"enable"`
	Offset    int `json:"offset"`
	BeginMon  int `json:"beginMon"`
	BeginWeek int `json:"beginWeek"`
	BeginDay  int `json:"beginDay"`
	BeginHour int `json:"beginHour"`
	EndMon    int `json:"endMon"`
	EndWeek   int `json:"endWeek"`
	EndDay    int `json:"endDay"`
	EndHour   int `json:"endHour"`
}

// Channel represents a camera channel
type Channel struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Online int    `json:"online"` // 1 = online, 0 = offline
	Status string `json:"status"`
}

// Schedule represents a time schedule configuration
type Schedule struct {
	Enable int        `json:"enable"`
	Table  [][]string `json:"table"` // 7x48 array representing week schedule
}

// StreamType represents video stream type
type StreamType string

const (
	StreamMain StreamType = "main" // Main stream (high quality)
	StreamSub  StreamType = "sub"  // Sub stream (low quality)
	StreamExt  StreamType = "ext"  // External stream
)

// StreamConfig represents video stream configuration
type StreamConfig struct {
	Channel    int    `json:"channel"`
	MainStream Stream `json:"mainStream"`
	SubStream  Stream `json:"subStream"`
}

// Stream represents individual stream settings
type Stream struct {
	VType     string `json:"vType"`     // Video codec: "h264" or "h265"
	Size      string `json:"size"`      // Resolution: "2560*1440", "1920*1080", etc.
	FrameRate int    `json:"frameRate"` // Frames per second
	BitRate   int    `json:"bitRate"`   // Bitrate in kbps
	GOP       int    `json:"gop"`       // Group of pictures
	Height    int    `json:"height"`    // Video height in pixels
	Width     int    `json:"width"`     // Video width in pixels
	Profile   string `json:"profile"`   // H.264/H.265 profile (Base, Main, High)
}

// Ability represents system capabilities
type Ability struct {
	AbilityInfo map[string]interface{} `json:"Ability"`
}

// AbilityValue wraps Ability for API response
type AbilityValue struct {
	Ability Ability `json:"Ability"`
}

// User represents a user account
type User struct {
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
	Level    string `json:"level"` // "admin" or "guest"
}

// UserValue wraps user array for API response
type UserValue struct {
	User []User `json:"User"`
}

// AddUserParam represents parameters for AddUser
type AddUserParam struct {
	User User `json:"User"`
}

// ModifyUserParam represents parameters for ModifyUser
type ModifyUserParam struct {
	User User `json:"User"`
}

// DelUserParam represents parameters for DelUser
type DelUserParam struct {
	User User `json:"User"`
}

// OnlineUser represents an online user session
type OnlineUser struct {
	UserName  string `json:"userName"`
	IP        string `json:"ip"`
	LoginTime string `json:"loginTime"`
}

// OnlineUserList represents a list of online users
type OnlineUserList struct {
	Users []OnlineUser `json:"User"`
}

// OnlineValue wraps OnlineUserList for API response
type OnlineValue struct {
	Online OnlineUserList `json:"Online"`
}

// DisconnectParam represents parameters for Disconnect
type DisconnectParam struct {
	User struct {
		UserName string `json:"userName"`
	} `json:"User"`
}

// HddInfo represents hard disk information
type HddInfo struct {
	Capacity int    `json:"capacity"` // Total capacity in MB
	Format   int    `json:"format"`   // Format status
	Mount    int    `json:"mount"`    // Mount status
	Size     int    `json:"size"`     // Used size in MB
	Status   string `json:"status"`   // "ok", "error", etc.
}

// HddInfoValue wraps HDD array for API response
type HddInfoValue struct {
	HddInfo []HddInfo `json:"HddInfo"`
}

// FormatParam represents parameters for Format command
type FormatParam struct {
	Hdd struct {
		ID int `json:"id"`
	} `json:"Hdd"`
}

// AutoMaint represents automatic maintenance configuration
type AutoMaint struct {
	Enable  int    `json:"enable"`
	WeekDay string `json:"weekDay"` // "Everyday", "Sunday", "Monday", etc.
	Hour    int    `json:"hour"`    // 0-23
	Min     int    `json:"min"`     // 0-59
	Sec     int    `json:"sec"`     // 0-59
}

// AutoMaintValue wraps AutoMaint for API response
type AutoMaintValue struct {
	AutoMaint AutoMaint `json:"AutoMaint"`
}

// AutoMaintParam represents parameters for SetAutoMaint
type AutoMaintParam struct {
	AutoMaint AutoMaint `json:"AutoMaint"`
}

// ChannelStatus represents status of a single channel
type ChannelStatus struct {
	Channel  int    `json:"channel"`
	Name     string `json:"name"`
	Online   int    `json:"online"`   // 0=offline, 1=online
	TypeInfo string `json:"typeInfo"` // Camera model/type
}

// ChannelStatusValue wraps channel status for API response
type ChannelStatusValue struct {
	Count  int             `json:"count"`
	Status []ChannelStatus `json:"status"`
}

// CertificateInfo represents SSL certificate information
type CertificateInfo struct {
	Enable  int    `json:"enable"`  // 0=disabled, 1=enabled
	CrtName string `json:"crtName"` // Certificate file name
	KeyName string `json:"keyName"` // Private key file name
}

// CertificateInfoValue wraps CertificateInfo for API response
type CertificateInfoValue struct {
	CertificateInfo CertificateInfo `json:"CertificateInfo"`
}
