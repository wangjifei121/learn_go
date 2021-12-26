package split

import (
	"fmt"
	"reflect"
	"testing"
)

/*
测试函数示例
	go test     //执行测试
	go test -v  //执行测试，查看测试函数名称和运行时间
	go test -v -run=RegExp //指定某些类型的方法测试
	go test -v -run=Split/simple只会运行simple对应的子测试用例
	go test -cover  //测试覆盖率 （测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。）
	go test -cover -coverprofile=c.out //将覆盖率相关的记录信息输出到一个文件
	go tool cover -html=c.out  //使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。
基准测试
	go test -bench=Split //基准测试命令
	go test -bench=Split -benchmem //获得内存分配的统计数据


*/

//测试函数示例
func TestSplit(t *testing.T) {
	// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         //方法实际得到的结果
	want := []string{"a", "b", "c"}    //预期结果
	if !reflect.DeepEqual(got, want) { //因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("exepcted:%v,got:%v\n", want, got) //错误信息
	}

}

func TestMoreSplit(t *testing.T) {
	got := Split("abcde", "cd")
	want := []string{"ab", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("exepcted:%v,got:%v\n", want, got)
	}
}

func TestChineseSplit(t *testing.T) {
	//定义一个测试用例类型
	type test struct {
		content string
		sep     string
		want    []string
	}
	// 定义一个存储测试用例的切片
	var my_test = []test{
		{content: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{content: "a:b:c", sep: ".", want: []string{"a:b:c"}},
		{content: "abcde", sep: "bc", want: []string{"a", "de"}},
		{content: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for _, tc := range my_test {
		got := Split(tc.content, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("exepcted:%#v,got:%#v", tc.want, got)
		}
	}
}

//子测试
func TestSubSplit(t *testing.T) {
	//定义一个测试的结构体
	type test struct {
		content string
		sep     string
		want    []string
	}
	//map类型测试用例
	tests := map[string]test{
		"simple":       {content: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong simple": {content: "a:b:c", sep: ".", want: []string{"a:b:c"}},
		"more sep":     {content: "a:b:c", sep: "b:", want: []string{"a:", "c"}},
		"chiness sep":  {content: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.content, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("exepcted:%#v,got:%#v", tc.want, got)
			}
		})
	}
}

//基准测试示例
//基准测试就是在一定的工作负载之下检测程序性能的一种方法

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

/*
go test -bench=Split

	(base) 192:split wangjifei$ go test -bench=Split
	goos: darwin
	goarch: amd64
	pkg: basic_lesson/test/split
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkSplit-8         5468158               216.6 ns/op
	PASS
	ok      basic_lesson/test/split 1.415s
	(base) 192:split wangjifei$

	- 其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
	10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

go test -bench=Split -benchmem

	(base) 192:split wangjifei$ go test -bench=Split -benchmem
	goos: darwin
	goarch: amd64
	pkg: basic_lesson/test/split
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkSplit-8         5384418               217.5 ns/op           112 B/op          3 allocs/op
	PASS
	ok      basic_lesson/test/split 1.406s
	(base) 192:split wangjifei$

	- 112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配
*/

//子测试的Setup与Teardown
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown\n")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown\n")
	}
}

func TestSplitSetupAndTeardown(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTestCase := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTestCase(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)         //执行测试
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%v got:%v", tc.want, got)
			}
		})
	}
}


//示例函数
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}