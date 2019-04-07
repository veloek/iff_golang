package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://golangcode.com/images/avatar.jpg"
	img := download(url)
	fmt.Printf("Downloaded %d bytes\n", len(img))
}

func download(url string) []byte {
	resp, _ := http.Get(url)
	data, _ := ioutil.ReadAll(resp.Body)
	return data
}
