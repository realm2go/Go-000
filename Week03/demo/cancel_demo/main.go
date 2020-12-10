package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():  // 收到关闭通道的后，该goroutine退出
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())

	//cancel closes c.done, cancels each of c's children
	defer cancel()  // 会去close c.Done,也会递归调用，关闭子context

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}


const shortDuration = 4 * time.Second

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(),d)

	defer cancel()

	select {
		case <- time.After(5 * time.Second):
			fmt.Println("over slept")
		case <- ctx.Done():
			fmt.Println(ctx.Err())
		}
}