package identity

import (
	"encoding/json"
	"net/http"
)

type Tendermint struct {
	url string
}

func New(url string) *Tendermint {
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
