.PHONY: dependencies run service ui

GO := $(shell which go)
NG := $(shell which ng)
NPM := $(shell which npm)
OS := $(shell uname -s)

run: dependencies
	@cd service && go run ./cmd/itracker

dependencies: go_dep ng_dep

go_dep:
ifeq ($(strip $(GO)),)
ifeq ($(strip $(OS)),Darwin)
	@brew install go
else
	@echo "Go must be installed first. Download the latest release from https://golang.org/doc/install" && exit 1
endif
endif

ng_dep:
ifeq ($(strip $(NG)),)
ifeq ($(strip $(NPM)),)
	@npm install -g @angular/cli
else
	@echo "Angular CLI must be install first. Download the latest release from https://cli.angular.io/" && exit 1
endif
endif

lint: dependencies
	@cd service && go vet ./...
	@cd ui && ng lint

service:
	@cd service && go build ./cmd/itracker

ui:
	@cd ui && ng build
