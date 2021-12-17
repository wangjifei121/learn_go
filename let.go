package main

import (
	"fmt"
)

/*
   知识点

       - go run go.go / go fmt
       - go语言的包或者字符串必须是双引号
       - 数据类型：
           bool、数字类型、字符串类型、派生类型（指针类型、数组类型、结构化类型、Channel类型、函数类型、切片类型、接口类型、Map类型）
       - 局部变量被定义后必须使用
       - & 获取内存地址

       - go module是go官方自带的go依赖管理库,在1.13版本正式推荐使用

       - go module可以将某个项目(文件夹)下的所有依赖整理成一个 go.mod 文件,里面写入了依赖的版本等

       - 使用go module之后我们可不用将代码放置在src下了

       - 使用 go module 管理依赖后会在项目根目录下生成两个文件 go.mod 和 go.sum。

	   - 关键字  25个
	   		    break        default      func         interface    select
				case         defer        go           map          struct
				chan         else         goto         package      switch
				const        fallthrough  if           range        type
				continue     for          import       return       var
	   - 保留字 37个
	   	    Constants:    true  false  iota  nil

			Types:    int  int8  int16  int32  int64
					uint  uint8  uint16  uint32  uint64  uintptr
					float32  float64  complex128  complex64
					bool  byte  rune  string  error

			Functions:   make  len  cap  new  append  copy  close  delete
						complex  real  imag
						panic  recover

*/

const (
	i = iota
	j = iota
	k
	l
)

func main() {

	// fmt.Println("Google" + "Runoob")

	// %d 表示整型数字，%s 表示字符串
	// var stockcode = 123
	// var enddate = "2020-12-31"
	// var url = "Code=%d&endDate=%s"
	// var target_url = fmt.Sprintf(url, stockcode, enddate)
	// fmt.Println(target_url)
	// fmt.Println(stockcode)

	//变量声明
	var num1 = 10
	var num2 = 10
	var a, b string = "中", "国"
	var c int
	d := 100

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(&num1) //获取变量内存地址
	fmt.Println(&num2)
	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("")

	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	fmt.Println(i, j, k, l)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// //字符串操作
	// //求长度
	// var my_name = "wangjifei"
	// var age = 30
	// fmt.Println("----len-----")
	// fmt.Println(len(my_name))
	// //拼接
	// fmt.Println("----+-----")
	// fmt.Println(my_name , " ", age)
	// //字符串切分
	// fmt.Println("----split-----")
	// str_list := strings.Split(my_name, "")
	// fmt.Println(str_list)
	// //contains
	// fmt.Println("----contains-----")
	// fmt.Println(strings.Contains(my_name, "wang"))
	// fmt.Println(strings.Contains(my_name, "kang"))
	// //前后缀判断
	// fmt.Println("----hasprefix hasSuffix-----")
	// fmt.Println(strings.HasPrefix(my_name, "wang"))
	// fmt.Println(strings.HasPrefix(my_name, "kang"))
	// fmt.Println(strings.HasSuffix(my_name, "fei"))
	// fmt.Println(strings.HasSuffix(my_name, "yue"))
	// //字符位置
	// fmt.Println("----index-----")
	// fmt.Println(strings.Index(my_name, "wang"))
	// fmt.Println(strings.Index(my_name, "k"))
	// fmt.Println(strings.LastIndex(my_name, "i"))
	// //join操作
	// fmt.Println("----join-----")
	// fmt.Println(strings.Join(str_list, "/"))

	var s = '中'
	fmt.Printf("%c \n%v \n", s, s)

	// fmt.Println("-----字符串遍历-----")
	// m := "hello王计飞"
	// for i := 0; i < len(m); i++ {
	// 	fmt.Printf("%v(%c)", m[i], m[i])

	// }
	// for index, r:= range m{
	// 	fmt.Println(index, r)
	// }
	// fmt.Println()

	// rune_m := []rune(m)
	// rune_m[0] = 'y'
	// fmt.Println(string(rune_m))

	t := "hello沙河小王子"
	var l = len(t)
	zw := 0
	for _, v :=range t {
		if v <= 'z'{
			zw += 1
		}
	}
	fmt.Printf("字符总长度：%d\n", l)
	fmt.Printf("汉子总长度：%d\n", zw)
	fmt.Printf("汉字占字符总数的百分比：%.2f\n", float64(zw) / float64(l))
	
	// q := 0
	// s1 := "hello沙河小王子"
	// for _, i := range s1 {
	// 	if i > 'z' {
	// 		q ++
	// 	}
	// }
	// fmt.Println(q)

}
