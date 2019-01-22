#! /bin/bash
set -e

wget -O slack-desktop-3.3.3-amd64.deb https://downloads.slack-edge.com/linux_releases/slack-desktop-3.3.3-amd64.deb
sudo dpkg --install slack-desktop-3.2.1-amd64.deb
rm, -f, slack-desktop-3.2.1-amd64.deb
apt --yes --fix-broken install
