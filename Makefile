run: 
	go run . -repo solo-io/gloo -asset glooctl-linux-amd64 -image desaegher/glooctl

test: 
	go test -v

integration:
	go test -tags=integration -v

build:
	go build

.PHONY: clean

clean:
	go clean
	rm -f dockerize-latest-release