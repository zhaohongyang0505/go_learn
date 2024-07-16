package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 条件循环
func main() {
	//game

	second := time.Now().Unix()
	source := rand.NewSource(second)
	rand := rand.New(source)
	target := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guess := 0; guess < 10; guess++ {
		fmt.Println("guess:", guess+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("dile")
		} else if guess > target {
			fmt.Println("gaole")
		} else {
			fmt.Println("duile")
			success = true
			break
		}
	}
	if !success {
		fmt.Println(target)
	}

}
