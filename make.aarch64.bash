#!/usr/bin/env bash

mkdir -p build

PKG_CONFIG="/usr/bin/aarch64-linux-gnu-pkg-config" \
CC="aarch64-linux-gnu-gcc" CXX="aarch64-linux-gnu-g++" \
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o build/cam2ip.linux.arm64 -ldflags "-linkmode external -s -w" github.com/gen2brain/cam2ip/cmd/cam2ip
