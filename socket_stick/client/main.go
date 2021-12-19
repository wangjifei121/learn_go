package main

import "fmt"
import "net"
import "basic_lesson/socket_stick/proto"

/*
主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。

“粘包”可发生在发送端也可发生在接收端：

由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

*/

//粘包现象代码
// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:30000")
// 	if err != nil {
// 		fmt.Println("dial failed, err", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for i := 0; i < 20; i++ {
// 		msg := `Hello, Hello. How are you?`
// 		conn.Write([]byte(msg))
// 	}
// }

//解决粘包代码
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//解决粘包代码 - 开始
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
		//解决粘包代码 - 结束
		// conn.Write([]byte(msg))
	}
}
