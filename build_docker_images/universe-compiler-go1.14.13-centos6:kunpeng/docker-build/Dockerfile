From docker-registry:5000/centos:centos6

# Yum source
RUN sed -i "s|enabled=1|enabled=0|g" /etc/yum/pluginconf.d/fastestmirror.conf
RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo https://www.xmpan.com/Centos-6-Vault-Aliyun.repo 
RUN yum clean all
RUN yum makecache

RUN yum install -y rpm-build git vim gcc wget

# Golang
RUN cd /tmp/ && curl -OL https://golang.google.cn/dl/go1.14.13.src.tar.gz
RUN cd /tmp/ && curl -OL  https://golang.google.cn/dl/go1.4.3.linux-amd64.tar.gz

RUN mkdir /root/go1.4 \
        && cat /tmp/go1.4.3.linux-amd64.tar.gz | tar zxC /root/go1.4 --strip-components 1 \
        && cat /tmp/go1.14.13.src.tar.gz | tar xzC /opt \
        && cd /opt/go/src \
        && GO_GCFLAGS=-N ./make.bash \
        && rm -rf /root/go1.4 \
        && mkdir -p /balm /ha /common \
        && rm -fr /tmp/go1.14.13.src.tar.gz \
        && rm -fr /tmp/go1.4.3.linux-amd64.tar.gz

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
RUN cd /tmp; curl -OL https://www.libarchive.org/downloads/libarchive-3.3.3.tar.gz && tar zxf libarchive-3.3.3.tar.gz
RUN cd /tmp/libarchive-3.3.3 && ./configure && make install
RUN cd /tmp; curl -OL http://ftp.rpm.org/releases/rpm-4.12.x/rpm-4.12.0.2.tar.bz2 && tar jxf rpm-4.12.0.2.tar.bz2
RUN cd /tmp/rpm-4.12.0.2; ./configure CPPFLAGS="-I/usr/include/nspr4 -I/usr/include/nss3 -I/usr/include/db4" LIBS='-L/usr/local/lib/' --with-external-db && make install
RUN rm -fr /tmp/libarchive-3.3.3.tar.gz && rm -fr /tmp/libarchive-3.3.3
RUN rm -fr /tmp/rpm-4.12.0.2.tar.bz2 && rm -fr /tmp/rpm-4.12.0.2

# Node
RUN cd /opt/ && curl -OL https://nodejs.org/dist/v10.13.0/node-v10.13.0-linux-x64.tar.xz
RUN cd /opt; tar zxf /opt/node-v10.13.0-linux-x64.tar.gz
RUN rm -fr /opt/node-v10.13.0-linux-x64.tar.xz
ENV PATH $PATH:/opt/node-v10.13.0-linux-x64/bin
ENV NODE_PATH /opt/node-v10.13.0-linux-x64/lib/node_modules

# Scspell
COPY scspell-master.tar.gz /opt/
RUN yum clean all
RUN yum install -y python-setuptools python-argparse --nogpgcheck
RUN cd /opt; tar zxf scspell-master.tar.gz
RUN cd /opt/scspell-master/; python setup.py install && cp /usr/bin/scspell /usr/local/bin/

# Gcc and Glibc
RUN cd /root && curl -OL https://releases.linaro.org/components/toolchain/binaries/7.3-2018.05/aarch64-linux-gnu/gcc-linaro-7.3.1-2018.05-x86_64_aarch64-linux-gnu.tar.xz && tar -xvf ./gcc-linaro-7.3.1-2018.05-x86_64_aarch64-linux-gnu.tar.xz
RUN cd /tmp/ && curl -OL http://ftp.gnu.org/gnu/glibc/glibc-2.14.tar.xz \
        && tar -xvf ./glibc-2.14.tar.xz && cd glibc-2.14 \
        && mkdir build && cd build && ../configure --prefix=/opt/glibc-2.14 && make && make install
RUN rm -fr /lib64/lib.so.6 \ 
        && LD_PRELOAD=/opt/glibc-2.14/lib/libc-2.14.so ln -s /opt/glibc-2.14/lib/libc-2.14.so /lib64/libc.so.6 -f
RUN rm -fr /tmp/glibc-2.14.tar.xz && rm -fr /tmp/glibc-2.14

VOLUME /universe
COPY * /docker-build/

ENTRYPOINT ["/bin/bash"]

