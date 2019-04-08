package main

var echo string
func main() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			echo = scanner.Text()
		}
	}()

	http.HandleFunc("/echo", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	flusher := w.(http.Flusher)

	fmt.Fprintln(w, echo)
	flusher.Flush()
}
