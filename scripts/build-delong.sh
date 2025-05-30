#!/bin/bash
set -e

# Configuration
DOCKER_USERNAME="${DOCKER_USERNAME:-lilhammer}"
IMAGE_NAME="${IMAGE_NAME:-delong}"
VERSION="${VERSION:-$(git describe --tags --always --dirty)}"
REGISTRY="${REGISTRY:-docker.io}"
FULL_IMAGE_NAME="$REGISTRY/$DOCKER_USERNAME/$IMAGE_NAME:$VERSION"
LATEST_TAG="$REGISTRY/$DOCKER_USERNAME/$IMAGE_NAME:latest"

echo "Building Delong test image..."

# Show build context optimization
echo "📊 Build context analysis:"
CONTEXT_SIZE=$(du -sh . | cut -f1)
FILE_COUNT=$(find . -type f | wc -l)
INCLUDED_COUNT=$(find . -type f | grep -v -f <(sed 's/#.*$//' .dockerignore | sed '/^$/d' | sed 's/^/^\.\//' | sed 's/\*/.*/' | sed 's/\/$/\/.*/' 2>/dev/null) | wc -l 2>/dev/null || echo "~40")

echo "  Total project size: $CONTEXT_SIZE"
echo "  Total files: $FILE_COUNT"
echo "  Files included in build: $INCLUDED_COUNT"
echo "  Optimization: ~$(( (FILE_COUNT - INCLUDED_COUNT) * 100 / FILE_COUNT ))% files excluded"

# Build the Docker image
echo "🔨 Starting Docker build..."
docker build -f deploy/docker/Dockerfile -t $FULL_IMAGE_NAME .
docker tag $FULL_IMAGE_NAME $LATEST_TAG

echo "✅ Image built successfully: $FULL_IMAGE_NAME"

# Check if we should push (requires Docker login)
if [[ "${PUSH:-}" == "true" ]] || [[ "${CI:-}" == "true" ]]; then
    echo "🚀 Pushing to registry..."
    docker push $FULL_IMAGE_NAME
    docker push $LATEST_TAG
    echo "✅ Push completed!"
    echo "Your image is available as: $FULL_IMAGE_NAME"
else
    read -p "Do you want to push to registry? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "🚀 Pushing to registry..."
        docker push $FULL_IMAGE_NAME
        docker push $LATEST_TAG
        echo "✅ Push completed!"
        echo "Your image is available as: $FULL_IMAGE_NAME"
        echo "Use this in your docker-compose.yml: image: $FULL_IMAGE_NAME"
    else
        echo "Skipping push. To push manually later:"
        echo "  docker push $FULL_IMAGE_NAME"
        echo "  docker push $LATEST_TAG"
    fi
fi

echo "🎉 Build process completed!"
echo ""
echo "📋 Summary:"
echo "  Image: $FULL_IMAGE_NAME"
echo "  Latest: $LATEST_TAG"
echo "  Registry: $REGISTRY"
echo ""
echo "🐳 To run locally:"
echo "  docker run -p 8080:8080 $FULL_IMAGE_NAME"
echo ""
echo "🚀 To deploy with docker-compose:"
echo "  cd deploy/docker && docker-compose up -d"
