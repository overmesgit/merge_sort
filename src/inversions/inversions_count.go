package main

import (
    "io/ioutil"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    data, _ := ioutil.ReadFile("src/inversions/IntegerArray.txt")
    string_data := string(data)
    split_data := strings.Split(string_data, "\n")
    input := make([]int, len(split_data))
    for i, str_num := range split_data {
        input[i], _ = strconv.Atoi(str_num)
    }

    fmt.Println("Result123: ")
}

func InversionsCounter(input []int) int {
    return 0
}
