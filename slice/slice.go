package main

import (
	"fmt"
	"sort"
)

//数组的长度固定导致了数组的局限性
//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容
func main() {
	// var a []int
	// var b []string
	// fmt.Println(a)
	// fmt.Printf("a type = %T\n", a)
	// fmt.Println(b)
	// fmt.Printf("b type = %T\n", b)

	// fmt.Println(a == nil) //切片定义但未赋值时为nil
	// fmt.Println(b == nil)

	// fmt.Println(len(a), cap(a))
	// fmt.Println(len(b), cap(b))

	// var a = [5]int{1, 2, 3, 4, 5}
	// s := a[2:4]
	// fmt.Println(s)
	// fmt.Printf("s type = %T len: %d cap: %d \n", s, len(s), cap(s))

	// a := [5]int{1, 2, 3, 4, 5}
	// t := a[3:4:4]
	// fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	//make 函数构造切片
	// var a = make([]int,3,10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	// var a []int
	// b := []int{}
	// c := make([]int, 0, 1)
	// fmt.Println(a == nil)
	// fmt.Println(b == nil)
	// fmt.Println(c == nil)
	// fmt.Println(len(a))
	// fmt.Println(len(b))
	// fmt.Println(len(c))

	//slice 赋值拷贝
	// a := []int{1, 2, 3, 4, 5}
	// b := a
	// a[1] = 10
	// fmt.Println(a)
	// fmt.Println(b)

	//slice 遍历
	// a := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// fmt.Println("---range----")
	// for index, val := range a {
	// 	fmt.Printf("index:%v val:%v\n", index, val)
	// }

	//slice append
	// var a []string
	// b :=[] string{"A","B","C"}
	// a = append(a,"beijing","sshanghai","guangzhou")
	// fmt.Println(a)
	// fmt.Printf("%T \n", a)
	// a = append(a, b...) //append
	// fmt.Println(a)

	//append slice 扩容
	//append()函数将元素追加到切片的最后并返回该切片。
	//切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	// var sumSlice [] int
	// for i := 0; i < 20; i++ {
	// 	sumSlice = append(sumSlice, i)
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p \n",sumSlice, len(sumSlice), cap(sumSlice), sumSlice)
	// }

	//slice copy
	// a := []int{1, 2, 3, 4, 5}
	// b := make([]int, 5, 5)
	// copy(b, a)
	// fmt.Println(a)
	// fmt.Println(b)
	// b[1]= 100
	// fmt.Println(a)
	// fmt.Println(b)

	//从切片中删除元素
	//Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// //删除索引为2的元素
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a)

	//练习题
	//请写出下面代码的输出结果
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)      //["","","","","","",0,1,2,3,4,5,6,7,8,9]
	// fmt.Println(len(a)) //15

	//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	a := [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
	fmt.Printf("%T \n",a)

}
