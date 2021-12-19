package main

import "net"
import "bufio"
import "io"
import "fmt"
import "basic_lesson/socket_stick/proto"

//粘包现象代码
// func process(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	var buf [1024]byte
// 	for {
// 		n, err := reader.Read(buf[:])
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("read from client failed, err:", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到client发来的数据：", recvStr)
// 	}
// }

//解决粘包代码
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}