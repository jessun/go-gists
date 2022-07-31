FROM ubuntu:16.04 
RUN apt-get update
RUN apt-get install -y rpm git vim gcc make 

# config
RUN apt-get install sudo
RUN echo "go ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
RUN sed -i "s/Defaults    requiretty/#Defaults    requiretty/g" /etc/sudoers

# Docker
RUN apt-get update
RUN apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN apt-get update
RUN add-apt-repository \
 "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
RUN apt-get update
RUN apt-get install -y docker-ce=18.06.1~ce~3-0~ubuntu
RUN echo '{"insecure-registries":["docker-registry:5000"]}' >> /etc/docker/daemon.json
RUN systemctl enable docker

# Golang
RUN cd /tmp/ && curl -OL https://golang.google.cn/dl/go1.14.13.src.tar.gz
RUN cd /tmp/ && curl -OL https://golang.google.cn/dl/go1.4.3.linux-amd64.tar.gz

RUN mkdir /root/go1.4 \
        && cat /tmp/go1.4.3.linux-amd64.tar.gz | tar zxC /root/go1.4 --strip-components 1 \
        && cat /tmp/go1.14.13.src.tar.gz | tar xzC /opt \
        && cd /opt/go/src \
        && GO_GCFLAGS=-N ./make.bash \
        && rm -rf /root/go1.4 \
        && rm -fr /tmp/go1.14.13.src.tar.gz \
        && rm -fr /tmp/go1.4.3.linux-amd64.tar.gz

RUN apt-get install -y dos2unix
ENV GOROOT /opt/go
ENV GOBIN /universe/bin
ENV PATH $PATH:$GOROOT/bin:$GOBIN



# Node
RUN cd /tmp/ && curl -OL https://nodejs.org/dist/v10.13.0/node-v10.13.0-linux-x64.tar.xz
RUN cd /tmp; tar zxf /tmp/node-v10.13.0-linux-x64.tar.gz && mv /tmp/node-v10.13.0-linux-x64 /opt/node-v10.13.0-linux-x64
ENV PATH $PATH:/opt/node-v10.13.0-linux-x64/bin
ENV NODE_PATH /opt/node-v10.13.0-linux-x64/lib/node_modules
RUN rm -fr /tmp/node-v10.13.0-linux-x64.tar.xz

# Glibc
RUN apt-get install libc6
RUN ln -s /lib/x86_64-linux-gnu/libc.so.6 /lib64/libc.so.6

RUN apt-get clean

VOLUME /universe
COPY * /docker-build/

ENTRYPOINT ["/bin/bash"]

