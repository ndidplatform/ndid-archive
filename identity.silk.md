# Identity
## GET /identity/cid/id/1234567890123

* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"
===
### Expected response
* Status: `200`
* Content-Type: "application/json; charset=UTF-8"
```json
{
  "result": "no"
}
```

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
