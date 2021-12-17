package main

import (
	"fmt"
	"sync"
	"time"
)

//lock锁
var x int
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	start := time.Now()
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Printf("耗时：%v \n", time.Now().Sub(start).Microseconds())
	fmt.Println(x)
}
