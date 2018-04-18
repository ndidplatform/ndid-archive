package idp

type Request struct {
	NameSpace      string `json:"namespace"`
	Identifier     string `json:"identifier"`
	RequestMessage string `json:"request_message"`
	RequestId      string `json:"request_id"`
}