/*
  1. 外层结构，【内嵌类型变量/方法】没有与【外部结构的变量/方法】重名的情况下，
     可“透明”地直接用选择表达式即“.”来访问内嵌结构体的字段及方法
  2. 《go语言核心36讲》中并没有提到这点，但通过了另外一个方式即：外层结构对象.内嵌结构类型.内嵌结构字段/方法 来访问.
      其实第一点会更加好理解和方便使用。
*/

package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}
