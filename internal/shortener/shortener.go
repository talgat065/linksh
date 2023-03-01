package shortener

import (
	"linksh/internal/entities"
	"linksh/internal/generator"
	"linksh/internal/valueobjects"
)

func NewShortener(repository entities.LinkRepository) Shortener {
	return Shortener{linkRepository: repository}
}

type Shortener struct {
	linkRepository entities.LinkRepository
}

func (s Shortener) Create(url string) (entities.Link, error) {
	link, err := s.linkRepository.FindByFullURL(url)

	shortLink := generator.CreateShortLink(url)

	switch err {
	case nil:
	case entities.ErrLinkNotFound:
		link, err = entities.Link{
			ID:       valueobjects.NewID(),
			FullURL:  valueobjects.NewURL(url),
			ShortURL: valueobjects.NewURL(shortLink),
		}, nil
	}
	return link, err
}
