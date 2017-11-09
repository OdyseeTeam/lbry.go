#!/bin/bash

set -euxo pipefail
echo "Building protobuf files"
rm -rf pb/*.pb.go
protoc --go_out=. pb/*.proto
go build ./...
go build ./cli/lbryschema-cli.go
go build -o lbryschema-python-binding.so -buildmode=c-shared ./binding/lbryschema-python-binding.go
