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
	funcs := []func([]int, ... int) []int{Merge, Quick}
    for _, f := range funcs {
        for i := 0; i < 100; i++ {
            rand.Seed(int64(time.Now().UnixNano()))
            init := rand.Perm(100)
            input := make([]int, 100)
            copy(input, init)
            result := f(input)
            sort.Sort(intArray(init))
            for i, v := range init {
                if v != result[i] {
                    t.Errorf("%v != %v", init, result)
                    t.FailNow()
                }
            }
	    }
    }



}

var result []int
func BenchmarkMergeSort(b *testing.B) {
	var input []int = rand.Perm(50000)
	b.ResetTimer()
	result = Merge(input)
    time.Sleep(1000*time.Millisecond)
}

func BenchmarkQuickSort(b *testing.B) {
	var input []int = rand.Perm(50000)
	b.ResetTimer()
	result = Quick(input)
    time.Sleep(1000*time.Millisecond)
}