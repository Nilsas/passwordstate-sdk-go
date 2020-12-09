package secret

import (
	"context"
	"encoding/json"
	"fmt"
)

// AddPassword Adds new password to Passwordstate client
func (r *Client) AddPassword(ctx context.Context, pass *Secret) error {
	j, err := json.Marshal(pass)
	if err != nil {
		return err
	}
	raw, code, err := r.post(ctx, "/api/passwords", nil, j)
	if err != nil {
		return err
	}
	var resp SecretResponse
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return err
	}
	if code != 201 {
		return fmt.Errorf("HTTP error %d: returns %s", code, resp.Status)
	}
	return err
}

// GetPassword Retrieves password from Passwordstate client
func (r *Client) GetPassword(ctx context.Context, id int) (*[]SecretResponse, error) {

	raw, code, err := r.get(ctx, fmt.Sprintf("/api/passwords/%d", id), nil)
	if err != nil {
		return nil, err
	}
	var resp []SecretResponse
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d", code)
	}
	return &resp, nil
}

// SetPassword Updates password in place
func (r *Client) SetPassword(ctx context.Context, pass *Secret) (*SecretResponse, error) {
	j, err := json.Marshal(pass)
	if err != nil {
		return nil, err
	}
	raw, code, err := r.put(ctx, "/api/passwords", nil, j)
	if err != nil {
		return nil, err
	}
	var resp SecretResponse
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, resp.Status)
	}
	return &resp, err
}

// DeletePassword removes password from the passwordlist specified
func (r *Client) DeletePassword(ctx context.Context, pass *Secret) error {
	var uri string
	if pass.MoveToRecycleBin == true {
		uri = fmt.Sprintf("/api/passwords/%d?MoveToRecycleBin=%s", pass.PasswordID, "True")
	} else {
		uri = fmt.Sprintf("/api/passwords/%d?MoveToRecycleBin=%s", pass.PasswordID, "False")
	}

	_, code, err := r.delete(ctx, uri)
	if err != nil {
		return err
	}
	if code != 204 {
		return fmt.Errorf("HTTP error %d: expected status 204 No content")
	}
}
