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
if [ "$OS" = "windows" ]; then
  EXT=".exe"
fi
VERSION=$(curl -sL https://raw.githubusercontent.com/ww24/lirc-web-api/master/wercker.yml | grep 'API_VERSION="v' | awk -F\" '{print $2}')
curl -Lo /tmp/lirc-web-api "https://github.com/ww24/lirc-web-api/releases/download/${VERSION}/api_${OS}_${ARCH}${EXT}"
chmod +x /tmp/lirc-web-api

if [ "$OS" = "linux" ] && hash systemctl; then
  INSTALL_DIR=/usr/local/bin
  mkdir -p $INSTALL_DIR
  INSTALL_PATH="$INSTALL_DIR/lirc-web-api"
  sudo mv /tmp/lirc-web-api $INSTALL_PATH

  cat <<EOF | sudo tee /lib/systemd/system/lirc-web-api.service
[Unit]
Description=lirc-web-api
After=network.target network-online.target

[Service]
Type=simple
ExecStart=$INSTALL_PATH
ExecReload=/bin/kill -HUP \$MAINPID
KillMode=control-group
Restart=on-failure

[Install]
WantedBy=multi-user.target
Alias=mackerel-agent.service
EOF

  sudo systemctl enable lirc-web-api
  sudo systemctl start lirc-web-api
  sudo systemctl status lirc-web-api
else
  mv /tmp/lirc-web-api .
  echo "installed at $(pwd)/lirc-web-api $(./lirc-web-api -v)"
fi
