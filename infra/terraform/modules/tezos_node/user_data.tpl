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

mkdir -p ~/.ssh && chmod 700 ~/.ssh && echo ${lambda_public_key} >>  ~/.ssh/authorized_keys

aws s3 cp s3://tzlink-blockchain-data-dev/${network}_node_data.tar.gz ${network}_node_data.tar.gz
tar xvf ${network}_node_data.tar.gz
mv archive ${network}_node_data
chown -R 100:65533 ${network}_node_data

curl -o /usr/local/bin/${network}.sh https://gitlab.com/tezos/tezos/raw/${computed_network}/scripts/alphanet.sh
chmod +x /usr/local/bin/${network}.sh

cd /home/ec2-user

${network}.sh node start --rpc-port 8000 --history-mode archive

rm -rf /var/lib/docker/volumes/${network}_node_data.tar.gz

cat > export-tezos-snap.sh << EOF
#!/bin/bash -e

mkdir .tezos-${network}
cp /.tezos-${network}/docker-compose.yml .tezos-${network}/docker-compose.yml 

echo "> Stop the node for snapshot"
${network}.sh stop

cd /var/lib/docker/volumes
echo "> Copy ${network}_node_data in archive"
sudo cp -r ${network}_node_data archive

echo "> Restart the node"
${network}.sh node start --rpc-port 8000 --history-mode archive

echo "> Remove files:"

echo -n "- peers.json "
if [ -e "./archive/_data/data/peers.json" ]; then
  sudo rm -f ./archive/_data/data/peers.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo -n "- identity.json "
if [ -e "./archive/_data/data/identity.json" ]; then
  sudo rm -f ./archive/_data/data/identity.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo -n "- config.json "
if [ -e "./archive/_data/data/config.json" ]; then
  sudo rm -f ./archive/_data/data/config.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo "> Generate ${network}_node_data.tar.gz from archive"
sudo tar zcvf ${network}_node_data.tar.gz ./archive

echo "> Send to S3 bucket ${network}_node_data.tar.gz"
aws s3 cp ./${network}_node_data.tar.gz s3://tzlink-blockchain-data-dev

echo "> Clear temporary files"
sudo rm -rf archive ${network}_node_data.tar.gz
EOF

chmod +x export-tezos-snap.sh
