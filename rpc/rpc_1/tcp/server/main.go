/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2022/6/05
 * @Description:
	一个最基本的golang rpc例子，本代码是服务端代码。
	- 与java的socket、tcp连接很相似。grpc同样也是注册服务、监听端口、处理请求并返回的逻辑
	- grpc使用的底层通信协议有 http版、tcp版、jsonrpc版。【本例使用tcp版】的。
	- **注意：上面TCP版和HTTP版,数据编码采用的都是默认的gob编码，而gob编码是Go特有的编码和解码的专用序列化方法，这也就意味着Gob无法跨语言使用。而JSON RPC则可以跨语言使用
*/
package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (h *HelloService) Say(request string, response *string) error {
	format := time.Now().Format("2006-01-02 15:04:05")
	*response = request + " -- " + format
	return nil
}

func main() {
	// 注册服务名称
	_ = rpc.RegisterName("HelloService", new(HelloService))
	// 监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	for {
		// 监听请求
		accept, err := listen.Accept()
		if err != nil {
			log.Fatalf("Accept Error: %s", err)
		}
		// 使用一个协程来建立链接，并处理该请求
		go rpc.ServeConn(accept)
	}
}
