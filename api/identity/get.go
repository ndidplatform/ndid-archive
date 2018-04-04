package identity

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ndidplatform/ndid/api/client/tendermint"
)

func GetIdentifier(c echo.Context) error {
	path := buildQueryPath(c.Param("ns"), c.Param("id"))

	var body Response
	t := tendermint.New(path)
	if err := t.Decode(&body); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, body)
}

func buildQueryPath(namespace, id string) (path string) {
	funcName := "GetIdentifier"
	tx := []byte("\"" + funcName + "," + namespace + "," + id + "\"")
	path = "/abci_query?data=" + string(tx)
	return
}
