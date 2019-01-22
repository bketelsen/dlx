#! /bin/bash
set -e

sudo systemctl disable --now snapd snapd.socket

sudo apt update
sudo sed -i "s/; enable-shm = yes/enable-shm = no/g" /etc/pulse/client.conf
sudo echo export PULSE_SERVER=unix:/tmp/.pulse-native | tee --append /home/ubuntu/.profile
sudo apt-get install --yes  x11-apps mesa-utils pulseaudio
