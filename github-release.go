package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func findLatestRelease(user string, repository string) {
	fmt.Println("Looking for " + repository + " at " + user)
	url := "https://api.github.com/repos/" + user + "/" + repository + "/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Printf(newStr)
}
