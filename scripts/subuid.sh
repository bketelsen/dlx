#!/bin/bash
echo "root:1000:1" | sudo tee -a /etc/subuid /etc/subgid
