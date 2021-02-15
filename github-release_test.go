// +build !integration

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubReleasePartialName(t *testing.T) {

	assetFound := readJSON("testdata/codefresh-cli.json", "macos-x64.tar.gz")

	assert.Equal(t, "https://github.com/codefresh-io/cli/releases/download/v0.74.9/codefresh-v0.74.9-macos-x64.tar.gz", assetFound, "Found incorrect asset URL")
}

func TestGitHubReleaseNoAssets(t *testing.T) {

	assetFound := readJSON("testdata/pulumi.json", "my-pulumi")

	assert.Equal(t, "", assetFound, "Found incorrect asset URL")
}

func TestGitHubReleaseExactName(t *testing.T) {

	assetFound := readJSON("testdata/solo-gloo.json", "glooctl-linux-amd64")

	assert.Equal(t, "https://github.com/solo-io/gloo/releases/download/v1.6.7/glooctl-linux-amd64", assetFound, "Found incorrect asset URL")
}

func readJSON(fileName string, assetName string) string {

	jsonFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	releaseResp := releaseResponse{}
	err = json.Unmarshal(jsonData, &releaseResp)
	if err != nil {
		log.Fatal(err)
	}
	return filterAssets(releaseResp, assetName)

}
