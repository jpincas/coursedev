#!/bin/bash
# Deploy coursedev to training.yagni.co.uk
set -e

SERVER="jon@training.yagni.co.uk"
REMOTE_DIR="/home/jon/src/github.com/yagniltd/coursedev"

echo "=== Building for Linux amd64 ==="
GOOS=linux GOARCH=amd64 go build -o coursedev-linux .

echo "=== Ensuring remote directory exists ==="
ssh "$SERVER" "mkdir -p $REMOTE_DIR"

echo "=== Syncing binary ==="
rsync -avz coursedev-linux "$SERVER:$REMOTE_DIR/coursedev"
ssh "$SERVER" "chmod +x $REMOTE_DIR/coursedev"

echo "=== Syncing content ==="
rsync -avz --delete content/ "$SERVER:$REMOTE_DIR/content/"

echo "=== Syncing static assets ==="
rsync -avz --delete static/ "$SERVER:$REMOTE_DIR/static/"

echo "=== Restarting service ==="
ssh -t "$SERVER" "sudo systemctl restart coursedev"

echo "=== Checking status ==="
ssh -t "$SERVER" "sudo systemctl status coursedev --no-pager"

# Clean up local cross-compiled binary
rm -f coursedev-linux

echo ""
echo "=== Deploy complete ==="
echo "https://training.yagni.co.uk"
