package srclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	URL string
	c   *http.Client
}

// SROpts used to initialize schema registry client
type SROpts struct {
	Timeout time.Duration
	URL     string
}

// NewSchemaRegistryClient returns a GitHelper
func NewSchemaRegistryClient(opts SROpts) *Client {
	return &Client{
		c: &http.Client{
			Timeout: opts.Timeout,
		},
		URL: opts.URL,
	}
}

func (sr *Client) getLatestSchemaURL(subject string) string {
	return fmt.Sprintf("%v/subjects/%v/versions/latest", sr.URL, subject)
}

// GetLatestSchema gets latest schema on a subject
func (sr *Client) GetLatestSchema(subject string) (*Schema, error) {
	schema := &Schema{}
	body, err := sr.get(sr.getLatestSchemaURL(subject))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &schema); err != nil {
		return nil, err
	}

	return schema, nil
}

// get returns bytes given a URL
func (sr *Client) get(uri string) (body []byte, err error) {
	req, err := http.NewRequest("GET", uri, &bytes.Buffer{})
	if err != nil {
		return body, fmt.Errorf("create new HTTP request: %v: %v", uri, err.Error())
	}

	data, err := sr.c.Do(req)
	if err != nil {
		return body, fmt.Errorf("make request error:%v: %v", uri, err.Error())
	}
	defer data.Body.Close()

	if data.StatusCode != 200 {
		data, err = sr.c.Do(req)
		if err != nil {
			return body, fmt.Errorf("make request retry error:%v: %v", uri, err.Error())
		}
		defer data.Body.Close()

		if data.StatusCode != 200 {
			return body, fmt.Errorf("make request retry error: unexpected response status %v", data.StatusCode)
		}
	}

	body, err = ioutil.ReadAll(data.Body)
	if err != nil {
		return body, fmt.Errorf("read body error: %v", err.Error())
	}

	return body, nil
}
