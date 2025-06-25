#!/bin/bash
# Image name
IMAGE_NAME="mt1976/frantic-mass"
VERSION_FILE="version.no"
VERSION_MAJORMINOR="0.0"
DOCKERHUB_USERNAME="mt1976"
DOCKERHUB_PASSWORD="Merc400350"

echo "🚀 Starting version increment: reading semantic version from '${VERSION_FILE}' and bumping patch version..."

# If version file doesn't exist, start with 0.0.1
if [ ! -f "$VERSION_FILE" ]; then
    echo "0.0.1" > "$VERSION_FILE"
fi

# Read the current semantic version
CURRENT_VERSION=$(cat "$VERSION_FILE")

# Validate format: must be like x.y.z
if [[ ! "$CURRENT_VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "❌ Error: version.no does not contain a valid semantic version (x.y.z)"
    exit 1
fi

# Split version into parts
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# Increment PATCH
PATCH=$((PATCH + 1))

# Reconstruct version
NEW_VERSION="${MAJOR}.${MINOR}.${PATCH}"

# Save it
echo "$NEW_VERSION" > "$VERSION_FILE"

# Output the version (for use in build scripts)
echo "✅ Version successfully incremented: ${CURRENT_VERSION} → ${NEW_VERSION}"
# Set your version here (change as needed or make dynamic)
VERSION="${NEW_VERSION}"

# Build the Docker image
echo "🚧 Starting build processes..."

docker build -t ${IMAGE_NAME}:latest -t ${IMAGE_NAME}:${VERSION} .

# Log in to Docker Hub (you can skip if already logged in)
docker login --username "${DOCKERHUB_USERNAME}" --password "${DOCKERHUB_PASSWORD}"
echo "✅ Logged In"
# Push both tags
echo "📤 Pushing Docker image ${IMAGE_NAME}:latest to Docker Hub..."
docker push ${IMAGE_NAME}:latest
echo "📤 Pushing Docker image ${IMAGE_NAME}:${VERSION} to Docker Hub..."
docker push ${IMAGE_NAME}:${VERSION}
echo "✅ Successfully pushed ${IMAGE_NAME} with tags: latest and ${VERSION}"