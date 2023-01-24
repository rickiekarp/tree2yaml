GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_PATH=build

buildTree2YamlAmd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build \
		-ldflags="-X git.rickiekarp.net/rickie/tree2yaml/generator.Version=$(shell git rev-parse HEAD)" \
		-o $(BUILD_PATH)/tree2yaml \
		tree2yaml/main.go
		
buildTree2YamlARM64v7:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 \
	go build \
		-ldflags="-X git.rickiekarp.net/rickie/tree2yaml/generator.Version=$(shell git rev-parse HEAD)" \
		-o $(BUILD_PATH)/tree2yaml \
		tree2yaml/main.go

deployTree2Yaml:
	rsync -rlvpt --delete build/tree2yaml pi:~/tools/

clean:
	rm -rf build