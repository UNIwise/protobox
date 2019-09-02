FROM golang:rc-alpine3.10 as builder

RUN apk add --no-cache git

RUN go get github.com/golang/protobuf/protoc-gen-go &&\
    go get github.com/uber/prototool/cmd/prototool

FROM alpine:3.10 

RUN apk --no-cache add protobuf

COPY --from=builder /go/bin/protoc-gen-go /usr/bin
COPY --from=builder /go/bin/prototool /usr/bin