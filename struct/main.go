package main

import (
	"fmt"
)
//struct是值类型
//使用type和struct关键字来定义结构体
type Person struct {
	name string
	age  int
}

type Man struct{
	*Person //通过嵌套匿名结构体实现继承
}

func main() {
	var p1 = Person{
		name: "wangjifei",
		age:  18,
	}
	fmt.Println("-----结构体的基本定义-------")
	fmt.Printf("T:%T  v:%v  #v%#v\n", p1, p1, p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	//匿名结构体
	var user  struct{
		name string
		age int
	}
	user.name = "wangjifei"
	user.age =  18
	fmt.Println("-----匿名结构体的定义-------")
	fmt.Printf("T:%T  v:%v  #v%#v\n", user, user, user)
	fmt.Println(user.name)
	fmt.Println(user.age)

	//创建指针类型结构体
	var p2 = new(Person)
	fmt.Println("-----创建指针类型的结构体-----")
	fmt.Printf("%T\n",p2)
	fmt.Printf("%#v\n",p2)
	p2.name = "wangjifei"
	p2.age = 18
	fmt.Println(p2.name)
	fmt.Println(p2.age)

	//取结构体的地址实例化
	fmt.Println("------取结构体的地址实例化-----")
	p3 := &Person{}
	p3.name = "wangjifei"
	p3.age = 18
	fmt.Printf("%T\n",p2)
	fmt.Printf("%#v\n",p2)
	fmt.Println(p3.name)
	fmt.Println(p3.age)

	//没有初始化的结构体，其成员变量都是对应其类型的零值
	fmt.Println("-----没有初始化的结构体，其成员变量都是对应其类型的零值----")
	var p4 Person
	fmt.Printf("%T\n",p4)
	fmt.Printf("%#v\n",p4)

	//练习题
	fmt.Println("------练习题------")
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu //{"小王子":&{name: "小王子", age: 18}}
	}
	fmt.Printf("%#v\n",m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)//小王子 大王八
	}

	//构造函数
	fmt.Println("------构造函数------")
	var p5 = NewPerson("wangjifei",18)
	fmt.Printf("%#v\n",p5)

	//Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self
	fmt.Println("-------方法和接收者-------")
	p5.walk()
	fmt.Println(p5.age)
	p5.setAge(10)
	fmt.Println(p5.age)

	//结构体的继承
	fmt.Println("------结构体的继承-------")
	new_m := Man{
		Person: &Person{
			name: "wangjifei",
			age: 18,
		},
	}
	fmt.Println(new_m.name)
	fmt.Println(new_m.age)
	new_m.walk()

}

type student struct {
	name string
	age  int
}

//构造函数
func NewPerson(name string,age int) *Person{
	return &Person{
		name: name,
		age: age,
	}
}
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型
//方法和接收者 值类型的接收者 copy 修改不了属性
func (p Person) walk(){
	fmt.Println("人可以行走")
}

//指针类型的接收者
func (p *Person) setAge(age int){
	p.age = age
}

//什么时候应该使用指针类型接收者
//1、需要修改接收者中的值
//2、接收者是拷贝代价比较大的大对象
//3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者

// 结构体字段的可见性
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。