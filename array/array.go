package main

import "fmt"

func main() {
	//数组的定义
	//长度一旦定义，不可变
	// var a [3]int
	// var b = [3]int{1, 2}
	// var c = [4]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	// var d = [...]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	// var e = [...]string{1: "beijing", 2: "shenzhen", 3: "shanghai", 4: "guangzhou"} //指定索引定义
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("type of a:%T\n", a)
	// fmt.Printf("type of a:%T\n", b)
	// fmt.Printf("type of a:%T\n", c)
	// fmt.Printf("type of a:%T\n", d)
	// fmt.Printf("type of a:%T\n", e)

	// //数组的遍历 方式一
	// fmt.Println(" 数组的遍历 common")
	// var a = [...]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// //数组的遍历 方式二
	// fmt.Println(" 数组的遍历 range")
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }
	// fmt.Println(" 数组的遍历 range index")
	// for index := range a {
	// 	fmt.Println(index)
	// }

	//多位数组 - 嵌套数组
	// a := [2][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// fmt.Println(a)
	// fmt.Println(a[1][1])

	// var b = [2][2]int{
	// 	{11, 12},
	// 	{13, 14},
	// }
	// fmt.Println(b)

	//二维数组的遍历
	// a := [2][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < len(a[i]); j++ {
	// 		fmt.Println(a[i][j])
	// 	}
	// }

	//数组的值类型
	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	// var a = [3] int{1, 2, 3}
	// modified_array(a)
	// c :=a
	// fmt.Println(a)
	// fmt.Println(c)

	//求数组[1, 3, 5, 7, 8]所有元素的和
	// var a = [...]int{1, 3, 5, 7, 8}
	// var sum int = 0
	// for _, x := range a {
	// 	sum += x
	// }
	// fmt.Println(sum)
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	var a = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}

}

// func modified_array(x [3] int){
// 	x[0] = 100
// 	fmt.Println("func modified =", x)
// }
