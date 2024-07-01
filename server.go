package todo

import (
	"context"
	"net/http"
	
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,	
		MaxHeaderBytes: 1 << 20,
		Handler: handler,
		ReadTimeout:    10,
		WriteTimeout:   10,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
