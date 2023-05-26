.PHONY: *
SHELL = /usr/bin/env bash

upgrade:
	./upgrade.sh

docs:
	go run cmd/docs/main.go

docsite:
	go run cmd/docs/main.go -out-dir ../speakeasy-registry/web/packages/docsite/docs/speakeasy-cli -doc-site