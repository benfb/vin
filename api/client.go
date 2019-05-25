package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Client is a wrapper aroudn an HTTP client for the MLB API
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

// NewClient is a constructor for HTTP clients
func NewClient() *Client {
	baseURL, _ := url.Parse("https://statsapi.mlb.com/api/v1/")

	return &Client{
		BaseURL:    baseURL,
		UserAgent:  "Vin/2.0 (https://github.com/benfb/vin)",
		httpClient: &http.Client{},
	}
}

// NewRequest creates a new request with query string parameters
func (c *Client) NewRequest(method, path string, params map[string]string) (*http.Request, error) {
	newURL, _ := url.Parse(c.BaseURL.String() + path)
	q := newURL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req, err := http.NewRequest(method, newURL.String(), nil)
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

// Do performs a request on the MLB API
func (c *Client) Do(req *http.Request, target interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(target)
	return resp, err
}
