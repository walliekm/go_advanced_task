package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func startHTTPSrv(srv *http.Server) error {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	return srv.ListenAndServe()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	httpSrv := &http.Server{Addr: ":8080"}

	sigChan := make(chan os.Signal, 1) //信号发送不会阻塞，如果chan缓冲区已满，则会被直接丢弃，因此此处buffer必须>0
	signal.Notify(sigChan)             //signal注册

	group.Go(func() error {
		return startHTTPSrv(httpSrv) //启动服务
	})

	group.Go(func() error {
		<-errCtx.Done()                 //阻塞，等待退出
		return httpSrv.Shutdown(errCtx) //关闭服务
	})

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): //发生错误，或者cancel被触发
				return errCtx.Err()
			case <-sigChan: //捕获signal信号，调用cancel退出
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		log.Println("group err:", err)
	}

	log.Println("all groutinue is done, exit!")
}
