package reolink

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNetworkAPI_GetNetPort(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetNetPort", "code": 0, "value": {"NetPort": {"httpPort": 80, "httpsPort": 443, "mediaPort": 9000, "onvifPort": 8000, "rtmpPort": 1935, "rtspPort": 554}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	netPort, err := client.Network.GetNetPort(ctx)
	if err != nil {
		t.Fatalf("GetNetPort failed: %v", err)
	}

	if netPort.HTTPPort != 80 {
		t.Errorf("Expected HTTPPort 80, got %d", netPort.HTTPPort)
	}
	if netPort.HTTPSPort != 443 {
		t.Errorf("Expected HTTPSPort 443, got %d", netPort.HTTPSPort)
	}
}

func TestNetworkAPI_SetNetPort(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetNetPort", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	netPort := NetPort{
		HTTPPort:  80,
		HTTPSPort: 443,
		MediaPort: 9000,
		OnvifPort: 8000,
		RTMPPort:  1935,
		RTSPPort:  554,
	}

	err := client.Network.SetNetPort(ctx, netPort)
	if err != nil {
		t.Fatalf("SetNetPort failed: %v", err)
	}
}

func TestNetworkAPI_GetLocalLink(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetLocalLink", "code": 0, "value": {"LocalLink": {"type": "DHCP", "static": {"ip": "192.168.1.100", "gateway": "192.168.1.1", "mask": "255.255.255.0"}, "dns": {"auto": 1, "dns1": "8.8.8.8", "dns2": "8.8.4.4"}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	localLink, err := client.Network.GetLocalLink(ctx)
	if err != nil {
		t.Fatalf("GetLocalLink failed: %v", err)
	}

	if localLink.Type != "DHCP" {
		t.Errorf("Expected Type DHCP, got %s", localLink.Type)
	}
	if localLink.Static.IP != "192.168.1.100" {
		t.Errorf("Expected Static.IP 192.168.1.100, got %s", localLink.Static.IP)
	}
}

func TestNetworkAPI_GetNtp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetNtp", "code": 0, "value": {"Ntp": {"enable": 1, "interval": 720, "port": 123, "server": "time.google.com"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ntp, err := client.Network.GetNtp(ctx)
	if err != nil {
		t.Fatalf("GetNtp failed: %v", err)
	}

	if ntp.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", ntp.Enable)
	}
	if ntp.Server != "time.google.com" {
		t.Errorf("Expected Server time.google.com, got %s", ntp.Server)
	}
}

func TestNetworkAPI_GetWifi(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetWifi", "code": 0, "value": {"Wifi": {"ssid": "MyNetwork", "password": "secret123"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	wifi, err := client.Network.GetWifi(ctx)
	if err != nil {
		t.Fatalf("GetWifi failed: %v", err)
	}

	if wifi.SSID != "MyNetwork" {
		t.Errorf("Expected SSID MyNetwork, got %s", wifi.SSID)
	}
	if wifi.Password != "secret123" {
		t.Errorf("Expected Password secret123, got %s", wifi.Password)
	}
}

func TestNetworkAPI_GetDdns(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetDdns", "code": 0, "value": {"Ddns": {"enable": 1, "type": "NO-IP", "domain": "mycamera.ddns.net"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ddns, err := client.Network.GetDdns(ctx)
	if err != nil {
		t.Fatalf("GetDdns failed: %v", err)
	}

	if ddns.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", ddns.Enable)
	}
	if ddns.Domain != "mycamera.ddns.net" {
		t.Errorf("Expected Domain mycamera.ddns.net, got %s", ddns.Domain)
	}
}

