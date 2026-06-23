#!/usr/bin/env bash

set -euo pipefail
echo 1
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
echo 2
case "$ARCH" in
  x86_64) ARCH=amd64;;
  aarch64|arm64) ARCH=arm64;;
  *)
  echo "Unsupported architecture: $ARCH"
  exit 1
  ;;esac
echo 3

VERSION=$(curl -fsSL \
  https://api.github.com/repos/gerritjvv/db/releases/latest|
  grep '"tag_name"' |
  cut -d '"' -f 4)

echo 4 "$VERSION"
FILE="db-${OS}-${ARCH}"


echo curl -L \
  -o db \
  "https://github.com/gerritjvv/db/releases/download/${VERSION}/${FILE}"
echo 5
chmod +x db

#sudo mv db /usr/local/bin/db
#
#echo "Installed $VERSION"