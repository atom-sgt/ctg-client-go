package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10, // TODO: Move to configs
		},
	}
}

func (c *Client) GetStudies() (*StudyList, error) {
	url := fmt.Sprintf("%s/studies", c.BaseURL)

	resp, err := c.doRequest("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var studies StudyList
	err = json.NewDecoder(resp.Body).Decode(&studies)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &studies, nil
}

func (c *Client) GetStudy(nctID string) (*Study, error) {
	// TODO
	return nil, fmt.Errorf("Unimplemented")
}

func (c *Client) doRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil) // nil body for GET
	if err != nil {
		return nil, fmt.Errorf("failed to create request.  %s %s: %w", method, url, err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make %s request.  %s: %w", method, url, err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("API returned bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return resp, nil
}
