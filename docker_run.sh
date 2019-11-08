#!/bin/bash

DOCKER_IMAGE=manigandanjeff/angago

if [[ "$(docker images -q $DOCKER_IMAGE 2> /dev/null)" == "" ]]; then
    echo "$DOCKER_IMAGE not found. installing it..."
    docker pull $DOCKER_IMAGE 
fi

echo "==> Running docker image  ..."
docker run -it -p 80:80 $DOCKER_IMAGE