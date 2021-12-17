package main

import "fmt"
//接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节
//接口就是类的抽象 - 
//接口是一种类型，一种抽象的类型
//接口类型不关心属性（数据），只关心行为（方法）
//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表
// 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui

type cat struct{}

func (c cat) say(){
	fmt.Println("---喵喵喵---")
}
type dog struct{}

func (d dog)say(){
	fmt.Println("----汪汪汪---")
}

type sayer interface{
	say()
}

func sayFunc(arg sayer){
	arg.say()
}

func main(){
	d:= dog{}
	d.say()
	c:= cat{}
	c.say()

	fmt.Println("----接口的dome----")
	sayFunc(d)
	sayFunc(c)

}