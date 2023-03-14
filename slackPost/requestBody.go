package slackPost

type RequestBody struct {
	Text string `json:"text"`
}

func NewRequestBodyFunc(text string) *RequestBody {
	return &RequestBody{
		Text: text,
	}
}
