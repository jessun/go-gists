From centos:7

LABEL maintainer="JESSSUN<jessun.zhang@gmail.com>"

COPY ./* /build/

ENV WORK_DIR=/workbench

# ENV https_proxy=
# ENV https_proxy=

# CENTOS
RUN mkdir -p ${WORK_DIR} \
    && cp -a /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak \
    && curl -o /etc/yum.repos.d/CentOS-Base.repo  /etc/yum.repos.d/CentOS-Base.repo https://repo.huaweicloud.com/repository/conf/CentOS-7-reg.repo \
    && yum clean all \
    && yum makecache \ 
    && yum update -y \
    && yum -y install wget ncurses ncurses-devel libaio-devel gmp gmp-devel mpfr mpfr-devel \
       libmpc libmpc-devel zlib-devel net-tools cmake openssl openssl-devel gcc gcc-c++ \
       numactl numactl-devel kernel-devel tcl bzip2 \
    && yum clean all 

# GCC
RUN cd $WORK_DIR \
   && curl -O https://mirrors.tuna.tsinghua.edu.cn/gnu/gcc/gcc-9.4.0/gcc-9.4.0.tar.gz \
       && tar -xvf gcc-9.4.0.tar.gz \
           && rm -fr gcc-9.4.0.tar.gz \
       && cd gcc-9.4.0 \
           && sh ./contrib/download_prerequisites \
           &&  ./configure --disable-multilib --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --enable-bootstrap && make -j 96 && make -j 96 install \
           && rm -fr ./gcc-9.4.0
