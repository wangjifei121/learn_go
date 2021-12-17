package main

import (
	// "errors"
	// "encoding/xml"
	"fmt"
	"strings"
)

//定义函数类型
type calculation func(int, int) int

func main() {
	s := sum(10, 15)
	fmt.Println(s)

	//调用
	ms := mult_sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(ms)

	//返回值命名
	// rs, bs := calc(10, 2)
	// fmt.Println(rs, bs)

	//给定义的类型赋值
	var c calculation
	c = sum
	ss := c(5, 10)
	fmt.Println(ss)

	//高阶函数
	sss := calc(2, 3, sum)
	fmt.Println(sss)

	//匿名函数 - 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(11, 22)

	//匿名函数 - 直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(11, 22)

	//闭包引用 - 初级
	fmt.Println("-----闭包初级-----")
	f := adder()
	fmt.Println(f(10))
	fmt.Println(f(7))
	fmt.Println(f(4))

	//闭包引用 - 进阶
	fmt.Println("-----闭包进阶-----")
	ff2 := adder2(4)
	fmt.Println(ff2(10))
	fmt.Println(ff2(7))
	fmt.Println(ff2(4))

	//闭包引用 - 进阶2
	fmt.Println("------闭包进阶2-----")
	csv_suff := makeSuffixFunc(".csv")
	fmt.Println(csv_suff("test"))
	fmt.Println(csv_suff("test.csv"))

	//闭包饮用 - 进阶3
	fmt.Println("------闭包进阶3----")
	fadd, fsub := fcalc(80)
	fmt.Println(fadd(1), fsub(1))
	fmt.Println(fadd(10), fsub(10))

	//defer语句
	fmt.Println("------defer----")
	my_defer()
}

//定义一个求和的函数
func sum(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//可变参数
func mult_sum(x ...int) int {
	var sum int = 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//返回值命名
// func calc(x, y int) (sum, sub int) {
// 	sum = x + y
// 	sub = x - y
// 	return
// }
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

//高阶函数
func calc(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//函数作为返回值
// func do_something(s string) (func(int, int) int, error) {
// 	switch s {
// 	case "+":
// 		return sum, nil
// 	case "_":
// 		return sub, nil
// 	default:
// 		error := errors.New("param s error")
// 		return nil, error
// 	}
// }

//闭包 - 初级
func adder() func(int) int {
	x := 0
	return func(y int) int {
		x = x + y
		return x
	}
}

//闭包进阶
func adder2(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//闭包进阶2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return fmt.Sprint(name, suffix)
		} else {
			return name
		}
	}
}

//闭包进阶3
func fcalc(base int) (func(int) int, func(int) int) {
	add := func(x int) int {
		base += x
		return base
	}
	sub := func(x int) int {
		base -= x
		return base
	}
	return add, sub

}

//defer语句
//由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
func my_defer() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
