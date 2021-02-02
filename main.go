package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Github repository argument missing. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	gitHubRepo := os.Args[1]
	slashesFound := strings.Count(gitHubRepo, "/")
	if slashesFound != 1 {
		fmt.Println("Wrong argument format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

	userAndRepo := strings.Split(gitHubRepo, "/")

	if len(userAndRepo) < 2 || userAndRepo[0] == "" || userAndRepo[1] == "" {
		fmt.Println("Wrong argument format. Use <github-user/github-repo>. Example `codefresh-io/cli`")
		os.Exit(1)
	}

}
