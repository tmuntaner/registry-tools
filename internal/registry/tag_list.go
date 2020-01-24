package registry

import (
	"encoding/json"
	"fmt"
	regHTTP "github.com/tmuntaner/registry-tools/internal/http"
	"github.com/tmuntaner/registry-tools/internal/parser"
)

type tagListResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// TagList returns a list of tags from a repository.
func TagList(image parser.DockerImage) ([]string, error) {

	url := fmt.Sprintf("%s/v2/%s/tags/list", image.Host, image.Image)

	client := regHTTP.RegistryHTTPClient{}
	_, _, body, err := client.Get(url)
	if err != nil {
		return []string{}, err
	}

	var tagListResponse tagListResponse
	err = json.Unmarshal(body, &tagListResponse)
	if err != nil {
		return []string{}, err
	}

	return tagListResponse.Tags, err
}
