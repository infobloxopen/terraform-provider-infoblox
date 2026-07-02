HOSTNAME=registry.terraform.io
NAMESPACE=infobloxopen
NAME=unified
BINARY=terraform-provider-$(NAME)
VERSION?=0.0.1
OS_ARCH=$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m)
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

default: install

build:
	go build -o $(BINARY) .

install: build
	mkdir -p ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)/$(VERSION)/$(OS_ARCH)
	cp $(BINARY) ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)/$(VERSION)/$(OS_ARCH)

clean:
	rm -f $(BINARY)

.PHONY: default build install clean

.PHONY: goimports
goimports: ## Check go imports
	@docker run --rm -v $(shell pwd):/data cytopia/goimports -w "$(GO_FILES)"
