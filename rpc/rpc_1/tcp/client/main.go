/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2022/6/05
 * @Description:
	一个最基本的golang rpc例子，本代码是客户端代码。
	与java的socket、tcp连接很相似。grpc的客户端同样也通过协议、地址：端口号进行拨号连接，
	然后请求具体的grpc服务方法，同时传递参数、等待服务返回响应。
*/
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	// 建立链接
	dial, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dial error ", err)
	}
	var result string
	for i := 0; i < 5; i++ {
		// 发起请求
		_ = dial.Call("HelloService.Say", "go", &result)
		fmt.Println("result:", result)
		time.Sleep(time.Second * 2)
	}
}
