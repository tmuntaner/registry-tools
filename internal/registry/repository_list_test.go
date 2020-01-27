package registry

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepositoryList(t *testing.T) {
	expected := []string{"foo", "bar"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(repoListResponse{Repositories: expected})
		w.WriteHeader(200)
		_, _ = w.Write(body)
	}))

	defer ts.Close()

	repositories, err := RepositoryList(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, expected, repositories)
}

func TestRepositoryListMultipage(t *testing.T) {
	expected := []string{"foo", "bar"}
	mux := http.NewServeMux()
	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(repoListResponse{Repositories: []string{"bar"}})
		w.WriteHeader(200)
		_, _ = w.Write(body)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(repoListResponse{Repositories: []string{"foo"}})
		w.Header().Add("Link", "</2>;")
		w.WriteHeader(200)
		_, _ = w.Write(body)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	repositories, err := RepositoryList(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, expected, repositories)
}

func TestRepositoryListMultipageWithError(t *testing.T) {
	expected := []string{"foo", "bar"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(repoListResponse{Repositories: expected})
		w.Header().Add("Link", "<foo")
		w.WriteHeader(200)
		_, _ = w.Write(body)
	}))

	defer ts.Close()

	repositories, err := RepositoryList(ts.URL)
	assert.NotNil(t, err)
	assert.Equal(t, expected, repositories)
}

func TestRepositoryListWithError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer ts.Close()

	_, err := RepositoryList(ts.URL)
	assert.NotNil(t, err)
}

func TestRepositoryListWithBadJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("foo"))
	}))

	defer ts.Close()

	_, err := RepositoryList(ts.URL)
	assert.NotNil(t, err)
}
