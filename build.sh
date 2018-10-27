#!/bin/sh

make build

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 

make pull-image 