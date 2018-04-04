package identity

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	cmn "github.com/tendermint/tmlibs/common"
)

type SID struct {
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
}

func CreateIdentity(c echo.Context) error {
	sid := new(SID)
	if err := c.Bind(sid); err != nil {
		return err
	}

	// prepare tx
	nonce := cmn.RandStr(12)
	fn := "CreateIdentity"
	tx := []byte("\"" + nonce + "," + fn + "," + sid.Namespace + "," + sid.Id + "\"")
	url := "http://" + tendermintAddr + "/broadcast_tx_commit?tx=" + string(tx)

	fmt.Println(string(tx))

	var body ResponseDeliver
	err := New(url, &body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, body)
}