func TestNetworkAPI_GetEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetEmail", "code": 0, "value": {"Email": {"smtpServer": "smtp.gmail.com", "smtpPort": 587, "addr1": "user@example.com", "schedule": {"enable": 1}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	email, err := client.Network.GetEmail(ctx)
	if err != nil {
		t.Fatalf("GetEmail failed: %v", err)
	}

	if email.SMTPServer != "smtp.gmail.com" {
		t.Errorf("Expected SMTPServer smtp.gmail.com, got %s", email.SMTPServer)
	}
	if email.SMTPPort != 587 {
		t.Errorf("Expected SMTPPort 587, got %d", email.SMTPPort)
	}
	if email.Addr1 != "user@example.com" {
		t.Errorf("Expected Addr1 user@example.com, got %s", email.Addr1)
	}
}

func TestNetworkAPI_GetFtp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetFtp", "code": 0, "value": {"Ftp": {"server": "ftp.example.com", "port": 21, "userName": "ftpuser", "schedule": {"enable": 1}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ftp, err := client.Network.GetFtp(ctx)
	if err != nil {
		t.Fatalf("GetFtp failed: %v", err)
	}

	if ftp.Server != "ftp.example.com" {
		t.Errorf("Expected Server ftp.example.com, got %s", ftp.Server)
	}
	if ftp.Port != 21 {
		t.Errorf("Expected Port 21, got %d", ftp.Port)
	}
	if ftp.UserName != "ftpuser" {
		t.Errorf("Expected UserName ftpuser, got %s", ftp.UserName)
	}
}

func TestNetworkAPI_GetPush(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetPush", "code": 0, "value": {"Push": {"schedule": {"enable": 1}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	push, err := client.Network.GetPush(ctx)
	if err != nil {
		t.Fatalf("GetPush failed: %v", err)
	}

	if push.Schedule.Enable != 1 {
		t.Errorf("Expected Schedule.Enable 1, got %d", push.Schedule.Enable)
	}
}

func TestNetworkAPI_GetP2p(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetP2p", "code": 0, "value": {"P2p": {"enable": 1, "uid": "ABCD1234EFGH5678"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	p2p, err := client.Network.GetP2p(ctx)
	if err != nil {
		t.Fatalf("GetP2p failed: %v", err)
	}

	if p2p.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", p2p.Enable)
	}
	if p2p.UID != "ABCD1234EFGH5678" {
		t.Errorf("Expected UID ABCD1234EFGH5678, got %s", p2p.UID)
	}
}

func TestNetworkAPI_GetUpnp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetUpnp", "code": 0, "value": {"Upnp": {"enable": 1}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	upnp, err := client.Network.GetUpnp(ctx)
	if err != nil {
		t.Fatalf("GetUpnp failed: %v", err)
	}

	if upnp.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", upnp.Enable)
	}
}

func TestNetworkAPI_TestEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "TestEmail", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	err := client.Network.TestEmail(ctx)
	if err != nil {
		t.Fatalf("TestEmail failed: %v", err)
	}
}

func TestNetworkAPI_TestFtp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "TestFtp", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	err := client.Network.TestFtp(ctx)
	if err != nil {
		t.Fatalf("TestFtp failed: %v", err)
	}
}

func TestNetworkAPI_ScanWifi(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "ScanWifi", "code": 0, "value": [{"ssid": "Network1", "signal": 80, "encrypt": 1}, {"ssid": "Network2", "signal": 60, "encrypt": 0}]}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	networks, err := client.Network.ScanWifi(ctx)
	if err != nil {
		t.Fatalf("ScanWifi failed: %v", err)
	}

	if len(networks) != 2 {
		t.Fatalf("Expected 2 networks, got %d", len(networks))
	}

	if networks[0].SSID != "Network1" {
		t.Errorf("Expected SSID Network1, got %s", networks[0].SSID)
	}
	if networks[0].Signal != 80 {
		t.Errorf("Expected Signal 80, got %d", networks[0].Signal)
	}
}

func TestNetworkAPI_GetWifiSignal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetWifiSignal", "code": 0, "value": {"signal": 75}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	signal, err := client.Network.GetWifiSignal(ctx)
	if err != nil {
		t.Fatalf("GetWifiSignal failed: %v", err)
	}

	if signal.Signal != 75 {
		t.Errorf("Expected Signal 75, got %d", signal.Signal)
	}
}

