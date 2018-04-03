# Identity
## GET /check_identifier

* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"
```json
{
 "Namespace":"cid",
 "Identifier":"1234567890123"
}
```
===
### Expected response
* Status: `200`
* Content-Type: "text/plain; charset=utf-8"
// * Content-Type: "application/json; charset=utf-8"
```json
{
  "result": "no"
}
```

## POST /create_identity_with_pub_key
* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"

```json
{
 "Namespace": "cid",
 "Identifier":"1234567890123"
}
```

===
### Expected response
// * Status: `201`
// * Content-Type: "application/json; charset=utf-8"
* Status: `200`
* Content-Type: "text/plain; charset=utf-8"

```json
{"result":"success"}
```
