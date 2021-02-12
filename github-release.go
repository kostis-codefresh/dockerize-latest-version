package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type gitHubAssetDetails struct {
	assetName    string
	assetURL     string
	assetVersion string
}

type Asset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}

type releaseResponse struct {
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
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

	releaseResp := releaseResponse{}
	json.Unmarshal(buf.Bytes(), &releaseResp)

	assetDetails.assetURL = releaseResp.Assets[0].DownloadURL
	assetDetails.assetVersion = releaseResp.Name

	return assetDetails, nil
}
