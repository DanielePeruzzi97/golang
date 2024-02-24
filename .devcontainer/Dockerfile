FROM ubuntu:22.04

RUN apt update && apt install -y wget
RUN wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin