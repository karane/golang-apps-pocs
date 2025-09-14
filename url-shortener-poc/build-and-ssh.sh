#!/bin/bash
set -e

if [ $# -lt 2 ]; then
  echo "Usage: $0 <Dockerfile path or directory> <image name>"
  exit 1
fi

INPUT_PATH=$1
IMAGE_NAME=$2
CONTAINER_NAME="${IMAGE_NAME}-container"

# If INPUT_PATH is a directory, assume Dockerfile inside
if [ -d "$INPUT_PATH" ]; then
  DOCKERFILE_PATH="$INPUT_PATH/Dockerfile"
  BUILD_CONTEXT="$INPUT_PATH"
else
  DOCKERFILE_PATH="$INPUT_PATH"
  BUILD_CONTEXT=$(dirname "$INPUT_PATH")
fi

# Build image
docker build -t "$IMAGE_NAME" -f "$DOCKERFILE_PATH" "$BUILD_CONTEXT"

# Remove old container if exists
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
  docker rm -f "$CONTAINER_NAME" >/dev/null 2>&1 || true
fi

# Run container (keep alive with tail)
docker run -d --name "$CONTAINER_NAME" "$IMAGE_NAME" tail -f /dev/null

# Try bash first, fallback to sh
if docker exec "$CONTAINER_NAME" bash -c "echo" >/dev/null 2>&1; then
  docker exec -it "$CONTAINER_NAME" bash
else
  docker exec -it "$CONTAINER_NAME" sh
fi
