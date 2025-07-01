#!/bin/bash

# Script to recursively replace import paths in Go files

# Replace the more specific path first to avoid partial replacements
find . -name "*.go" -type f -exec sed -i '' 's|github\.com/microsoft/typescript-go/internal|github.com/pagpeter/typescript-go/external|g' {} \;

# Replace the general path
find . -name "*.go" -type f -exec sed -i '' 's|github\.com/microsoft/typescript-go|github.com/pagpeter/typescript-go|g' {} \;

echo "Replacement complete!"
echo "Files modified:"
find . -name "*.go" -type f -exec grep -l "github.com/pagpeter/typescript-go" {} \;