package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Fetch(url string, method string, params map[string]string, headers map[string]string) (*http.Response, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
