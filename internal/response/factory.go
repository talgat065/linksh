package response

type Factory struct {
}

func (f Factory) LinkCreatedResponse(id string, fullURL string, shortURL string) LinkCreatedResponse {
	return LinkCreatedResponse{
		ID:       id,
		FullURL:  fullURL,
		ShortURL: shortURL,
	}
}
