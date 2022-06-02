#!/bin/sh

# Author : Saurav Kumar
# Copyright (c) Saurav on 27/05/2022.

APP_NAME=fossapp
VERSION=v3
TAG=$APP_NAME":"$VERSION
DOCKER_FILENAME=Dockerfile_prod
HOST_PORT=8088
CONTAINER_PORT=8081

#docker image rm $(docker images)
docker build -t $TAG -f $DOCKER_FILENAME .
docker run --rm -v go-volume:/var/opt/goapp \
  -dp $HOST_PORT:$CONTAINER_PORT \
  --name $APP_NAME \
  $TAG



