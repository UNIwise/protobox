FROM golang:rc-alpine3.10 as builder

RUN apk add --no-cache git
RUN go get github.com/golang/protobuf/protoc-gen-go

FROM node:12-alpine 

RUN apk --no-cache add protobuf && mkdir gen
COPY --from=builder /go/bin/protoc-gen-go /usr/bin

WORKDIR gen

RUN npm install -g ts-protoc-gen