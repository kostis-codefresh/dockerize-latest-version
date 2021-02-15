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

	assetFound := readJSON("testdata/codefresh-cli.json")

	assert.Equal(t, assetFound, "dfdf", "Found incorrect asset URL")
}

func TestRealExample2(t *testing.T) {

	assetFound := readJSON("testdata/pulumi.json")

	assert.Equal(t, assetFound, "dfdf", "Found incorrect asset URL")
}

func TestNoAssets(t *testing.T) {

	assetFound := readJSON("testdata/solo-gloo.json")

	assert.Equal(t, assetFound, "dfdf", "Found incorrect asset URL")
}

func readJSON(fileName string) string {

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
	return filterAssets(releaseResp)

}
