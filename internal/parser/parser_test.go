package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var gunTests = []struct {
	gun         string
	registryURL string
	expected    DockerImage
}{
	{
		"foo",
		"",
		DockerImage{
			Host:  "registry-1.docker.io",
			Image: "library/foo",
			Tag:   "",
		},
	},
	{
		"foo:bar",
		"",
		DockerImage{
			Host:  "registry-1.docker.io",
			Image: "library/foo",
			Tag:   "bar",
		},
	},
	{
		"tmuntan1/bar",
		"",
		DockerImage{
			Host:  "registry-1.docker.io",
			Image: "tmuntan1/bar",
			Tag:   "",
		},
	},
	{
		"tmuntan1/bar:foo",
		"",
		DockerImage{
			Host:  "registry-1.docker.io",
			Image: "tmuntan1/bar",
			Tag:   "foo",
		},
	},
	{
		"registry.example.com/tmuntan1/bar",
		"",
		DockerImage{
			Host:  "registry.example.com",
			Image: "tmuntan1/bar",
			Tag:   "",
		},
	},
	{
		"registry.example.com/tmuntan1/bar:foo",
		"",
		DockerImage{
			Host:  "registry.example.com",
			Image: "tmuntan1/bar",
			Tag:   "foo",
		},
	},
	{
		"bar:foo",
		"registry.example.com",
		DockerImage{
			Host:  "registry.example.com",
			Image: "bar",
			Tag:   "foo",
		},
	},
	{
		"tmuntan1/bar:foo",
		"registry.example.com",
		DockerImage{
			Host:  "registry.example.com",
			Image: "tmuntan1/bar",
			Tag:   "foo",
		},
	},
	{
		"registry.example.com/tmuntan1/bar:foo",
		"registry.example.com",
		DockerImage{
			Host:  "registry.example.com",
			Image: "tmuntan1/bar",
			Tag:   "foo",
		},
	},
}

func TestGunToImage(t *testing.T) {

	for _, test := range gunTests {
		result, err := GunToImage(test.gun, test.registryURL)

		assert.Nil(t, err)
		assert.Equal(t, result.Image, test.expected.Image)
		assert.Equal(t, result.Host, test.expected.Host)
		assert.Equal(t, result.Tag, test.expected.Tag)
	}
}
