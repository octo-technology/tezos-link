#!/bin/bash -ex

sudo dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
sudo dnf install docker-ce unzip --nobest -y
sudo systemctl enable --now docker
sudo usermod -aG docker ec2-user

sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

sudo mkfs.ext4 -E discard /dev/nvme0n1
sudo mount /dev/nvme0n1 /var/lib/docker/volumes

aws s3 cp s3://tzlink-blockchain-data-dev/QmXinCsMZHrBwvHbjdUbETFA4Y26LmM5BLCYKQzJ4GcwJx.gz /var/lib/docker/volumes/QmXinCsMZHrBwvHbjdUbETFA4Y26LmM5BLCYKQzJ4GcwJx.gz

cd /var/lib/docker/volumes
tar xvf QmXinCsMZHrBwvHbjdUbETFA4Y26LmM5BLCYKQzJ4GcwJx.gz

mkdir -p mainnet_node_data/_data/data
echo '2018-06-30T16:07:32Z-betanet' > mainnet_node_data/_data/alphanet_version
echo '{ "version": "0.0.4" }' > mainnet_node_data/_data/data/version.json
mv data/tezos/mainnet/context mainnet_node_data/_data/data/context
mv data/tezos/mainnet/store mainnet_node_data/_data/data/store
chown -R 100:65533 mainnet_node_data

curl -o /usr/local/bin/mainnet.sh https://gitlab.com/tezos/tezos/raw/babylonnet/scripts/alphanet.sh
sudo chmod +x /usr/local/bin/mainnet.sh

mainnet.sh node start --rpc-port 8000 --history-mode archive