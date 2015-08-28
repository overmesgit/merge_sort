package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main() {
	fmt.Println("Hello, Артем!")
	fmt.Println(time.Now())
	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println(rand.Perm(100))

	a := make([]int, 0, 0)
	fmt.Println(a[:2])
}
