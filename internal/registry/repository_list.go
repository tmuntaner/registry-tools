package registry

import (
	"encoding/json"
	"fmt"
	regHTTP "github.com/tmuntaner/registry-tools/internal/http"
	"regexp"
)

type repoListResponse struct {
	Repositories []string `json:"repositories"`
}

// RepositoryList returns a list of repositories from a docker registry.
func RepositoryList(repo string) ([]string, error) {

	client := regHTTP.RegistryHTTPClient{}
	url := fmt.Sprintf("%s/v2/_catalog?n=500", repo)
	var repositories []string

	for {
		_, headers, body, err := client.Get(url)
		if err != nil {
			return repositories, err
		}

		var repoListResponse repoListResponse
		err = json.Unmarshal(body, &repoListResponse)
		if err != nil {
			return repositories, err
		}

		repositories = append(repositories, repoListResponse.Repositories...)
		linkHeader := headers.Get("Link")
		if linkHeader == "" {
			return repositories, nil
		}
		next, err := parseLink(linkHeader)
		if err != nil {
			return repositories, err
		}
		url = repo + next
	}
}

func parseLink(link string) (string, error) {

	regExp := regexp.MustCompile(`\<(?P<url>.*)\>`)
	result := make(map[string]string)
	matches := regExp.FindStringSubmatch(link)
	if len(matches) < 1 {
		return "", fmt.Errorf("could not parse next page with header: %s", link)
	}

	for i, name := range regExp.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}

	return result["url"], nil
}
