package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type assetToBeDockerizedDetails struct {
	Version  string `json:"version"`
	Filename string `json:"filename"`
}

//OutputFile is the file that contains the http response as is
const OutputFile string = "asset-details.json"

func main() {
	baseDockerImage := flag.String("image", "", "<base container image>")
	gitHubRepo := flag.String("repo", "", "<GitHub repository>")
	assetName := flag.String("asset", "", "<GitHub release asset name>")
	mode := flag.String("mode", "check-and-download", "[check] only for latest version, [download] always latest version, [check-and-download] to download only if newer")

	flag.Parse()

	if *baseDockerImage == "" || *gitHubRepo == "" || *assetName == "" {
		log.Println("Missing arguments. -repo, -asset and -image are required. Use -h for full syntax")
		os.Exit(1)
	}

	slashesFound := strings.Count(*gitHubRepo, "/")
	if slashesFound != 1 {
		log.Println("Wrong GitHub repository format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	userAndRepo := strings.Split(*gitHubRepo, "/")

	if len(userAndRepo) < 2 || userAndRepo[0] == "" || userAndRepo[1] == "" {
		log.Println("Wrong GitHub repository format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	latestReleaseDetails, err := findLatestRelease(userAndRepo[0], userAndRepo[1], *assetName)

	if err != nil {
		log.Fatal(err)
	}

	if latestReleaseDetails.assetURL == "" {
		log.Fatal("No asset found for version " + latestReleaseDetails.assetVersion)
	}

	log.Printf("Latest release is %s, %s\n", latestReleaseDetails.assetVersion, latestReleaseDetails.assetURL)

	//We were asked to only check for last github release and do nothing else
	if *mode == "check" {
		os.Exit(0)
	}

	//We were asked to download the latest release regardless of whether a Docker image exists for it.
	if *mode == "download" {
		_ = fetchURLToLocalFile(latestReleaseDetails.assetURL)
		os.Exit(0)
	}

	//Since we are here we will download release only if there isn't a Docker image for it
	foundInRegistry := containerTagExists(*baseDockerImage, latestReleaseDetails.assetVersion)

	if foundInRegistry {
		log.Printf("Found existing container image %s:%s. Nothing to do, exiting", *baseDockerImage, latestReleaseDetails.assetVersion)
		os.Exit(0)
	}

	log.Printf("Missing container image %s:%s\n", *baseDockerImage, latestReleaseDetails.assetVersion)

	localFileName := fetchURLToLocalFile(latestReleaseDetails.assetURL)

	details := assetToBeDockerizedDetails{
		Version:  latestReleaseDetails.assetVersion,
		Filename: localFileName,
	}
	jsonValue, _ := json.Marshal(details)

	err = ioutil.WriteFile(OutputFile, jsonValue, 0700)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved response to " + OutputFile)

}

func fetchURLToLocalFile(url string) (localFileName string) {
	localFilePath := path.Base(url)
	log.Printf("Downloading %s to ./app/%s\n", url, localFilePath)

	err := downloadFile(url, "./app/"+localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Downloaded: " + url)
	return localFilePath

}

func downloadFile(url string, targetFilepath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(targetFilepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
