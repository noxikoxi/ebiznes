#!/bin/bash
IMAGE="scala-app"
TAG="latest"


if docker images --format "{{.Repository}}:{{.Tag}}" | grep -q -w "^$IMAGE:$TAG"; then
    echo "Image $IMAGE:$TAG already exists"
else
    echo "Building image $IMAGE:$TAG..."
    docker build -t $IMAGE:$TAG .
fi

echo "Running container..."
CONTAINER_ID=$(docker run -d --rm -p 9000:9000 "$IMAGE:$TAG")

if [ -z "$CONTAINER_ID" ]; then
    echo "Failed to start container"
    exit 1
fi

cleanup() {
    echo "Exit, stopping container $CONTAINER_ID..."
    docker stop "$CONTAINER_ID"
    echo "Container stopped"
}

trap cleanup EXIT

sleep 3

echo "Starting ngrok..."
ngrok http 9000