package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool,1)

	var mu sync.Mutex

	var g1c,g2c int

	// g1
	go func() {
		for {
			select {
				case <-done:
					return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
				g1c++
			}
		}
	}()

	// g2  // 锁饥饿
	go func() {
		for i :=0; i <10; i++ {
			time.Sleep(100 * time.Microsecond)
			mu.Lock()
			mu.Unlock()
			g2c++
		}

	}()

	time.Sleep(1 * time.Second)

	done <- true

	fmt.Printf("g1锁的次数:%d,g2获得锁的次数:%d\n",g1c,g2c)
}
