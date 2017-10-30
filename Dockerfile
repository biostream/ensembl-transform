FROM golang:1.9.1-stretch

COPY run.go /etc
COPY src/bmeg /etc/src/bmeg

ENV GOPATH /etc

RUN go get github.com/blachlylab/gff3
RUN go get github.com/golang/protobuf/jsonpb
RUN go get github.com/golang/protobuf/proto

WORKDIR /opt
