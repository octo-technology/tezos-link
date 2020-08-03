#!/bin/bash -ex

# Setup Docker, aws CLI and utilitary tools

dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
dnf install docker-ce unzip jq --nobest -y
systemctl enable --now docker
usermod -aG docker ec2-user

curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/bin/docker-compose
chmod +x /usr/bin/docker-compose

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
./aws/install

# Setup lambda ssh key

mkdir -p ~/.ssh && chmod 700 ~/.ssh
bash -c '
echo "${lambda_public_key}" >> /home/ec2-user/.ssh/authorized_keys
'

# Install Tezos docker-compose wrapper

curl -o /usr/bin/${network}.sh https://gitlab.com/tezos/tezos/raw/latest-release/scripts/tezos-docker-manager.sh
chmod +x /usr/bin/${network}.sh

cd /home/ec2-user

# Stop the autoscaling alarm based on CPU

aws autoscaling suspend-processes --auto-scaling-group-name tzlink-${network}-rolling --scaling-processes AlarmNotification

# Import snapshot from the S3 bucket

aws s3 cp s3://tzlink-blockchain-data-dev/${network}_rolling-snapshot.tar.gz ${network}_rolling-snapshot.tar.gz
tar xvf ${network}_rolling-snapshot.tar.gz
rm ${network}_rolling-snapshot.tar.gz

# Dry start to load the snapshot

${network}.sh node start --rpc-port 8000 --history-mode experimental-rolling
${network}.sh stop

if [ -f "/var/lib/docker/volumes/${network}_node_data/_data/data/lock" ]; then
  sudo rm -f /var/lib/docker/volumes/${network}_node_data/_data/data/lock
fi

if [ -f "/var/lib/docker/volumes/${network}_node_data/_data/data/context" ]; then
  sudo rm -rf /var/lib/docker/volumes/${network}_node_data/_data/data/context
fi

if [ -f "/var/lib/docker/volumes/${network}_node_data/_data/data/store" ]; then
  sudo rm -rf /var/lib/docker/volumes/${network}_node_data/_data/data/store
fi

${network}.sh snapshot import /home/ec2-user/snapshot.rolling

# Start tezos node in rolling mode

${network}.sh node start --rpc-port 8000 --history-mode experimental-rolling

rm snapshot.rolling

# Setup cronjob to avoid the "timed out" problem on the node

cat > health-logs-analysis.sh << 'EOF'
#!/bin/bash

if [[ $(curl localhost:8000/chains/main/blocks/head -w %%{time_total} -o /dev/null --silent) > 2 ]]; then
  echo "$(date -R) - health-logs-analysis - Warning : Dysfunctionment detected. Restarting the node."
  mkdir /root/.tezos-mainnet/
  cp -r /.tezos-mainnet/docker-compose.yml /root/.tezos-mainnet/docker-compose.yml
  ${network}.sh stop

  echo -n "Remove peers.json file"
  if [ -f "/var/lib/docker/volumes/mainnet_node_data/_data/data/peers.json" ]; then
    rm -f /var/lib/docker/volumes/mainnet_node_data/_data/data/peers.json
    echo "removed"
  else
    echo "absent. (doing nothing)"
  fi

  sleep 40s

  ${network}.sh node start --rpc-port 8000 --history-mode experimental-rolling

else

  echo "$(date -R) - health-logs-analysis - Node healthy, doing nothing"

fi
EOF
chmod 755 health-logs-analysis.sh
echo "*/2 * * * * /home/ec2-user/health-logs-analysis.sh" | crontab -

# Restart autoscaling CPU alarm
date -R
sleep 4m
aws autoscaling resume-processes --auto-scaling-group-name tzlink-${network}-rolling