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

## Run ABCI app and API
1. open 4 terminal window
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run abci/server.go tcp://127.0.0.1:46658`
1. `tendermint init && tendermint unsafe_reset_all && tendermint node --consensus.create_empty_blocks=false`
1. `cd $GOPATH/src/github.com/ndidplatform/ndid` and then `go run api/server.go`

## Create and check identity
### Check identity
1. `curl http://127.0.0.1:8000/identity/cid/1234567890123` should got result `"{... "log":"does not exists" ...}`
### Create identity
1. `curl http://127.0.0.1:8000/identity -X POST -H 'Content-Type: application/json' -d "{ \"namespace\":\"cid\", \"id\":\"1234567890123\"}"` should got result `{... "log":"success" ...}`
### Check identity after created
1. `curl http://127.0.0.1:8000/identity/cid/1234567890123` should got result `{... "log":"exists" ...}`



## WIP
// Scenario: Authentication only
 ```sh
curl -XPOST -H "Content-type: application/json" -d '{
  "data_request_list": [],
  "request_message": "ขอยืนยันตัวตน",
  "min_ial": 2,
  "min_aal": 1,
  "min_idp": 1,
  "timeout": 4320,
  "reference_id": "e3cb44c9-8848-4dec-98c8-8083f373b1f7",
  "call_back_url": ""
}' 'http://127.0.0.1:8000/rp/requests/cid/1234567890123'
 ```

// Scenario: ขอ Statement จาก 1 bank เพื่อขอ**วีซ่า**
 ```sh
curl -XPOST -H "Content-type: application/json" -d '{
  "data_request_list": [
    {
      "service_id": "bank_statement",
      "as_id": "AS1",
      "request_params": {
        "format": "pdf",
        "language": "en"
      }
    }
  ],
  "request_message": "ขอ Bank statement เพื่อทำ VISA ที่สถานฑูตลาว",
  "min_ial": 2,
  "min_aal": 1,
  "min_idp": 1,
  "timeout": 4320,
  "reference_id": "e3cb44c9-8848-4dec-98c8-8083f373b1f7",
  "call_back_url": ""
}' 'http://127.0.0.1:8000/rp/requests/cid/1234567890123'
 ```
