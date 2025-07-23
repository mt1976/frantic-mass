#!/bin/bash

# Reset script to clear specific directories
clear_script="./clear.sh"
if [ -f "$clear_script" ]; then
    echo "Running clear script..."
    bash "$clear_script"
else
    echo "Clear script not found: $clear_script"
fi

# Move data from /defaults to /data/defaults
defaults_dir="defaults"
if [ -d "$defaults_dir" ]; then
    echo "Copying data from $defaults_dir to /data/defaults"
    cp "$defaults_dir"/* data/defaults/
else
    echo "Defaults directory not found: $defaults_dir"
fi
