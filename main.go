package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"localhost/backend/config"
	"localhost/backend/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// go routine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// 建立一个 os.Signal 的通道
	quit := make(chan os.Signal, 1)
	// 接收到 os.Interrupt 信号后，关闭 server，并退出 （signal.Notify 函数用于设置通道 quit 来监听 os.Interrupt 信号，这个信号通常在用户按下 Ctrl+C 时发送。当 os.Interrupt 信号发生时，quit 通道将接收到这个信号。）
	signal.Notify(quit, os.Interrupt)
	// 这是一个阻塞操作，程序会在这里等待直到 quit 通道接收到信号
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// defer 关键字用于延迟函数的执行直到包含它的函数返回
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v", err)
	}

	log.Println("Server exiting")

}

// channel demo

// func main() {
// 	channel := make(chan int, 0)

// 	// channel 在 go routine 中写
// 	go func() {
// 		channel <- 114
// 		channel <- 114
// 	}()

// 	x := <-channel

// 	fmt.Println(x)
// }
