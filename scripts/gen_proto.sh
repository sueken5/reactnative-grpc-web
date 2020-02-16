#!/bin/bash

set -euxo pipefail

# gen
mkdir -p "server/apis/"
protoc --go_out=plugins=grpc:./server/apis "proto/ping/ping.proto"
mkdir -p "view/src/apis/"
protoc --js_out=import_style=commonjs:./view/src/apis --grpc-web_out=import_style=typescript,mode=grpcwebtext:./view/src/apis "proto/ping/ping.proto"
