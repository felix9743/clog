#!/bin/bash
set -e

REPO="hellotimking/clog"
INSTALL_DIR="/usr/local/bin"
BINARY="clog"

echo "Installing CLOG - High-Visibility Caddy Logs"
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux|darwin) ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64)
        ASSET="clog-${OS}-amd64"
        ;;
    aarch64|arm64)
        ASSET="clog-${OS}-arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "Detected: $OS/$ARCH"
echo "Downloading $ASSET..."

# Get latest release download URL
DOWNLOAD_URL=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep "browser_download_url" \
    | grep "$ASSET" \
    | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not find download URL for $ASSET"
    exit 1
fi

# Download binary
curl -L "$DOWNLOAD_URL" -o "/tmp/$BINARY"

# Install
chmod +x "/tmp/$BINARY"
sudo mv "/tmp/$BINARY" "$INSTALL_DIR/$BINARY"

echo ""
echo "CLOG installed successfully!"
echo "Run: clog /path/to/caddy/access.log"
