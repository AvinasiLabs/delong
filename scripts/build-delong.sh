#!/usr/bin/env bash
#
# Build and push delong Docker image using buildx with architecture selection
# Usage: ./build-delong.sh [OPTIONS]
#

set -euo pipefail

# Configuration
IMAGE_NAME="lilhammer/delong"
DOCKERFILE_PATH="deploy/docker/Dockerfile"
BUILD_CONTEXT="."

# Default values
ARCH="amd64"
UPDATE_LATEST=false
SHOW_HELP=false

# Parse command line options
while getopts "a:lh" opt; do
    case $opt in
        a)
            ARCH="$OPTARG"
            ;;
        l)
            UPDATE_LATEST=true
            ;;
        h)
            SHOW_HELP=true
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            echo "Use -h for help"
            exit 1
            ;;
    esac
done

# Show usage if help requested
if [[ "$SHOW_HELP" == "true" ]]; then
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -a ARCH         Target architecture (default: amd64)"
    echo "                  Examples: amd64, arm64, arm64/v7, arm64/v8, windows/amd64"
    echo "  -l              Also tag as 'latest' (default: false)"
    echo "  -h              Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                    # Build amd64, git-hash tag only"
    echo "  $0 -a arm64           # Build arm64, git-hash tag only"
    echo "  $0 -a amd64 -l        # Build amd64, with latest tag"
    echo "  $0 -l                 # Build amd64, with latest tag"
    echo ""
    echo "Tagging strategy:"
    echo "  - Always creates: \${ARCH}-\${GIT_HASH}"
    echo "  - With -l flag:   \${ARCH}-latest (overwrites existing)"
    echo ""
    echo "Note: Platform support depends on your Docker buildx configuration."
    echo "Run 'docker buildx inspect' to see supported platforms."
    exit 0
fi

# Check if Dockerfile exists
if [[ ! -f "$DOCKERFILE_PATH" ]]; then
    echo "❌ Error: Dockerfile not found at '$DOCKERFILE_PATH'"
    echo ""
    echo "Please make sure you are running this script from the project root directory."
    echo "The expected project structure should be:"
    echo "  project-root/"
    echo "  ├── deploy/docker/Dockerfile"
    echo "  └── scripts/build-delong.sh"
    echo ""
    echo "Current directory: $(pwd)"
    exit 1
fi

# Generate git-based tag
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
if [[ "$GIT_HASH" == "unknown" ]]; then
    echo "⚠️  Warning: Not in a git repository, using 'unknown' as hash"
fi

# Generate platform string and safe tag name
PLATFORM="linux/$ARCH"
# Handle special cases for platform
case "$ARCH" in
    windows/*)
        PLATFORM="$ARCH"  # Keep as-is for Windows
        ;;
    *)
        PLATFORM="linux/$ARCH"  # Default to Linux
        ;;
esac

# Convert arch to tag-safe format (replace / and : with -)
ARCH_SAFE=$(echo "$ARCH" | tr '/:' '-')
PRIMARY_TAG="${ARCH_SAFE}-${GIT_HASH}"
PRIMARY_IMAGE="${IMAGE_NAME}:${PRIMARY_TAG}"

# Prepare build tags
BUILD_TAGS=("-t" "$PRIMARY_IMAGE")
TAG_LIST=("$PRIMARY_IMAGE")

if [[ "$UPDATE_LATEST" == "true" ]]; then
    LATEST_TAG="${ARCH_SAFE}-latest"
    LATEST_IMAGE="${IMAGE_NAME}:${LATEST_TAG}"
    BUILD_TAGS+=("-t" "$LATEST_IMAGE")
    TAG_LIST+=("$LATEST_IMAGE")
fi

echo "Building with:"
echo "  IMAGE: ${IMAGE_NAME}"
echo "  ARCHITECTURE: ${ARCH}"
echo "  PLATFORM: ${PLATFORM}"
echo "  PRIMARY TAG: ${PRIMARY_TAG}"
if [[ "$UPDATE_LATEST" == "true" ]]; then
    echo "  LATEST TAG: ${LATEST_TAG} (will be updated)"
else
    echo "  LATEST TAG: not updated (use -l to update)"
fi
echo "  GIT HASH: ${GIT_HASH}"
echo "  DOCKERFILE: ${DOCKERFILE_PATH}"
echo ""

# Switch to default builder and bootstrap
echo "Switching to default builder and bootstrapping..."
docker context use default
docker buildx inspect default --bootstrap

echo "Current builder info:"
docker buildx ls | grep '\*'
echo ""
echo "Available platforms:"
docker buildx inspect default | grep "Platforms:" -A1
echo ""

# Build images
echo "Building images:"
for tag in "${TAG_LIST[@]}"; do
    echo "  📦 $tag"
done
echo "  Platform: ${PLATFORM}"
echo ""

# Let Docker buildx handle platform validation - it will fail gracefully if unsupported
if docker buildx build \
    --builder default \
    --platform="$PLATFORM" \
    -f "$DOCKERFILE_PATH" \
    "${BUILD_TAGS[@]}" \
    --load \
    --push \
    "$BUILD_CONTEXT"; then

    echo ""
    echo "🎉 Images built and pushed successfully:"
    for tag in "${TAG_LIST[@]}"; do
        echo "  ✅ $tag"
    done
    echo ""
    echo "Usage in docker-compose:"
    echo "  For this specific build: image: ${PRIMARY_IMAGE}"
    if [[ "$UPDATE_LATEST" == "true" ]]; then
        echo "  For latest version:      image: ${LATEST_IMAGE}"
    fi
    echo ""
    echo "Development workflow tips:"
    echo "  For testing: $0 -a ${ARCH}      # No latest update"
    echo "  For release: $0 -a ${ARCH} -l   # Update latest tag"
else
    echo ""
    echo "❌ Build failed!"
    echo ""
    echo "This could be due to:"
    echo "  1. Platform '${PLATFORM}' not supported by current builder"
    echo "  2. Dockerfile not compatible with target platform"
    echo "  3. Base image not available for target platform"
    echo ""
    echo "To check supported platforms:"
    echo "  docker buildx inspect"
    echo ""
    echo "To see available builders:"
    echo "  docker buildx ls"
    exit 1
fi
