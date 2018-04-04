package identity

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var tendermintAddr = "127.0.0.1:46657"

func GetIdentifier(c echo.Context) error {
	namespace := c.Param("ns")
	id := c.Param("id")

	// prepare data
	funcName := "GetIdentifier"
	tx := []byte("\"" + funcName + "," + namespace + "," + id + "\"")
	url := "http://" + tendermintAddr + "/abci_query?data=" + string(tx)

	fmt.Println(string(tx))

	var body Response
	t := New(url)
	err := t.Decode(&body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}
