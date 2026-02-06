#!/bin/bash
# Restart the coursedev server
# Usage: ./restart.sh

set -e

# Kill any existing server on port 8080
lsof -ti:8080 | xargs kill -9 2>/dev/null || true

# Build
echo "Building..."
go build -o coursedev .

# Run in background
echo "Starting server..."
./coursedev &

# Wait a moment for startup
sleep 1

# Check if it's running
if lsof -ti:8080 > /dev/null 2>&1; then
    echo "Server running at http://localhost:8080"
else
    echo "Failed to start server"
    exit 1
fi
