package webserver

import (
	"linksh/internal/response"
	"linksh/internal/shortener"
	"linksh/internal/storage/memory"
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) initializeRoutes() {
	linkRepository := memory.LinkRepository{}
	shortenerService := shortener.NewShortener(linkRepository)
	responseFactory := response.Factory{}

	http.HandleFunc("/create", MakeShortLinkHandler(shortenerService, responseFactory))
}

func (s *Server) Run() error {
	s.initializeRoutes()
	return http.ListenAndServe("127.0.0.1"+s.Port, nil)
}
