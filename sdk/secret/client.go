package secret

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Nilsas/passwordstate-sdk-go/sdk/request"
)

type Client interface {
	AddSecret(ctx context.Context, pass *Secret) (*SecretResponse, error)
	GetSecret(ctx context.Context, id int) (*[]SecretResponse, error)
	UpdateSecret(ctx context.Context, pass *Secret) (*[]SecretResponse, error)
}

type secretClient struct {
	r *request.Client
}

func NewClient(r *request.Client) Client {
	return &secretClient{r}
}

// AddSecret Adds new password to Passwordstate client
func (secret *secretClient) AddSecret(ctx context.Context, data *Secret) (*SecretResponse, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	raw, code, err := secret.r.Post(ctx, "/api/passwords", nil, j)
	if err != nil {
		return nil, err
	}
	var resp SecretResponse
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if code != 201 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, resp.Status)
	}
	return &resp, nil
}

// GetSecret Retrieves password from Passwordstate client
func (secret *secretClient) GetSecret(ctx context.Context, id int) (*[]SecretResponse, error) {
	raw, code, err := secret.r.Get(ctx, fmt.Sprintf("/api/passwords/%d", id), nil)
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

// UpdateSecret Updates secret values in Passwordstate client
func (secret *secretClient) UpdateSecret(ctx context.Context, data *Secret) (*[]SecretResponse, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	raw, _, err := secret.r.Put(ctx, "/api/passwords", nil, j)
	if err != nil {
		return nil, err
	}
	var resp []SecretResponse
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecret Removed the secret from Passwordstate
func (secret *secretClient) DeleteSecret(ctx context.Context, id int) error {
	raw, code, err := secret.r.Delete(ctx, fmt.Sprintf("/api/passwords/%d", id))
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, err)
	if err != nil {
		return err
	}
	if code != 204 {
		return fmt.Errorf("HTTP error %d: %s", code, err)
	}
	return nil
}