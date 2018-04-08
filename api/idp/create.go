package idp

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type Response struct {
	Status     string    `json:"status"`
	RequestID  uuid.UUID `json:"request_id"`
	NameSpace  string    `json:"namespace"`
	Identifier string    `json:"identifier"`
	Secret     string    `json:"secret"`
	Loa        int       `json:"loa"`
	Approval   string    `json:"approval"`
	Signature  string    `json:"signature"`
	AccessorId uuid.UUID `json:"accessor_id"`
}

func CreateResponse(c echo.Context) error {
	// Initialize response
	response := new(Response)
	// convert some uuid
	request_id, _ := uuid.FromString("ef6f4c9c-818b-42b8-8904-3d97c4c520f6")
	response.RequestID = request_id
	response.NameSpace = "citizenid"
	response.Identifier = "01234567890123"
	response.Secret = "MAGIC"
	response.Loa = 3
	response.Approval = "CONFIRM"
	response.Signature = "<RSA signature of signing the request_message(ขอ Bank statement...) with accessor type>"
	accessor_id, _ := uuid.FromString("12a8f328-53da-4d51-a927-3cc6d3ed3feb")
	response.AccessorId = accessor_id
	// save response message to blockchain

	// return api response
	return c.JSON(http.StatusOK, response)
}
