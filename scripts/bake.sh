#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <language>"
    exit
fi

language=$1

pushd runner
    go get
    go build
popd

mkdir -p boj/dockerfiles/$language
cp runner/runner boj/dockerfiles/$language
jinja2 --format yaml\
    boj/template/Dockerfile.jinja\
    boj/runners/$language.yaml > boj/dockerfiles/$language/Dockerfile
