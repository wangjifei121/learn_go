package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func onlyRead() {
	//只读的方式打开当前目录的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file error :", err)
		return
	}
	//关闭文件
	defer file.Close()

	//使用read方法读取文件内容
	var tmp = make([]byte, 1280)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("file read end")
		return
	}
	if err != nil {
		fmt.Println("read file fail:", err)
		return
	}
	fmt.Printf("读取%d的字节数据\n", n)
	fmt.Println("读取数据内容为：", string(tmp[:n])) //将读取到的bety类型转化为string
}

//循环读取所有文件内容
func rangeRead() {
	//只读的方式打开当前目录的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file error :", err)
		return
	}
	//关闭文件
	defer file.Close()

	//使用read方法循环读取文件内容
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("file read end")
			break
		}
		if err != nil {
			fmt.Println("read file fail:", err)
			break
		}
		content = append(content, tmp[:n]...)
	}

	fmt.Printf("读取%d的字节数据\n", len(content))
	fmt.Println("读取数据内容为：", string(content)) //将读取到的bety类型转化为string
}

//bufio是在file的基础上封装了一层API，支持更多的功能
//bufio按行读取
func bufioReadLine() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

//ioutil.ReadFile读取整个文件
//io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入
func ioutilRead() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("ioutil read fill error:", err)
		return
	}
	fmt.Println(string(content))
}

//Write和WriteString
func write() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据
}

//bufio.NewWriter
func bufioWrite() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

//ioutil.WriteFile
func ioutilWrite() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

//借助io.Copy()实现一个拷贝文件函数
func copyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func main() {
	// onlyRead()
	// rangeRead()
	// bufioReadLine()
	// ioutilRead()

	// write()
	// bufioWrite()
	// ioutilWrite()
	copyFile("xx.txt", "c.txt")
}
