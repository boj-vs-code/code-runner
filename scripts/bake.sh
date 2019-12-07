#!/bin/bash

if [ $# -lt 2 ]; then
    echo "Usage: $0 <platform> <language>"
    exit
fi

platform=$1
language=$2

dockerfile_directory=$platform/dockerfiles/$language
template_path=$platform/template/Dockerfile.jinja
runner_configuration_path=$platform/runners/$language.yaml

pushd runner
    go get
    go build
popd

mkdir -p $dockerfile_directory
cp runner/runner $dockerfile_directory
jinja2 --format yaml\
    $template_path\
    $runner_configuration_path > $dockerfile_directory/Dockerfile

image_name="moreal/code-runner-$platform-$language"
docker build $dockerfile_directory -t "$image_name"
docker push "$image_name"
