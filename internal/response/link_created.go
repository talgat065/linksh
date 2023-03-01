package response

type LinkCreatedResponse struct {
	ID       string `json:"id"`
	FullURL  string `json:"full_url"`
	ShortURL string `json:"short_url"`
}

func (response LinkCreatedResponse) Wrap() Wrapper {
	return Wrapper{
		ErrorCode: 0,
		Message:   "success",
		Data:      response,
	}
}
