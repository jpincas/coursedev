#!/bin/bash
# Restart the coursedev server
# Usage: ./restart.sh

set -e

export COURSEDEV_ADMIN_KEY="${COURSEDEV_ADMIN_KEY:-admin-test-key}"

# Kill any existing server on port 8080
lsof -ti:8080 | xargs kill -9 2>/dev/null || true

# Build
echo "Building..."
go build -o coursedev .

# Run in background
echo "Starting server..."
./coursedev &

# Wait for server to start (up to 5 seconds)
for i in 1 2 3 4 5; do
    if lsof -ti:8080 > /dev/null 2>&1; then
        echo "Server running at http://localhost:8080"
        exit 0
    fi
    sleep 1
done

echo "Failed to start server"
exit 1
