#!/usr/bin/env bash

set -euo pipefail

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
 Linux) OS=linux;;
 Darwin) OS=darwin;;
 *)
   echo "Unsupported OS: $OS"
   exit 1
   ;;
esac

case "$ARCH" in
  x86_64) ARCH=amd64;;
  aarch64|arm64) ARCH=arm64;;
  *)
  echo "Unsupported architecture: $ARCH"
  exit 1
  ;;esac

VERSION=$(curl -fsSL \
  https://github.com/mainsail-partners/db/releases/latest |
  grep '"tag_name"' |
  cut -d '"' -f 4)

FILE="db-${OS}-${ARCH}"

curl -L \
  -o db \
  "https://github.com/yourorg/db/releases/download/${VERSION}/${FILE}"

chmod +x db

sudo mv db /usr/local/bin/db

echo "Installed $VERSION"