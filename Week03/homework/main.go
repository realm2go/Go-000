package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	// 启动http1
	http1 := NewHttpServer(":8111")
	g.Go(func() error {
		if err := http1.start(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	// 启动http2
	http2 := NewHttpServer(":8112")
	g.Go(func() error {
		if err := http2.start(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	// 监听sig信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			select {
			case s := <-c:
				switch s {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					cancel()
				default:
				}
			}
		}
	}()

	// context取消后，关闭http server
	go func() {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			go func() {
				if err := http1.shutdown(ctx); err != nil {
					log.Println("http1 shutdown err: ", err)
				}
			}()
			go func() {
				if err := http2.shutdown(ctx); err != nil {
					log.Println("http2 shutdown err: ", err)
				}
			}()
		}

	}()


	// main函数阻塞在此
	if err := g.Wait(); err != nil {
		log.Println("all exit: ", err)
	}
}

type httpServer struct {
	server http.Server
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		server: http.Server{
			Addr: addr,
		},
	}
}

func (h *httpServer) start() error {
	return h.server.ListenAndServe()
}

func (h *httpServer) shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
