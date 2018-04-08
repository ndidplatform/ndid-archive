package idp

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	cmn "github.com/tendermint/tmlibs/common"
)

type Request struct {
	RequestID  uuid.UUID `json:"request_id"`
	NameSpace  string    `json:"namespace"`
	Identifier string    `json:"identifier"`
	Secret     string    `json:"secret"`
	Loa        int       `json:"loa"`
	Approval   string    `json:"approval"`
	Signature  string    `json:"signature"`
	AccessorID uuid.UUID `json:"accessor_id"`
}

type Response struct {
	RequestID     uuid.UUID `json:"request_id"`
	AAL           int       `json:"aal"`
	IAL           int       `json:"ial"`
	Approval      string    `json:"approval"`
	Signature     string    `json:"signature"`
	AccessorID    uuid.UUID `json:"accessor_id"`
	IdentityProof string    `json:"identity_proof"`
}

type ResponseDeliver struct {
	Result struct {
		Height  int `json:"height"`
		CheckTx struct {
			Fee struct{} `json:"fee"`
		} `json:"check_tx"`
		DeliverTx struct {
			Log string   `json:"log"`
			Fee struct{} `json:"fee"`
		} `json:"deliver_tx"`
		Hash string `json:"hash"`
	} `json:"result"`
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
}

func CreateResponse(c echo.Context) error {
	// TODO: Validate some input & business logic
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return err
	}

	// TODO: Create blockchain request
	response := createResponse()
	/*
		path := buildBroadcastPath(response)
		var body ResponseDeliver
		t := tendermint.New(path)
		if err := t.Decode(&body); err != nil {
			return err
		}
	*/
	// return api response
	return c.JSON(http.StatusOK, response)
}

func createResponse() (response Response) {
	// convert some uuid
	requestID, _ := uuid.FromString("ef6f4c9c-818b-42b8-8904-3d97c4c520f6")
	response.RequestID = requestID
	response.AAL = 3
	response.IAL = 2
	response.Approval = "CONFIRM"
	response.Signature = "<RSA signature of signing the request_message(ขอ Bank statement...) with accessor type>"
	accessorID, _ := uuid.FromString("12a8f328-53da-4d51-a927-3cc6d3ed3feb")
	response.AccessorID = accessorID
	response.IdentityProof = "very very long number"
	return response
}

func buildBroadcastPath(response Response) (path string) {
	nonce := cmn.RandStr(12)
	fn := "CreateIDPResponse"
	dataBytes, _ := json.Marshal(response)
	tx := []byte("\"" + nonce + "," + fn + "," + string(dataBytes) + "\"")
	path = "/broadcast_tx_commit?tx=" + string(tx)
	return
}
