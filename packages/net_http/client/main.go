package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，
//所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

//基本get请求
func getDemo(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error:", err)
	}
	defer resp.Body.Close() //程序在使用完response后必须关闭回复的主体。

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil readall error:", err)
	}
	fmt.Println(string(body))
}

//带参数的get请求
func paramsGetDemo() {
	//借助net/url来处理url
	apiUrl := "http://127.0.0.1:8080/user/search"
	//URL params
	data := url.Values{}
	data.Set("username", "wangjifei")
	data.Set("address", "beijing")
	//parse url
	url, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("parse url error:", err)
	}
	url.RawQuery = data.Encode() //url encode
	fmt.Println(url.String())    //输出url

	//二、直接拼接url
	// urlString := "http://127.0.0.1:8080/user/search?address=beijing&username=wangjifei"
	// getDemo(urlString)
	getDemo(url.String())
}

//POST请求示例
func postDemo(url string, contentType string, data string) {
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

//表单数据
func formPostDemo() {
	contentType := "application/x-www-form-urlencoded"
	data := "username=wangjifei&address=shanghai"
	url := "http://127.0.0.1:9090/"
	postDemo(url, contentType, data)
}

// json数据
func jsonPostDemo() {
	contentType := "application/json"
	data := `{"username":"wangjifei","address":"shenzhen"}`
	url := "http://127.0.0.1:9090/"
	postDemo(url, contentType, data)
}

/*
	//自定义Client
	//要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http://example.com")
	// ...
	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...

*/

/*
	//自定义Transport
	//要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
*/

func main() {
	// u :="http://www.baidu.com"
	// getDemo(u)
	// paramsGetDemo()
	// formPostDemo()
	jsonPostDemo()
}
