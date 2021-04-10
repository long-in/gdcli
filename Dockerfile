FROM ubuntu:20.04

RUN apt update && \
    apt list --upgradable && \
    apt -y upgrade && \
    apt -y install wget vim

RUN cd /usr/local/src/ && \
    wget https://golang.org/dl/go1.16.3.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf ./go1.16.3.linux-amd64.tar.gz

ENV PATH $PATH:/usr/local/go/bin

WORKDIR /root/go/src/github.com

CMD ["/usr/bin/bash"]
