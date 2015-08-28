from operator import lt, gt


def merge_sort(input_array, reverse=False):
    input_len = len(input_array)
    operator = lt if reverse else gt

    if input_len > 2:
        slice_index = int(input_len / 2)
        left_part = merge_sort(input_array[:slice_index], reverse)
        right_part = merge_sort(input_array[slice_index:], reverse)

        l, r = 0, 0
        for k in range(input_len):
            if r == len(right_part) or l != len(left_part) and not operator(left_part[l], right_part[r]):
                input_array[k] = left_part[l]
                l += 1
            else:
                input_array[k] = right_part[r]
                r += 1
    else:
        if input_len <= 1:
            return input_array
        else:
            if operator(input_array[0], input_array[1]):
                input_array.reverse()

    return input_array


def merge_sort_not_recursion(input_array, reverse=False):
    input_copy = input_array.copy()
    input_length = len(input_copy)
    operator = lt if reverse else gt
    merge_length = 1

    # optimization for merge_length 1
    # for start in range(0, input_length, 2):
    #     end = start + 1
    #     if end < input_length and operator(input_array[start], input_array[end]):
    #         input_array[start], input_array[end] = input_array[end], input_array[start]
    # merge_length = 2

    while merge_length < input_length:
        for l in range(0, input_length, merge_length*2):
            r = min(l + merge_length, input_length)
            l_last = r
            r_last = min(r + merge_length, input_length)

            for k in range(l, r_last):
                if r == r_last or l != l_last and not operator(input_array[l], input_array[r]):
                    input_copy[k] = input_array[l]
                    l += 1
                else:
                    input_copy[k] = input_array[r]
                    r += 1

        merge_length *= 2
        input_array, input_copy = input_copy, input_array

    return input_array
