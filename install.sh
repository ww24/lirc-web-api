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
INSTALL_TMP_DIR=/tmp/lirc-web-api
mkdir -p $INSTALL_TMP_DIR
tar xzvf /tmp/lirc-web-api.tar.gz -C $INSTALL_TMP_DIR
rm -f /tmp/lirc-web-api.tar.gz
chmod +x "/tmp/lirc-web-api/api_${OS}_${ARCH}${EXT}"

if [ "$OS" = "linux" ] && hash systemctl; then
  install_dir=/usr/local/bin
  share_dir=/usr/local/share/lirc-web-api
  sudo mkdir -p $install_dir $share_dir
  install_path="$install_dir/lirc-web-api"
  document_root="$share_dir/frontend"
  sudo mv "$INSTALL_TMP_DIR/api_${OS}_${ARCH}${EXT}" $install_path
  sudo rsync -ru $INSTALL_TMP_DIR/frontend $share_dir
  rm -rf $INSTALL_TMP_DIR

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

  sudo systemctl daemon-reload
  sudo systemctl enable lirc-web-api
  sudo systemctl restart lirc-web-api
  sudo systemctl status lirc-web-api
else
  rsync -ru $INSTALL_TMP_DIR .
  rm -r $INSTALL_TMP_DIR
  echo "installed at $(pwd)/lirc-web-api $("./lirc-web-api/api_${OS}_${ARCH}${EXT}" -v)"
  echo "Usage: ./lirc-web-api/api_${OS}_${ARCH}${EXT} -p 3000 -f ./lirc-web-api/frontend"
fi
