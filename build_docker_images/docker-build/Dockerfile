FROM ubuntu:16.04 
RUN apt-get update
RUN apt-get install -y rpm git vim wget gcc make 

COPY go1.4.3.linux-amd64.tar.gz /opt/
COPY go1.14.13.src.tar.gz /opt/

RUN mkdir /root/go1.4 \
        && cat /opt/go1.4.3.linux-amd64.tar.gz | tar zxC /root/go1.4 --strip-components 1 \
        && cat /opt/go1.14.13.src.tar.gz | tar xzC /opt \
        && cd /opt/go/src \
        && GO_GCFLAGS=-N ./make.bash \
        && rm -rf /root/go1.4

RUN apt-get install -y dos2unix
ENV GOROOT /opt/go
ENV GOBIN /universe/bin

# sudoers
RUN apt-get install sudo
RUN echo "go ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
RUN sed -i "s/Defaults    requiretty/#Defaults    requiretty/g" /etc/sudoers

ENV PATH $PATH:$GOROOT/bin:$GOBIN
# install docker
RUN echo "1"
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
RUN echo "2"
RUN apt-get update
RUN apt-get install -y docker-ce=18.06.1~ce~3-0~ubuntu
COPY docker-daemon.json /etc/docker/daemon.json
RUN systemctl enable docker
VOLUME /universe
RUN echo "3"
COPY * /docker-build/
RUN mkdir /docker-build/node
RUN mv /docker-build/node-v10.13.0 /docker-build/node/node-v10.13.0

# libc
RUN apt-get install libc6
RUN ln -s /lib/x86_64-linux-gnu/libc.so.6 /lib64/libc.so.6

ENV PATH $PATH:/docker-build/node/node-v10.13.0/bin
ENV NODE_PATH /docker-build/node/node-v10.13.0/lib/node_modules

ENTRYPOINT ["/bin/bash"]

