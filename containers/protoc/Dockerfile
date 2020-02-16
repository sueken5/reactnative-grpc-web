FROM golang:1.13.6-buster

ENV GO111MODULE on
RUN apt-get update && apt-get install -y --no-install-recommends \
            autoconf=2.69-11 \
            automake=1:1.16.1-4 \
            libtool=2.4.6-9 \
            && rm -rf /var/lib/apt/lists/*
ENV PROTOBUF_VERSION 3.9.1
RUN wget -O /usr/local/src/protobuf.tar.gz "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
WORKDIR /usr/local/src
RUN tar --no-same-permissions --no-same-owner -xzf protobuf.tar.gz
WORKDIR /usr/local/src/protobuf-${PROTOBUF_VERSION}
RUN ./autogen.sh && ./configure && make && make install && ldconfig

## install go protobuf
ENV GO_PROTOBUF_VERSION 1.3.2
RUN git clone https://github.com/golang/protobuf /go/src/github.com/golang/protobuf
WORKDIR /go/src/github.com/golang/protobuf
RUN git checkout "v${GO_PROTOBUF_VERSION}"
RUN go install github.com/golang/protobuf/protoc-gen-go

## install grpc gateway
ENV GRPC_GATEWAY_VERSION 1.11.1
RUN git clone https://github.com/grpc-ecosystem/grpc-gateway /go/src/github.com/grpc-ecosystem/grpc-gateway
WORKDIR /go/src/github.com/grpc-ecosystem/grpc-gateway
RUN git checkout "v${GRPC_GATEWAY_VERSION}"
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

## install grpc web
ENV GRPC_WEB_VERSION 1.0.7
RUN git clone https://github.com/grpc/grpc-web /grpc-web/
WORKDIR /grpc-web/
RUN git checkout "${GRPC_WEB_VERSION}"
RUN make install-plugin

WORKDIR /