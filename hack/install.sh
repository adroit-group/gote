#!/usr/bin/env bash

GITHUB_REPO="https://github.com/adroit-group/gote.git"

# Initialize variables
DEST_DIR=""
NEW_PACKAGE=""

# Function to show usage
show_usage() {
  echo "Usage: $0 [-d|--dest DESTINATION] [-p|--package PACKAGE_NAME] [-h|--help]"
  echo ""
  echo "Options:"
  echo "  -d, --dest DESTINATION      Destination directory for the project"
  echo "  -p, --package PACKAGE_NAME  New package name (e.g., github.com/user/project)"
  echo "  -h, --help                  Show this help message"
  echo ""
  echo "If flags are not provided, you will be prompted for the values."
}

# Function to parse command line arguments
parse_arguments() {
  while [[ $# -gt 0 ]]; do
    case $1 in
      -d|--dest)
        DEST_DIR="$2"
        shift 2
        ;;
      -p|--package)
        NEW_PACKAGE="$2"
        shift 2
        ;;
      -h|--help)
        show_usage
        exit 0
        ;;
      *)
        echo "Unknown option: $1"
        show_usage
        exit 1
        ;;
    esac
  done
}

# Function to check dependencies
check_dependencies() {
  echo "Checking dependencies"
  if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
  fi
}

# Function to collect input (prompts for missing values)
collect_input() {
  # Prompt for destination directory only if not provided via flag
  if [ -z "$DEST_DIR" ]; then
    echo "Enter destination directory (press Tab for autocompletion):"
    read -e -r -p "> " DEST_DIR
  fi

  # Prompt for the new package name only if not provided via flag
  if [ -z "$NEW_PACKAGE" ]; then
    echo "Enter new package name (currently github.com/adroit-group/gote):"
    read -r NEW_PACKAGE
  fi
}

# Function to setup sed command based on OS
setup_sed() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    function do_sed() {
      sed -i '' "$1" "$2"
    }
  else
    function do_sed() {
      sed -i "$1" "$2"
    }
  fi
}

# Function to clone repository and clean up unwanted files
setup_repository() {
  echo "Setting up repository in $DEST_DIR"

  # Create directory if it doesn't exist
  mkdir -p "$DEST_DIR"
  cd "$DEST_DIR" || exit 1

  # Clone repository
  git clone -q --depth 1 --branch main "$GITHUB_REPO" . || exit 1

  # Remove unwanted files
  echo "./pkg .git go.sum CONTRIBUTIONS.md LICENSE README.md ./hack ./.github/PULL_REQUEST_TEMPLATE.md ./.github/ISSUE_TEMPLATE" | xargs rm -rf
}

# Function to replace package names throughout the codebase
replace_package_names() {
  local OLD_PACKAGE="github.com/adroit-group/gote"

  echo "Updating package names from $OLD_PACKAGE to $NEW_PACKAGE"

  # Update go.mod
  do_sed "s|$OLD_PACKAGE|$NEW_PACKAGE|g" go.mod

  # Process Go files
  find . -type f -name "*.go" | while read -r file; do
    # Replace package imports (but only for internal imports)
    do_sed "s|import \"$OLD_PACKAGE/internal|import \"$NEW_PACKAGE/internal|g" "$file"
    do_sed "s|import \"$OLD_PACKAGE/cmd|import \"$NEW_PACKAGE/cmd|g" "$file"
    do_sed "s|import ($|import (|" "$file" # Preserve multiline imports
    do_sed "s|\t\"$OLD_PACKAGE/internal|\t\"$NEW_PACKAGE/internal|g" "$file"
    do_sed "s|\t\"$OLD_PACKAGE/cmd|\t\"$NEW_PACKAGE/cmd|g" "$file"

    # Fix package declarations ONLY for files in the internal directory
    if [[ "$file" == ./internal/* ]]; then
      # Extract the relative path from the internal directory
      REL_PATH=$(echo "$file" | sed "s|^./internal/||" | xargs dirname)
      if [ "$REL_PATH" = "." ]; then
        # For files directly in the internal directory
        do_sed "s|^package .*$|package internal|g" "$file"
      else
        # For files in subdirectories of internal
        PACKAGE_NAME=$(basename "$REL_PATH")
        do_sed "s|^package .*$|package $PACKAGE_NAME|g" "$file"
      fi
    fi
  done

  # Update Dockerfile if it exists
  if [ -f "build/Dockerfile" ]; then
    do_sed "s|$OLD_PACKAGE/|$NEW_PACKAGE/|g" "build/Dockerfile"
  fi

  echo "Package name updated from $OLD_PACKAGE to $NEW_PACKAGE"
}

# Function to finalize the setup
finalize_setup() {
  echo "Finalizing setup"

  # Tidy go modules
  go mod tidy

  # Initialize git repository
  git init -b main

  echo "Setup complete!"
}

# Main function to orchestrate the installation
main() {
  echo "Installing GoTe..."

  parse_arguments "$@"
  check_dependencies
  collect_input
  setup_sed
  setup_repository
  replace_package_names
  finalize_setup
}

# Run main function with all arguments
main "$@"
