# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
PACKAGE_NAME=function.zip
BUILD_DIR=$(shell pwd)/build

all:
build: clean
				mkdir -p $(BUILD_DIR)
				cd ./src && GOOS=linux $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v .
clean: 
				rm -rf $(BUILD_DIR)
