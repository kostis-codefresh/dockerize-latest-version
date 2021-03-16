# Create Dockerhub images from Github releases



[![Go Report Card](https://goreportcard.com/badge/github.com/kostis-codefresh/dockerize-latest-version)](https://goreportcard.com/report/github.com/kostis-codefresh/dockerize-latest-version)
[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/kostis-codefresh/my-plugins%2Fdockerize-github-release?type=cf-1&key=eyJhbGciOiJIUzI1NiJ9.NWIwZmYzYmE1ODAzMWUwMDAxYjJlOGUw.dFYNhKzaLSj6l3LoOWe0DlGiuY0McdrmrgHWtWNC9WE)]( https://g.codefresh.io/pipelines/edit/new/builds?id=6026be2d0ec1b47e060d3980&pipeline=dockerize-github-release&projects=my-plugins&projectId=6026be05d65b217adde97a71)

This is a mini CLI + Container image to check if a github release is newer than the respective Dockerhub image.
If it is then it downloads the github asset so that you can create a Docker image out of it.

## How to build

Run:

 *  `go build` to get the executable OR
 *  `docker build . -t dockerize-latest-version` to create a container image if you prefer docker instead or don't have access to a Go dev environment

A prebuilt image is already available at [https://hub.docker.com/r/kostiscodefresh/dockerize-latest-version](https://hub.docker.com/r/kostiscodefresh/dockerize-latest-version)

## How to use the CLI

This CLI can be used like

Run `dockerize-latest-release --repo argoproj/argo-rollouts -asset kubectl-argo-rollouts-linux-amd64 -image kostiscodefresh/argo-rollouts-cli `. This will fetch the github details and save them to a file called `asset-details.json`. It will also download `kubectl-argo-rollouts-linux-amd64` locally

You can use this CLI in any CI/CD system and the Dockerhub image in any container based pipeline.

## Codefresh example

See an example for Codefresh at [codefresh-example.yml](codefresh-example.yml)

Currently it runs every 6 hours and pushes images to [https://hub.docker.com/r/kostiscodefresh/glooctl/](https://hub.docker.com/r/kostiscodefresh/glooctl/) and [https://hub.docker.com/r/kostiscodefresh/kubectl-argo-rollouts](https://hub.docker.com/r/kostiscodefresh/kubectl-argo-rollouts)
