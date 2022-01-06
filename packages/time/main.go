package main

import (
	"fmt"
	"time"
)

//获取当前时间
func testTimeNow() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//获取时间戳
func testTimestampDemo() {
	now := time.Now()             //获取当前时间
	timestamp1 := now.Unix()      //时间戳
	timestamp2 := now.UnixMilli() //毫秒时间戳
	timestamp3 := now.UnixMicro() //微秒时间戳
	timestamp4 := now.UnixNano()  //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	fmt.Printf("current timestamp3:%v\n", timestamp3)
	fmt.Printf("current timestamp4:%v\n", timestamp4)
}

//将时间戳转成时间对象
func testTimeUnix(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间间隔
func testTimeDuration() {
	nan := time.Nanosecond
	mic := time.Microsecond
	mil := time.Millisecond
	s := time.Second
	m := time.Minute
	h := time.Hour
	fmt.Printf("time.Nanosecond:%v\n", nan)
	fmt.Printf("time.Microsecond:%v\n", mic)
	fmt.Printf("time.Millisecond:%v\n", mil)
	fmt.Printf("time.Second:%v\n", s)
	fmt.Printf("time.Minute:%v\n", m)
	fmt.Printf("time.Hour:%v\n", h)
}

//时间操作
func testTimeOpt() {
	// 当前时间加1小时后的时间
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	//求时间耗时
	start := time.Now()
	time.Sleep(time.Second)
	fmt.Println(time.Now().Sub(start).Milliseconds())

	//两个时间是否相等
	t1 := time.Now()
	fmt.Println(time.Now().Equal(t1))

	//两个时间前后
	t2 := time.Now()
	fmt.Println(time.Now().Before(t2))

	t3 := time.Now()
	fmt.Println(time.Now().After(t3))
}

//定时器 - 定时器的本质上是一个通道（channel）
func testTimeTick(){
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}

//时间格式化
func testTimeFormat(){
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

//解析字符串格式的时间
func testTimeParse(){
	now := time.Now()
	fmt.Println("当前时间：",now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("输入时间：",timeObj)
	fmt.Println("时间差：",now.Sub(timeObj))

}

func main() {

	// testTimeNow()
	// testTimestampDemo()
	// testTimeUnix(time.Now().Unix())
	// testTimeDuration()
	// testTimeOpt()
	// testTimeTick()
	// testTimeFormat()
	testTimeParse()
}
