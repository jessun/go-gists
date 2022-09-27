#! /bin/bash

set -e

# check if centos
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
fi

if [ "$OS" != "CentOS Linux" ]; then
	echo "only support centos, but $OS"
        exit 1
fi


function download_ob(){
curl -C - -O ftp://ftpuser:ftpuser@10.186.18.90//ocean_base_rpms/libobclient-2.0.2-2.el7.x86_64.rpm
curl -C - -O ftp://ftpuser:ftpuser@10.186.18.90//ocean_base_rpms/oceanbase-ce-libs-3.1.4-10000092022071511.el7.x86_64.rpm

curl -C - -O ftp://ftpuser:ftpuser@10.186.18.90//ocean_base_rpms/ob_admin.el7.x86_64
curl -C - -O ftp://ftpuser:ftpuser@10.186.18.90//ocean_base_rpms/ob-deploy-1.5.0-12.el7.x86_64.rpm
curl -C - -O ftp://ftpuser:ftpuser@10.186.18.90//ocean_base_rpms/oceanbase-ce-3.1.4-10000092022071511.el7.x86_64.rpm
rpm -Uvh https://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
yum install mysql-community-client.x86_64

#rpm -ivh libobclient-2.0.2-2.el7.x86_64.rpm
#rpm -ivh oceanbase-ce-libs-3.1.4-10000092022071511.el7.x86_64.rpm
#rpm -ivh ob-deploy-1.5.0-12.el7.x86_64.rpm
#rpm -ivh oceanbase-ce-3.1.4-10000092022071511.el7.x86_64.rpm

cp /home/admin/oceanbase/lib/*  /usr/lib64/
mv ob_admin.el7.x86_64 /home/admin/oceanbase/bin/ob_admin && chmod +x /home/admin/oceanbase/bin/ob_admin
}

function admin_user(){
	useradd -U admin -d /home/admin -s /bin/bash
	mkdir -p /home/admin
	sudo chown -R admin:admin /home/admin
	echo 'admin       ALL=(ALL)       NOPASSWD: ALL' >> /etc/sudoers
}

function common_pkg(){
sudo yum install -y vim chrony wget nfs-utils
}

function update_sysctl (){
echo '# for oceanbase
## 修改内核异步 I/O 限制
fs.aio-max-nr=1048576

## 网络优化
net.core.somaxconn = 2048
net.core.netdev_max_backlog = 10000 
net.core.rmem_default = 16777216 
net.core.wmem_default = 16777216 
net.core.rmem_max = 16777216 
net.core.wmem_max = 16777216

net.ipv4.ip_local_port_range = 3500 65535 
net.ipv4.ip_forward = 0 
net.ipv4.conf.default.rp_filter = 1 
net.ipv4.conf.default.accept_source_route = 0 
net.ipv4.tcp_syncookies = 0 
net.ipv4.tcp_rmem = 4096 87380 16777216 
net.ipv4.tcp_wmem = 4096 65536 16777216 
net.ipv4.tcp_max_syn_backlog = 16384 
net.ipv4.tcp_fin_timeout = 15 
net.ipv4.tcp_max_syn_backlog = 16384 
net.ipv4.tcp_tw_reuse = 1 
net.ipv4.tcp_tw_recycle = 1 
net.ipv4.tcp_slow_start_after_idle=0
sunrpc.tcp_max_slot_table_entries=128

vm.swappiness = 0
vm.min_free_kbytes = 2097152

# 此处为 OceanBase 数据库的 data 目录
kernel.core_pattern = /data/core-%e-%p-%t
' >> /etc/sysctl.conf
sysctl -p
}

function update_limits_conf(){
echo '
root soft nofile 655350
root hard nofile 655350
* soft nofile 655350
* hard nofile 655350
* soft stack 20480
* hard stack 20480
* soft nproc 655360
* hard nproc 655360
* soft core unlimited
* hard core unlimited
'  >> /etc/security/limits.conf
}

function nfs(){
mkdir -p /data
 sudo mount -tnfs4 -o rw,nfsvers=4.1,sync,lookupcache=positive,hard,timeo=600,wsize=1048576,rsize=1048576,namlen=255 10.186.60.75:/data/nfs_server /data
}

common_pkg
#admin_user
#update_sysctl
#update_limits_conf
#download_ob
nfs
