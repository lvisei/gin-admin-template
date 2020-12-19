.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v7.0.0

APP = gin-admin-template
SERVER_BIN = ./cmd/${APP}
RELEASE_ROOT = release
RELEASE_SERVER = release/${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

all: start

build:
	@go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN) ./cmd/${APP}

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN) ./cmd/${APP}

start:
	go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" cmd/${APP}/main.go web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml

swagger:
	swag init --parseDependency --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger

wire:
	wire gen ./internal/app

test:
	@go test -v $(shell go list ./...)

clean:
	rm -rf data release $(SERVER_BIN) internal/app/test/data cmd/${APP}/data

out-packages:
	go list -m -u all

pack: build
	rm -rf $(RELEASE_ROOT) && mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && tar -cvf $(APP).tar ${APP} && rm -rf ${APP}

pack-linux: build-linux
	rm -rf $(RELEASE_ROOT) && mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && tar -cvf $(APP).tar ${APP} && rm -rf ${APP}
