package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	baseDockerImage := flag.String("image", "", "<base container image>")
	gitHubRepo := flag.String("repo", "", "<GitHub repository>")
	assetName := flag.String("asset", "", "<GitHub release asset name>")
	mode := flag.String("mode", "check-and-download", "[check] only for latest version, [download] always latest version, [check-and-download] to download only if newer")

	flag.Parse()

	if *baseDockerImage == "" || *gitHubRepo == "" || *assetName == "" {
		fmt.Println("Missing arguments. -repo, -asset and -image are required. Use -h for full syntax")
		os.Exit(1)
	}

	slashesFound := strings.Count(*gitHubRepo, "/")
	if slashesFound != 1 {
		fmt.Println("Wrong GitHub repository format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	userAndRepo := strings.Split(*gitHubRepo, "/")

	if len(userAndRepo) < 2 || userAndRepo[0] == "" || userAndRepo[1] == "" {
		fmt.Println("Wrong GitHub repository format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	fmt.Printf("Docker is %s, %s\n", *baseDockerImage, *mode)
	latestReleaseDetails, err := findLatestRelease(userAndRepo[0], userAndRepo[1], *assetName)

	if err != nil {
		log.Fatal(err)
	}

	if latestReleaseDetails.assetURL == "" {
		log.Fatal("No asset found for version " + latestReleaseDetails.assetVersion)
	}

	fmt.Printf("Latest release is %s, %s\n ", latestReleaseDetails.assetVersion, latestReleaseDetails.assetURL)

	foundInRegistry := containerTagExists(*baseDockerImage, latestReleaseDetails.assetVersion)

	if foundInRegistry {
		fmt.Printf("Found existing container image %s:%s. Nothing to do, exiting", *baseDockerImage, latestReleaseDetails.assetVersion)
		os.Exit(0)
	}

	fmt.Printf("Missing container image %s:%s\n", *baseDockerImage, latestReleaseDetails.assetVersion)
	localFilePath := path.Base(latestReleaseDetails.assetURL)
	fmt.Printf("Downloading %s to ./%s\n", latestReleaseDetails.assetURL, localFilePath)

	err = downloadFile(latestReleaseDetails.assetURL, localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Downloaded: " + latestReleaseDetails.assetURL)

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
