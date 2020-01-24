package parser

import (
	"strings"
)

// DockerImage is a representation of a docker image with its host, image, and tag
type DockerImage struct {
	Host  string
	Image string
	Tag   string
}

// GunToImage will convert an input string of a GUN and convert it to a DockerImage struct
func GunToImage(input string, registryURL string) (DockerImage, error) {

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
		if registryURL == "" {
			image.Host = "registry-1.docker.io"
			image.Image = "library/" + gun
		} else {
			image.Host = registryURL
			image.Image = gun
		}
	} else if !strings.Contains(stringParts[0], ".") &&
		!strings.Contains(stringParts[0], ":") &&
		!strings.Contains(stringParts[0], "localhost") {

		image.Image = gun
		if registryURL == "" {
			image.Host = "registry-1.docker.io"
		} else {
			image.Host = registryURL
		}
	} else {
		image.Host = stringParts[0]
		image.Image = stringParts[1]
	}

	image.Tag = tag

	return image, nil
}
