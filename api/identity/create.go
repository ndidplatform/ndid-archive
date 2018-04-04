package identity

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ndidplatform/ndid/api/client/tendermint"
	cmn "github.com/tendermint/tmlibs/common"
)

type Sid struct {
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
}

func CreateIdentity(c echo.Context) error {
	sid := new(Sid)
	if err := c.Bind(sid); err != nil {
		return err
	}

	path := buildBroadcastPath(sid)

	var body ResponseDeliver
	t := tendermint.New(path)
	if err := t.Decode(&body); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, body)
}

func buildBroadcastPath(sid *SID) (path string) {
	nonce := cmn.RandStr(12)
	fn := "CreateIdentity"
	tx := []byte("\"" + nonce + "," + fn + "," + sid.Namespace + "," + sid.Id + "\"")
	path = "/broadcast_tx_commit?tx=" + string(tx)
	return
}
