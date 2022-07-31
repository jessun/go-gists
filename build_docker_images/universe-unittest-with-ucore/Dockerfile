From reg.actiontech.com/actiontech/centos:centos7
RUN echo "nameserver 114.114.114.114" > /etc/resolv.conf
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo
RUN sed -i -e '/mirrors.cloud.aliyuncs.com/d' -e '/mirrors.aliyuncs.com/d' /etc/yum.repos.d/CentOS-Base.repo
RUN yum makecache
RUN yum install -y rpm rpm-build sudo wget
RUN rpm -iv ftp://ftpuser:ftpuser@10.186.18.90/actiontech-ucore/qa/9.9.9.9/ucore-9.9.9.9.standard.qa.el7.x86_64.rpm
RUN echo 'actiontech-universe ALL=NOPASSWD: ALL' >> /etc/sudoers
ENV PATH /usr/local/go/bin:$PATH
COPY ./Dockerfile /docker-build/
COPY ./Makefile /docker-build/
