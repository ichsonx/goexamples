package main

/*
解释，当接收者方法的接收者分别为 【指针】、【值】 的情况下会有什么结果
无论是【指针】，还是【值】接收者，都不影响方法的使用。可以互相使用。【值】接收者，调用【指针】接收者方法，仍然可以达到【指针】效果。
*/
import "fmt"

type receiver struct {
	Name string
}

func (r receiver) change() {
	r.Name = "new name!"
}
func (r *receiver) change2() {
	r.Name = "new name!"
}

func main() {
	//==============================指针、值接受者方法的区别================================
	fmt.Println("定义一个  非指针receiver  变量，分别调用change、change2，结果如下：")
	r := receiver{"old name..."}               //定义一个  值的receiver  变量，名为 old name
	r.change()                                 //调用  值接收者方法  change，修改名字
	fmt.Printf("call change : %s \n", r.Name)  //打印的仍然是old name
	r.change2()                                //调用  指针接收者方法  change2，修改名字
	fmt.Printf("call change2 : %s \n", r.Name) //修改成功，打印的是new name
	//由此可发现  值接收者  都可以调用  指针接收者方法、值接收者方法，最后效果会跟调用的方法一致。

	fmt.Println("定义一个  指针receiver  变量，分别调用change、change2，结果如下：")
	rr := new(receiver) //定义一个  值的receiver  变量，名为 rr old name
	rr.Name = "rr old name..."
	rr.change()                                 //调用  值接收者方法  change，修改名字
	fmt.Printf("call change : %s \n", rr.Name)  //打印的仍然是old name
	rr.change2()                                //调用  指针接收者方法  change2，修改名字
	fmt.Printf("call change2 : %s \n", rr.Name) //修改成功，打印的是new name
	//由此可发现  指针接收者  都可以调用  指针接收者方法、值接收者方法，最后效果会跟调用的方法一致。

	//结论，无论是指针，还是值接收者，都不影响方法的使用。可以互相使用。值接收者，调用指针接收者方法，仍然可以达到指针效果。
}
