#!/bin/bash

TARGET=$1
if [ -z $TARGET ]
then
    TARGET="build"
fi

CGO_ENABLED=0
#GOOS=linux

case $TARGET in
    build)
        rm -rf dist
        go build -a -installsuffix cgo -o dist/goGate main.go
        ;;
    buildstatic)
        rm -rf dist
        mkdir -p dist/static
        go build -ldflags "-s" -a -tags netgo -installsuffix netgo -o dist/static/goGate main.go
        ;;
    builddocker)
        rm -rf dist
        docker build -t "winkingzhang/gogate" -f Dockerfile .
        ;;
    builddockeralpine)
        rm -rf dist
        mkdir -p dist/alpine
        go build -a -tags netgo -installsuffix netgo -o dist/alpine/goGate main.go
        docker build -t "winkingzhang/gogate:alpine" -f Dockerfile.alpine .
        ;;
    builddockerstatic)
        rm -rf dist
        mkdir -p dist/static
        go build -ldflags "-s" -a -tags netgo -installsuffix netgo -o dist/static/goGate main.go
        docker build -t "winkingzhang/gogate:static" -f Dockerfile.static .
        ;;
    *)
        echo $"Usage: $0 {build|buildstatic|builddocker|builddockeralpine|builddockerstatic}"
        exit 1
esac