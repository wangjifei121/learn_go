package main

import (
	"fmt"
	// "strings"
	// "math/rand"
	// "sort"
)

func main() {
	//map 定义方式
	// a := make(map[int]string, 8) //定义map
	// a[1] = "beijing"
	// a[2] = "shanghai"
	// fmt.Println(a)
	// fmt.Printf("%#v\n", a)
	// fmt.Printf("a type of %T\n", a)

	//map也支持在声明的时候填充元素
	// userInfo := map[string]string{
	// 	"username":"wangjifei",
	// 	"age":"30",
	// }
	// fmt.Println(userInfo)

	// var userInfo = make(map[string]string, 8)
	// userInfo["username"] = "wangjifei"
	// userInfo["age"] = "30"
	// fmt.Println(userInfo)

	//map 判断key是否存在

	// userInfo := make(map[string]string,8)
	// userInfo["username"] = "wangjifei"
	// userInfo["age"] = "30"
	// fmt.Println(userInfo)
	// val, ok := userInfo["sex"]
	// if ok{//判断key是否存在
	// 	fmt.Println(val)
	// }else{
	// 	userInfo["sex"] = "nan"
	// }
	// fmt.Println(userInfo)

	//map的遍历
	// studentScoreMap := make(map[string]int, 100)
	// for i := 0; i < 50; i++ {
	// 	studentName := fmt.Sprintf("student_%02d", i)
	// 	score := rand.Intn(100)
	// 	studentScoreMap[studentName] = score
	// }
	// for key, val := range studentScoreMap {
	// 	fmt.Printf("key:%v val:%v\n", key, val)
	// }

	//map指定顺序遍历
	// studentScoreMap := make(map[string]int, 100)
	// for i := 0; i < 50; i++ {
	// 	studentName := fmt.Sprintf("student_%02d", i)
	// 	score := rand.Intn(100)
	// 	studentScoreMap[studentName] = score
	// }
	// nameSlice := make([]string, 0, 100)
	// for k := range studentScoreMap {
	// 	nameSlice = append(nameSlice, k)
	// }
	// fmt.Println(nameSlice)
	// sort.Strings(nameSlice)
	// fmt.Println(nameSlice)
	// for _, name := range nameSlice {
	// 	fmt.Printf("name=%v value=%d\n", name, studentScoreMap[name])
	// }

	//元素为map的切片
	// mapSlice := make([]map[string]string, 10)
	// mapSlice[0] = make(map[string]string, 5)
	// mapSlice[0]["name"] = "wangjifei"
	// mapSlice[0]["age"] = "30"
	// fmt.Println(mapSlice)
	// for _,v := range mapSlice{
	// 	fmt.Println(v)
	// }

	//值为slice的map
	// sliceMap := make(map[string][]string,10)
	// sliceMap["name"] = []string{
	// 	"wangjifei",
	// 	"kanghuayue",
	// }
	// fmt.Println(sliceMap)

	// sliceMap := make(map[string][]string,10)
	// s := make([]string, 0, 10)
	// s = append(s, "wangjifei","kanghuayue")
	// sliceMap["name"] = s

	// fmt.Println(sliceMap)

	//练习题
	//1、写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
	// var s string = "how do you do"
	// ss := strings.Split(s, " ") //字符串切片
	// fmt.Println(ss)
	// wordCount := make(map[string]int)
	// for i := 0; i < len(ss); i++ {
	// 	v, ok := wordCount[ss[i]]
	// 	if ok {
	// 		wordCount[ss[i]] = v + 1
	// 	} else {
	// 		wordCount[ss[i]] = 1
	// 	}
	// }
	// fmt.Println(wordCount)

	//2、观察下面代码，写出最终的打印结果
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //[1,2,3]
	m["q1mi"] = s
	fmt.Println("--", m)
	s1 := s[:1]
	s2 := s[2:]
	fmt.Println("--s1--", s1)
	fmt.Println("--s2--", s2)
	s = append(s1, s2...)
	fmt.Printf("%+v\n", s)         //[1,3]
	fmt.Printf("%+v\n", m["q1mi"]) //[1,3]????
}
