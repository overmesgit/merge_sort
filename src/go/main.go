package main
import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"merge"
)

func main() {
	rand.Seed(int64(time.Now().UnixNano()))
	var input []int = rand.Perm(5000)
	start := time.Now().UnixNano()
	fmt.Println(merge.Merge(input)[:10])
	end := time.Now().UnixNano()

	fmt.Println(float64(end - start)/math.Pow(10, 9))

}