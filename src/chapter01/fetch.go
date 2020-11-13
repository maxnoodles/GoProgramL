package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//func main(){
//	for _, url := range os.Args[1:]{
//		resp, err := http.Get(url)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//			os.Exit(1)
//		}
//		b, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
//			os.Exit(1)
//		}
//		fmt.Printf("%s", b)
//	}
//}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch)
	}
	fmt.Printf(" %.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
