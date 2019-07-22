# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=dyndns-route53-util
PACKAGE_NAME=function.zip
BUILD_DIR=$(shell pwd)/build

all:
build: clean
				mkdir -p $(BUILD_DIR)
				cd ./src && GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v .
clean: 
				rm -rf $(BUILD_DIR)
