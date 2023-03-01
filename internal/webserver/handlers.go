package webserver

import (
	"fmt"
	"linksh/internal/response"
	"linksh/internal/shortener"
	"net/http"
)

func MakeShortLinkHandler(service shortener.Shortener, responseFactory response.Factory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error occured while trying to parse request data: " + err.Error())
			return
		}
		if !r.URL.Query().Has("url") {
			err := response.WriteError(w, http.StatusBadRequest, "param 'url' has not been provided")
			if err != nil {
				fmt.Println("error occured while trying to send error response: " + err.Error())
				return
			}
			return
		}

		link, err := service.Create(r.URL.Query().Get("url"))
		if err != nil {
			fmt.Println(fmt.Errorf("error occured while trying to create new short link: %s", err.Error()))
			err := response.WriteError(w, http.StatusInternalServerError, "internal error")
			if err != nil {
				fmt.Println(fmt.Errorf("error occured while trying to create new short link: %s", err.Error()))
			}
			return
		}

		err = response.WriteSuccess(w, responseFactory.LinkCreatedResponse(
			link.ID.Val,
			link.FullURL.Val,
			link.ShortURL.Val,
		))
		if err != nil {
			fmt.Println(fmt.Errorf("error occured while trying to create new short link: %s", err.Error()))
			return
		}
	}
}
