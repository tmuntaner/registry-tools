package http

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type RegistryHttpClient struct {}

func (c RegistryHttpClient) Get(url string) (int, http.Header, []byte, error) {

	statusCode, headers, body, err := c.httpGet(url, "")
	if err != nil {
		return statusCode, headers, body, err
	} else if statusCode == 200 {
		return statusCode, headers, body, nil
	} else if statusCode != 401 {
		return statusCode, headers, body, errors.New(fmt.Sprintf("request failed with status code: %d", statusCode))
	}

	token, err := c.auth(url)
	if err != nil {
		return statusCode, headers, body, err
	}

	return c.httpGet(url, token)
}

func (RegistryHttpClient) httpGet(url string, token string) (int, http.Header, []byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, nil, err
	}

	if token != "" {
		authHeader := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", authHeader)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}

	body := buf.String()

	return resp.StatusCode, resp.Header, []byte(body), nil
}

