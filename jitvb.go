package main
import (
	"fmt"
	"time"
)
func main() {
	for {
		fmt.Println("Jusqu'ici tout va bien...")
		time.Sleep(500 * time.Millisecond)
	}
}
