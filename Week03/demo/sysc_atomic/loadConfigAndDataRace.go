package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

func main() {
	cfg := &Config{}

	// 使用原子操作
	var v atomic.Value
	v.Store(cfg) // 存储的结构类型

	// 模拟1人写
	go func(){
		i := 0
		for {
			i++
			cfg.a = []int{i,i+1,i+2,i+3,i+4,i+5}
			v.Store(cfg)  // 用原子保存
		}
	}()

	var wg sync.WaitGroup
	// 模拟4人读
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func(){
			//for n := 0;n < 100; n++{
			cfg := v.Load().(*Config)
			fmt.Printf("%v\n",cfg)
			//}
			wg.Done()
		}()
	}

	wg.Wait()

}
// go test -bentch=.
