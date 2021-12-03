package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (ser *Server) Run(port string, handler http.Handler) error {
	ser.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return ser.httpServer.ListenAndServe()
}

func (ser *Server) Shutdown(ctx context.Context) error {
	return ser.httpServer.Shutdown(ctx)
}
