#!/usr/bin/env bash

# MacOS
GOOS=darwin GOARCH=amd64 go build -o gif-macos main.go
# Linux 32 bit
GOOS=linux GOARCH=386 go build -o gif-linux32 main.go
# Linux 64 bit
GOOS=linux GOARCH=amd64 go build -o gif-linux64 main.go