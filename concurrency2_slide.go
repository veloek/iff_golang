func main() {
	url := "https://golangcode.com/images/avatar.jpg"
	var img []byte

	done := make(chan []byte)
	go func() {
		done <- download(url)
	}()
	fmt.Printf("Downloading ")
loop:
	for {
		select {
		case img = <-done:
			fmt.Println()
			break loop
		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	fmt.Printf("Downloaded %d bytes\n", len(img))
}
