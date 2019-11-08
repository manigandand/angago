#!/bin/bash

DOCKER_IMAGE=manigandanjeff/angago:latest
DOCKER_MNT_VOIUME_ENV=""

if [[ "$(docker images -q $DOCKER_IMAGE 2> /dev/null)" == "" ]]; then
    echo "$DOCKER_IMAGE not found. installing it..."
    docker pull $DOCKER_IMAGE 
fi

# validate the config file if provided
if [ ! -z "$1" ] ; then
    if [ ! -f $1 ]; then
        echo "config file $1 is not dound"
        exit 0;
    fi
    DOCKER_MNT_VOIUME_ENV="-v $1:/mnt$1 -e 'ANGAGO_CONFIG_PATH=/mnt$1'"
fi

echo "==> Running docker image  ..."
DOCKER_RUN_CMD="run $DOCKER_MNT_VOIUME_ENV -it $DOCKER_IMAGE"
echo $DOCKER_RUN_CMD
docker $DOCKER_RUN_CMD