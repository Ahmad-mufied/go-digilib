#!/bin/bash

# Function to display error messages
function error_exit {
    echo "$1" 1>&2
    exit 1
}

# Search for all .env* files in the current directory
# shellcheck disable=SC2207
env_files=($(ls .env* 2>/dev/null))

# Check if any .env files are found
if [ ${#env_files[@]} -eq 0 ]; then
    error_exit "No .env files found in the current directory!"
fi

# Add an option to cancel the script
env_files+=("Cancel")

# Display the select menu
echo "Select the .env file to load or choose 'Cancel' to exit:"
select env_file in "${env_files[@]}"; do
    if [ "$env_file" == "Cancel" ]; then
        echo "Script execution canceled."
        exit 0
    elif [ -n "$env_file" ]; then
        echo "You have selected: $env_file"
        break
    else
        echo "Invalid selection. Please try again."
    fi
done

# Check if the selected file exists and is readable
if [ ! -f "$env_file" ] || [ ! -r "$env_file" ]; then
    error_exit "The selected file does not exist or is not readable."
fi

# Load environment variables from the selected .env file
set -o allexport
while IFS='=' read -r key value; do
    # Skip lines that are empty or start with a comment
    [[ -z "$key" || "$key" == \#* ]] && continue
    # Remove possible surrounding quotes from the value
    value="${value%\"}"
    value="${value#\"}"
    export "$key=$value"
done < "$env_file"
set +o allexport

# Loop through each variable in the selected .env file and set it in Fly.io
while IFS='=' read -r key value; do
    # Skip lines that are empty or start with a comment
    [[ -z "$key" || "$key" == \#* ]] && continue
    # Remove possible surrounding quotes from the value
    value="${value%\"}"
    value="${value#\"}"
    # Set the variable in Fly.io
    fly secrets set "$key=$value" || error_exit "Failed to set $key in Fly.io."
    # Add a 5-second delay after setting each variable
    sleep 5
done < "$env_file"

echo "All environment variables from $env_file have been set in Fly.io."

# Optionally, restart the Fly.io app
# flyctl apps restart <your-app-name> || error_exit "Failed to restart the Fly.io app."

echo "Fly.io app environment variables set successfully."
