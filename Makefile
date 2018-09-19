override BINARY=gorest
VERSION=V1.0.0
API_ENV?=prod

clean: ## Clean up all the resources
	rm -rf build
	go clean -r

compile: ## Compile the project to generate the binary in the build folder
ifndef VERSION
	$(error VERSION environment variable has to be set)
endif
	GO15VENDOREXPERIMENT=1 go build -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=`git rev-parse HEAD` -X main.BuildTime=`date +%FT%T%z`" -o build/${BINARY}

run: ## run the application
	./build/${BINARY}

setup: logdir config ## Installs gorest

config: ## setup configuration file for the expected environment
ifndef API_ENV
	$(error "API_ENV variable not defined, valid value : local,prod,test")
endif
	cp  ${CURDIR}/config.json.${API_ENV} ${CURDIR}/config.json

logdir: ## creates log directory for  logs
	mkdir -p /var/log/${BINARY}


help: ## You can always run this command to see whau options are available to you while running the make command
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help compile config clean logdir run
.DEFAULT_GOAL := help