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
	oldRequestId, err := db.Get([]byte(request.ReferenceID), nil)
	if err != nil {
		newRequestId := uuid.NewV4()

		// Store (reference ID → request ID) in node’s database
		err = db.Put([]byte(request.ReferenceID), newRequestId.Bytes(), nil)
		if err != nil {
			return err
		}

		// TODO: Create blockchain request

		// TODO: Create messaging request to node_ids

		return c.JSON(http.StatusCreated, &Response{newRequestId})
	}

	id, err := uuid.FromBytes(oldRequestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, &Response{id})
}
