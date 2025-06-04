#!/usr/bin/env bash
#
# Build and push delong Docker image using buildx with multi-platform support
# Usage: ./build-delong.sh [OPTIONS]
#

set -euo pipefail

# Configuration
IMAGE_NAME="lilhammer/delong"
DOCKERFILE_PATH="deploy/docker/Dockerfile"
BUILD_CONTEXT="."
BUILDER_NAME="delong-builder"

# Default values
ARCH=""
MULTI_PLATFORM=false
UPDATE_LATEST=false
SHOW_HELP=false

# Predefined platform combinations - using function instead of associative array for better compatibility
get_platform_preset() {
    case "$1" in
        "common")
            echo "linux/amd64,linux/arm64"
            ;;
        "full")
            echo "linux/amd64,linux/arm64,linux/arm/v7"
            ;;
        "intel")
            echo "linux/amd64"
            ;;
        "arm")
            echo "linux/arm64,linux/arm/v7"
            ;;
        "apple")
            echo "linux/arm64"
            ;;
        *)
            echo ""
            ;;
    esac
}

get_available_presets() {
    echo "common full intel arm apple"
}

# Parse command line options
while getopts "a:m:p:lh" opt; do
    case $opt in
        a)
            ARCH="$OPTARG"
            ;;
        m)
            MULTI_PLATFORM=true
            PLATFORMS="$OPTARG"
            ;;
        p)
            MULTI_PLATFORM=true
            PLATFORMS=$(get_platform_preset "$OPTARG")
            if [[ -z "$PLATFORMS" ]]; then
                echo "❌ Unknown preset: $OPTARG"
                echo "Available presets: $(get_available_presets)"
                exit 1
            fi
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
    echo "Build Modes:"
    echo "  Single Platform:  $0 -a ARCH"
    echo "  Multi Platform:   $0 -m PLATFORMS"
    echo "  Platform Preset:  $0 -p PRESET"
    echo ""
    echo "Options:"
    echo "  -a ARCH         Single target architecture"
    echo "                  Examples: amd64, arm64, arm/v7"
    echo "  -m PLATFORMS    Multi-platform build (comma-separated)"
    echo "                  Example: linux/amd64,linux/arm64"
    echo "  -p PRESET       Use predefined platform combination:"
    echo "                    common: $(get_platform_preset common)"
    echo "                    full:   $(get_platform_preset full)"
    echo "                    intel:  $(get_platform_preset intel)"
    echo "                    arm:    $(get_platform_preset arm)"
    echo "                    apple:  $(get_platform_preset apple)"
    echo "  -l              Also tag as 'latest' (default: false)"
    echo "  -h              Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 -a amd64             # Single platform build"
    echo "  $0 -m linux/amd64,linux/arm64  # Multi-platform build"
    echo "  $0 -p common            # Common platforms (amd64 + arm64)"
    echo "  $0 -p common -l         # Common platforms + latest tag"
    echo ""
    echo "Tagging strategy:"
    echo "  Single platform: {arch}-{git-hash}, optionally {arch}-latest"
    echo "  Multi platform:  {git-hash}, optionally latest"
    echo ""
    echo "Note: Multi-platform builds are pushed directly to registry."
    echo "Single-platform builds are loaded locally AND pushed."
    exit 0
fi

# Validate input
if [[ -n "$ARCH" && "$MULTI_PLATFORM" == "true" ]]; then
    echo "❌ Error: Cannot use both -a (single arch) and -m/-p (multi-platform) options"
    exit 1
fi

if [[ -z "$ARCH" && "$MULTI_PLATFORM" == "false" ]]; then
    echo "❌ Error: Must specify either -a ARCH or -m PLATFORMS or -p PRESET"
    echo "Use -h for help"
    exit 1
fi

# Check if Dockerfile exists
if [[ ! -f "$DOCKERFILE_PATH" ]]; then
    echo "❌ Error: Dockerfile not found at '$DOCKERFILE_PATH'"
    echo ""
    echo "Please make sure you are running this script from the project root directory."
    exit 1
