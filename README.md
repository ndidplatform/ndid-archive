# White Paper:
[Digital Identity Platform Whitepaper (Draft)](https://docs.google.com/document/d/1SKydNM-Nyox62m3vuvYgFYCr8ABVQV8RhjwiMjdCpQ8/edit)

## Prerequire
- Go version >= 1.9.2
- [Install Go](https://golang.org/dl/) follow by [installation instructions.](https://golang.org/doc/install)
- Tendermint >= 0.16.0
- [Install Tendermint](http://tendermint.readthedocs.io/projects/tools/en/master/index.html) follow by [installation instructions.](http://tendermint.readthedocs.io/projects/tools/en/master/install.html)
- [Go dependency management tool](https://github.com/golang/dep)

## Setup
1. `go get github.com/ndidplatform/ndid`
1. `go get -u github.com/tendermint/abci/cmd/abci-cli`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go get ./...`

## Run IdP node
1. open 4 terminal window
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run abci/server.go tcp://127.0.0.1:46000`
1. `tendermint --home ./config/tendermint/IDP unsafe_reset_all && tendermint --home ./config/tendermint/IDP node --consensus.create_empty_blocks=false`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run api/server.go -port :8000 -tenderm 127.0.0.1:45000`
1. test call API `curl http://127.0.0.1:8000/identity/cid/1234567890123`

## Run RP node
1. open 4 terminal window
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run abci/server.go tcp://127.0.0.1:46001`
1. `tendermint --home ./config/tendermint/RP unsafe_reset_all && tendermint --home ./config/tendermint/RP node --consensus.create_empty_blocks=false`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run api/server.go -port :8001 -tenderm 127.0.0.1:45001`
1. test call API `curl http://127.0.0.1:8001/identity/cid/1234567890123`


## Run in Docker
Required
- Docker CE [Install docker](https://docs.docker.com/install/)
- docker-compose [Install docker-compose](https://docs.docker.com/compose/install/)

```
cd docker
docker-compose up
```

## API specs

### Testing
1. run test api for rp `silk -silk.url="http://127.0.0.1:8001" api/specs/relying_party/*.silk.md`
