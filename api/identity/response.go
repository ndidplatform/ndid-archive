package identity

type Response struct {
	ID     string `json:"id"`
	Result struct {
		Response struct {
			Value string `json:"value"`
			Log   string `json:"log"`
			Key   string `json:"key"`
		} `json:"response"`
	} `json:"result"`
}

type ResponseDeliver struct {
	Result struct {
		Height  int `json:"height"`
		CheckTx struct {
			Fee struct{} `json:"fee"`
		} `json:"check_tx"`
		DeliverTx struct {
			Log string   `json:"log"`
			Fee struct{} `json:"fee"`
		} `json:"deliver_tx"`
		Hash string `json:"hash"`
	} `json:"result"`
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
}
