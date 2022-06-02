#!/bin/sh

#
# Copyright (c)  by Saurav from 2022
#

APP_NAME=fossapp
VERSION=v12
TAG=$APP_NAME":"$VERSION
DOCKER_FILENAME=Dockerfile_dev
HOST_PORT=8081
CONTAINER_PORT=8081

#docker image rm $(docker images)
docker build -t $TAG -f $DOCKER_FILENAME .
docker run -it --rm -v go-vol:/var/opt/fossvol \
  -dp $HOST_PORT:$CONTAINER_PORT \
  --network go-network \
  --name $APP_NAME \
  $TAG