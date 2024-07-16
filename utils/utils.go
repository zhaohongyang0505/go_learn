// 包注释，会展示在go doc的输出中
package utils

import "fmt"

// 包函数注释，会展示在go doc的输出中
func Hello() {
	fmt.Println("hello")
}

func Hi() {
	fmt.Println("hi")
}

type MyStruct struct {
	Number      float64
	Word        string
	Toggle      bool
	HomeAddress Address
	Address
}

type Address struct {
	City  string
	Sheng string
}
