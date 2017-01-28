#!/bin/sh

OS=$(uname | tr '[:upper:]' '[:lower:]')

if [ "$OS" = "linux" ] && hash systemctl; then
  install_dir=/usr/local/bin
  share_dir=/usr/local/share/lirc-web-api
  install_path="$install_dir/lirc-web-api"
  document_root="$share_dir/frontend"

  sudo systemctl stop lirc-web-api
  sudo systemctl disable lirc-web-api
  sudo rm /lib/systemd/system/lirc-web-api.service
  sudo systemctl daemon-reload
  sudo rm $install_path
  sudo rm -r $document_root
else
  rm -r ./lirc-web-api
fi
