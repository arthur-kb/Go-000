// server
// @author: Laba Zhang
package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	*grpc.Server
	addr string
}

func NewServer(addr string) *Server {
	srv := grpc.NewServer()
	return &Server{Server: srv, addr: addr}
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	log.Printf("grpc server start: %s", s.addr)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
		log.Printf("grpc server gracefull stop")
	}()

	return s.Serve(l)
}
