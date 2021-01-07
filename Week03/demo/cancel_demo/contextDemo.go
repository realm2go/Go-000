package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	messages := make(chan int,10)

	// 生产者
	for i :=0; i<10 ;i++  {
		messages <- i
	}

	ctx,cancelFun := context.WithTimeout(context.Background(),6 * time.Second)

	// 消费者
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
				case <-ctx.Done():
					//time.Sleep(3 * time.Second)
					fmt.Println("子goroutine 退出")
					return
				default:
					fmt.Printf("send message :%d\n",<-messages)
				}
		}
	}(ctx)

	defer close(messages)

	defer cancelFun()
	//time.Sleep(5 * time.Second)
	//cancel()

	select {
		case <-ctx.Done():
			time.Sleep(1 * time.Second)
			fmt.Println("Main process exit!")
	}
}
