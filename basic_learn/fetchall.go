// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	/* 获得当前时间 */
	start := time.Now()
	/* 生成一个channel管道？ */
	ch := make(chan string)
	/* 对每一个参数都执行"go"前缀 协程 fetch */
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	/* 并将协程结果打印 */
	for range os.Args[1:] {
		/* 接收 */
		fmt.Println(<-ch) // receive from channel ch
	}
	/* 打印总时间 */
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	/* 获得当前时间 */
	start := time.Now()
	/* 执行get */
	resp, err := http.Get(url)
	/* 错误信息返回 */
	if err != nil {
		/* 发送 */
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	/* 这是将body写到 discard 丢弃 */
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	/* 打印url和字节数量和时间 */
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/*
2.33s     4154  http://gopl.io
Get "https://golang.org": dial tcp 216.239.37.1:443: i/o timeout
Get "https://pkg.go.dev/?utm_source=godoc": dial tcp 172.217.24.19:443: i/o timeout
30.45s elapsed
*/
