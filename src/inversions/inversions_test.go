package main
import (
    "testing"
    "fmt"
)

func TestSort(t *testing.T) {
    input := []map[string][]int{}
    input = append(input, map[string][]int{"test": {1, 2}, "res": {0}})
    input = append(input, map[string][]int{"test": {2, 1}, "res": {1}})
    input = append(input, map[string][]int{"test": {1, 2, 3}, "res": {0}})
    input = append(input, map[string][]int{"test": {3, 2, 1}, "res": {3}})
    input = append(input, map[string][]int{"test": {1, 3, 5, 6, 2, 7}, "res": {3}})
    input = append(input, map[string][]int{"test": {1, 3, 5, 2, 4, 6}, "res": {3}})
    input = append(input, map[string][]int{"test": {1, 1, 1, 1, 1, 1}, "res": {0}})

    for _, testCase := range input {
        array, result := InversionsCounter(testCase["test"])

        if result != testCase["res"][0] {
            t.Errorf("input: %v inversions: %v != %v", array, testCase["res"][0], result)
        }
    }
    fmt.Println("Done")
}
