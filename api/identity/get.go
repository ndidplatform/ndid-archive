package identity

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ndidplatform/ndid/api/client/tendermint"
)

func GetIdentifier(c echo.Context) error {
	namespace := c.Param("ns")
	id := c.Param("id")

	// prepare data
	funcName := "GetIdentifier"
	tx := []byte("\"" + funcName + "," + namespace + "," + id + "\"")
	path := "/abci_query?data=" + string(tx)

	fmt.Println(string(tx))

	var body Response
	t := tendermint.New(path)
	if err := t.Decode(&body); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}
