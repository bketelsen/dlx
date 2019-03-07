#! /bin/bash
set -e


sudo apt update
sudo apt install --yes gconf2 gconf-service libgtk2.0-0 libnotify4 libnss3 python gvfs-bin xdg-utils libxss1 libappindicator1 libsecret-1-0 libasound2
wget -O slack-desktop-3.3.7-amd64.deb https://downloads.slack-edge.com/linux_releases/slack-desktop-3.3.7-amd64.deb
sudo dpkg --install slack-desktop-3.3.7-amd64.deb
rm -f slack-desktop-3.3.7-amd64.deb