func TestNetworkAPI_GetEmailV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetEmailV20", "code": 0, "value": {"Email": {"smtpServer": "smtp.gmail.com", "smtpPort": 465, "userName": "user@example.com", "schedule": {"enable": 1, "table": {"MD": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	email, err := client.Network.GetEmailV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetEmailV20 failed: %v", err)
	}

	if email.SMTPServer != "smtp.gmail.com" {
		t.Errorf("Expected SMTPServer smtp.gmail.com, got %s", email.SMTPServer)
	}
	if email.SMTPPort != 465 {
		t.Errorf("Expected SMTPPort 465, got %d", email.SMTPPort)
	}
}

func TestNetworkAPI_SetEmailV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetEmailV20", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	email := Email{
		SMTPServer: "smtp.gmail.com",
		SMTPPort:   465,
		UserName:   "user@example.com",
		Schedule: EmailSchedule{
			Enable: 1,
		},
	}

	err := client.Network.SetEmailV20(ctx, 0, email)
	if err != nil {
		t.Fatalf("SetEmailV20 failed: %v", err)
	}
}

func TestNetworkAPI_GetFtpV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetFtpV20", "code": 0, "value": {"Ftp": {"server": "ftp.example.com", "port": 21, "userName": "ftpuser", "schedule": {"enable": 1, "table": {"MD": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ftp, err := client.Network.GetFtpV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetFtpV20 failed: %v", err)
	}

	if ftp.Server != "ftp.example.com" {
		t.Errorf("Expected Server ftp.example.com, got %s", ftp.Server)
	}
	if ftp.Port != 21 {
		t.Errorf("Expected Port 21, got %d", ftp.Port)
	}
}

func TestNetworkAPI_SetFtpV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetFtpV20", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ftp := Ftp{
		Server:   "ftp.example.com",
		Port:     21,
		UserName: "ftpuser",
		Schedule: FtpSchedule{
			Enable: 1,
		},
	}

	err := client.Network.SetFtpV20(ctx, 0, ftp)
	if err != nil {
		t.Fatalf("SetFtpV20 failed: %v", err)
	}
}

func TestNetworkAPI_GetPushV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetPushV20", "code": 0, "value": {"Push": {"schedule": {"enable": 1, "table": {"MD": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}}}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	push, err := client.Network.GetPushV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetPushV20 failed: %v", err)
	}

	if push.Schedule.Enable != 1 {
		t.Errorf("Expected Schedule.Enable 1, got %d", push.Schedule.Enable)
	}
}

func TestNetworkAPI_SetPushV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetPushV20", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	push := Push{
		Schedule: PushSchedule{
			Enable: 1,
		},
	}

	err := client.Network.SetPushV20(ctx, 0, push)
	if err != nil {
		t.Fatalf("SetPushV20 failed: %v", err)
	}
}

func TestNetworkAPI_GetPushCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetPushCfg", "code": 0, "value": {"PushCfg": {"enable": 1, "token": "test-push-token"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	pushCfg, err := client.Network.GetPushCfg(ctx)
	if err != nil {
		t.Fatalf("GetPushCfg failed: %v", err)
	}

	if pushCfg.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", pushCfg.Enable)
	}
	if pushCfg.Token != "test-push-token" {
		t.Errorf("Expected Token test-push-token, got %s", pushCfg.Token)
	}
}

func TestNetworkAPI_SetPushCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetPushCfg", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	pushCfg := PushCfg{
		Enable: 1,
		Token:  "test-push-token",
	}

	err := client.Network.SetPushCfg(ctx, pushCfg)
	if err != nil {
		t.Fatalf("SetPushCfg failed: %v", err)
	}
}

func TestNetworkAPI_TestWifi(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "TestWifi", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	err := client.Network.TestWifi(ctx)
	if err != nil {
		t.Fatalf("TestWifi failed: %v", err)
	}
}

