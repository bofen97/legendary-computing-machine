package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://news.163.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code %v \n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error for read body  %v \n", err)
		return
	}
	fmt.Println("body:", string(body))

}
