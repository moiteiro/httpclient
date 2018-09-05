package httpclient

import (
	"encoding/json"
	"net/http"
)

// PostJSON performs a basic http POST request and decodes
// the JSON response into the out interface
func (c *Client) PostJSON(path string, headers map[string]string, in, out interface{}) error {
	// Marshal the in interface to a byte slice
	body, err := json.Marshal(in)
	if err != nil {
		return err
	}

	// Retrieve the bytes and decode the response
	res, err := c.PostBytes(path, headers, body)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, out)
}

// PostBytes performs a POST request using the passed path
// and body
func (c *Client) PostBytes(path string, headers map[string]string, in []byte) ([]byte, error) {
	// Execute the request and return the response
	res, err := c.Do(NewRequest(http.MethodPost, path, headers, in))
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
