package webserver

import (
	"encoding/json"
	"fmt"
	"linksh/internal/response"
	"linksh/internal/shortener"
	"net/http"
)

func MakeShortLinkHandler(service shortener.Shortener, responseFactory response.Factory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {

			return
		}
		if !r.URL.Query().Has("url") {
			fmt.Println("param 'url' not provided")
			return
		}

		link, err := service.Create(r.URL.Query().Get("url"))
		if err != nil {
			fmt.Println(fmt.Errorf("error occured while trying to create new short link: %s", err.Error()))
			return
		}

		resp := responseFactory.CreateNewLinkResponse(
			link.ID.Val,
			link.FullURL.Val,
			link.ShortURL.Val,
		).Wrap()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			fmt.Println(fmt.Errorf("error occured while trying to encode response: %s", err.Error()))
		}
	}
}
