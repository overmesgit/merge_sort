package merge
import "math/rand"


func Quick(input []int, opt ...int) []int {
    start := 0
    end := len(input) - 1
    if len(opt) > 0 {
        start = opt[0]
        end = opt[1]
    }

    if length := end - start; length > 0 {
        pivot_index := rand.Intn(length) + start
        pivot_number := input[pivot_index]
        l, r := start, end

        for l < r {
            for input[l] < pivot_number {
                l++
            }
            for input[r] > pivot_number {
                r--
            }
            if l <= r {
                input[l], input[r] = input[r], input[l]
                l++
                r--
            }

        }
        Quick(input, start, r)
        Quick(input, l, end)
    }

    return input
}
