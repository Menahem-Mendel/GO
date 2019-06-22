// найдите веб сайт, который содержит большое количество данных.
// Исследуйте работу кеширования путем двукратного запуска fetchall и сравнения времени запросов.
// Получаете ли вы каждый раз одно и то де содержимое?
// Измените fetchall так,
// чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(f, <-ch)
	}
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
