/**
 * @Author: sonic
 * @File:  gointerface_1.go
 * @Date: 2019/9/26 15:26
 * @Description: golang中 【接口】 nil != nil问题。从中可明白interface的基础结构原理。
		1、接口在golang内部实现中保存了2个值，一个是类型T，一个是值V。因此我们在讨论接口时，永远是在讨论接口的【类型】及【类型的值】。
		2、只有【类型】与【类型的值】同为nil的时候，接口才是nil。( T=nil, V=nil) == nil interface。
		3、在golang中，接口是隐式实现。因此，当我们用一个接口类型去接收一个nil结构体的时候，那么这个接口将【不会是nil】。
		   因为此时的接口的值为( T=struct, V = nil )，不满足接口为nil条件。
*/
package main

import "fmt"

// 接口 nil != nil 经典问题的前置环境  start
type DetailedError struct {
	code    int
	message string
}

//DetailedError实现error接口方法Error()。
func (e DetailedError) Error() string {
	return fmt.Sprintf("Error occured at (%d,%s)", e.code, e.message)
}
func handle(x int) *DetailedError {
	if x != 1 {
		return &DetailedError{code: 1000, message: "who am i ?"}
	}
	return nil
}

// 接口 nil != nil 经典问题的前置环境  end

func main() {
	var err error
	err = handle(0)
	if err != nil {
		fmt.Printf("i am Error 1 of %s\n", err)
	}
	err = handle(1)
	if err != nil {
		fmt.Printf("i am Error 2 of %s\n", err)
	}
	/*
		结果打印如下：
		i am Error 1 of Error occured at (1000,who am i ?)
		i am Error 2 of <nil>
	*/
	/*
		1、接口在golang内部实现中保存了2个值，一个是类型T，一个是值V。因此我们在讨论接口时，永远是在讨论接口的【类型】及【类型的值】。
		2、只有【类型】与【类型的值】同为nil的时候，接口才是nil。( T=nil, V=nil) == nil interface。
		3、在golang中，接口是隐式实现。因此，当我们用一个接口类型去接收一个nil结构体的时候，那么这个接口将【不会是nil】。
		   因为此时的接口的值为( T=struct, V = nil )，不满足接口为nil条件。
	*/
}
