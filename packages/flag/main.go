package main

import "os"
import "fmt"
import "flag"
import "time"

//Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。

/*
	有以下两种常用的定义命令行flag参数的方法。

	flag.Type()
	基本格式如下：

	flag.Type(flag名, 默认值, 帮助信息)*Type 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	需要注意的是，此时name、age、married、delay均为对应类型的指针。

	flag.TypeVar()
	基本格式如下： flag.TypeVar(Type指针, flag名, 默认值, 帮助信息) 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
*/

/*
	使用
		命令行参数使用提示：

		$ ./flag_demo -help
		Usage of ./flag_demo:
		-age int
				年龄 (default 18)
		-d duration
				时间间隔
		-married
				婚否
		-name string
				姓名 (default "张三")
		正常使用命令行flag参数：

		$ ./flag_demo -name 沙河娜扎 --age 28 -married=false -d=1h30m
		沙河娜扎 28 false 1h30m0s
		[]
		0
		4
		使用非flag命令行参数：

		$ ./flag_demo a b c
		张三 18 false 0s
		[a b c]
		3
		0
*/

//os.Args
func testArgs() {
	//os.Args是一个[]string
	//os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

//flag demo
func flagDemo() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

func main() {
	// testArgs()
	flagDemo()

}
