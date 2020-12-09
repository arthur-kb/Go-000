package main

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

func server(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	return errors.Wrap(s.ListenAndServe(), "httpServer Quit!")
}
