package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	// "strings"
)

func testPrint() {
	//Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，
	//Printf函数支持格式化输出字符串，
	//Println函数会在输出内容的结尾添加一个换行符
	fmt.Print("小王子")
	fmt.Printf("%v\n", "小王子")
	fmt.Println("小王子")
}

func testFprint() {
	//Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

func testSprint() {
	//Sprint系列函数会把传入的数据生成并返回一个字符串
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func testEorrorf() {
	//Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)

}

func testFormat() {
	/*
		占位符	说明
		%v	值的默认格式表示
		%+v	类似%v，但输出结构体时会添加字段名
		%#v	值的Go语法表示
		%T	打印值的类型
		%%	百分号
		%t	true或false
	*/
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	fmt.Printf("%t\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

}

func testInt() {
	/*
		占位符	说明
		%b	表示为二进制
		%c	该值对应的unicode码值
		%d	表示为十进制
		%o	表示为八进制
		%x	表示为十六进制，使用a-f
		%X	表示为十六进制，使用A-F
		%U	表示为Unicode格式：U+1234，等价于”U+%04X”
		%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	*/

	n := 65
	fmt.Printf("%%b=%b\n", n)
	fmt.Printf("%%c=%c\n", n)
	fmt.Printf("%%d=%d\n", n)
	fmt.Printf("%%o=%o\n", n)
	fmt.Printf("%%x=%x\n", n)
	fmt.Printf("%%X=%X\n", n)
	fmt.Printf("%%U=%U\n", n)
	fmt.Printf("%%q=%q\n", n)
}

func testFloat() {
	/*
		占位符	说明
		%b	无小数部分、二进制指数的科学计数法，如-123456p-78
		%e	科学计数法，如-1234.456e+78
		%E	科学计数法，如-1234.456E+78
		%f	有小数部分但无指数部分，如123.456
		%F	等价于%f
		%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	*/
	f := 12.34
	fmt.Printf("%%b=%b\n", f)
	fmt.Printf("%%e=%e\n", f)
	fmt.Printf("%%E=%E\n", f)
	fmt.Printf("%%f=%f\n", f)
	fmt.Printf("%%F=%F\n", f)
	fmt.Printf("%%g=%g\n", f)
	fmt.Printf("%%G=%G\n", f)
}

func testByte() {
	/*
		占位符	说明
		%s	直接输出字符串或者[]byte
		%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x	每个字节用两字符十六进制数表示（使用a-f)
		%X	每个字节用两字符十六进制数表示（使用A-F）
	*/
	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}

func testPoint() {
	/*
		占位符	说明
		%p	表示为十六进制，并加上前导的0x
	*/
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}

func testPrecision() {
	/*
		占位符	说明
		%f	默认宽度，默认精度
		%9f	宽度9，默认精度
		%.2f	默认宽度，精度2
		%9.2f	宽度9，精度2
		%9.f	宽度9，精度0
	*/
	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)
}

func testOther(){
	/*
		占位符	说明
		’+’	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
		’ ‘	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
		’-’	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
		’#’	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
		‘0’	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	*/
	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)

}

func testScan(){
	/*
		Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
		本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因
	*/
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func testScanf(){
	/*
		fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，
		fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。
	*/
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func testScanln(){

	/*
		fmt.Scanln遇到回车就结束扫描了，这个比较常用
	*/
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}


func testBufio(){
		//有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。示例代码如下
		reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
		fmt.Print("请输入内容：")
		text, _ := reader.ReadString('\n') // 读到换行
		text = strings.TrimSpace(text)
		fmt.Printf("%v\n", text)
		fmt.Printf("%#v\n", text)
}



func main() {
	// testPrint()
	// testFprint()
	// testSprint()
	// testEorrorf()
	// testFormat()
	// testInt()
	// testFloat()
	// testByte()
	// testPoint()
	// testPrecision()
	// testOther()
	// testScan()
	// testScanf()
	// testScanln()
	// testBufio()
}
