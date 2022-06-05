/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2022/6/05
 * @Description:
	一个最基本的golang rpc例子，本代码是服务端代码。
	- 与java的socket、tcp连接很相似。grpc同样也是注册服务、监听端口、处理请求并返回的逻辑
	  这里相对rpc版本代码上更简单点，一旦监听并自动交由httphandle处理。
	- grpc使用的底层通信协议有 http版、tcp版、jsonrpc版。【本例使用http版】的。
	- **注意：上面TCP版和HTTP版,数据编码采用的都是默认的gob编码，而gob编码是Go特有的编码和解码的专用序列化方法，这也就意味着Gob无法跨语言使用。而JSON RPC则可以跨语言使用
*/
package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

// 相乘方法
func (u *MathService) Multi(a int, sum *int) error {
	*sum = a * a
	return nil
}

func main() {
	userService := new(MathService)
	// 注册服务
	err := rpc.Register(userService)
	if err != nil {
		return
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
