FROM golang:1.24.3-bullseye

RUN apt-get update && apt-get upgrade -y && apt-get install -y unzip

ARG protoc_version=3.19.4
ARG protoc_arch=x86_64
ARG protoc_release=https://github.com/protocolbuffers/protobuf/releases/download/v${protoc_version}/protoc-${protoc_version}-linux-${protoc_arch}.zip

ADD $protoc_release /opt/protoc.zip
RUN unzip /opt/protoc.zip

WORKDIR ./src/github.com/sensu/core
