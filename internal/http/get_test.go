package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_,_ = w.Write([]byte("foo"))
	}))

	defer ts.Close()

	statusCode, _, body, err := Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "foo", string(body))
}

func TestGetWithMalformedUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_,_ = w.Write([]byte("foo"))
	}))

	defer ts.Close()

	_, _, _, err := Get("ht\ntp://www.foo.com")
	assert.NotNil(t, err)
}

func TestGetWithServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("foo"))
	}))

	defer ts.Close()

	statusCode, _, body, err := Get(ts.URL)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, "foo", string(body))
}

func TestGetWithError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("foo"))
	}))
	// we would need to uncomment the following line to get this test to pass. We rely on the DefaultClient not having access to the tls certs.
	// http.DefaultClient = ts.Client()

	defer ts.Close()

	_, _, _, err := Get(ts.URL)
	assert.NotNil(t, err)
}

func TestDefaultHttps(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("foo"))
	}))
	http.DefaultClient = ts.Client()

	defer ts.Close()

	url := strings.Replace(ts.URL, "https://", "", -1)
	statusCode, _, body, err := Get(url)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "foo", string(body))
}

func TestRegistryHttpClientWithToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("foo"))
	}))

	defer ts.Close()

	statusCode, _, body, err := httpGet(ts.URL, "foobar")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, "foo", string(body))
}

func TestRegistryHttpClientWith401(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(TokenResponse{Token: "foo"})
		w.Write(body)
	}))
	regServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Www-Authenticate",  `Bearer realm="`+ authServer.URL + `",service="foo-service",scope="foo-scope"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("foo"))
	}))

	defer regServer.Close()
	defer authServer.Close()

	statusCode, _, body, err := Get(regServer.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, statusCode)
	assert.Equal(t, "foo", string(body))
}

func TestRegistryHttpClientWith401Error(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("foo"))
	}))
	regServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Www-Authenticate",  `Bearer realm="`+ authServer.URL + `",service="foo-service",scope="foo-scope"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("foo"))
	}))

	defer regServer.Close()
	defer authServer.Close()

	statusCode, _, body, err := Get(regServer.URL)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnauthorized, statusCode)
	assert.Equal(t, "foo", string(body))
}
