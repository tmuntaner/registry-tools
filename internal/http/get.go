package http

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type RegistryHTTPClient struct {
	Token string
}
func (c *RegistryHTTPClient) Get(url string) (int, http.Header, []byte, error) {

	statusCode, headers, body, err := httpGet(url, c.Token)
	if err != nil {
		return statusCode, headers, body, err
	} else if statusCode == 200 {
		return statusCode, headers, body, nil
	} else if statusCode != 401 {
		return statusCode, headers, body, errors.New(fmt.Sprintf("request failed with status code: %d", statusCode))
	}

	c.Token, err = tryAuth(headers)
	if err != nil {
		return statusCode, headers, body, err
	}

	return httpGet(url, c.Token)
}

func httpGet(url string, token string) (int, http.Header, []byte, error) {

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, nil, err
	}

	if token != "" {
		authHeader := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	return resp.StatusCode, resp.Header, []byte(buf.String()), err
}

