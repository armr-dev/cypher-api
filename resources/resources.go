package Resources

type Request struct {
	Text      string `json:"text"`
	Algorithm string `json:"algorithm"`
}

type Data struct {
	Text string `json:"text"`
}

type Response struct {
	Data `json:"data"`
}