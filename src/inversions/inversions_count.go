package main

import (
    "io/ioutil"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("src/inversions/IntegerArray.txt")
    fmt.Println(err)
    string_data := string(data)
    split_data := strings.Split(string_data, "\r\n")
    split_data = split_data[:len(split_data) - 1]
    input := make([]int, len(split_data))

    fmt.Println(split_data[0], split_data[len(input) - 1])
    fmt.Println(len(input))
    for i, str_num := range split_data {
        input[i], _ = strconv.Atoi(str_num)
    }
    _, res := InversionsCounter(input)
    fmt.Println("Result: ", res)
}

func InversionsCounter(input []int) ([]int, int) {
    inputLength := len(input)
    leftSplits, rightSplits, commonSplits := 0, 0, 0

    if inputLength > 1 {
        splitLength := inputLength/2
        left, right := make([]int, splitLength), make([]int, inputLength - splitLength)
        copy(left, input[:splitLength])
        copy(right, input[splitLength:])

        left, leftSplits = InversionsCounter(left)
        right, rightSplits = InversionsCounter(right)
        l, r := 0, 0
        for k := range input {
            if r == len(right) || l != len(left) && left[l] <= right[r] {
                input[k] = left[l]
                l++
            } else {
                input[k] = right[r]
                r++
                commonSplits += len(left) - l
            }
        }

    }

    return input, leftSplits + rightSplits + commonSplits
}
