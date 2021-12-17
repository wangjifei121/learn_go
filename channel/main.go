package main

import (
	"fmt"
)

/*
channel
	- 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
	- Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
	- 通道是引用类型，通道类型的空值是nil,声明的通道后需要使用make函数初始化之后才能使用。
	- 通道有发送（send）、接收(receive）和关闭（close）三种操作
	- 关闭后的通道有以下特点：
		对一个关闭的通道再发送值就会导致panic。
		对一个关闭的通道进行接收会一直获取值直到通道为空。
		对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		关闭一个已经关闭的通道会导致panic。
	- 无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送
	- 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道
	- 我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做

*/

func main() {

	// //channel的声明
	// fmt.Println("------channel的声明------")
	// var ch1 chan int // 声明一个传递int类型的通道
	// var ch2 chan string //声明一个传递string类型的通道
	// var ch3 chan bool //声明一个传递bool类型的通道
	// fmt.Println(ch1) //nil
	// fmt.Println(ch2) //nil
	// fmt.Println(ch3) //nil

	// fmt.Println("-------channel的初始化------")
	// ch1 := make(chan int)
	// ch2 := make(chan string)
	// ch3 := make(chan bool)
	// fmt.Printf("ch1:%T\n", ch1) //ch1:chan int
	// fmt.Printf("ch2:%T\n", ch2) //ch2:chan string
	// fmt.Printf("ch3:%T\n", ch3) //	ch3:chan bool

	// //channel的操作 - 需要指定缓冲区大小 否则会报错：fatal error: all goroutines are asleep - deadlock!
	// //无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送
	// fmt.Println("-----channel的操作----")
	// ch := make(chan int, 1) //初始化channel
	// ch <- 10             //发送一个值到channel
	// num := <-ch          //接收一个值
	// fmt.Println(num)

	//练习题：定义两个channel 先向第一个存100个数，在将100个数取出来放到第二个channel中，在打印第二个channel的值
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	//向ch1存值
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//从ch1取存到ch2
	go func() {
		for {
			n, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- n * n
		}
		close(ch2) //ch2不关闭的话最后会报错：fatal error: all goroutines are asleep - deadlock!
	}()
	//从ch2取出并打印
	for n := range ch2 {
		fmt.Println(n)
	}

}
