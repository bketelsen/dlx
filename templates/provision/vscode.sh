#! /bin/bash
set -e

sudo apt update
wget -O code.deb https://go.microsoft.com/fwlink/?LinkID=760865
sudo dpkg -i code.deb
sudo apt-get install --yes -f
