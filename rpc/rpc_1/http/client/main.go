/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2022/6/05
 * @Description:
	一个最基本的golang rpc例子，本代码是客户端【http版本】代码。
	与java的socket、tcp连接很相似。grpc的客户端同样也通过协议、地址：端口号进行拨号连接，
	然后请求具体的grpc服务方法，同时传递参数、等待服务返回响应。
*/
package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立链接
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	// 返回
	var result int
	//  请求方法
	for i := 1; i < 10; i++ {
		err = client.Call("MathService.Multi", i, &result)
		fmt.Printf("i:%v result:%v \n", i, result)
	}
}
