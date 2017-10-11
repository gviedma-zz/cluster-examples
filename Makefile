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

build-member:
	cd member && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . && \
	docker build -t cluster-example-member -f Dockerfile.scratch .	

build-all : build-seed build-member
	echo done

run-seed:
	docker run --env CONSUL_HTTP_ADDR=172.18.0.2:8500 --net "ad" --rm --name cluster-example-seed cluster-example-seed

run-member:
	docker run --env CONSUL_HTTP_ADDR=172.18.0.3:8500 --net "ad" --rm --name cluster-example-member cluster-example-member

run-infra:
	@echo "Setting up mysql/consul" 
	./infrasetup.sh

run-all: run-infra run-seed run-member



