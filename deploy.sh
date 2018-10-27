#!/bin/sh

make docker-build-image

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 

make docker-push-image 