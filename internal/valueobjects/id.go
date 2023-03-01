package valueobjects

import "github.com/google/uuid"

func NewID() ID {
	return ID{val: uuid.NewString()}
}

type ID struct {
	val string
}

func NewURL(url string) URL {
	return URL{val: url}
}

type URL struct {
	val string
}
