#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <platform>"
    exit
fi

platform=$1
prefix_length=`expr length $platform/runners/`

for yaml in $platform/runners/*.yaml; do
    language=`expr match "$yaml" '.*\/\([a-zA-Z0-9]*\).yaml'`
    if [ $language == "template" ]; then
        continue
    fi
    echo Build $language
    ./scripts/bake.sh $platform $language
done