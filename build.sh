#!/bin/sh

# Author : Saurav Kumar
# Copyright (c) Saurav on 27/05/2022.

APP_NAME=fossapp
VERSION=latest
TAG=$APP_NAME":"$VERSION
DOCKER_FILENAME=Dockerfile
HOST_PORT=8083
CONTAINER_PORT=8083
DEBUG_PORT=2345
APP_VOL=govol
APP_NET=myapp

#docker image rm $(docker images)
docker build -t $TAG -f $DOCKER_FILENAME .
docker run --rm -it \
  -v $APP_VOL:/var/opt/goapp \
  -d -p $HOST_PORT:$CONTAINER_PORT \
  -p $DEBUG_PORT:$DEBUG_PORT \
  --name $APP_NAME \
  -e APP_ADDRESS=spring:8081 \
  --net $APP_NET \
  $TAG



