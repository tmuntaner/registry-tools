package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	url2 "net/url"
	"regexp"
)

type TokenResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}

func tryAuth(headers http.Header) (string, error) {

	realm, service, scope := "", "", ""

	if headers == nil {
		// auth was needed, but there were no request headers
		return "", errors.New("auth required, but request headers were empty")
	} else {
		authHeader := headers.Get("Www-Authenticate")
		realm, service, scope = parseAuthHeader(authHeader)
	}

	if realm == "" || service == "" {
		return "", errors.New("auth required, but realm/service is null")
	}

	return getAuthToken(realm, service, scope)
}

func getAuthToken(realm string, service string, scope string) (string, error) {

	url := fmt.Sprintf("%s?service=%s&scope=%s", realm, url2.QueryEscape(service), url2.QueryEscape(scope))

	_, _, body, err := httpGet(url, "")
	if err != nil {
		return "", err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	if tokenResponse.Token != "" {
		return tokenResponse.Token, nil
	}

	return tokenResponse.AccessToken, nil
}

func parseAuthHeader(header string) (string, string, string) {

	regExp := regexp.MustCompile(`Bearer realm="(?P<realm>[^"]*)",service="(?P<service>[^"]*)"(?:,scope=")?(?P<scope>[^"]*)(?:")?`)
	result := make(map[string]string)
	matches := regExp.FindStringSubmatch(header)
	if len(matches) < 3 {
		return "", "", ""
	}

	for i, name := range regExp.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}

	return result["realm"], result["service"], result["scope"]
}
