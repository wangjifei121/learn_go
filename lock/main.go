package main

import (
	"fmt"
	"sync"
	"time"
)

//lock锁
var x int
var y int
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁
var rwlock sync.RWMutex //读写互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

//写数据方法
func write(){
	// rwlock.Lock()
	lock.Lock()
	y +=1
	lock.Unlock()
	// rwlock.Unlock()
	wg.Done()
}

//读数据方法
func read(){
	// rwlock.RLock()
	lock.Lock()
	time.Sleep(time.Millisecond)//sleep 1毫秒
	// rwlock.RUnlock()
	lock.Unlock()
	wg.Done()
}

func main() {

	// fmt.Println("-----互斥锁----")
	// start := time.Now()
	// wg.Add(2)
	// go add()
	// go add()

	// wg.Wait()
	// fmt.Printf("耗时：%v \n", time.Now().Sub(start).Microseconds())
	// fmt.Println(x)

	fmt.Println("-------读写互斥锁------")
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Printf("当前执行耗时：%f\n", time.Now().Sub(start).Seconds())
}
