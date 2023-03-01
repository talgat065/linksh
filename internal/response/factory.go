package response

type Factory struct {
}

func (f Factory) CreateNewLinkResponse(id string, fullURL string, shortURL string) LinkCreatedResponse {
	return LinkCreatedResponse{
		ID:       id,
		FullURL:  fullURL,
		ShortURL: shortURL,
	}
}