fi

# Generate git-based tag
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
if [[ "$GIT_HASH" == "unknown" ]]; then
    echo "⚠️  Warning: Not in a git repository, using 'unknown' as hash"
fi

# Setup build configuration
if [[ "$MULTI_PLATFORM" == "true" ]]; then
    # Multi-platform build
    PLATFORM_LIST="$PLATFORMS"
    PRIMARY_TAG="$GIT_HASH"
    PRIMARY_IMAGE="${IMAGE_NAME}:${PRIMARY_TAG}"

    BUILD_TAGS=("-t" "$PRIMARY_IMAGE")
    TAG_LIST=("$PRIMARY_IMAGE")

    if [[ "$UPDATE_LATEST" == "true" ]]; then
        LATEST_IMAGE="${IMAGE_NAME}:latest"
        BUILD_TAGS+=("-t" "$LATEST_IMAGE")
        TAG_LIST+=("$LATEST_IMAGE")
    fi

    BUILD_MODE="Multi-platform"
    LOAD_FLAG=""  # Cannot use --load with multi-platform
    BUILDER_TO_USE="$BUILDER_NAME"
else
    # Single platform build
    PLATFORM_LIST="linux/$ARCH"
    ARCH_SAFE=$(echo "$ARCH" | tr '/:' '-')
    PRIMARY_TAG="${ARCH_SAFE}-${GIT_HASH}"
    PRIMARY_IMAGE="${IMAGE_NAME}:${PRIMARY_TAG}"

    BUILD_TAGS=("-t" "$PRIMARY_IMAGE")
    TAG_LIST=("$PRIMARY_IMAGE")

    if [[ "$UPDATE_LATEST" == "true" ]]; then
        LATEST_TAG="${ARCH_SAFE}-latest"
        LATEST_IMAGE="${IMAGE_NAME}:${LATEST_TAG}"
        BUILD_TAGS+=("-t" "$LATEST_IMAGE")
        TAG_LIST+=("$LATEST_IMAGE")
    fi

    BUILD_MODE="Single-platform"
    LOAD_FLAG="--load"  # Can load single-platform builds locally
    BUILDER_TO_USE="default"  # Use default builder for single platform
fi

echo "🔧 Build Configuration:"
echo "  MODE: $BUILD_MODE"
echo "  IMAGE: $IMAGE_NAME"
echo "  PLATFORMS: $PLATFORM_LIST"
echo "  PRIMARY TAG: $PRIMARY_TAG"
if [[ "$UPDATE_LATEST" == "true" ]]; then
    echo "  LATEST TAG: ${LATEST_IMAGE##*:} (will be updated)"
else
    echo "  LATEST TAG: not updated (use -l to update)"
fi
echo "  GIT HASH: $GIT_HASH"
echo "  DOCKERFILE: $DOCKERFILE_PATH"
echo "  BUILDER: $BUILDER_TO_USE"
echo ""

# Setup builder
echo "🔧 Setting up builder..."

if [[ "$MULTI_PLATFORM" == "true" ]]; then
    # For multi-platform builds, create or use a docker-container builder
    if ! docker buildx inspect "$BUILDER_NAME" >/dev/null 2>&1; then
        echo "📦 Creating multi-platform builder: $BUILDER_NAME"
        docker buildx create --name "$BUILDER_NAME" --driver docker-container --bootstrap >/dev/null
    else
        echo "♻️  Using existing builder: $BUILDER_NAME"
    fi

    # Ensure builder is running
    docker buildx inspect "$BUILDER_NAME" --bootstrap >/dev/null 2>&1
else
    # For single-platform builds, use default builder
    echo "📦 Using default builder for single-platform build"
    docker context use default >/dev/null 2>&1 || true
    docker buildx inspect default --bootstrap >/dev/null 2>&1 || true
