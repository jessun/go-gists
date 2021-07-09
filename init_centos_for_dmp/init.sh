# yum
yum install -y vim curl wget unzip yum-utils git python-devel python3-devel python3-pip openssl-devel gcc libffi-devel libevent-devel gevent docker-compose
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
wget  ftp://ftpuser:ftpuser@10.186.18.90/housekeep/udp-quick-usage.tar.gz
tar -xvf ./udp-quick-usage.tar.gz
cd quick-usage/deploy && make clean && make install

