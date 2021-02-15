package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
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

	url := "https://api.github.com/repos/" + gitHubUser + "/" + gitHubRepository + "/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return assetDetails, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		return assetDetails, errors.New("Could not access " + url)
	}

	log.Println("Response status of api.github.com:", resp.Status)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return assetDetails, err
	}

	releaseResp := releaseResponse{}
	err = json.Unmarshal(buf.Bytes(), &releaseResp)
	if err != nil {
		return assetDetails, err
	}

	assetDetails.assetURL = filterAssets(releaseResp, assetDetails.assetName)

	assetDetails.assetVersion = releaseResp.Name

	return assetDetails, nil
}

func filterAssets(releasesFound releaseResponse, assetName string) (assetURL string) {

	//No assets were published in GitHub for this project. No need to look further
	if len(releasesFound.Assets) == 0 {
		return ""
	}

	for _, possibleAsset := range releasesFound.Assets {
		if strings.Contains(possibleAsset.Name, assetName) {
			return possibleAsset.DownloadURL
		}
	}

	return ""
}
