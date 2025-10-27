package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// NetworkAPI provides methods for network configuration
type NetworkAPI struct {
	client *Client
}

// NetPort represents network port configuration
type NetPort struct {
	HTTPEnable  int `json:"httpEnable"`  // 0=disabled, 1=enabled
	HTTPPort    int `json:"httpPort"`    // HTTP port (default: 80)
	HTTPSEnable int `json:"httpsEnable"` // 0=disabled, 1=enabled
	HTTPSPort   int `json:"httpsPort"`   // HTTPS port (default: 443)
	MediaPort   int `json:"mediaPort"`   // Media port (default: 9000)
	OnvifEnable int `json:"onvifEnable"` // 0=disabled, 1=enabled
	OnvifPort   int `json:"onvifPort"`   // ONVIF port (default: 8000)
	RTMPEnable  int `json:"rtmpEnable"`  // 0=disabled, 1=enabled
	RTMPPort    int `json:"rtmpPort"`    // RTMP port (default: 1935)
	RTSPEnable  int `json:"rtspEnable"`  // 0=disabled, 1=enabled
	RTSPPort    int `json:"rtspPort"`    // RTSP port (default: 554)
}

// NetPortValue represents the response value for GetNetPort
type NetPortValue struct {
	NetPort NetPort `json:"NetPort"`
}

// LocalLink represents local network configuration
type LocalLink struct {
	Type   string    `json:"type"`   // "DHCP" or "Static"
	Static StaticIP  `json:"static"` // Static IP configuration
	DNS    DNSConfig `json:"dns"`    // DNS configuration
}

// StaticIP represents static IP configuration
type StaticIP struct {
	IP      string `json:"ip"`      // IP address
	Mask    string `json:"mask"`    // Subnet mask
	Gateway string `json:"gateway"` // Gateway address
}

// DNSConfig represents DNS configuration
type DNSConfig struct {
	Auto int    `json:"auto"` // 0=manual, 1=auto
	DNS1 string `json:"dns1"` // Primary DNS server
	DNS2 string `json:"dns2"` // Secondary DNS server
}

// LocalLinkValue represents the response value for GetLocalLink
type LocalLinkValue struct {
	LocalLink LocalLink `json:"LocalLink"`
}

// Ntp represents NTP configuration
type Ntp struct {
	Enable   int    `json:"enable"`   // 0=disabled, 1=enabled
	Server   string `json:"server"`   // NTP server address
	Port     int    `json:"port"`     // NTP server port (default: 123)
	Interval int    `json:"interval"` // Sync interval in seconds (0=immediate, 10-65535)
}

// NtpValue represents the response value for GetNtp
type NtpValue struct {
	Ntp Ntp `json:"Ntp"`
}

// GetNetPort gets network port configuration
func (n *NetworkAPI) GetNetPort(ctx context.Context) (*NetPort, error) {
	n.client.logger.Debug("getting network port configuration")

	req := []Request{{
		Cmd:    "GetNetPort",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get network port configuration: %v", err)
		return nil, fmt.Errorf("GetNetPort request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetNetPort")
		n.client.logger.Error("failed to get network port configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get network port configuration: %v", err)
		return nil, err
	}

	var value NetPortValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse network port configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetNetPort response: %w", err)
	}

	n.client.logger.Info("successfully retrieved network port configuration: httpPort=%d httpsPort=%d",
		value.NetPort.HTTPPort, value.NetPort.HTTPSPort)
	return &value.NetPort, nil
}

