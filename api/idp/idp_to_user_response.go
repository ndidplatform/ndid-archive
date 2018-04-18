package idp

type Reponse struct {
	RequestId  string `json:"request_id"`
	Namespace  string `json:"namespace"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	Loa        int    `json:"loa"`
	Approval   string `json:"approval"`
	Signature  []byte `json:"signature"`
	AccessorId string `json:"accessor_id"`
}
