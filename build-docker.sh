#!/bin/sh

tag="$1"

if [ -z "$1" ]
then
    echo "Must specify a tag"
    exit 1
fi

# build for docker scratch image
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./app .

# build the image
docker build -t kubetest:latest -t kubetest:$tag .

# clean up
rm ./app