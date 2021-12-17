package main

import (
	"fmt"
)

/*
worker pool

*/

func worker(id int, job <-chan int, result chan<- int) {
	for n := range job {
		fmt.Printf("worker:%d start job %d\n", id, n)
		result <- n * n
	}
}

func main() {
	job := make(chan int, 100)
	result := make(chan int, 100)

	//创建3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, job, result)
	}
	//创建10个任务
	for i := 0; i < 10; i++ {
		job <- i
	}
	close(job)
	// for r:= range result{
	// 	fmt.Printf("任务结果：%d\n",r)
	// }
	// 输出结果
	for a := 1; a <= 10; a++ {
		r := <-result
		fmt.Printf("任务结果：%d\n", r)
	}
}
