#!/bin/bash -ex

dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
dnf install docker-ce unzip --nobest -y
systemctl enable --now docker
usermod -aG docker ec2-user

curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
./aws/install

mkfs.ext4 -E discard /dev/nvme0n1
mount /dev/nvme0n1 /var/lib/docker/volumes

cd /var/lib/docker/volumes

aws s3 cp s3://tzlink-blockchain-data-dev/mainnet_node_data.tar.gz mainnet_node_data.tar.gz
tar xvf mainnet_node_data.tar.gz
mv archive mainnet_node_data
chown -R 100:65533 mainnet_node_data

curl -o /usr/local/bin/mainnet.sh https://gitlab.com/tezos/tezos/raw/babylonnet/scripts/alphanet.sh
chmod +x /usr/local/bin/mainnet.sh

cd /home/ec2-user

mainnet.sh node start --rpc-port 8000 --history-mode archive

rm -rf /var/lib/docker/volumes/mainnet_node_data.tar.gz

cat > tezos-snap.sh << EOF
#!/bin/bash -ex

cp /.tezos-mainnet/docker-compose.yml .tezos-mainnet/docker-compose.yml 

mainnet.sh stop

cd /var/lib/docker/volumes
sudo cp -r mainnet_node_data archive

mainnet.sh node start --rpc-port 8000 --history-mode archive

sudo tar zcvf mainnet_node_data.tar.gz ./archive

sudo aws s3 cp ./mainnet_node_data.tar.gz s3://tzlink-blockchain-data-dev

sudo rm -rf archive mainnet_node_data.tar.gz
EOF

chmod +x tezos-snap.sh