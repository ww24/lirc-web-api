#!/bin/sh

ARCH=$(uname -m)
case $ARCH in
  arm*) ARCH="arm";;
  aarch64) ARCH="arm";;
  x86) ARCH="386";;
  x86_64) ARCH="amd64";;
  i686) ARCH="386";;
i386) ARCH="386";;
esac
OS=$(uname | tr '[:upper:]' '[:lower:]')
if [ $OS = "windows" ]; then
  EXT=".exe"
fi
VERSION=$(curl -sL https://raw.githubusercontent.com/ww24/lirc-web-api/master/wercker.yml | grep 'API_VERSION="v' | awk -F\" '{print $2}')
curl -sLo api "https://github.com/ww24/lirc-web-api/releases/download/${VERSION}/api_${OS}_${ARCH}${EXT}"
chmod +x api

echo "installed at $(pwd)/api $(./api -v)"
