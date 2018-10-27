#!/bin/sh

make docker-build-image

make docker-login

make docker-push-image 