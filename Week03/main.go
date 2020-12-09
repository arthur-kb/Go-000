package main

import (
	"context"
	"fmt"
	"errors"
	"golang.org/x/sync/errgroup"
	xerrors "github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
)

var interruptError = errors.New("quit signal")

//http服务
func httpServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello,QCon!")
	})

	return server(ctx, "0.0.0.0:8080", mux)
}

//监听系统信号服务
func listenSystemSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	//监听指定信号，ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill)

	select {
	case <-c:
		return xerrors.Wrap(interruptError, "")
	case <-ctx.Done():
		return nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, groupCtx:= errgroup.WithContext(ctx)

	group.Go(func() error {
		 if err := httpServer(groupCtx); err != nil && errors.Is(err, http.ErrServerClosed) {
		 	fmt.Println("http server完美退出")
		  	return nil
		 }else {
			 cancel()
			 return err
		 }
	})

	group.Go(func() error {
		if err := listenSystemSignal(groupCtx); err != nil {
			if errors.Is(err, interruptError) {
				fmt.Println("send cancel signal to httpserver")
				cancel()
				return err
			} 
		}
		fmt.Println("signal 完美退出")
		return nil
	})


	if err := group.Wait(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
