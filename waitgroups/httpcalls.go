package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://midnightprogrammer.net",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homePage() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()

}

func main() {
	homePage()
}
