#!/usr/bin/env sh

# GOOS=js GOARCH=wasm go build -o main.wasm

# Build tofus

echo "Building Tofus..."
go mod tidy
go mod vendor
go build -o tofus
echo "Installing Tofus ..."
sudo rm /usr/local/bin/tofus
sudo cp tofus /usr/local/bin/tofus
echo "Tofus installed successfully!"
sudo rm tofus