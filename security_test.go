package reolink

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityAPI_GetUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetUser",
			Code: 0,
			Value: json.RawMessage(`{
				"User": [
					{
						"userName": "admin",
						"level": "admin"
					},
					{
						"userName": "guest",
						"level": "guest"
					}
				]
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	users, err := client.Security.GetUsers(ctx)
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}

	if users[0].UserName != "admin" {
		t.Errorf("expected first user 'admin', got '%s'", users[0].UserName)
	}

	if users[1].UserName != "guest" {
		t.Errorf("expected second user 'guest', got '%s'", users[1].UserName)
	}
}

func TestSecurityAPI_AddUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "AddUser" {
			t.Errorf("expected AddUser command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "AddUser",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	user := User{
		UserName: "newuser",
		Password: "password123",
		Level:    "guest",
	}
	err := client.Security.AddUser(ctx, user)
	if err != nil {
		t.Fatalf("AddUser failed: %v", err)
	}
}

func TestSecurityAPI_ModifyUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "ModifyUser" {
			t.Errorf("expected ModifyUser command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "ModifyUser",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	user := User{
		UserName: "existinguser",
		Password: "newpassword",
		Level:    "admin",
	}
	err := client.Security.ModifyUser(ctx, user)
	if err != nil {
		t.Fatalf("ModifyUser failed: %v", err)
	}
}

func TestSecurityAPI_DeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "DelUser" {
			t.Errorf("expected DelUser command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "DelUser",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.Security.DeleteUser(ctx, "userToDelete")
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
}

func TestSecurityAPI_GetOnlineUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetOnline",
			Code: 0,
			Value: json.RawMessage(`{
				"Online": {
					"User": [
						{
							"userName": "admin",
							"ip": "192.168.1.100",
							"loginTime": "2024-10-27 14:30:00"
						}
					]
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	users, err := client.Security.GetOnlineUsers(ctx)
	if err != nil {
		t.Fatalf("GetOnlineUsers failed: %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 online user, got %d", len(users))
	}

	if users[0].UserName != "admin" {
		t.Errorf("expected user 'admin', got '%s'", users[0].UserName)
	}

	if users[0].IP != "192.168.1.100" {
		t.Errorf("expected IP '192.168.1.100', got '%s'", users[0].IP)
	}
}

func TestSecurityAPI_DisconnectUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "Disconnect" {
			t.Errorf("expected Disconnect command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "Disconnect",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.Security.DisconnectUser(ctx, "userToDisconnect")
	if err != nil {
		t.Fatalf("DisconnectUser failed: %v", err)
	}
}

func TestSecurityAPI_GetCertificateInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetCertificateInfo",
			Code: 0,
			Value: json.RawMessage(`{
				"CertificateInfo": {
					"enable": 1,
					"crtName": "server.crt",
					"keyName": "server.key"
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	certInfo, err := client.Security.GetCertificateInfo(ctx)
	if err != nil {
		t.Fatalf("GetCertificateInfo failed: %v", err)
	}

	if certInfo.Enable != 1 {
		t.Errorf("expected enable 1, got %d", certInfo.Enable)
	}
	if certInfo.CrtName != "server.crt" {
		t.Errorf("expected crtName 'server.crt', got '%s'", certInfo.CrtName)
	}
}

func TestSecurityAPI_CertificateClear(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		json.NewDecoder(r.Body).Decode(&req)

		if len(req) != 1 || req[0].Cmd != "CertificateClear" {
			t.Errorf("expected CertificateClear command, got %v", req)
		}

		resp := []Response{{
			Cmd:  "CertificateClear",
			Code: 0,
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	err := client.Security.CertificateClear(ctx)
	if err != nil {
		t.Fatalf("CertificateClear failed: %v", err)
	}
}

func TestSecurityAPI_GetSysCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetSysCfg",
			Code: 0,
			Value: json.RawMessage(`{
				"version": "1.0",
				"deviceName": "TestCamera",
				"settings": {
					"key1": "value1",
					"key2": "value2"
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	config, err := client.Security.GetSysCfg(ctx, 0)
	if err != nil {
		t.Fatalf("GetSysCfg failed: %v", err)
	}

	if config == nil {
		t.Fatal("expected non-nil config")
	}

	if version, ok := config["version"].(string); !ok || version != "1.0" {
		t.Errorf("expected version '1.0', got %v", config["version"])
	}

	if deviceName, ok := config["deviceName"].(string); !ok || deviceName != "TestCamera" {
		t.Errorf("expected deviceName 'TestCamera', got %v", config["deviceName"])
	}
}

func TestSecurityAPI_SetSysCfg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req []Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request: %v", err)
		}

		if len(req) == 0 || req[0].Cmd != "SetSysCfg" {
			t.Errorf("expected SetSysCfg command, got %v", req)
		}

		resp := []Response{{
			Cmd:   "SetSysCfg",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := context.Background()
	config := map[string]interface{}{
		"version":    "1.0",
		"deviceName": "TestCamera",
		"settings": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := client.Security.SetSysCfg(ctx, config)
	if err != nil {
		t.Fatalf("SetSysCfg failed: %v", err)
	}
}
