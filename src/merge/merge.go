package merge

func Merge(input_array []int, opt ...int) []int {

	input_length := len(input_array)
	if input_length > 1 {
		split_length := int(input_length/2)

		left_array := make([]int, split_length)
		right_array := make([]int, input_length - split_length)

		copy(left_array, input_array[:split_length])
		copy(right_array, input_array[split_length:])

		Merge(left_array)
		Merge(right_array)

		l, r := 0, 0
		for i := range input_array {
			if r == len(right_array) || l != len(left_array) && left_array[l] < right_array[r] {
				input_array[i] = left_array[l]
				l++
			} else {
				input_array[i] = right_array[r]
				r++
			}
		}
	}

	return input_array
}
