package main

import (
	"fmt"
	"time"
)


func test(n int)int{
	sum:=0
	for i := 0; i < n; i++ {
		sum+=i
	}
	return sum
}


func main(){
	start :=time.Now()
	ret:= test(100000000)
	fmt.Println(ret)
	fmt.Printf("当前程序执行耗时：%v\n", time.Now().Sub(start).Seconds())

}