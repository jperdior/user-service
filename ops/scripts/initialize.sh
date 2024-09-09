#!/bin/bash

# Define the initialized file name
INIT_FILE=".initialized"

# Check if the project has already been initialized
if [ -f "$INIT_FILE" ]; then
  echo "Project has already been initialized. Exiting."
  exit 0
fi

# Prompt for the new project name
read -p "This is the first time you run this project. Enter the new project name, this will be the module name: " project_name

# Check if the input is empty
if [ -z "$project_name" ]; then
  echo "Project name cannot be empty."
  exit 1
fi

# Find and replace 'golang-template' with the new project name in all files
find . -type f -exec sed -i '' "s/golang-template/${project_name}/g" {} +

# Optional: Rename the project folder if it contains 'golang-template' in the name
if [[ $(basename "$PWD") == *"golang-template"* ]]; then
  new_dir=$(echo "$PWD" | sed "s/golang-template/${project_name}/")
  cd ..
  mv "$(basename "$PWD")" "$new_dir"
  cd "$new_dir"
fi

# Create the initialized file to mark the project as initialized
touch "$INIT_FILE"

echo "Project name updated to '$project_name' and marked as initialized."
