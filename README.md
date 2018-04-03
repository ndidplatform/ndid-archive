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

## Run ABCI app and API
1. open 4 terminal window
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run ABCI/abci-server.go tcp://127.0.0.1:46658`
1. `tendermint init` and `tendermint unsafe_reset_all` and then `tendermint node --consensus.create_empty_blocks=false`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run API/api-server.go 127.0.0.1:46657`

## Create and check identity
### Check identity
1. `curl http://127.0.0.1:8000/check_identifier -X GET -H 'Content-Type: application/json' -d "{ \"Namespace\":\"cid\", \"Identifier\":\"1234567890123\"}"` should got result `{"result":"no"}`
### Create identity
1. `curl http://localhost:8000/create_identity_with_pub_key -X POST -H 'Content-Type: application/json' -d "{ \"Namespace\":\"cid\", \"Identifier\":\"1234567890123\"}"` should got result `{"result":"success"}`
### Check identity after created
1. `curl http://127.0.0.1:8000/check_identifier -X GET -H 'Content-Type: application/json' -d "{ \"Namespace\":\"cid\", \"Identifier\":\"1234567890123\"}"` should got result `{"result":"yes"}`
