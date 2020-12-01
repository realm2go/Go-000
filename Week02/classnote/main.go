package main

import "fmt"

func main() {
	Go(func() {
		fmt.Println("xxxxxxxx")
		panic("00000000")
	})

}

// 防止野生的goroutine panic
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}()
	x()
}
