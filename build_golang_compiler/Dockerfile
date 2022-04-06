FROM reg.actiontech.com/actiontech/centos:centos6

ENV WORK_DIR=/root

ENV http_proxy "http://172.17.0.1:1087"
ENV HTTP_PROXY "http://172.17.0.1:1087"
ENV https_proxy "http://172.17.0.1:1087"
ENV HTTPS_PROXY "http://172.17.0.1:1087"

ARG go_version=go_version

RUN yum install -y rpm-build git vim gcc gcc-c++ zlib-devel nss-devel file-devel popt-devel libarchive-devel \
    wget ncurses ncurses-devel libaio-devel gmp gmp-devel mpfr mpfr-devel \
    libmpc libmpc-devel net-tools cmake openssl openssl-devel numactl numactl-devel kernel-devel tcl bzip2 \
    db4-devel lua-devel elfutils-devel --nogpgcheck && \
    yum clean all

RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    wget 'https://go.dev/dl/go1.4.3.linux-amd64.tar.gz' && \
    mkdir -p /root/go1.4 &&  tar -xvf ./go1.4.3.linux-amd64.tar.gz -C /root/go1.4 --strip-components 1 && \
    rm -fr /root/download/go1.4.3.linux-amd64.tar.gz

RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    curl -O https://mirrors.tuna.tsinghua.edu.cn/gnu/gcc/gcc-9.4.0/gcc-9.4.0.tar.gz && \
    tar -xvf gcc-9.4.0.tar.gz && rm -fr gcc-9.4.0.tar.gz && \
    cd gcc-9.4.0 && sh ./contrib/download_prerequisites  && \
    ./configure --disable-multilib --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --enable-bootstrap && \
    make -j 96 && make -j 96 install && \
    rm -fr ./gcc-9.4.0

RUN mkdir -p ${WORK_DIR}/download && cd ${WORK_DIR}/download && \
    wget https://go.dev/dl/go${go_version}.src.tar.gz -O go_src.tar.gz && \
    tar -xvf ./go_src.tar.gz -C /opt && \
    cd /opt/go/src && GO_GCFLAGS=-N ./make.bash && \
    rm -fr /root/go1.4 && rm -fr ${WORK_DIR}/download/go_src.tar.gz

ENV GOROOT /opt/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH

RUN rpm --rebuilddb && yum install -y sudo && \
    echo "go ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers && \
    sed -i "s/Defaults    requiretty/#Defaults    requiretty/g" /etc/sudoers && \
    yum clean all

RUN cd /opt; wget https://www.libarchive.org/downloads/libarchive-3.3.3.tar.gz && tar zxf libarchive-3.3.3.tar.gz && \
    cd /opt/libarchive-3.3.3; ./configure && make install && \
    cd /opt; wget http://ftp.rpm.org/releases/rpm-4.12.x/rpm-4.12.0.2.tar.bz2 && tar jxf rpm-4.12.0.2.tar.bz2 && \
    cd /opt/rpm-4.12.0.2; ./configure CPPFLAGS="-I/usr/include/nspr4 -I/usr/include/nss3 -I/usr/include/db4" LIBS='-L/usr/local/lib/' --with-external-db && make install
