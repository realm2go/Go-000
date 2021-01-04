package main

import (
	"container/ring"
	"fmt"
	"time"
)

//var head *ring.Ring     // 环形队列（链表）

func main() {

	head := ring.New(10)

	for i := 0; i < 10; i++ {
		head.Value = i
		head = head.Next()
	}

	for i := 0; i < head.Len(); i++ {
		fmt.Println(head.Value)
		head = head.Next()
	}

	fmt.Println("-------------")
	for i := 0; i < head.Len()-5; i++ {
		//fmt.Println(head.Value)
		head.Value = 0
		head = head.Next()
	}

	fmt.Println("---------------")
	for i := 0; i < head.Len(); i++ {
		fmt.Println(head.Value)
		head = head.Next()
	}

	go func() {
		timer := time.NewTicker(time.Second * 1)
		for range timer.C {
			fmt.Println(head.Value)
			head = head.Next()
		}
	}()

	time.Sleep(time.Second * 30)

}
