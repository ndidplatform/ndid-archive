# White Paper:
[Digital Identity Platform Whitepaper (Draft)](https://docs.google.com/document/d/1SKydNM-Nyox62m3vuvYgFYCr8ABVQV8RhjwiMjdCpQ8/edit)

## Prerequire
- Go version >= 1.9.2
- [Install Go](https://golang.org/dl/) follow by [installation instructions.](https://golang.org/doc/install)
- Tendermint >= 0.16.0
- [Install Tendermint](http://tendermint.readthedocs.io/projects/tools/en/master/index.html) follow by [installation instructions.](http://tendermint.readthedocs.io/projects/tools/en/master/install.html)

## Setup
1. `go get github.com/ndidplatform/ndid`
1. `go get -u github.com/tendermint/abci/cmd/abci-cli`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go get ./...`

## Run ABCI app and API
1. open 4 terminal window
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run abci/server.go tcp://127.0.0.1:46658`
1. `tendermint init && tendermint unsafe_reset_all && tendermint node --consensus.create_empty_blocks=false`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run api/server.go`

## Create and check identity
### Check identity
1. `curl http://127.0.0.1:8000/identity/cid/id/1234567890123` should got result `"{... "log":"does not exists" ...}`
### Create identity
1. `curl http://127.0.0.1:8000/identity -X POST -H 'Content-Type: application/json' -d "{ \"namespace\":\"cid\", \"id\":\"1234567890123\"}"` should got result `{... "log":"success" ...}`
### Check identity after created
1. `curl http://127.0.0.1:8000/identity/cid/id/1234567890123` should got result `{... "log":"exists" ...}`
