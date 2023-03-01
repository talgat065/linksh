package valueobjects

import "github.com/google/uuid"

func NewID() ID {
	return ID{Val: uuid.NewString()}
}

type ID struct {
	Val string `json:"value"`
}

func NewURL(url string) URL {
	return URL{Val: url}
}

type URL struct {
	Val string `json:"value"`
}

func (u URL) GetVal() string {
	return u.Val
}
