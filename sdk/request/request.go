package request

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DefaultHTTPClient initialized Passwordstate with appropriate conditions.
// It allows you globally redefine HTTP client.
var DefaultHTTPClient = http.DefaultClient

// Client uses Passwordstate REST API for interacting with Passwordstate server.
type Client struct {
	baseURL   string
	Key       string
	basicAuth bool
	client    *http.Client
}

// NewClient initializes client for interacting with an instance of Passwordstate server;
// apiKey accepts Passwordstate API key
func NewClient(apiURL, apiKey string, client *http.Client) *Client {
	return &Client{
		baseURL: apiURL,
		Key:     apiKey,
		client:  client,
	}
}

func (r *Client) Get(ctx context.Context, query string, params url.Values) ([]byte, int, error) {
	return r.doRequest(ctx, "GET", query, params, nil)
}

func (r *Client) Patch(ctx context.Context, query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest(ctx, "PATCH", query, params, bytes.NewBuffer(body))
}

func (r *Client) Put(ctx context.Context, query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest(ctx, "PUT", query, params, bytes.NewBuffer(body))
}

func (r *Client) Post(ctx context.Context, query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest(ctx, "POST", query, params, bytes.NewBuffer(body))
}

func (r *Client) Delete(ctx context.Context, query string) ([]byte, int, error) {
	return r.doRequest(ctx, "DELETE", query, nil, nil)
}

func (r *Client) doRequest(ctx context.Context, method, query string, params url.Values, buf io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest(method, r.baseURL+query, buf)
	if err != nil {
		return nil, 0, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("APIKey", r.Key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autograf")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}
