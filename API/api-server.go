package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	cmn "github.com/tendermint/tmlibs/common"
)

func main() {
	runAPIServer(os.Args)
}

var tendermintAddr = "localhost:45000"

func runAPIServer(args []string) error {
	tendermintAddr = args[1]
	router := mux.NewRouter()

	// IDP Onboarding
	router.HandleFunc("/create_identity_with_pub_key", create_identity_with_pub_key).Methods("POST")
	router.HandleFunc("/check_identifier", check_identifier).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
	return nil
}

type SID struct {
	Namespace  string
	Identifier string
}

func create_identity_with_pub_key(w http.ResponseWriter, r *http.Request) {

	result := map[string]interface{}{}
	var sid SID
	if err := json.NewDecoder(r.Body).Decode(&sid); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// prepare tx
	nonce := cmn.RandStr(12)
	funcName := "create_identity_with_pub_key"
	tx := []byte("\"" + nonce + "," + funcName + "," + sid.Namespace + "," + sid.Identifier + "\"")
	url := "http://" + tendermintAddr + "/broadcast_tx_commit?tx=" + string(tx)

	fmt.Println(string(tx))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		var body interface{}
		json.NewDecoder(resp.Body).Decode(&body)
		errCode := (body.(map[string]interface{})["result"]).(map[string]interface{})["deliver_tx"]
		errCode = errCode.(map[string]interface{})["code"]

		if errCode == nil {
			result["result"] = "success"
		} else {
			result["result"] = "fail"
		}
	}

	json.NewEncoder(w).Encode(result)
}

func check_identifier(w http.ResponseWriter, r *http.Request) {

	result := map[string]interface{}{}
	var sid SID
	if err := json.NewDecoder(r.Body).Decode(&sid); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// prepare data
	funcName := "check_identifier"
	tx := []byte("\"" + funcName + "," + sid.Namespace + "," + sid.Identifier + "\"")
	url := "http://" + tendermintAddr + "/abci_query?data=" + string(tx)

	fmt.Println(string(tx))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		var body interface{}
		json.NewDecoder(resp.Body).Decode(&body)
		log := (body.(map[string]interface{})["result"]).(map[string]interface{})["response"]
		log = log.(map[string]interface{})["log"]

		if log == "exists" {
			result["result"] = "yes"
		} else {
			result["result"] = "no"
		}
	}

	json.NewEncoder(w).Encode(result)
}
