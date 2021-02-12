package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type gitHubAssetDetails struct {
	assetName    string
	assetURL     string
	assetVersion string
}

func findLatestRelease(gitHubUser string, gitHubRepository string, gitHubAssetName string) (latestVersion gitHubAssetDetails, err error) {
	assetDetails := gitHubAssetDetails{
		assetName: gitHubAssetName,
	}

	fmt.Println("Looking for " + gitHubRepository + " at " + gitHubUser)
	url := "https://api.github.com/repos/" + gitHubUser + "/" + gitHubRepository + "/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return assetDetails, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		return assetDetails, errors.New("Could not access " + url)
	}

	fmt.Println("Response status:", resp.Status)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Println(newStr)

	assetDetails.assetURL = "lala"
	assetDetails.assetVersion = "1.2.3"

	return assetDetails, nil
}
