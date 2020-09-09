package main

import (
	"fmt"
	"time"
)

// MyError is a struct
type MyError struct {
	When time.Time
	What string
}

// MyError结构体实现了error接口的Error方法，通过指针向Error方法中传入MyError结构体实例。
func (e *MyError) Error() string {

	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run函数的返回类型是error接口
// run函数中实例化了一个MyError结构体，并返回结构体实例的引用。
// 注意：如果MyError结构体没有实现error接口的Error方法，run函数中返回MyError结构体实例的引用将引起语法错误
func run() error {

	temp := &MyError{
		time.Now(),
		"it didn't work",
	}

	fmt.Printf("MyError instance addr is: %p\n", temp)

	return temp
}

func main() {

	// 初始化err变量，
	if err := run(); err != nil {

		fmt.Printf("err addr is: %p\n", err)
		fmt.Println(err)
	}
}
