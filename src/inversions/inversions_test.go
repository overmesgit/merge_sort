package main
import (
    "testing"
    "fmt"
)

func TestSort(t *testing.T) {
    input := []map[string][]int{}
    input = append(input, map[string][]int{"test": {1, 2, 3}, "res": {0}})
    input = append(input, map[string][]int{"test": {3, 2, 1}, "res": {3}})

    for _, testCase := range input {
        result := InversionsCounter(testCase["test"])

        if result != testCase["res"][0] {
            t.Errorf("input: %v inversions: %v != %v", testCase["test"], testCase["res"][0], result)
        }
    }
    fmt.Println("Done")
}
