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

docker_tag=latest
docker_image_name="moreal/code-runner-$platform-$language:$docker_tag"

docker build $dockerfile_directory -t "$docker_image_name"
docker push "$docker_image_name"