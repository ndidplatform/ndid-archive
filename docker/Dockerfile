FROM golang:1.10 as builder
LABEL maintainer="Narate Ketram <rate@dome.cloud>"

ENV TERM=xterm-256color
ENV CGO_ENABLED=0
RUN go get github.com/golang/dep/cmd/dep
RUN go get -d github.com/ndidplatform/ndid/...

WORKDIR $GOPATH/src/github.com/ndidplatform/ndid
RUN dep ensure
RUN go install github.com/ndidplatform/ndid/api github.com/ndidplatform/ndid/abci

FROM alpine:3.7
COPY --from=builder /go/bin/api /api-server
COPY --from=builder /go/bin/abci /abci-server
