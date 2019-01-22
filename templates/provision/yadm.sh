#! /bin/bash
set -e

sudo apt update
sudo apt install --yes yadm
yadm clone --bootstrap git@github.com:bketelsen/dotfiles
