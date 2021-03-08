package blockbook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	conn          *http.Client
	baseURL       string
	Verbose       bool
	BeforeRequest func(module, action string, param map[string]interface{}) error
	AfterRequest  func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func NewClient(baseURL string) BlockBook {
	return NewCustomized(Customization{
		Timeout: 30 * time.Second,
		BaseURL: baseURL,
	})
}

type Customization struct {
	Timeout       time.Duration
	BaseURL       string
	Verbose       bool
	BeforeRequest func(module, action string, param map[string]interface{}) error
	AfterRequest  func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func NewCustomized(config Customization) *client {
	return &client{
		conn: &http.Client{
			Timeout: config.Timeout,
		},
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

func (c *client) GetStatus() (*Status, error) {
	resp, err := c.conn.Get(c.baseURL + "/api")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var status Status
	json.Unmarshal([]byte(body), &status)
	return &status, nil
}
