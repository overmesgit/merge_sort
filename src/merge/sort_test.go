package merge

import (
	"testing"
	"math/rand"
	"time"
	"sort"
)

type intArray []int
func (s intArray) Len() int { return len(s) }
func (s intArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

func TestSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		rand.Seed(int64(time.Now().UnixNano()))
		var input []int = rand.Perm(100)
		result := Merge(input)
		sort.Sort(intArray(input))
		for i, v := range input {
			if v != result[i] {
				t.Errorf("%v != %v", input, result)
				t.FailNow()
			}
		}
	}

}

var result []int
func BenchmarkMergeSort(b *testing.B) {
	var input []int = rand.Perm(50000)
	b.ResetTimer()
	result = Merge(input)
}