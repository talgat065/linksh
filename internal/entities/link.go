package entities

import (
	"errors"
	"linksh/internal/valueobjects"
)

var ErrLinkNotFound = errors.New("link not found")

type Link struct {
	ID       valueobjects.ID
	FullURL  valueobjects.URL
	ShortURL valueobjects.URL
}

type LinkRepository interface {
	FindByFullURL(url string) (Link, error)
	Save(link Link) error
}
