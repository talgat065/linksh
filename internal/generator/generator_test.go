package generator

import "github.com/stretchr/testify/assert"
import "testing"

func TestGenerateShortLink(t *testing.T) {
	link1 := "https://en.wikipedia.org/wiki/John_Lennon"
	shortLink1 := CreateShortLink(link1)
	assert.Equal(t, "3n3AvHib", shortLink1)

	link2 := "https://habr.com/ru/company/intel/blog/156863/"
	shortLink2 := CreateShortLink(link2)
	assert.Equal(t, "759Fn9df", shortLink2)

	link3 := "https://datatracker.ietf.org/doc/html/rfc9112"
	shortLink3 := CreateShortLink(link3)
	assert.Equal(t, "XeaTkEW8", shortLink3)
}
