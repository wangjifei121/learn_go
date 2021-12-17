package main

import "fmt"

func main() {
	// recover()必须搭配defer使用。
	// defer一定要在可能引发panic的语句之前定义。
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("func A")
}

// func funcB() {
// 	panic("panic in B")
// }
func funcB() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
