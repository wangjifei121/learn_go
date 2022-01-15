package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//使用Go语言中的net/http包来编写一个简单的接收HTTP请求的Server端示例，net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据
//parse form
func postFormHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))

	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

//parse json
func postJsonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	// 开启server
	// http.HandleFunc("/", postFormHandler)
	http.HandleFunc("/", postJsonHandler)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

/*
	自定义Server
	要管理服务端的行为，可以创建一个自定义的Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

*/