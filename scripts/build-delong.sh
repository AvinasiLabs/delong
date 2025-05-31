#!/usr/bin/env bash
#
# Build and push delong Docker image using buildx
# Usage: ./build-delong.sh [TAG] [DOCKERFILE_PATH]
#

set -euo pipefail

# Configuration with defaults
IMAGE_NAME="lilhammer/delong"
TAG="${1:-amd64-latest}"
DOCKERFILE_PATH="${2:-deploy/docker/Dockerfile}"
BUILD_CONTEXT="."
USE_LOAD=true

# Show usage if help requested
if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
    echo "Usage: $0 [TAG] [DOCKERFILE_PATH]"
    echo ""
    echo "Arguments:"
    echo "  TAG             Docker image tag (default: amd64-latest)"
    echo "  DOCKERFILE_PATH Path to Dockerfile (default: deploy/docker/Dockerfile)"
    echo ""
    echo "Examples:"
    echo "  $0                          # Use defaults"
    echo "  $0 v1.0.0                   # Custom tag"
    echo "  $0 v1.0.0 docker/Dockerfile # Custom tag and dockerfile"
    exit 0
fi

echo "Building with:"
echo "  IMAGE: ${IMAGE_NAME}"
echo "  TAG: ${TAG}"
echo "  DOCKERFILE: ${DOCKERFILE_PATH}"
echo ""

# Switch to default builder and bootstrap
echo "Switching to default builder and bootstrapping"
docker buildx use default
docker buildx inspect default --bootstrap

echo "Current builder:"
docker buildx ls | grep '\*'
echo "Supported platforms:"
docker buildx inspect default | grep "Platforms:" -A1

# Build and push image
FULL_IMAGE="${IMAGE_NAME}:${TAG}"
echo "Building image: ${FULL_IMAGE}"

if [ "$USE_LOAD" = true ]; then
  docker buildx build \
    --builder default \
    --platform=linux/amd64 \
    -f "${DOCKERFILE_PATH}" \
    -t "${FULL_IMAGE}" \
    --load \
    --push \
    "${BUILD_CONTEXT}"
else
  docker buildx build \
    --builder default \
    --platform=linux/amd64 \
    -f "${DOCKERFILE_PATH}" \
    -t "${FULL_IMAGE}" \
    --push \
    "${BUILD_CONTEXT}"
fi

echo "Image built and pushed successfully: ${FULL_IMAGE}"
