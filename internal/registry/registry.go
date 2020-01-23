package registry

import (
	"encoding/json"
	"fmt"
	http "github.com/tmuntaner/registry-tools/internal/http"
	"github.com/tmuntaner/registry-tools/internal/parser"
)

var httpClient http.RegistryHttpClient

type Client struct {}

type RepoListResponse struct {
	Repositories []string `json:"repositories"`
}

func (c Client) RepositoryList(repo string) ([]string, error) {

	url := fmt.Sprintf("%s/v2/_catalog", repo)

	_, _, body, err := httpClient.Get(url)
	if err != nil {
		return []string{}, err
	}

	var repoListResponse RepoListResponse
	err = json.Unmarshal(body, &repoListResponse)
	if err != nil {
		return []string{}, err
	}

	return repoListResponse.Repositories, err
}

type TagListResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (c Client) TagList(image parser.DockerImage) ([]string, error) {

	url := fmt.Sprintf("%s/v2/%s/tags/list", image.Host, image.Image)

	_, _, body, err := httpClient.Get(url)
	if err != nil {
		return []string{}, err
	}

	var tagListResponse TagListResponse
	err = json.Unmarshal(body, &tagListResponse)
	if err != nil {
		return []string{}, err
	}

	return tagListResponse.Tags, err
}
