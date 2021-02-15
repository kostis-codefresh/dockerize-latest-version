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

type asset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}

type releaseResponse struct {
	Name   string  `json:"name"`
	Assets []asset `json:"assets"`
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
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return assetDetails, err
	}
	newStr := buf.String()

	fmt.Println(newStr)

	releaseResp := releaseResponse{}
	err = json.Unmarshal(buf.Bytes(), &releaseResp)
	if err != nil {
		return assetDetails, err
	}

	//No assets were published in GitHub for this project. No need to look further
	if len(releaseResp.Assets) == 0 {
		return assetDetails, nil
	}

	assetDetails.assetURL = filterAssets(releaseResp)

	// releaseResp.Assets[0].DownloadURL
	assetDetails.assetVersion = releaseResp.Name

	return assetDetails, nil
}

func filterAssets(releasesFound releaseResponse) (assetURL string) {
	releasesFound.Name = "lala"
	return "dff"
}
