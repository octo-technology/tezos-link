#!/bin/bash -ex

# Setup Docker, aws CLI and utilitary tools

dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
dnf install docker-ce unzip jq --nobest -y
systemctl enable --now docker
usermod -aG docker ec2-user

curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
./aws/install

# Setup lambda ssh key

mkdir -p ~/.ssh && chmod 700 ~/.ssh
bash -c '
echo "${lambda_public_key}" >> /home/ec2-user/.ssh/authorized_keys
'

# Mount I3 special volume

mkfs.ext4 -E discard /dev/nvme0n1
mount /dev/nvme0n1 /var/lib/docker/volumes

# Install Tezos docker-compose wrapper

if [ "${network}" == "mainnet" ]; then
  curl -o /usr/local/bin/${network}.sh https://gitlab.com/tezos/tezos/raw/${computed_network}/scripts/alphanet.sh
else 
  curl -o /usr/local/bin/${network}.sh https://gitlab.com/tezos/tezos/raw/latest-release/scripts/tezos-docker-manager.sh
fi
chmod +x /usr/local/bin/${network}.sh

# Import archive from the S3 bucket

cd /var/lib/docker/volumes

aws s3 cp s3://tzlink-blockchain-data-dev/${network}_node_data.tar.gz ${network}_node_data.tar.gz
tar xvf ${network}_node_data.tar.gz
mv archive ${network}_node_data
chown -R 100:65533 ${network}_node_data
rm -rf ${network}_node_data.tar.gz

cd /home/ec2-user

# Start tezos node in archive mode

${network}.sh node start --rpc-port 8000 --history-mode archive

# Setup export snapshot backup script

cat > export-tezos-snap.sh << 'EOF'
#!/bin/bash -e

aws autoscaling suspend-processes --auto-scaling-group-name tzlink-${network}-archive --scaling-processes ReplaceUnhealthy

mkdir -p .tezos-${network}
cp /.tezos-${network}/docker-compose.yml .tezos-${network}/docker-compose.yml 

echo "> Snapshot rolling-mode node"

current_hash=$(curl -s localhost:8000/chains/main/blocks/head | jq .hash)
echo ">>> Using the block $${current_hash}"
echo ">>> Generate the snapshot.rolling file"
docker exec ${network}_node_1 sh -c "tezos-node snapshot export snapshot.rolling --block $${current_hash} --data-dir /var/run/tezos/node/data --rolling && mv snapshot.rolling /var/run/tezos/client/snapshot.rolling"

cd /var/lib/docker/volumes/${network}_client_data/_data/

echo ">>> Generate ${network}_rolling-snapshot.tar.gz from the snapshot"
sudo tar zcvf ${network}_rolling-snapshot.tar.gz snapshot.rolling

echo ">>> Send to S3 bucket ${network}_rolling-snapshot.tar.gz"
aws s3 cp ${network}_rolling-snapshot.tar.gz s3://tzlink-blockchain-data-dev/${network}_rolling-snapshot.tar.gz

echo ">>> Clear temporary files"
sudo rm ${network}_rolling-snapshot.tar.gz snapshot.rolling

cd -

echo "> Snapshot archive node"

echo ">>> Stop the node for snapshot"
${network}.sh stop

cd /var/lib/docker/volumes
echo ">>> Copy ${network}_node_data in archive"
sudo cp -r ${network}_node_data archive

echo ">>> Restart the node"
${network}.sh node start --rpc-port 8000 --history-mode archive

echo ">>> Remove files:"

echo -n "- peers.json "
if [ -f "./archive/_data/data/peers.json" ]; then
  sudo rm -f ./archive/_data/data/peers.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo -n "- identity.json "
if [ -f "./archive/_data/data/identity.json" ]; then
  sudo rm -f ./archive/_data/data/identity.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo -n "- config.json "
if [ -f "./archive/_data/data/config.json" ]; then
  sudo rm -f ./archive/_data/data/config.json
  echo "removed"
else
  echo "absent. (doing nothing)"
fi

echo ">>> Generate ${network}_node_data.tar.gz from archive"
sudo tar zcvf ${network}_node_data.tar.gz ./archive

echo ">>> Send to S3 bucket ${network}_node_data.tar.gz"
aws s3 cp ./${network}_node_data.tar.gz s3://tzlink-blockchain-data-dev

echo ">>> Clear temporary files"
sudo rm -rf archive ${network}_node_data.tar.gz

aws autoscaling resume-processes --auto-scaling-group-name tzlink-${network}-archive --scaling-processes ReplaceUnhealthy
EOF

chmod +x export-tezos-snap.sh

# Setup export snapshot backup service

cat > /etc/systemd/system/tezos-backup.service << EOF
[Unit]
Description=Backup service for Tezos node to S3

[Service]
Type=simple
Restart=no
ExecStart=/bin/bash /home/ec2-user/export-tezos-snap.sh

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload