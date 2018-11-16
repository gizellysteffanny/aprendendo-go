package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	canal := make(chan string)
	for _, url := range urls {
		go func(url string) { // func anônima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			canal <- r.FindStringSubmatch(string(html))[1]
		}(url) // chamando a func anônima
	}

	return canal
}

func main() {
	titulo1 := titulo("https://www.github.com", "https://www.youtube.com")
	titulo2 := titulo("https://www.google.com", "https://www.slack.com")

	fmt.Println("Primeiros: ", <-titulo1, " | ", <-titulo2)
	fmt.Println("Segundos: ", <-titulo1, " | ", <-titulo2)
}
