// 程序启动入口
// @author: Laba Zhang
package userapp_admin

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	userapp_v1 "week04/api/userapp/v1"
	"week04/internal/userapp/server"
	"week04/internal/userapp/service"
)

const (
	addr = ":2365"
)

func main() {
	// init service api
	us := InitUserBiz()
	userService := service.NewUserService(us)

	// register grpc userService
	s := server.NewServer(addr)
	userapp_v1.RegisterUserServer(s.Server, userService)

	// context
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// start grpc server
	g.Go(func() error {
		return s.Start(ctx)
	})

	// catch signals
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	// wait stop
	if err := g.Wait(); err != nil {
		log.Printf("error: %v", err)
	}
}
