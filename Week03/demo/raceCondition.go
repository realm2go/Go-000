package main

import (
	"fmt"
	"sync"
	"time"
)

var WG sync.WaitGroup
var Counter = 0

func main() {
	for routineID := 1; routineID <= 2; routineID++ {

		WG.Add(1)
		go Routine(routineID)
	}

	WG.Wait()

	fmt.Printf("Final Counter: %d\n",Counter)
}

func Routine(id int){
	for count :=0; count < 2; count++{
		value := Counter
		// time.sleep 触发go的上下文切换
		time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
	}
	WG.Done()
}

//  go build -race raceCondition.go
