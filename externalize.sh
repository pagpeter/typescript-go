#!/bin/bash

# Script to recursively replace import paths in Go files

# Replace the more specific path first to avoid partial replacements
find . -name "*.go" -type f -exec sed -i '' 's|github\.com/microsoft/typescript-go/internal|github.com/pagpeter/typescript-go/external|g' {} \;

# Replace the general path
find . -name "*.go" -type f -exec sed -i '' 's|github\.com/microsoft/typescript-go|github.com/pagpeter/typescript-go|g' {} \;

echo "Replacement complete!"
echo "Files modified:"
find . -name "*.go" -type f -exec grep -l "github.com/pagpeter/typescript-go" {} \;

# Rename the internal directory to external
if [ -d "./internal" ]; then
    mv ./internal ./external
    echo "Renamed ./internal to ./external"
else
    echo "Directory ./internal not found"
fi

# Delete the cmd directory
if [ -d "./cmd" ]; then
    rm -rf ./cmd
    echo "Deleted ./cmd directory"
else
    echo "Directory ./cmd not found"
fi

# Update go.mod file
if [ -f "./go.mod" ]; then
    sed -i '' 's|github\.com/microsoft/typescript-go|github.com/pagpeter/typescript-go|g' ./go.mod
    echo "Updated ./go.mod"
else
    echo "File ./go.mod not found"
fi