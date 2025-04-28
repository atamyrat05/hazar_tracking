package server

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func (s *Server) Run(add string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           ":" + add,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.Server.ListenAndServe()
}