func TestNetworkAPI_GetRtspUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetRtspUrl", "code": 0, "value": {"rtspUrl": {"channel": 1, "mainStream": "rtsp://192.168.1.100:554/Preview_01_main", "subStream": "rtsp://192.168.1.100:554/Preview_01_sub"}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	rtspUrl, err := client.Network.GetRtspUrl(ctx, 1)
	if err != nil {
		t.Fatalf("GetRtspUrl failed: %v", err)
	}

	if rtspUrl.Channel != 1 {
		t.Errorf("Expected Channel 1, got %d", rtspUrl.Channel)
	}
	if rtspUrl.MainStream != "rtsp://192.168.1.100:554/Preview_01_main" {
		t.Errorf("Expected MainStream rtsp://192.168.1.100:554/Preview_01_main, got %s", rtspUrl.MainStream)
	}
	if rtspUrl.SubStream != "rtsp://192.168.1.100:554/Preview_01_sub" {
		t.Errorf("Expected SubStream rtsp://192.168.1.100:554/Preview_01_sub, got %s", rtspUrl.SubStream)
	}
}

func TestNetworkAPI_SetLocalLink(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetLocalLink", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	localLink := LocalLink{
		Type: "Static",
		Static: StaticIP{
			IP:      "192.168.1.100",
			Gateway: "192.168.1.1",
			Mask:    "255.255.255.0",
		},
		DNS: DNSConfig{
			Auto: 0,
			DNS1: "8.8.8.8",
			DNS2: "8.8.4.4",
		},
	}

	err := client.Network.SetLocalLink(ctx, localLink)
	if err != nil {
		t.Fatalf("SetLocalLink failed: %v", err)
	}
}

func TestNetworkAPI_SetNtp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetNtp", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ntp := Ntp{
		Enable:   1,
		Interval: 720,
		Port:     123,
		Server:   "time.google.com",
	}

	err := client.Network.SetNtp(ctx, ntp)
	if err != nil {
		t.Fatalf("SetNtp failed: %v", err)
	}
}

func TestNetworkAPI_SetWifi(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetWifi", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	wifi := Wifi{
		SSID:     "TestNetwork",
		Password: "testpassword",
	}

	err := client.Network.SetWifi(ctx, wifi)
	if err != nil {
		t.Fatalf("SetWifi failed: %v", err)
	}
}

func TestNetworkAPI_SetDdns(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetDdns", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ddns := Ddns{
		Enable: 1,
		Type:   "NO-IP",
	}

	err := client.Network.SetDdns(ctx, ddns)
	if err != nil {
		t.Fatalf("SetDdns failed: %v", err)
	}
}

func TestNetworkAPI_SetEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetEmail", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	email := Email{
		SMTPServer: "smtp.gmail.com",
	}

	err := client.Network.SetEmail(ctx, email)
	if err != nil {
		t.Fatalf("SetEmail failed: %v", err)
	}
}

func TestNetworkAPI_SetFtp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetFtp", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	ftp := Ftp{
		Server: "ftp.example.com",
	}

	err := client.Network.SetFtp(ctx, ftp)
	if err != nil {
		t.Fatalf("SetFtp failed: %v", err)
	}
}

func TestNetworkAPI_SetPush(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetPush", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	push := Push{
		Schedule: PushSchedule{
			Enable: 1,
		},
	}

	err := client.Network.SetPush(ctx, push)
	if err != nil {
		t.Fatalf("SetPush failed: %v", err)
	}
}

func TestNetworkAPI_SetP2p(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetP2p", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	p2p := P2p{
		Enable: 1,
	}

	err := client.Network.SetP2p(ctx, p2p)
	if err != nil {
		t.Fatalf("SetP2p failed: %v", err)
	}
}

func TestNetworkAPI_SetUpnp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetUpnp", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	upnp := Upnp{
		Enable: 1,
	}

	err := client.Network.SetUpnp(ctx, upnp)
	if err != nil {
		t.Fatalf("SetUpnp failed: %v", err)
	}
}
