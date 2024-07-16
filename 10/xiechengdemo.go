package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func respSize(url string, channel chan Page) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
func main() {
	pageChan := make(chan Page)
	urls := []string{"https://www.baidu.com", "https://www.qq.com", "https://www.taobao.com"}
	for _, url := range urls {
		go respSize(url, pageChan)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pageChan
		fmt.Println(page.URL, page.Size)
	}
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("接收协程", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

func send(myChannel chan string) {
	reportNap("发送协程", 2)
	fmt.Println("发送数据")
	myChannel <- "a"
	fmt.Println("发送数据")
	myChannel <- "b"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "睡眠")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "恢复")
}
