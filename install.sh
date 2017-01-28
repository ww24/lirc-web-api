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

DOWNLOAD_URL="https://github.com/ww24/lirc-web-api/releases/download/${VERSION}/api_${OS}_${ARCH}.tar.gz"
echo "Downloading $DOWNLOAD_URL"
curl -Lo /tmp/lirc-web-api.tar.gz "$DOWNLOAD_URL"
mkdir -p /tmp/licr-web-api
tar xzvf /tmp/lirc-web-api.tar.gz -C /tmp/lirc-web-api
chmod +x "/tmp/lirc-web-api/api_${OS}_${ARCH}${EXT}"

if [ "$OS" = "linux" ] && hash systemctl; then
  install_dir=/usr/local/bin
  share_dir=/usr/local/share
  mkdir -p $install_dir $share_dir
  install_path="$install_dir/lirc-web-api"
  document_root="$share_dir/lirc-web-frontend"
  sudo mv /tmp/lirc-web-api/lirc-web-api $install_path
  sudo mv /tmp/lirc-web-api/frontend $share_dir

  cat <<EOF | sudo tee /lib/systemd/system/lirc-web-api.service
[Unit]
Description=lirc-web-api
After=network.target network-online.target

[Service]
Type=simple
ExecStart=$install_path -p 3000 -f $document_root
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
  echo "installed at $(pwd)/lirc-web-api $("./lirc-web-api/api_${OS}_${ARCH}${EXT}" -v)"
  echo "Usage: ./lirc-web-api/api_${OS}_${ARCH}${EXT} -p 3000 -f ./lirc-web-api/frontend"
fi
