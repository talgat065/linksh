package main

import (
	"fmt"
	"linksh/internal/shortener"
	"linksh/internal/storage/memory"
	"os"
)

func main() {
	args := os.Args[1:]
	service := shortener.NewShortener(memory.LinkRepository{})

	for _, v := range args {
		lnk, err := service.Create(v)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Full: %s, Short: %s", lnk.FullURL, lnk.ShortURL))
		}
	}
}
