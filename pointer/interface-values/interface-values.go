package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {

	fmt.Println(t.S)
}

func (f F) M() {

	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	// 没有实现接口的数据类型无法给接口初始化，如下面注释的代码，因为float64没有继承接口I（没有实现接口的方法），所以这些写会产生语法错误提示
	// i = float64(3.14.15926)
	i = F(3.1415926)
	describe(i)
	i.M()
}
