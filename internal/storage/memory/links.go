package memory

import "linksh/internal/entities"

type LinkRepository struct {
}

func (repository LinkRepository) FindByFullURL(_ string) (entities.Link, error) {
	return entities.Link{}, entities.ErrLinkNotFound
}

func (repository LinkRepository) Save(_ entities.Link) error {
	return nil
}
