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
            if r == len(right_part) or l < len(left_part) and not operator(left_part[l], right_part[r]):
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
