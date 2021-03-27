# yum
yum install -y vim curl wget unzip yum-utils git python-devel python3-devel
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
yum install -y docker-ce docker-ce-cli containerd.io

# hosts
echo "10.186.18.20 docker-registry" >> /etc/hosts

# docker
mkdir -p /etc/docker
echo '{"insecure-registries":["docker-registry:5000"]}' >> /etc/docker/daemon.json
systemctl start docker
systemctl enable docker

# quick-usage
wget  ftp://ftp:ftp@10.186.18.20/housekeep/udp-quick-usage.zip
unzip udp-quick-usage.zip
sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
cd quick-usage/deploy && make install