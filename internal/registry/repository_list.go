package registry

import (
	"encoding/json"
	"fmt"
	regHTTP "github.com/tmuntaner/registry-tools/internal/http"
)

type RepoListResponse struct {
	Repositories []string `json:"repositories"`
}

func RepositoryList(repo string) ([]string, error) {

	url := fmt.Sprintf("%s/v2/_catalog", repo)

	_, _, body, err := regHTTP.Get(url)
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
