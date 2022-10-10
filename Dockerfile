FROM golang:1.18-bullseye

RUN apt-get update && apt-get upgrade && apt-get install -y unzip

ARG protoc_release=https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

ADD $protoc_release /opt/protoc.zip
RUN unzip /opt/protoc.zip

WORKDIR ./src/github.com/sensu/core
