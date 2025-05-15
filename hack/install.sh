#!/usr/bin/env bash

GITHUB_REPO="https://github.com/adroit-group/gote.git"

echo "Installing GoTe..."

# Prompt for destination directory with autocompletion
echo "Enter destination directory (press Tab for autocompletion):"
read -e -r -p "> " DEST_DIR

# Create directory if it doesn't exist
mkdir -p "$DEST_DIR"
cd "$DEST_DIR" || exit 1

git clone -q --depth 1 --branch main "$GITHUB_REPO" .

rm -rf ./pkg .git go.sum CONTRIBUTIONS.md LICENSE README.md ./hack ./.github/PULL_REQUEST_TEMPLATE.md ./.github/ISSUE_TEMPLATE

if [[ "$OSTYPE" == "darwin"* ]]; then
  function do_sed() {
    sed -i '' "$1" "$2"
  }
else
  function do_sed() {
    sed -i "$1" "$2"
  }
fi

# Prompt for the new package name
echo "Enter new package name (currently github.com/adroit-group/gote):"
read -r new_package

# Store the old package name
old_package="github.com/adroit-group/gote"

do_sed "s|$old_package|$new_package|g" go.mod

# Process Go files in the internal directory
find . -type f -name "*.go" | while read -r file; do
  # Replace package imports (but only for internal imports)
  do_sed "s|import \"$old_package/internal|import \"$new_package/internal|g" "$file"
  do_sed "s|import \"$old_package/cmd|import \"$new_package/cmd|g" "$file"
  do_sed "s|import ($|import (|" "$file" # Preserve multiline imports
  do_sed "s|\t\"$old_package/internal|\t\"$new_package/internal|g" "$file"
  do_sed "s|\t\"$old_package/cmd|\t\"$new_package/cmd|g" "$file"
  # Fix package declarations
  # Extract the relative path from the internal directory
  rel_path=$(echo "$file" | sed "s|internal/||" | xargs dirname)
  if [ "$rel_path" = "." ]; then
    # For files directly in the internal directory
    do_sed "s|^package .*$|package internal|g" "$file"
  else
    # For files in subdirectories of internal
    package_name=$(basename "$rel_path")
    do_sed "s|^package .*$|package $package_name|g" "$file"
  fi
done

if [ -f "build/Dockerfile" ]; then
  do_sed "s|$old_package/|$new_package/|g" "build/Dockerfile"
fi


echo "Package name updated from $old_package to $new_package"

go mod tidy

git init -b main
