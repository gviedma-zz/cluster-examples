GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

# Just builds
all: test build

dep: glide.yaml
	glide install --strip-vendor

dep-up:
	glide up --strip-vendor

build-seed:
	cd seed && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . && \
	docker build -t cluster-example-seed -f Dockerfile.scratch .

run-seed:
	docker run --env CONSUL_HTTP_ADDR=172.18.0.2:8500 --net "ad" --rm --name cluster-example-seed cluster-example-seed

build-member:
	cd seed && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .  && \
	docker build -t cluster-example-seed -f Dockerfile.scratch .

build:  $(GOFILES)
	go build
