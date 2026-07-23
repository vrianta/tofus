#!/usr/bin/env sh

# GOOS=js GOARCH=wasm go build -o main.wasm

# Build tofus

echo "Building Tofus..."
go mod tidy
go mod vendor
go build -o tofus
echo "Installing Tofus ..."
sudo cp tofus /usr/local/bin/tofus
echo "Tofus installed successfully!"
sudo rm tofus