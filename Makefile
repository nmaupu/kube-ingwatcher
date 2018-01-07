BIN=bin
BIN_NAME=kube-ingwatcher

IMAGE_NAME=$(BIN_NAME)
IMAGE_VERSION=1.0
REMOTE_NAME=$(DOCKER_ID_USER)/$(IMAGE_NAME)

.PHONY: all fmt vendor clean

all: $(BIN)/$(BIN_NAME)

image:
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) -f Dockerfile.scratch .

tag: image
	docker tag $(IMAGE_NAME):$(IMAGE_VERSION) $(REMOTE_NAME):$(IMAGE_VERSION)

push: tag
	docker push $(REMOTE_NAME):$(IMAGE_VERSION)

fmt:
	go fmt ./...

vendor:
	glide update -v

$(BIN)/$(BIN_NAME) build-linux: $(BIN) $(shell find . -name "*.go")
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BIN)/$(BIN_NAME) .

$(BIN)/$(BIN_NAME)-darwin build-darwin: $(BIN) $(shell find . -name "*.go")
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BIN)/$(BIN_NAME)-darwin .

clean:
	go clean -i
	rm -rf $(BIN)
	rm -rf vendor

$(BIN):
	mkdir -p $(BIN)

