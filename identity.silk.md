# Identity
## GET /identity/cid/id/1234567890123

* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"
===
### Expected response
* Status: `200`
* Content-Type: "application/json; charset=UTF-8"
* Data.id: ""
* Data.result.response.value: ""
* Data.result.response.log: "does not exist"

## POST /identity
* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"

```json
{
 "namespace": "cid",
 "id":"1234567890123"
}
```

===
### Expected response
* Status: `201`
* Content-Type: "application/json; charset=UTF-8"

```json
{"result":"success"}
```

## GET /identity/cid/id/1234567890123

* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"
===
### Expected response
* Status: `200`
* Content-Type: "application/json; charset=UTF-8"
* Data.id: ""
* Data.result.response.value: /(.*)/
* Data.result.response.log: "exists"
* Data.result.response.key: /(.*)/
