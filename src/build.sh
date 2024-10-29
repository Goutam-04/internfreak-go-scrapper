#!/bin/bash
mkdir -p build

echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o ../build/scrapper-windows-amd64.exe ./main.go

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o ../build/scrapper-linux-amd64 ./main.go

echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o ../build/scrapper-macos-amd64 ./main.go

echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o ../build/scrapper-macos-arm64 ./main.go

echo "All builds completed."
