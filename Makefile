BIN=bin
BIN_RELEASE=bin_release
BIN_NAME=kube-ingwatcher

IMAGE_NAME=$(BIN_NAME)
IMAGE_VERSION ?= 1.5
REMOTE_NAME=$(DOCKER_ID_USER)/$(IMAGE_NAME)

.PHONY: all fmt clean

all: $(BIN)/$(BIN_NAME)

image: tmp
	wget -O tmp/kube-ingwatcher https://github.com/nmaupu/kube-ingwatcher/releases/download/v${IMAGE_VERSION}/kube-ingwatcher_linux-amd64
	chmod +x tmp/kube-ingwatcher
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) -f Dockerfile.scratch .

tag: image
	docker tag $(IMAGE_NAME):$(IMAGE_VERSION) $(REMOTE_NAME):$(IMAGE_VERSION)

push: tag
	docker push $(REMOTE_NAME):$(IMAGE_VERSION)

push-test: tmp
	make $(BIN)/$(BIN_NAME)
	cp $(BIN)/$(BIN_NAME) tmp/kube-ingwatcher
	chmod +x tmp/kube-ingwatcher
	docker build -t $(IMAGE_NAME):test -f Dockerfile.scratch .
	docker tag $(IMAGE_NAME):test $(REMOTE_NAME):test
	docker push $(REMOTE_NAME):test

fmt:
	go fmt ./...

$(BIN)/$(BIN_NAME) build-linux: $(BIN) $(shell find . -name "*.go")
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BIN)/$(BIN_NAME) .

$(BIN)/$(BIN_NAME)-darwin build-darwin: $(BIN) $(shell find . -name "*.go")
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BIN)/$(BIN_NAME)-darwin .

clean:
	go clean -i
	rm -rf $(BIN)
	rm -rf tmp
	rm -rf bin_release

tmp:
	mkdir -p tmp

$(BIN):
	mkdir -p $(BIN)

$(BIN_RELEASE):
	mkdir -p $(BIN_RELEASE)

prepare-release: $(BIN_RELEASE)
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\" -X main.AppVersion=v$(IMAGE_VERSION)" -o bin_release/kube-ingwatcher_linux-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\" -X main.AppVersion=v$(IMAGE_VERSION)" -o bin_release/kube-ingwatcher_darwin-amd64
