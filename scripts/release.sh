#!/usr/bin/env bash

DISCOVERY_VERSION=v0.0.1

go generate ./cmd/discovery/main.go

env GOOS=linux GOARCH=amd64 go build -o discovery-linux-$DISCOVERY_VERSION ./cmd/discovery
# env GOOS=freebsd GOARCH=amd64 go build -o discovery-freebsd-$DISCOVERY_VERSION ./cmd/discovery
# env GOOS=darwin GOARCH=amd64 go build -o discovery-macos-$DISCOVERY_VERSION ./cmd/discovery
# env GOOS=windows GOARCH=amd64 go build -o discovery-windows-$DISCOVERY_VERSION.exe ./cmd/discovery

env GOOS=linux GOARCH=amd64 go build -o discovery-cli-linux-$DISCOVERY_VERSION ./cmd/cli
# env GOOS=freebsd GOARCH=amd64 go build -o discovery-cli-freebsd-$DISCOVERY_VERSION ./cmd/cli
# env GOOS=darwin GOARCH=amd64 go build -o discovery-cli-macos-$DISCOVERY_VERSION ./cmd/cli
# env GOOS=windows GOARCH=amd64 go build -o discovery-cli-windows-$DISCOVERY_VERSION.exe ./cmd/cli