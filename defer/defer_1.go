/**
 * @Author: sonic
 * @File:  godefer.go
 * @Date: 2019/10/12 14:47
 * @Description: defer

	被推迟函数的实参（如果该函数为方法则还包括接收者）在推迟执行时就会被求值，而不是在调用执行时才求值。
	这样不仅无需担心变量值在函数执行时被改变，同时还意味着单个被推迟的调用可推迟多个函数的执行。
*/
package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	// 被推迟函数un()，在defer执行时，里面的trace("a")已经被求值（执行），因此会先打印【 entering: a】
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	// 被推迟函数un()，在defer执行时，里面的trace("b")已经被求值（执行），因此会先打印【 entering: b】
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {

	//输出结果： 4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	// 利用这个特点，即被推迟函数的实参在 defer 执行时就会被求值。跟踪例程可针对反跟踪例程设置实参。
	b()
	/*
		输出结果为
		entering: b
		in b
		entering: a
		in a
		leaving: a
		leaving: b
	*/
}
