package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	baseDockerImage := flag.String("image", "", "<base container image>")
	gitHubRepo := flag.String("repo", "", "<GitHub repository>")
	checkOnly := flag.Bool("check", false, "[Do not download the release just check if there is a new one]")
	flag.Parse()
	if *baseDockerImage == "" || *gitHubRepo == "" {
		fmt.Println("Missing arguments. -image and -repo are required. Use -h for full syntax")
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

	fmt.Printf("Docker is %s, %v\n", *baseDockerImage, *checkOnly)
	findLatestRelease(userAndRepo[0], userAndRepo[1])

}
