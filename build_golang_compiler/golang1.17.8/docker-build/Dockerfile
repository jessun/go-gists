FROM reg.actiontech.com/actiontech/centos:centos6

ENV http_proxy "http://172.17.0.1:1087"
ENV HTTP_PROXY "http://172.17.0.1:1087"
ENV https_proxy "http://172.17.0.1:1087"
ENV HTTPS_PROXY "http://172.17.0.1:1087"

ENV WORK_DIR=/root

RUN sed -i "s|enabled=1|enabled=0|g" /etc/yum/pluginconf.d/fastestmirror.conf && \
    mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup && \
    curl -o /etc/yum.repos.d/CentOS-Base.repo https://www.xmpan.com/Centos-6-Vault-Aliyun.repo  && \
    yum makecache

RUN yum install -y rpm-build git vim wget gcc make sudo && yum clean all

## sudo 
RUN rpm --rebuilddb && yum install -y sudo && \
    echo "go ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers && \
    sed -i "s/Defaults    requiretty/#Defaults    requiretty/g" /etc/sudoers

# gcc
RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    curl -O https://mirrors.tuna.tsinghua.edu.cn/gnu/gcc/gcc-9.4.0/gcc-9.4.0.tar.gz && \
    tar -xvf gcc-9.4.0.tar.gz && rm -fr gcc-9.4.0.tar.gz && \
    cd gcc-9.4.0 && sh ./contrib/download_prerequisites  && \
    ./configure --disable-multilib --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --enable-bootstrap && \
    make -j 96 && make -j 96 install && \
    rm -fr ./gcc-9.4.0

# go1.4
COPY go1.4.3.linux-amd64.tar.gz /opt/
RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    mkdir -p /root/go1.4 &&  tar -xvf ./go1.4.3.linux-amd64.tar.gz -C /root/go1.4 --strip-components 1 && \
    rm -fr /root/download/go1.4.3.linux-amd64.tar.gz


RUN yum install -y zlib-devel nss-devel file-devel popt-devel libarchive-devel db4-devel lua-devel elfutils-devel --nogpgcheck && \
    cd /opt; wget https://www.libarchive.org/downloads/libarchive-3.3.3.tar.gz && tar zxf libarchive-3.3.3.tar.gz && \
    cd /opt/libarchive-3.3.3; ./configure && make install && \
    cd /opt; wget http://ftp.rpm.org/releases/rpm-4.12.x/rpm-4.12.0.2.tar.bz2 && tar jxf rpm-4.12.0.2.tar.bz2 && \
    cd /opt/rpm-4.12.0.2; ./configure CPPFLAGS="-I/usr/include/nspr4 -I/usr/include/nss3 -I/usr/include/db4" LIBS='-L/usr/local/lib/' --with-external-db && make install && \
    yum clean all

COPY node-v10.13.0-linux-x64.tar.gz /opt/node-v10.13.0-linux-x64.tar.gz
RUN cd /opt; tar zxf /opt/node-v10.13.0-linux-x64.tar.gz

COPY scspell-master.tar.gz /opt/
RUN yum clean all
RUN yum install -y python-setuptools python-argparse --nogpgcheck
RUN cd /opt; tar zxf scspell-master.tar.gz
RUN cd /opt/scspell-master/; python setup.py install && cp /usr/bin/scspell /usr/local/bin/

VOLUME /universe
COPY * /docker-build/

ENV PATH $PATH:/opt/node-v10.13.0-linux-x64/bin
ENV NODE_PATH /opt/node-v10.13.0-linux-x64/lib/node_modules

# go
ARG go_version=go_version
RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    wget https://go.dev/dl/go${go_version}.src.tar.gz -O go_src.tar.gz && \
    tar -xvf ./go_src.tar.gz -C /opt && \
    cd /opt/go/src && GO_GCFLAGS=-N ./make.bash && \
    rm -fr /root/go1.4 && rm -fr ${WORK_DIR}/download/go_src.tar.gz

ENV GOROOT /opt/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH

ENTRYPOINT ["/bin/bash"]
