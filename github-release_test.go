package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealExample1(t *testing.T) {

	assetFound := readJSON("testdata/codefresh-cli.json", "macos-x64.tar.gz")

	assert.Equal(t, "https://github.com/codefresh-io/cli/releases/download/v0.74.9/codefresh-v0.74.9-macos-x64.tar.gz", assetFound, "Found incorrect asset URL")
}

func TestRealExample2(t *testing.T) {

	assetFound := readJSON("testdata/pulumi.json", "my-pulumi")

	assert.Equal(t, "", assetFound, "Found incorrect asset URL")
}

func TestNoAssets(t *testing.T) {

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
	json.Unmarshal(jsonData, &releaseResp)
	return filterAssets(releaseResp, assetName)

}
