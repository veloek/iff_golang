package main
import "fmt"

func main() {
	foobar(new(output))
}
type logger interface {
	log(msg string)
}
func foobar(l logger) {
	l.log("Hello, World!")
}

// output kunne like gjerne vært i en annen pakke.
type output struct{}
func (o *output) log(msg string) {
	fmt.Println(msg)
}

