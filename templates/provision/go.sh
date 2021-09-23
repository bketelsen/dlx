#! /bin/bash
set -e
sudo rm -rf /usr/local/go
sudo mkdir -p /usr/local/go
sudo sh -c 'curl -Ls https://storage.googleapis.com/golang/go1.17.1.linux-amd64.tar.gz | tar xvzf - -C /usr/local/go --strip-components=1'
