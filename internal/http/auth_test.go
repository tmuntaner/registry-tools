package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseAuthHeader(t *testing.T) {

	parseAuthHeaderTests := []struct {
		Header          string
		ExpectedRealm   string
		ExpectedService string
		ExpectedScope   string
	}{
		{
			Header:          `Bearer realm="http://foo.com",service="foo-service",scope="foo-scope"`,
			ExpectedRealm:   "http://foo.com",
			ExpectedService: "foo-service",
			ExpectedScope:   "foo-scope",
		},
		{
			Header:          `Bearer realm="http://foo.com",service="foo-service"`,
			ExpectedRealm:   "http://foo.com",
			ExpectedService: "foo-service",
			ExpectedScope:   "",
		},
		{
			Header:          `Bearer`,
			ExpectedRealm:   "",
			ExpectedService: "",
			ExpectedScope:   "",
		},
	}

	for _, test := range parseAuthHeaderTests {
		realm, service, scope := parseAuthHeader(test.Header)
		assert.Equal(t, test.ExpectedRealm, realm)
		assert.Equal(t, test.ExpectedService, service)
		assert.Equal(t, test.ExpectedScope, scope)
	}
}


func TestGetAuthTokenWithError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(TokenResponse{Token: "foo"})
		w.Write(body)
	}))

	defer ts.Close()

	_, err := getAuthToken(ts.URL, "", "")
	assert.NotNil(t, err)
}

func TestGetAuthTokenWithToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(TokenResponse{Token: "foo"})
		w.Write(body)
	}))

	defer ts.Close()

	token, err := getAuthToken(ts.URL, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "foo", token)
}

func TestGetAuthTokenWithAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(TokenResponse{AccessToken: "foo"})
		w.Write(body)
	}))

	defer ts.Close()

	token, err := getAuthToken(ts.URL, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "foo", token)
}

func TestTryAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(TokenResponse{AccessToken: "foo"})
		w.Write(body)
	}))

	defer ts.Close()

	headers := http.Header{}
	headers.Add("Www-Authenticate",  `Bearer realm="`+ ts.URL + `",service="foo-service",scope="foo-scope"`)

	token, err := tryAuth(headers)
	assert.Nil(t, err)
	assert.Equal(t, "foo", token)
}

func TestTryAuthWithEmptyHeaders(t *testing.T) {
	headers := http.Header{}
	_, err := tryAuth(headers)
	assert.NotNil(t, err)
}

func TestTryAuthWithNilHeaders(t *testing.T) {
	_, err := tryAuth(nil)
	assert.NotNil(t, err)
}
