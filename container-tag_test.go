// +build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDockerImageWithValidImage(t *testing.T) {

	foundInRegistry := containerTagExists("codefresh/cf-sendmail", "latest")

	assert.True(t, foundInRegistry, "Should be found in docker registry")
}

func TestCheckDockerImageWithFullDomain(t *testing.T) {

	foundInRegistry := containerTagExists("docker.io/codefresh/cf-sendmail", "latest")

	assert.True(t, foundInRegistry, "Should be found in docker registry")
}

func TestCheckDockerImageWithInvalidImage(t *testing.T) {

	foundInRegistry := containerTagExists("foo", "bar")

	assert.False(t, foundInRegistry, "Should be found in docker registry")
}

func TestCheckGCRImageWithValidImage(t *testing.T) {

	foundInRegistry := containerTagExists("gcr.io/cloud-builders/mvn", "3.5.0-jdk-8")

	assert.True(t, foundInRegistry, "Should be found in docker registry")
}

func TestCheckGCRImageWithoutTag(t *testing.T) {

	imageAndTag := dockerImageName{
		BaseImage: "gcr.io/cloud-builders/git",
		HasTag:    false,
		Tag:       "",
	}

	registryConnection := connectToRegistryOfImage(&imageAndTag)
	foundInRegistry := checkDockerImage(registryConnection, imageAndTag)

	assert.True(t, foundInRegistry, "Should be found in docker registry")

	assert.True(t, foundInRegistry, "Should be found in docker registry")
}

func TestUnknownRegistry(t *testing.T) {

	foundInRegistry := containerTagExists("r.cfcr.io/jbadeau/gauge-typescript-plugin", "")

	assert.False(t, foundInRegistry, "Should be found in docker registry")
}
