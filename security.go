package reolink

import (
	"context"
	"encoding/json"
	"fmt"
)

// SecurityAPI provides access to security and user management API endpoints
type SecurityAPI struct {
	client *Client
}

// GetUsers retrieves the list of users
func (s *SecurityAPI) GetUsers(ctx context.Context) ([]User, error) {
	s.client.logger.Debug("getting users")

	req := []Request{{
		Cmd:    "GetUser",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get users: %v", err)
		return nil, fmt.Errorf("GetUser request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get users: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get users: %v", apiErr)
		return nil, apiErr
	}

	var value UserValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse users response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	s.client.logger.Info("successfully retrieved users: count=%d", len(value.User))
	return value.User, nil
}

// AddUser adds a new user
func (s *SecurityAPI) AddUser(ctx context.Context, user User) error {
	s.client.logger.Info("adding user: username=%s", user.UserName)

	req := []Request{{
		Cmd: "AddUser",
		Param: AddUserParam{
			User: user,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to add user: %v", err)
		return fmt.Errorf("AddUser request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to add user: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to add user: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully added user")
	return nil
}

// ModifyUser modifies an existing user
func (s *SecurityAPI) ModifyUser(ctx context.Context, user User) error {
	s.client.logger.Info("modifying user: username=%s", user.UserName)

	req := []Request{{
		Cmd: "ModifyUser",
		Param: ModifyUserParam{
			User: user,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to modify user: %v", err)
		return fmt.Errorf("ModifyUser request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to modify user: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to modify user: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully modified user")
	return nil
}

// DeleteUser deletes a user
func (s *SecurityAPI) DeleteUser(ctx context.Context, username string) error {
	s.client.logger.Warn("deleting user (destructive): username=%s", username)

	req := []Request{{
		Cmd: "DelUser",
		Param: DelUserParam{
			User: User{
				UserName: username,
			},
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to delete user: %v", err)
		return fmt.Errorf("DelUser request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to delete user: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to delete user: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully deleted user")
	return nil
}

// GetOnlineUsers retrieves the list of currently online users
func (s *SecurityAPI) GetOnlineUsers(ctx context.Context) ([]OnlineUser, error) {
	s.client.logger.Debug("getting online users")

	req := []Request{{
		Cmd:    "GetOnline",
		Action: 0,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get online users: %v", err)
		return nil, fmt.Errorf("GetOnline request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get online users: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get online users: %v", apiErr)
		return nil, apiErr
	}

	var value OnlineValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse online users response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	s.client.logger.Info("successfully retrieved online users: count=%d", len(value.Online.Users))
	return value.Online.Users, nil
}

// DisconnectUser disconnects a user session
func (s *SecurityAPI) DisconnectUser(ctx context.Context, username string) error {
	s.client.logger.Warn("disconnecting user: username=%s", username)

	req := []Request{{
		Cmd: "Disconnect",
		Param: DisconnectParam{
			User: struct {
				UserName string `json:"userName"`
			}{
				UserName: username,
			},
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to disconnect user: %v", err)
		return fmt.Errorf("disconnect request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to disconnect user: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to disconnect user: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully disconnected user")
	return nil
}

// GetSysCfg exports system configuration
func (s *SecurityAPI) GetSysCfg(ctx context.Context, channel int) (map[string]interface{}, error) {
	s.client.logger.Debug("getting system configuration export")

	req := []Request{{
		Cmd: "GetSysCfg",
		Param: map[string]interface{}{
			"channel": channel,
		},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get system configuration export: %v", err)
		return nil, fmt.Errorf("GetSysCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get system configuration export: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get system configuration export: %v", apiErr)
		return nil, apiErr
	}

	var value map[string]interface{}
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse system configuration export response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return value, nil
}

// SetSysCfg imports system configuration
func (s *SecurityAPI) SetSysCfg(ctx context.Context, config map[string]interface{}) error {
	s.client.logger.Info("importing system configuration")

	req := []Request{{
		Cmd:   "SetSysCfg",
		Param: config,
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to import system configuration: %v", err)
		return fmt.Errorf("SetSysCfg request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to import system configuration: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to import system configuration: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully imported system configuration")
	return nil
}

// GetCertificateInfo gets SSL certificate information
func (s *SecurityAPI) GetCertificateInfo(ctx context.Context) (*CertificateInfo, error) {
	s.client.logger.Debug("getting certificate info")

	req := []Request{{
		Cmd:    "GetCertificateInfo",
		Action: 0,
		Param:  map[string]interface{}{},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to get certificate info: %v", err)
		return nil, fmt.Errorf("GetCertificateInfo request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to get certificate info: %v", err)
		return nil, err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to get certificate info: %v", apiErr)
		return nil, apiErr
	}

	var value CertificateInfoValue
	if err := json.Unmarshal(resp[0].Value, &value); err != nil {
		s.client.logger.Error("failed to parse certificate info response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &value.CertificateInfo, nil
}

// CertificateClear clears SSL certificate
func (s *SecurityAPI) CertificateClear(ctx context.Context) error {
	s.client.logger.Warn("clearing SSL certificate (destructive)")

	req := []Request{{
		Cmd:    "CertificateClear",
		Action: 0,
		Param:  map[string]interface{}{},
	}}

	var resp []Response
	if err := s.client.do(ctx, req, &resp); err != nil {
		s.client.logger.Error("failed to clear certificate: %v", err)
		return fmt.Errorf("CertificateClear request failed: %w", err)
	}

	if len(resp) == 0 {
		err := fmt.Errorf("empty response")
		s.client.logger.Error("failed to clear certificate: %v", err)
		return err
	}

	if apiErr := resp[0].ToAPIError(); apiErr != nil {
		s.client.logger.Error("failed to clear certificate: %v", apiErr)
		return apiErr
	}

	s.client.logger.Info("successfully cleared SSL certificate")
	return nil
}
