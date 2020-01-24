package registry

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tmuntaner/registry-tools/internal/parser"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTagList(t *testing.T) {
	expected := []string{"foo", "bar"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(tagListResponse{Tags: expected})
		w.WriteHeader(200)
		_, _ = w.Write(body)
	}))

	defer ts.Close()

	image := parser.DockerImage{Host: ts.URL}
	tags, err := TagList(image)
	assert.Nil(t, err)
	assert.Equal(t, expected, tags)
}

func TestTagListWithError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer ts.Close()

	image := parser.DockerImage{Host: ts.URL}
	_, err := TagList(image)
	assert.NotNil(t, err)
}

func TestTagListWithBadJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("foo"))
	}))

	defer ts.Close()

	image := parser.DockerImage{Host: ts.URL}
	_, err := TagList(image)
	assert.NotNil(t, err)
}
