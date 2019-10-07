FROM centos

RUN yum update -y \
    && yum install -y epel-release \
    && yum install -y git mercurial subversion wget vim gcc gcc-c++ make bind-utils whois

RUN cd /usr/local/src \
    && wget https://dl.google.com/go/go1.12.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.12.linux-amd64.tar.gz \
    && mkdir -p /root/go/{src,bin,pkg} \
    && mkdir -p /root/go/src/github.com

ENV HOME /root
ENV GOPATH $HOME/go
ENV PATH $PATH:/usr/local/go/bin
ENV GO111MODULE on

WORKDIR /root/go/src/github.com

CMD ["/bin/bash"]
