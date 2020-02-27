#!/bin/bash -x

sudo dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
sudo dnf install docker-ce --nobest -y
sudo systemctl enable --now docker
sudo usermod -aG docker ec2-user

sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

curl -o /usr/local/bin/mainnet.sh https://gitlab.com/tezos/tezos/raw/babylonnet/scripts/alphanet.sh
sudo chmod +x /usr/local/bin/mainnet.sh

mainnet.sh node start --rpc-port 8000 --history-mode archive