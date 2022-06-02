/*
本片段摘自《go语言36讲》第13讲
  1.一个自定义数据类型的方法集合中仅会包含它的所有【值方法】，而该类型的【指针类型】的方法集合却囊括了前者的所有方法，包括所有【值方法】和所有【指针方法】。
严格来讲，我们在这样的基本类型的值上只能调用到它的值方法。但是，Go 语言会适时地为我们进行自动地转译，使得我们在这样的值上也能调用到它的指针方法。
比如，在Cat类型的变量cat之上，之所以我们可以通过cat.SetName("monster")修改猫的名字，是因为 Go 语言把它自动转译为了(&cat).SetName("monster")，即：先取cat的指针值，然后在该指针值上调用SetName方法。

  2.在后边你会了解到，一个类型的方法集合中有哪些方法与它能实现哪些接口类型是息息相关的。如果一个【基本类型】和它的【指针类型】的方法集合是不同的，
那么它们具体实现的接口类型的数量就也会有差异，除非这两个数量都是零。
  比如，一个【指针类型】实现了某某接口类型，但它的【基本类型】却不一定能够作为该接口的实现类型。
能够体现【值方法】和【指针方法】之间差异的小例子我放在 demo30.go 文件里了，你可以参

3.以上叙述以理论为基础。如果从【调用】、【包含】的角度出发：
  ((包含角度)). struct的【值类型】只包含【值方法】，而【指针类型】包含了【值方法】+【指针方法】。这个包含关系将直接影响到是否实现了某些【接口】。
  ((调用角度)). struct的【值类型】、【指针类型】，都可以调用对方所包含的所有方法（go自动做了转换）。

4.同时，无论【值类型】还是【指针类型】调用【值方法】or【指针方法】，最终的效果，还是根据落在方法上的。不会因为你是什么类型而有所改变。（即依旧只有【指针方法】会改变变量的值）
*/
package main

import "fmt"

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster") // (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}
