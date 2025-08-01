#!/bin/bash

echo "Starting the application..."
echo "Resetting the environment..."
# Reset script to clear specific directories
./reset.sh

echo "Reset completed. Starting the application..."
go run ./...
