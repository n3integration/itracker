.PHONY: dependencies run service ui network go_dep ng_dep fabric_dep netup netdown run

BIN := $(shell test -f network/bin/cryptogen && echo "exists")
GO  := $(shell which go)
NG  := $(shell which ng)
NPM := $(shell which npm)
OS  := $(shell uname -s)
URL := "http://localhost:8020"

run: dependencies ui
	@echo
	@echo "Open a browser window to => $(URL)"
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

fabric_dep:
ifeq ($(strip $(BIN)),)
	@cd network && curl -sfL https://raw.githubusercontent.com/hyperledger/fabric/master/scripts/bootstrap.sh | bash -s -- -s -d
	@cd network && test -e bin/configtxlator || (echo "failed to download bootstrap.sh" && exit 1)
endif

vendor:
	@cd chaincode/itracker && go mod vendor

netup: fabric_dep vendor
	@cd network && echo -- -y | xargs ./byfn.sh up -c inventory -s couchdb #-n

netdown: fabric_dep
	@cd network && echo -- -y | xargs ./byfn.sh down

lint: dependencies
	@cd service && go vet ./...
	@cd ui && ng lint

service:
	@cd service && go build ./cmd/itracker

ui:
	@cd ui && ng build

clean: netdown
	@rm -rf network/{bin,config,crypto-config,connection-*} chaincode/itracker/vendor
