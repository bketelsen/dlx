#! /bin/bash
set -e

echo Waiting for cloud-init to finish.
cloud-init status --wait
sudo systemctl disable --now snapd snapd.socket
sudo apt update
sudo apt install --yes git mercurial yadm build-essential
