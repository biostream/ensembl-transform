FROM golang:1.9.1-stretch

RUN mkdir /command /in /out
COPY run.go /command
COPY src/bmeg /command/src/bmeg

ENV GOPATH /command

RUN go get github.com/blachlylab/gff3
RUN go get github.com/golang/protobuf/jsonpb
RUN go get github.com/golang/protobuf/proto

WORKDIR /out
