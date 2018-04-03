package identity

import (
	"encoding/json"
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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var body ResponseDeliver
	json.NewDecoder(resp.Body).Decode(&body)
	return c.JSON(http.StatusCreated, body)
}
