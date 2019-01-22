#! /bin/bash
set -e

sudo systemctl disable --now snapd snapd.socket
sudo apt update
sudo apt install --yes git mercurial yadm build-essential
