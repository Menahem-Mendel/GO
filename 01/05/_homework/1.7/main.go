// вызов функции io.Copy(dst, src) выполняет чтение src и запись в dst. 
// Воспользуйтесь ею вместо ioutil.ReadAll для копирования тела ответа в поток os.Stdout 
// без необходимости выделения достаточно большого для хранения всего ответа буфера. 
// Нt забудьте проверить, не произошла ли ошибка при вызове io.Copy
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}