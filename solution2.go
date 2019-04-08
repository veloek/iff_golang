package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var echo = make(chan string)
var clients []chan string

func main() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			echo <- scanner.Text()
		}
	}()

	go fanout()

	http.HandleFunc("/echo", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	flusher := w.(http.Flusher)

	client := make(chan string)
	clients = append(clients, client)

	for e := range client {
		fmt.Fprintln(w, e)
		flusher.Flush()
	}
}

func fanout() {
	for x := range echo {
		for _, o := range clients {
			o <- x
		}
	}
}