fi

echo "📋 Builder info:"
docker buildx ls | grep -E "(NAME|$BUILDER_TO_USE|\*)"
echo ""
echo "🎯 Available platforms for $BUILDER_TO_USE:"
docker buildx inspect "$BUILDER_TO_USE" | grep "Platforms:" -A1
echo ""

# Validate platforms before building
echo "🔍 Validating platforms..."
AVAILABLE_PLATFORMS=$(docker buildx inspect "$BUILDER_TO_USE" | grep "Platforms:" | cut -d: -f2 | tr -d ' ')
IFS=',' read -ra REQUESTED_PLATFORMS <<< "$PLATFORM_LIST"
for platform in "${REQUESTED_PLATFORMS[@]}"; do
    platform=$(echo "$platform" | tr -d ' ')
    if [[ "$AVAILABLE_PLATFORMS" != *"$platform"* ]]; then
        echo "⚠️  Warning: Platform '$platform' may not be supported"
    else
        echo "✅ Platform '$platform' is supported"
    fi
done
echo ""

# Build images
echo "🚀 Building images:"
for tag in "${TAG_LIST[@]}"; do
    echo "  📦 $tag"
done
echo "  🎯 Platforms: $PLATFORM_LIST"
echo "  🔧 Builder: $BUILDER_TO_USE"
echo ""

# Construct build command
BUILD_CMD=(
    docker buildx build
    --builder "$BUILDER_TO_USE"
    --platform="$PLATFORM_LIST"
    -f "$DOCKERFILE_PATH"
    "${BUILD_TAGS[@]}"
    --push
)

# Add --load for single-platform builds only
if [[ -n "$LOAD_FLAG" ]]; then
    BUILD_CMD+=("$LOAD_FLAG")
fi

BUILD_CMD+=("$BUILD_CONTEXT")

# Execute build
if "${BUILD_CMD[@]}"; then
    echo ""
    echo "🎉 Build completed successfully!"
    echo ""
    echo "📦 Images built:"
    for tag in "${TAG_LIST[@]}"; do
        echo "  ✅ $tag"
    done
    echo ""
    echo "🎯 Platforms: $PLATFORM_LIST"
    echo "🔧 Builder used: $BUILDER_TO_USE"
    echo ""
    echo "📋 Usage in docker-compose:"
    echo "  image: ${PRIMARY_IMAGE}"
    if [[ "$UPDATE_LATEST" == "true" && "$MULTI_PLATFORM" == "true" ]]; then
        echo "  image: ${IMAGE_NAME}:latest  # Latest multi-platform"
    elif [[ "$UPDATE_LATEST" == "true" ]]; then
        echo "  image: ${LATEST_IMAGE}  # Latest for ${ARCH}"
    fi
    echo ""
    echo "💡 Development tips:"
    echo "  Quick test:    $0 -a amd64"
    echo "  Multi-platform: $0 -p common"
    echo "  Release:       $0 -p common -l"
    echo ""
    if [[ "$MULTI_PLATFORM" == "true" ]]; then
        echo "🗑️  To clean up builder: docker buildx rm $BUILDER_NAME"
    fi
else
    echo ""
    echo "❌ Build failed!"
    echo ""
    echo "This could be due to:"
    echo "  1. Platform(s) not supported: $PLATFORM_LIST"
    echo "  2. Dockerfile not compatible with target platform(s)"
    echo "  3. Base image not available for target platform(s)"
    echo "  4. Network issues with Docker registry"
    echo ""
    echo "🔧 Troubleshooting:"
    echo "  Check platforms: docker buildx inspect $BUILDER_TO_USE"
    echo "  Check builders:  docker buildx ls"
    echo "  Try single:      $0 -a amd64"
    if [[ "$MULTI_PLATFORM" == "true" ]]; then
        echo "  Reset builder:   docker buildx rm $BUILDER_NAME"
    fi
    exit 1
fi
