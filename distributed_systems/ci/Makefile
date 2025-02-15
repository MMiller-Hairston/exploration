#Inspired by Google Open-Match

BASE_VERSION = 0.1.0
SHORT_SHA = $(shell git rev-parse --short=7 HEAD | tr -d [:punct:])
VERSION = $(BASE_VERSION)-$(SHORT_SHA)
# Abs local dir of the project
REPO_ROOT := $(patsubst %/,%,$(dir $(abspath $(MAKEFILE_LIST))))
BUILD_DIR = $(REPO_ROOT)/build
REGISTRY ?= ghcr.io/mmiller-hairston
TAG = $(VERSION)

# list of all folders in cmd/
CMDS = $(notdir $(wildcard cmd/*))

build/cmd: $(foreach CMD,$(CMDS),build/cmd/$(CMD)) 

$(foreach CMD,$(CMDS),build/cmd/$(CMD)): build/cmd/%: build/cmd/%/BUILD_PHONY

build/cmd/%/BUILD_PHONY:
	mkdir -p $(BUILD_DIR)/cmd/$*
	CGO_ENABLED=0 go build -v -installsuffix cgo -o $(BUILD_DIR)/cmd/$*/run ./cmd/$*/...

build/docker: $(foreach CMD, $(CMDS),build/docker/$(CMD)/BUILD_PHONY)
$(foreach CMD,$(CMDS),build/docker/$(CMD)): build/docker/%: build/docker/%/BUILD_PHONY:

build/docker/%/BUILD_PHONY:
	docker build \
		--build-arg IMAGE_TITLE=$* \
		-t $(REGISTRY)/$*:$(TAG) \
		-t $(REGISTRY)/$*:$(BASE_VERSION) \
		.