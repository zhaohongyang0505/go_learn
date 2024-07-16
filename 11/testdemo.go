package main

import (
	"fmt"
	"learn/testdemo"
)

func main() {
	phrases := []string{"a", "b"}
	fmt.Println(testdemo.JoinWithCommas(phrases))
	phrases = []string{"a", "b", "c"}
	fmt.Println(testdemo.JoinWithCommas(phrases))
}
