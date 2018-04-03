package identity

import (
	"encoding/json"
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

	var body Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}
