#!/bin/sh

#
# Copyright (c)  by Saurav from 2022
#

APP_NAME=go-app
VERSION=latest
TAG=$APP_NAME":"$VERSION
DOCKER_FILENAME=Dockerfile_dev
HOST_PORT=8083
CONTAINER_PORT=8083
DEBUGGER_PORT=2345

#docker image rm $(docker images)
docker build -t $TAG -f $DOCKER_FILENAME .
docker run --rm -it -v go-vol:/usr/storage \
  -d -p $HOST_PORT:$CONTAINER_PORT \
  -p $DEBUGGER_PORT:$DEBUGGER_PORT \
  -e APP_ADDRESS=java-app:8081 \
  --network myapp \
  --name $APP_NAME \
  $TAG