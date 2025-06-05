#!/bin/bash

# Setup clean .env file from template
TEMPLATE="${1:-deploy/docker/.env.template}"
ENV_FILE="$(dirname "$TEMPLATE")/.env"

if [ ! -f "$TEMPLATE" ]; then
    echo "Error: Template file not found: $TEMPLATE"
    exit 1
fi

grep -v '^#' "$TEMPLATE" | grep -v '^[[:space:]]*$' > "$ENV_FILE"
echo "Created $ENV_FILE from $TEMPLATE (removed comments and empty lines)"