package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	c := &Class{
		Title:    "清华班",
		Students: make([]*Student, 0, 100),
	}
	for i := 0; i < 100; i++ {
		stu := Student{
			Name: fmt.Sprintf("student_%02d", i),
			ID:   i,
		}
		c.Students = append(c.Students, &stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, error := json.Marshal(c)
	if error != nil {
		fmt.Println("----json-error-----")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	jsonStr := data
	c1 := &Class{}
	err := json.Unmarshal([]byte(jsonStr), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}

//学生结构体
type Student struct {
	Name string
	ID   int
}

//班级结构体

type Class struct {
	Title    string
	Students []*Student
}
