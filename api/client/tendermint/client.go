package tendermint

import (
	"encoding/json"
	"net/http"
)

var tendermintAddr = "127.0.0.1:46657"

type Tendermint struct {
	url string
}

func New(path string) *Tendermint {
	url := "http://" + tendermintAddr + path
	return &Tendermint{
		url: url,
	}
}

func (t *Tendermint) Decode(body interface{}) error {
	req, err := http.NewRequest("GET", t.url, nil)
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

	return json.NewDecoder(resp.Body).Decode(&body)
}
