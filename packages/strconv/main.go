package main

import (
	"fmt"
	"strconv"
)

//Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。

//Atoi()函数用于将字符串类型的整数转换为int类型
func testAtoi() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("strconv atoi error:", err)
	}
	fmt.Printf("type=%T, num=%d\n", i1, i1)
}

//Itoa()函数用于将int类型数据转换为对应的字符串表示
func testItoa() {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

//Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
func testParse() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("type=%T, value=%v\n", b, b)
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("type=%T, value=%v\n", f, f)
	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("type=%T, value=%v\n", i, i)
	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("type=%T, value=%v\n", u, u)
}

//Format系列函数实现了将给定类型数据格式化为string类型数据的功能
func testFormat() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("type=%T,value=%v\n", s1, s1)
	fmt.Printf("type=%T,value=%v\n", s2, s2)
	fmt.Printf("type=%T,value=%v\n", s3, s3)
	fmt.Printf("type=%T,value=%v\n", s4, s4)
}

func main() {
	// testAtoi()
	// testItoa()
	// testParse()
	testFormat()
}
