package parser

import (
	"strings"
)

type DockerImage struct {
	Host  string
	Image string
	Tag   string
}

type RegistryParser struct {}

func (c RegistryParser) GunToImage(input string, registryUrl string) (DockerImage, error) {

	tagParts := strings.SplitN(input, ":", 2)
	var gun string
	var tag string

	if len(tagParts) > 1 {
		gun = tagParts[0]
		tag = tagParts[1]
	} else {
		gun = input
		tag = ""
	}

	stringParts := strings.SplitN(gun, "/", 2)
	image := DockerImage{}

	if len(stringParts) == 1 {
		if registryUrl == "" {
			image.Host = "registry-1.docker.io"
			image.Image = "library/" + gun
		} else {
			image.Host = registryUrl
			image.Image = gun
		}
	} else if !strings.Contains(stringParts[0], ".") &&
		!strings.Contains(stringParts[0], ":") &&
		!strings.Contains(stringParts[0], "localhost") {

		image.Image = gun
		if registryUrl == "" {
			image.Host = "registry-1.docker.io"
		} else {
			image.Host = registryUrl
		}
	} else {
		image.Host = stringParts[0]
		image.Image = stringParts[1]
	}

	image.Tag = tag

	return image, nil
}
