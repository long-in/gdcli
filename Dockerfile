FROM centos:7

RUN yum update -y \
    && yum install -y epel-release \
    && yum install -y git mercurial subversion wget vim gcc gcc-c++ make bind-utils whois

RUN cd /usr/local/src \
    && wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz \
    && mkdir -p /root/go/{src,bin,pkg} \
    && mkdir -p /root/go/src/github.com

ENV HOME /root
ENV GOPATH $HOME/go
ENV PATH $PATH:/usr/local/go/bin

WORKDIR /root/go/src/github.com

CMD ["/bin/bash"]
