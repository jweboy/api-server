#!/bin/sh

# Build new image
sudo docker build -t api-server .

# Push image to my dockerhub repository
# sudo docker push jweboy/apiserver:latest