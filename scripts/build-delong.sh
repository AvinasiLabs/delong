#!/usr/bin/env bash
#
# Build and push delong Docker image using buildx with multi-tagging
# Usage: ./build-delong.sh [TAG] [DOCKERFILE_PATH]
#

set -euo pipefail

# Configuration with defaults
IMAGE_NAME="lilhammer/delong"
DOCKERFILE_PATH="${2:-deploy/docker/Dockerfile}"
BUILD_CONTEXT="."
USE_LOAD=true

# Generate git-based tag
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DEFAULT_TAG="amd64-${GIT_HASH}"
TAG="${1:-$DEFAULT_TAG}"

# Show usage if help requested
if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
    echo "Usage: $0 [TAG] [DOCKERFILE_PATH]"
    echo ""
    echo "Arguments:"
    echo "  TAG             Docker image tag (default: amd64-\${GIT_HASH})"
    echo "  DOCKERFILE_PATH Path to Dockerfile (default: deploy/docker/Dockerfile)"
    echo ""
    echo "Examples:"
    echo "  $0                          # Use git hash: amd64-abc1234"
    echo "  $0 v1.0.0                   # Custom tag"
    echo "  $0 v1.0.0 docker/Dockerfile # Custom tag and dockerfile"
    echo ""
    echo "Multi-tagging strategy:"
    echo "  - Primary tag: \${TAG}"
    echo "  - Latest tag: amd64-latest (always points to newest build)"
    exit 0
fi

echo "Building with:"
echo "  IMAGE: ${IMAGE_NAME}"
echo "  PRIMARY TAG: ${TAG}"
echo "  GIT HASH: ${GIT_HASH}"
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

# Build primary tagged image
PRIMARY_IMAGE="${IMAGE_NAME}:${TAG}"
LATEST_IMAGE="${IMAGE_NAME}:amd64-latest"

echo "Building images:"
echo "  Primary: ${PRIMARY_IMAGE}"
echo "  Latest:  ${LATEST_IMAGE}"

if [ "$USE_LOAD" = true ]; then
  # Build with multiple tags
  docker buildx build \
    --builder default \
    --platform=linux/amd64 \
    -f "${DOCKERFILE_PATH}" \
    -t "${PRIMARY_IMAGE}" \
    -t "${LATEST_IMAGE}" \
    --load \
    --push \
    "${BUILD_CONTEXT}"
else
  docker buildx build \
    --builder default \
    --platform=linux/amd64 \
    -f "${DOCKERFILE_PATH}" \
    -t "${PRIMARY_IMAGE}" \
    -t "${LATEST_IMAGE}" \
    --push \
    "${BUILD_CONTEXT}"
fi

echo "Images built and pushed successfully:"
echo "  ✅ ${PRIMARY_IMAGE}"
echo "  ✅ ${LATEST_IMAGE}"
echo ""
echo "Usage in docker-compose:"
echo "  For stable deployments: image: ${PRIMARY_IMAGE}"
echo "  For latest version:     image: ${LATEST_IMAGE}"
