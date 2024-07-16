package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoNumSum([]int{2, 7, 11, 15}, 9))
}

func twoNumSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
