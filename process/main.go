package main

import (
	"fmt"
)

func main() {
	//if条件
	// var score int = 70
	// if score < 60{
	// 	fmt.Println("成绩不合格")
	// } else if score < 90{
	// 	fmt.Println("成绩合格")
	// }else{
	// 	fmt.Println("成绩优秀")
	// }

	//if条件简写
	// if score := 80; score > 90 {
	// 	fmt.Println("A")
	// }else if score > 60{
	// 	fmt.Println("B")
	// }else{
	// 	fmt.Println("C")
	// }

	//for循环
	// for i:=0;i < 20; i++{
	// 	fmt.Println(i)
	// }
	//for初始语句简写
	// i:=20
	// for ;i > 1; i--{
	// 	fmt.Println(i)
	// }
	//for循环初始语句和结束语句都简写
	// i:=0
	// for i<10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//for range 键值循环
	// ss := "我是王计飞先生"
	// for index, s :=range ss{
	// 	fmt.Printf("key:%d value:%c\n", index, s)
	// }

	//switch
	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	// finger := 8
	// switch finger {
	// case 1, 6:
	// 	fmt.Println("大拇指")
	// case 2, 7:
	// 	fmt.Println("食指")
	// case 3, 8:
	// 	fmt.Println("中指")
	// 	fallthrough
	// case 4, 9:
	// 	fmt.Println("无名指")
	// case 5, 0:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("输入有误")
	// }

	//goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时
	// for i:=0; i< 20; i++{
	// 	if i > 10{
	// 		goto breakgo
	// 	}else{
	// 		fmt.Println(i)
	// 	}
	// }
	// breakgo:
	// 	fmt.Println("for循环结束")

	//break
	// for i := 0; i < 20; i++ {
	// 	if i > 10 {
	// 		break
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
	//break 指定标签
	// breaktag:
	// 	for i := 0; i < 20; i++ {
	// 		if i > 10 {
	// 			break breaktag
	// 		} else {
	// 			fmt.Println(i)
	// 		}
	// 	}

	//continue
	// for i := 0; i < 20; i++ {
	// 	if i == 10 {
	// 		continue
	// 	}
	// 	fmt.Println("continue",i)
	// }

	//编写9*9
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

}
