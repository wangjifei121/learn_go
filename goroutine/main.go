package main

import (
	"fmt"
	"runtime"
	"sync"
)
/*
goruntine
	- Go语言在语言层面天生支持并发，这也是Go语言流行的一个很重要的原因
    - Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 		goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制
    - Go语言中的操作系统线程和goroutine的关系：
		一个操作系统线程对应用户态多个goroutine。
		go程序可以同时使用多个操作系统线程。
		goroutine和OS线程是多对多的关系，即m:n。

*/


//定义全局的变量
var wg sync.WaitGroup

func hello(){
	defer wg.Done() //标记当前goroutine执行完毕
	fmt.Println("hello go")
}

func worckerA(){
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("worcker A:%v\n",i)
	}
}

func worckerB(){
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("worcker B:%v\n",i)
	}
}

func main(){
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) //登记一个goroutine
	// 	go hello()
	// }
	// fmt.Println("go main process done!")
	// wg.Wait() //等待全部登记的goroutine执行完

	// fmt.Println("------runtime.GOMAXPROCS(1)------")
	// //为当前执行分配一个物理核
	// //现象 - 两个任务只能一个执行完另一个在执行 
	// runtime.GOMAXPROCS(1)
	// wg.Add(2)
	// go worckerA()
	// go worckerB()
	// wg.Wait()
	// fmt.Println("---main done---")

	fmt.Println("------runtime.GOMAXPROCS(4)------")
	//为当前执行分配多个物理核
	//现象 - 两个任务可以同时在执行 
	runtime.GOMAXPROCS(4)
	wg.Add(2)
	go worckerA()
	go worckerB()
	wg.Wait()
	fmt.Println("---main done---")
	

}