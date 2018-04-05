package rp

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/syndtr/goleveldb/leveldb"
	// "github.com/ndidplatform/ndid/api/client/tendermint"
	// cmn "github.com/tendermint/tmlibs/common"
)

type Request struct {
	ServiceID     string                 `json:"service_id"`
	AsID          string                 `json:"as_id"`
	RequestParams map[string]interface{} `json:"request_params"`
}

type Requests struct {
	DataRequestList []Request `json:"data_request_list"`
	RequestMessage  string    `json:"request_message"`
	MinIal          int       `json:"min_ial"`
	MinAal          int       `json:"min_aal"`
	MinIdp          int       `json:"min_idp"`
	Timeout         float64   `json:"timeout"`
	ReferenceID     string    `json:"reference_id"`
	CallBackURL     string    `json:"call_back_url"`
}

type Response struct {
	RequestID uuid.UUID `json:"request_id"`
}

func CreateRequest(c echo.Context) error {

	response := new(Response)
	request := new(Requests)
	if err := c.Bind(request); err != nil {
		return err
	}

	db, err := leveldb.OpenFile("localDB", nil)
	if err != nil {
		return err
	}
	defer db.Close()

	//Input validation
	//Check Ref ID is already exists
	old_request_id_bytes, err := db.Get([]byte(request.ReferenceID), nil)
	if err != nil {
		new_request_id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		// Store (reference ID → request ID) in node’s database
		err = db.Put([]byte(request.ReferenceID), new_request_id.Bytes(), nil)
		if err != nil {
			return err
		}

		// TODO: Create blockchain request

		// TODO: Create messaging request to node_ids

		// Return request_id
		response.RequestID = new_request_id
		return c.JSON(http.StatusCreated, response)
	} else {
		// TODO: return the result of that request
		old_request_id, _ := uuid.FromBytes(old_request_id_bytes)
		response.RequestID = old_request_id
		return c.JSON(http.StatusCreated, response)
	}
}