// SetNetPort sets network port configuration
func (n *NetworkAPI) SetNetPort(ctx context.Context, netPort NetPort) error {
	n.client.logger.Info("setting network port configuration: httpPort=%d httpsPort=%d",
		netPort.HTTPPort, netPort.HTTPSPort)

	req := []Request{{
		Cmd: "SetNetPort",
		Param: map[string]interface{}{
			"NetPort": netPort,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set network port configuration: %v", err)
		return fmt.Errorf("SetNetPort request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetNetPort")
		n.client.logger.Error("failed to set network port configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set network port configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set network port configuration")
	return nil
}

// GetLocalLink gets local network configuration
func (n *NetworkAPI) GetLocalLink(ctx context.Context) (*LocalLink, error) {
	n.client.logger.Debug("getting local network configuration")

	req := []Request{{
		Cmd:    "GetLocalLink",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get local network configuration: %v", err)
		return nil, fmt.Errorf("GetLocalLink request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetLocalLink")
		n.client.logger.Error("failed to get local network configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get local network configuration: %v", err)
		return nil, err
	}

	var value LocalLinkValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse local network configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetLocalLink response: %w", err)
	}

	n.client.logger.Info("successfully retrieved local network configuration: type=%s",
		value.LocalLink.Type)
	return &value.LocalLink, nil
}

// SetLocalLink sets local network configuration
func (n *NetworkAPI) SetLocalLink(ctx context.Context, localLink LocalLink) error {
	n.client.logger.Info("setting local network configuration: type=%s",
		localLink.Type)

	req := []Request{{
		Cmd: "SetLocalLink",
		Param: map[string]interface{}{
			"LocalLink": localLink,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set local network configuration: %v", err)
		return fmt.Errorf("SetLocalLink request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetLocalLink")
		n.client.logger.Error("failed to set local network configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set local network configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set local network configuration")
	return nil
}

// GetNtp gets NTP configuration
func (n *NetworkAPI) GetNtp(ctx context.Context) (*Ntp, error) {
	n.client.logger.Debug("getting NTP configuration")

	req := []Request{{
		Cmd:    "GetNtp",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get NTP configuration: %v", err)
		return nil, fmt.Errorf("GetNtp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetNtp")
		n.client.logger.Error("failed to get NTP configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get NTP configuration: %v", err)
		return nil, err
	}

	var value NtpValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse NTP configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetNtp response: %w", err)
	}

	n.client.logger.Info("successfully retrieved NTP configuration: server=%s enable=%d", value.Ntp.Server, value.Ntp.Enable)
	return &value.Ntp, nil
}

// SetNtp sets NTP configuration
func (n *NetworkAPI) SetNtp(ctx context.Context, ntp Ntp) error {
	n.client.logger.Info("setting NTP configuration: server=%s enable=%d", ntp.Server, ntp.Enable)

	req := []Request{{
		Cmd: "SetNtp",
		Param: map[string]interface{}{
			"Ntp": ntp,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set NTP configuration: %v", err)
		return fmt.Errorf("SetNtp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetNtp")
		n.client.logger.Error("failed to set NTP configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set NTP configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set NTP configuration")
	return nil
}

// Wifi represents WiFi configuration
type Wifi struct {
	SSID     string `json:"ssid"`     // WiFi network name
	Password string `json:"password"` // WiFi password
}

// WifiValue represents the response value for GetWifi
type WifiValue struct {
	Wifi Wifi `json:"Wifi"`
}

// Ddns represents DDNS configuration
type Ddns struct {
	Enable   int    `json:"enable"`   // 0=disabled, 1=enabled
	Type     string `json:"type"`     // "3322" or "Dyndns"
	UserName string `json:"userName"` // DDNS username
	Password string `json:"password"` // DDNS password
	Domain   string `json:"domain"`   // Domain name
}

// DdnsValue represents the response value for GetDdns
type DdnsValue struct {
	Ddns Ddns `json:"Ddns"`
}

// Email represents email configuration
type Email struct {
	SMTPServer string        `json:"smtpServer"` // SMTP server address
	SMTPPort   int           `json:"smtpPort"`   // SMTP port (default: 25, 465 for SSL)
	UserName   string        `json:"userName"`   // Email username
	Password   string        `json:"password"`   // Email password
	Addr1      string        `json:"addr1"`      // Recipient email 1
	Addr2      string        `json:"addr2"`      // Recipient email 2
	Addr3      string        `json:"addr3"`      // Recipient email 3
	Interval   int           `json:"interval"`   // Email interval in seconds
	Schedule   EmailSchedule `json:"schedule"`   // Email schedule
}

// EmailSchedule represents email schedule configuration
type EmailSchedule struct {
	Enable int         `json:"enable"` // 0=disabled, 1=enabled
	Table  interface{} `json:"table"`  // string for v1, EmailScheduleTable for v2.0
}

// EmailScheduleTable represents v2.0 email schedule with multiple alarm types
type EmailScheduleTable struct {
	MD        string `json:"MD,omitempty"`         // Motion detection schedule
	TIMING    string `json:"TIMING,omitempty"`     // Timing schedule
	AIPeople  string `json:"AI_PEOPLE,omitempty"`  // AI people detection schedule
	AIVehicle string `json:"AI_VEHICLE,omitempty"` // AI vehicle detection schedule
	AIDogCat  string `json:"AI_DOG_CAT,omitempty"` // AI dog/cat detection schedule
}

// EmailValue represents the response value for GetEmail
type EmailValue struct {
	Email Email `json:"Email"`
}

// Ftp represents FTP configuration
type Ftp struct {
	Server    string      `json:"server"`              // FTP server address
	Port      int         `json:"port"`                // FTP port (default: 21)
	UserName  string      `json:"userName"`            // FTP username
	Password  string      `json:"password"`            // FTP password
	RemoteDir string      `json:"remoteDir,omitempty"` // Remote directory
	Schedule  FtpSchedule `json:"schedule"`            // FTP schedule
}

// FtpSchedule represents FTP schedule configuration
type FtpSchedule struct {
	Enable int         `json:"enable"` // 0=disabled, 1=enabled
	Table  interface{} `json:"table"`  // string for v1, FtpScheduleTable for v2.0
}

// FtpScheduleTable represents v2.0 FTP schedule with multiple alarm types
type FtpScheduleTable struct {
	MD        string `json:"MD,omitempty"`         // Motion detection schedule
	TIMING    string `json:"TIMING,omitempty"`     // Timing schedule
	AIPeople  string `json:"AI_PEOPLE,omitempty"`  // AI people detection schedule
	AIVehicle string `json:"AI_VEHICLE,omitempty"` // AI vehicle detection schedule
	AIDogCat  string `json:"AI_DOG_CAT,omitempty"` // AI dog/cat detection schedule
}

// FtpValue represents the response value for GetFtp
type FtpValue struct {
	Ftp Ftp `json:"Ftp"`
}

// Push represents push notification configuration
type Push struct {
	Schedule PushSchedule `json:"schedule"` // Push schedule
}

// PushSchedule represents push schedule configuration
type PushSchedule struct {
	Enable int         `json:"enable"` // 0=disabled, 1=enabled
	Table  interface{} `json:"table"`  // string for v1, PushScheduleTable for v2.0
}

// PushScheduleTable represents v2.0 push schedule with multiple alarm types
type PushScheduleTable struct {
	MD        string `json:"MD,omitempty"`         // Motion detection schedule
	TIMING    string `json:"TIMING,omitempty"`     // Timing schedule
	AIPeople  string `json:"AI_PEOPLE,omitempty"`  // AI people detection schedule
	AIVehicle string `json:"AI_VEHICLE,omitempty"` // AI vehicle detection schedule
	AIDogCat  string `json:"AI_DOG_CAT,omitempty"` // AI dog/cat detection schedule
}

// PushValue represents the response value for GetPush
type PushValue struct {
	Push Push `json:"Push"`
}

// P2p represents P2P configuration
type P2p struct {
	Enable int    `json:"enable"` // 0=disabled, 1=enabled
	UID    string `json:"uid"`    // P2P UID
}

// P2pValue represents the response value for GetP2p
type P2pValue struct {
	P2p P2p `json:"P2p"`
}

// Upnp represents UPnP configuration
type Upnp struct {
	Enable int `json:"enable"` // 0=disabled, 1=enabled
}

// UpnpValue represents the response value for GetUpnp
type UpnpValue struct {
	Upnp Upnp `json:"Upnp"`
}

// GetWifi gets WiFi configuration
func (n *NetworkAPI) GetWifi(ctx context.Context) (*Wifi, error) {
	n.client.logger.Debug("getting WiFi configuration")

	req := []Request{{
		Cmd:    "GetWifi",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get WiFi configuration: %v", err)
		return nil, fmt.Errorf("GetWifi request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetWifi")
		n.client.logger.Error("failed to get WiFi configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get WiFi configuration: %v", err)
		return nil, err
	}

	var value WifiValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse WiFi configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetWifi response: %w", err)
	}

	n.client.logger.Info("successfully retrieved WiFi configuration: ssid=%s", value.Wifi.SSID)
	return &value.Wifi, nil
}

// SetWifi sets WiFi configuration
func (n *NetworkAPI) SetWifi(ctx context.Context, wifi Wifi) error {
	n.client.logger.Info("setting WiFi configuration: ssid=%s", wifi.SSID)

	req := []Request{{
		Cmd: "SetWifi",
		Param: map[string]interface{}{
			"Wifi": wifi,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set WiFi configuration: %v", err)
		return fmt.Errorf("SetWifi request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetWifi")
		n.client.logger.Error("failed to set WiFi configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set WiFi configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set WiFi configuration")
	return nil
}

// GetDdns gets DDNS configuration
func (n *NetworkAPI) GetDdns(ctx context.Context) (*Ddns, error) {
	n.client.logger.Debug("getting DDNS configuration")

	req := []Request{{
		Cmd:    "GetDdns",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get DDNS configuration: %v", err)
		return nil, fmt.Errorf("GetDdns request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetDdns")
		n.client.logger.Error("failed to get DDNS configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get DDNS configuration: %v", err)
		return nil, err
	}

	var value DdnsValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse DDNS configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetDdns response: %w", err)
	}

	n.client.logger.Info("successfully retrieved DDNS configuration: enable=%d type=%s", value.Ddns.Enable, value.Ddns.Type)
	return &value.Ddns, nil
}

// SetDdns sets DDNS configuration
func (n *NetworkAPI) SetDdns(ctx context.Context, ddns Ddns) error {
	n.client.logger.Info("setting DDNS configuration: enable=%d type=%s", ddns.Enable, ddns.Type)

	req := []Request{{
		Cmd: "SetDdns",
		Param: map[string]interface{}{
			"Ddns": ddns,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set DDNS configuration: %v", err)
		return fmt.Errorf("SetDdns request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetDdns")
		n.client.logger.Error("failed to set DDNS configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set DDNS configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set DDNS configuration")
	return nil
}

// GetEmail gets email configuration
func (n *NetworkAPI) GetEmail(ctx context.Context) (*Email, error) {
	n.client.logger.Debug("getting email configuration")

	req := []Request{{
		Cmd:    "GetEmail",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get email configuration: %v", err)
		return nil, fmt.Errorf("GetEmail request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetEmail")
		n.client.logger.Error("failed to get email configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get email configuration: %v", err)
		return nil, err
	}

	var value EmailValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse email configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetEmail response: %w", err)
	}

	n.client.logger.Info("successfully retrieved email configuration: server=%s", value.Email.SMTPServer)
	return &value.Email, nil
}

// SetEmail sets email configuration
func (n *NetworkAPI) SetEmail(ctx context.Context, email Email) error {
	n.client.logger.Info("setting email configuration: server=%s", email.SMTPServer)

	req := []Request{{
		Cmd: "SetEmail",
		Param: map[string]interface{}{
			"Email": email,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set email configuration: %v", err)
		return fmt.Errorf("SetEmail request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetEmail")
		n.client.logger.Error("failed to set email configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set email configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set email configuration")
	return nil
}

// GetFtp gets FTP configuration
func (n *NetworkAPI) GetFtp(ctx context.Context) (*Ftp, error) {
	n.client.logger.Debug("getting FTP configuration")

	req := []Request{{
		Cmd:    "GetFtp",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get FTP configuration: %v", err)
		return nil, fmt.Errorf("GetFtp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetFtp")
		n.client.logger.Error("failed to get FTP configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get FTP configuration: %v", err)
		return nil, err
	}

	var value FtpValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse FTP configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetFtp response: %w", err)
	}

	n.client.logger.Info("successfully retrieved FTP configuration: server=%s", value.Ftp.Server)
	return &value.Ftp, nil
}

// SetFtp sets FTP configuration
func (n *NetworkAPI) SetFtp(ctx context.Context, ftp Ftp) error {
	n.client.logger.Info("setting FTP configuration: server=%s", ftp.Server)

	req := []Request{{
		Cmd: "SetFtp",
		Param: map[string]interface{}{
			"Ftp": ftp,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set FTP configuration: %v", err)
		return fmt.Errorf("SetFtp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetFtp")
		n.client.logger.Error("failed to set FTP configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set FTP configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set FTP configuration")
	return nil
}

// GetPush gets push notification configuration
func (n *NetworkAPI) GetPush(ctx context.Context) (*Push, error) {
	n.client.logger.Debug("getting push notification configuration")

	req := []Request{{
		Cmd:    "GetPush",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get push notification configuration: %v", err)
		return nil, fmt.Errorf("GetPush request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetPush")
		n.client.logger.Error("failed to get push notification configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get push notification configuration: %v", err)
		return nil, err
	}

	var value PushValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse push notification configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetPush response: %w", err)
	}

	n.client.logger.Info("successfully retrieved push notification configuration")
	return &value.Push, nil
}

// SetPush sets push notification configuration
func (n *NetworkAPI) SetPush(ctx context.Context, push Push) error {
	n.client.logger.Info("setting push notification configuration")

	req := []Request{{
		Cmd: "SetPush",
		Param: map[string]interface{}{
			"Push": push,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set push notification configuration: %v", err)
		return fmt.Errorf("SetPush request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetPush")
		n.client.logger.Error("failed to set push notification configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set push notification configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set push notification configuration")
	return nil
}

// GetP2p gets P2P configuration
func (n *NetworkAPI) GetP2p(ctx context.Context) (*P2p, error) {
	n.client.logger.Debug("getting P2P configuration")

	req := []Request{{
		Cmd:    "GetP2p",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get P2P configuration: %v", err)
		return nil, fmt.Errorf("GetP2p request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetP2p")
		n.client.logger.Error("failed to get P2P configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get P2P configuration: %v", err)
		return nil, err
	}

	var value P2pValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse P2P configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetP2p response: %w", err)
	}

	n.client.logger.Info("successfully retrieved P2P configuration: enable=%d", value.P2p.Enable)
	return &value.P2p, nil
}

// SetP2p sets P2P configuration
func (n *NetworkAPI) SetP2p(ctx context.Context, p2p P2p) error {
	n.client.logger.Info("setting P2P configuration: enable=%d", p2p.Enable)

	req := []Request{{
		Cmd: "SetP2p",
		Param: map[string]interface{}{
			"P2p": p2p,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set P2P configuration: %v", err)
		return fmt.Errorf("SetP2p request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetP2p")
		n.client.logger.Error("failed to set P2P configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set P2P configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set P2P configuration")
	return nil
}

// GetUpnp gets UPnP configuration
func (n *NetworkAPI) GetUpnp(ctx context.Context) (*Upnp, error) {
	n.client.logger.Debug("getting UPnP configuration")

	req := []Request{{
		Cmd:    "GetUpnp",
		Action: 0,
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get UPnP configuration: %v", err)
		return nil, fmt.Errorf("GetUpnp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetUpnp")
		n.client.logger.Error("failed to get UPnP configuration: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get UPnP configuration: %v", err)
		return nil, err
	}

	var value UpnpValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse UPnP configuration response: %v", err)
		return nil, fmt.Errorf("failed to parse GetUpnp response: %w", err)
	}

	n.client.logger.Info("successfully retrieved UPnP configuration: enable=%d", value.Upnp.Enable)
	return &value.Upnp, nil
}

// SetUpnp sets UPnP configuration
func (n *NetworkAPI) SetUpnp(ctx context.Context, upnp Upnp) error {
	n.client.logger.Info("setting UPnP configuration: enable=%d", upnp.Enable)

	req := []Request{{
		Cmd: "SetUpnp",
		Param: map[string]interface{}{
			"Upnp": upnp,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set UPnP configuration: %v", err)
		return fmt.Errorf("SetUpnp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetUpnp")
		n.client.logger.Error("failed to set UPnP configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set UPnP configuration: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set UPnP configuration")
	return nil
}

// TestEmail sends a test email
func (n *NetworkAPI) TestEmail(ctx context.Context) error {
	n.client.logger.Info("testing email configuration")

	req := []Request{{
		Cmd: "TestEmail",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to test email: %v", err)
		return fmt.Errorf("TestEmail request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from TestEmail")
		n.client.logger.Error("failed to test email: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to test email: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully tested email configuration")
	return nil
}

// TestFtp tests FTP connection
func (n *NetworkAPI) TestFtp(ctx context.Context) error {
	n.client.logger.Info("testing FTP configuration")

	req := []Request{{
		Cmd: "TestFtp",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to test FTP: %v", err)
		return fmt.Errorf("TestFtp request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from TestFtp")
		n.client.logger.Error("failed to test FTP: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to test FTP: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully tested FTP configuration")
	return nil
}

// WifiNetwork represents a WiFi network from scan results
type WifiNetwork struct {
	SSID    string `json:"ssid"`    // Network name
	Signal  int    `json:"signal"`  // Signal strength
	Encrypt int    `json:"encrypt"` // Encryption type
}

// ScanWifi scans for available WiFi networks
func (n *NetworkAPI) ScanWifi(ctx context.Context) ([]WifiNetwork, error) {
	n.client.logger.Info("scanning for WiFi networks")

	req := []Request{{
		Cmd: "ScanWifi",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to scan WiFi networks: %v", err)
		return nil, fmt.Errorf("ScanWifi request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from ScanWifi")
		n.client.logger.Error("failed to scan WiFi networks: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to scan WiFi networks: %v", err)
		return nil, err
	}

	var networks []WifiNetwork
	if err := json.Unmarshal(resp[0].Value, &networks); err != nil {
		n.client.logger.Error("failed to parse WiFi scan response: %v", err)
		return nil, fmt.Errorf("failed to parse ScanWifi response: %w", err)
	}

	n.client.logger.Info("successfully scanned WiFi networks: found %d networks", len(networks))
	return networks, nil
}

// WifiSignal represents WiFi signal strength
type WifiSignal struct {
	Signal int `json:"signal"` // Signal strength (0-100)
}

// GetWifiSignal gets current WiFi signal strength
func (n *NetworkAPI) GetWifiSignal(ctx context.Context) (*WifiSignal, error) {
	n.client.logger.Debug("getting WiFi signal strength")

	req := []Request{{
		Cmd: "GetWifiSignal",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get WiFi signal strength: %v", err)
		return nil, fmt.Errorf("GetWifiSignal request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetWifiSignal")
		n.client.logger.Error("failed to get WiFi signal strength: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get WiFi signal strength: %v", err)
		return nil, err
	}

	var signal WifiSignal
	if err := json.Unmarshal(resp[0].Value, &signal); err != nil {
		n.client.logger.Error("failed to parse WiFi signal strength response: %v", err)
		return nil, fmt.Errorf("failed to parse GetWifiSignal response: %w", err)
	}

	n.client.logger.Info("successfully retrieved WiFi signal strength: %d", signal.Signal)
	return &signal, nil
}

// GetEmailV20 gets email configuration (v2.0 with enhanced features)
func (n *NetworkAPI) GetEmailV20(ctx context.Context, channel int) (*Email, error) {
	n.client.logger.Debug("getting email configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd: "GetEmailV20",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get email configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetEmailV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetEmailV20")
		n.client.logger.Error("failed to get email configuration (v2.0): %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get email configuration (v2.0): %v", err)
		return nil, err
	}

	var value EmailValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse email configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse GetEmailV20 response: %w", err)
	}

	n.client.logger.Info("successfully retrieved email configuration (v2.0): server=%s", value.Email.SMTPServer)
	return &value.Email, nil
}

// SetEmailV20 sets email configuration (v2.0 with enhanced features)
func (n *NetworkAPI) SetEmailV20(ctx context.Context, channel int, email Email) error {
	n.client.logger.Info("setting email configuration (v2.0): channel=%d server=%s", channel, email.SMTPServer)

	req := []Request{{
		Cmd: "SetEmailV20",
		Param: map[string]interface{}{
			"Email": email,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set email configuration (v2.0): %v", err)
		return fmt.Errorf("SetEmailV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetEmailV20")
		n.client.logger.Error("failed to set email configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set email configuration (v2.0): %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set email configuration (v2.0)")
	return nil
}

// GetFtpV20 gets FTP configuration (v2.0 with enhanced features)
func (n *NetworkAPI) GetFtpV20(ctx context.Context, channel int) (*Ftp, error) {
	n.client.logger.Debug("getting FTP configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd: "GetFtpV20",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get FTP configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetFtpV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetFtpV20")
		n.client.logger.Error("failed to get FTP configuration (v2.0): %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get FTP configuration (v2.0): %v", err)
		return nil, err
	}

	var value FtpValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse FTP configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse GetFtpV20 response: %w", err)
	}

	n.client.logger.Info("successfully retrieved FTP configuration (v2.0): server=%s", value.Ftp.Server)
	return &value.Ftp, nil
}

// SetFtpV20 sets FTP configuration (v2.0 with enhanced features)
func (n *NetworkAPI) SetFtpV20(ctx context.Context, channel int, ftp Ftp) error {
	n.client.logger.Info("setting FTP configuration (v2.0): channel=%d server=%s", channel, ftp.Server)

	req := []Request{{
		Cmd: "SetFtpV20",
		Param: map[string]interface{}{
			"Ftp": ftp,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set FTP configuration (v2.0): %v", err)
		return fmt.Errorf("SetFtpV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetFtpV20")
		n.client.logger.Error("failed to set FTP configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set FTP configuration (v2.0): %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set FTP configuration (v2.0)")
	return nil
}

// GetPushV20 gets push notification configuration (v2.0 with enhanced features)
func (n *NetworkAPI) GetPushV20(ctx context.Context, channel int) (*Push, error) {
	n.client.logger.Debug("getting push notification configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd: "GetPushV20",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get push notification configuration (v2.0): %v", err)
		return nil, fmt.Errorf("GetPushV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetPushV20")
		n.client.logger.Error("failed to get push notification configuration (v2.0): %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get push notification configuration (v2.0): %v", err)
		return nil, err
	}

	var value PushValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse push notification configuration (v2.0) response: %v", err)
		return nil, fmt.Errorf("failed to parse GetPushV20 response: %w", err)
	}

	n.client.logger.Info("successfully retrieved push notification configuration (v2.0)")
	return &value.Push, nil
}

// SetPushV20 sets push notification configuration (v2.0 with enhanced features)
func (n *NetworkAPI) SetPushV20(ctx context.Context, channel int, push Push) error {
	n.client.logger.Info("setting push notification configuration (v2.0): channel=%d", channel)

	req := []Request{{
		Cmd: "SetPushV20",
		Param: map[string]interface{}{
			"Push": push,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set push notification configuration (v2.0): %v", err)
		return fmt.Errorf("SetPushV20 request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetPushV20")
		n.client.logger.Error("failed to set push notification configuration (v2.0): %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set push notification configuration (v2.0): %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set push notification configuration (v2.0)")
	return nil
}

// PushCfg represents push configuration details
type PushCfg struct {
	Enable int    `json:"enable"` // 0=disabled, 1=enabled
	Token  string `json:"token"`  // Push token
}

// PushCfgValue represents the response value for GetPushCfg
type PushCfgValue struct {
	PushCfg PushCfg `json:"PushCfg"`
}

// GetPushCfg gets push configuration details
func (n *NetworkAPI) GetPushCfg(ctx context.Context) (*PushCfg, error) {
	n.client.logger.Debug("getting push configuration details")

	req := []Request{{
		Cmd: "GetPushCfg",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get push configuration details: %v", err)
		return nil, fmt.Errorf("GetPushCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetPushCfg")
		n.client.logger.Error("failed to get push configuration details: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get push configuration details: %v", err)
		return nil, err
	}

	var value PushCfgValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse push configuration details response: %v", err)
		return nil, fmt.Errorf("failed to parse GetPushCfg response: %w", err)
	}

	n.client.logger.Info("successfully retrieved push configuration details: enable=%d", value.PushCfg.Enable)
	return &value.PushCfg, nil
}

// SetPushCfg sets push configuration details
func (n *NetworkAPI) SetPushCfg(ctx context.Context, pushCfg PushCfg) error {
	n.client.logger.Info("setting push configuration details: enable=%d", pushCfg.Enable)

	req := []Request{{
		Cmd: "SetPushCfg",
		Param: map[string]interface{}{
			"PushCfg": pushCfg,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to set push configuration details: %v", err)
		return fmt.Errorf("SetPushCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from SetPushCfg")
		n.client.logger.Error("failed to set push configuration details: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to set push configuration details: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully set push configuration details")
	return nil
}

// TestWifi tests WiFi connection
func (n *NetworkAPI) TestWifi(ctx context.Context) error {
	n.client.logger.Info("testing WiFi configuration")

	req := []Request{{
		Cmd: "TestWifi",
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to test WiFi: %v", err)
		return fmt.Errorf("TestWifi request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from TestWifi")
		n.client.logger.Error("failed to test WiFi: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		n.client.logger.Error("failed to test WiFi: %v", apiErr)
		return apiErr
	}

	n.client.logger.Info("successfully tested WiFi configuration")
	return nil
}

// RtspUrl represents RTSP URL information
type RtspUrl struct {
	Channel    int    `json:"channel"`    // Channel number
	MainStream string `json:"mainStream"` // RTSP URL for main stream
	SubStream  string `json:"subStream"`  // RTSP URL for sub stream
}

// RtspUrlValue represents the response value for GetRtspUrl
type RtspUrlValue struct {
	RtspUrl RtspUrl `json:"rtspUrl"`
}

// GetRtspUrl gets RTSP streaming URL from the camera
func (n *NetworkAPI) GetRtspUrl(ctx context.Context, channel int) (*RtspUrl, error) {
	n.client.logger.Debug("getting RTSP URL: channel=%d", channel)

	req := []Request{{
		Cmd: "GetRtspUrl",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := n.client.do(ctx, req, &resp); err != nil {
		n.client.logger.Error("failed to get RTSP URL: %v", err)
		return nil, fmt.Errorf("GetRtspUrl request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response from GetRtspUrl")
		n.client.logger.Error("failed to get RTSP URL: %v", err)
		return nil, err
	}

	if err := resp[0].ToAPIError(); err != nil {
		n.client.logger.Error("failed to get RTSP URL: %v", err)
		return nil, err
	}

	var value RtspUrlValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		n.client.logger.Error("failed to parse RTSP URL response: %v", err)
		return nil, fmt.Errorf("failed to parse GetRtspUrl response: %w", err)
	}

	n.client.logger.Info("successfully retrieved RTSP URL: channel=%d", value.RtspUrl.Channel)
	return &value.RtspUrl, nil
}
