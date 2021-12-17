package main

import "fmt"

func main() {

	//每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等
	// a :=10
	// b := &a //取变量的地址
	// fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	// fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	// fmt.Println(&b)                    // 0xc00000e018
	// fmt.Println(*b) //指针取值          //10

	//
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100

	//new与make的区别
	//二者都是用来做内存分配的。
    //make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
    //而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针

}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
