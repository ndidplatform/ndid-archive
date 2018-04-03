package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	cmn "github.com/tendermint/tmlibs/common"
)

//var tendermintAddr = "localhost:45000"
var tendermintAddr = "127.0.0.1:46657"

func main() {
	// TODO
	//tendermintAddr = args[1]
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/identity/:ns/id/:id", GetIdentifier)
	e.POST("/identity", CreateIdentity)

	e.Logger.Fatal(e.Start(":8000"))
}

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

	var body interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	errCode := (body.(map[string]interface{})["result"]).(map[string]interface{})["deliver_tx"]
	errCode = errCode.(map[string]interface{})["code"]

	result := map[string]string{}
	if errCode == nil {
		result["result"] = "success"
	} else {
		result["result"] = "fail"
	}

	return c.JSON(http.StatusCreated, result)
}

func GetIdentifier(c echo.Context) error {
	namespace := c.Param("ns")
	id := c.Param("id")

	// prepare data
	funcName := "Identifier"
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
