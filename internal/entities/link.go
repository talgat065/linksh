package entities

import (
	"errors"
	"linksh/internal/valueobjects"
)

var ErrLinkNotFound = errors.New("link not found")

type Link struct {
	ID       valueobjects.ID  `json:"id"`
	FullURL  valueobjects.URL `json:"full_url"`
	ShortURL valueobjects.URL `json:"short_url"`
}

type LinkRepository interface {
	FindByFullURL(url string) (Link, error)
	Save(link Link) error
}
