From docker-registry:5000/centos:centos6

# Yum srouce
RUN sed -i "s|enabled=1|enabled=0|g" /etc/yum/pluginconf.d/fastestmirror.conf
RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo https://www.xmpan.com/Centos-6-Vault-Aliyun.repo 
RUN yum clean all
RUN yum makecache

RUN yum install -y rpm-build git vim gcc wget

# Golang
RUN cd /tmp && curl -OL https://golang.google.cn/dl/go1.14.13.src.tar.gz
RUN cd /tmp && curl -OL https://golang.google.cn/dl/go1.4.3.linux-amd64.tar.gz
RUN mkdir /root/go1.4 \
        && cat /tmp/go1.4.3.linux-amd64.tar.gz | tar zxC /root/go1.4 --strip-components 1 \
        && cat /tmp/go1.14.13.src.tar.gz | tar xzC /opt \
        && cd /opt/go/src \
        && GO_GCFLAGS=-N ./make.bash \
        && rm -rf /root/go1.4 \
        && mkdir -p /balm /ha /common
RUN rm -fr /tmp/go1.14.13.src.tar.gz
RUN rm -fr /tmp/go1.4.3.linux-amd64.tar.gz

RUN yum clean all
RUN yum install -y make

ENV GOROOT /opt/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH

# Config
RUN yum clean all
RUN rpm --rebuilddb && yum install -y sudo
RUN echo "go ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
RUN sed -i "s/Defaults    requiretty/#Defaults    requiretty/g" /etc/sudoers

# Rpm
RUN yum clean all
RUN yum install -y zlib-devel nss-devel file-devel popt-devel libarchive-devel db4-devel lua-devel elfutils-devel --nogpgcheck
RUN cd /tmp && curl -OL https://www.libarchive.org/downloads/libarchive-3.3.3.tar.gz && tar zxf libarchive-3.3.3.tar.gz
RUN cd /tmp/libarchive-3.3.3 && ./configure && make install
RUN rm -fr /tmp/libarchive-3.3.3.tar.gz && rm -fr /tmp/libarchive-3.3.3
RUN cd /tmp; curl -OL http://ftp.rpm.org/releases/rpm-4.12.x/rpm-4.12.0.2.tar.bz2 && tar jxf rpm-4.12.0.2.tar.bz2
RUN cd /tmp/rpm-4.12.0.2 && ./configure CPPFLAGS="-I/usr/include/nspr4 -I/usr/include/nss3 -I/usr/include/db4" LIBS='-L/usr/local/lib/' --with-external-db && make install
RUN rm -fr /tmp/rpm-4.12.0.2.tar.bz2 && rm -fr /tmp/rpm-4.12.0.2


# Node
RUN cd /tmp/ && curl -OL https://nodejs.org/dist/v10.13.0/node-v10.13.0-linux-x64.tar.xz
RUN cd /tmp/ && tar -xvf /tmp/node-v10.13.0-linux-x64.tar.xz && mv /tmp/node-v10.13.0-linux-x64 /opt/node-v10.13.0-linux-x64
ENV PATH $PATH:/opt/node-v10.13.0-linux-x64/bin
ENV NODE_PATH /opt/node-v10.13.0-linux-x64/lib/node_modules
RUN rm -fr /tmp/node-v10.13.0-linux-x64.tar.xz

# Scspell
COPY scspell-master.tar.gz /tmp/
RUN yum clean all
RUN yum install -y python-setuptools python-argparse --nogpgcheck
RUN cd /tmp; tar zxf scspell-master.tar.gz
RUN cd /tmp/scspell-master && python setup.py install && cp /usr/bin/scspell /usr/local/bin/
RUN rm -fr /tmp/scspell-master

VOLUME /universe
COPY * /docker-build/

ENTRYPOINT ["/bin/bash"]

