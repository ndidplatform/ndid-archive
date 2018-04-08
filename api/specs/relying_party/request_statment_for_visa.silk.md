# Scenario: Requesting a bank statement for Visa application
# RP → Platform
## POST /rp/requests/cid/1234567890123
* Content-Type: "application/json; charset=UTF-8"
* Accept: "application/json"

```json
{
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
}
```
===
### Expected response
* Status: `201`
* Content-Type: "application/json; charset=UTF-8"
* Data.request_id: /(.+)/
