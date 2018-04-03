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
