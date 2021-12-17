package main

import "fmt"

//定义结构体
type Animal struct{
	name string
}
//定义接收者
func (a Animal)Move(){
	fmt.Printf("%v 会动\n", a.name)
}

//定义dog
type Dog struct{
	Feet int
	*Animal //通过嵌套匿名结构体实现继承
}

func (d Dog)Wang(){
	fmt.Printf("%v 汪汪\n", d.name)
}

func main(){
	//结构体继承
	d := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d.Animal.Move()
	d.Wang()

	//练习题
	fmt.Println("-------struct 练习题-------")
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // {"吃饭", "不睡觉", "打豆豆"}

}

type Person struct {
	name   string
	age    int8
	dreams []string
}

//因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
// func (p *Person) SetDreams(dreams []string) {
// 	p.dreams = dreams
// }

//正确的写法
func (p *Person) SetDreams(dreams []string) {
	// p.dreams = make([]string, 0, len(dreams)) //先初始化
	// p.dreams = append(p.dreams, dreams...)
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}