FROM golang:1.10 as builder
LABEL maintainer="Narate Ketram <rate@dome.cloud>"
ENV TERM=xterm-256color
ENV CGO_ENABLED=0
WORKDIR /go/src
RUN go get github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/ndidplatform && \
	git clone https://github.com/ndidplatform/ndid /go/src/github.com/ndidplatform/ndid
RUN cd /go/src/github.com/ndidplatform/ndid && \
	/go/bin/dep ensure
RUN go build -o $GOPATH/bin/api-server $GOPATH/src/github.com/ndidplatform/ndid/api/server.go
RUN go build -o $GOPATH/bin/abci-server $GOPATH/src/github.com/ndidplatform/ndid/abci/server.go

FROM alpine:3.7
COPY --from=builder /go/bin/api-server /api-server
COPY --from=builder /go/bin/abci-server /abci-server
