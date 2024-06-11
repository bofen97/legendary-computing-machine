package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

type BrowserFetch struct {
}

func (b *BrowserFetch) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error: %v ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code %v \n", resp.StatusCode)
		return nil, err
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
func main() {

	url := "https://news.163.com/"
	B := &BrowserFetch{}
	body, err := B.Get(url)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	numlinks := strings.Count(string(body), "<a")
	fmt.Println("numlinks : ", numlinks)

	exits := strings.Contains(string(body), "黄金")
	fmt.Println("has 黄金？ : ", exits)

	numlinks = bytes.Count(body, []byte("<a"))
	fmt.Println("numlinks (bytes): ", numlinks)

	exits = bytes.Contains(body, []byte("黄金"))
	fmt.Println("has 黄金？ (byte) : ", exits)

}
