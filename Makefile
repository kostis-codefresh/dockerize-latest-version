run: 
	go run . -repo solo-io/gloo -asset glooctl-linux-amd64 -image desaegher/glooctl

test: *_test.go
	go test -v

integration:
	go test -tags=integration -v

build: *.go
	go build

.PHONY: clean

clean:
	go clean
	rm -f dockerize-latest-release