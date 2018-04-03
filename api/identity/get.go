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

	var body interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	log := (body.(map[string]interface{})["result"]).(map[string]interface{})["response"]
	log = log.(map[string]interface{})["log"]

	result := map[string]string{}
	if log == "exists" {
		result["result"] = "yes"
	} else {
		result["result"] = "no"
	}

	return c.JSON(http.StatusOK, result)
}
