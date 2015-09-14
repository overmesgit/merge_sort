import operator
import random


def quick_sort(array, reverse=False, sub_array=None):
    if not sub_array:
        sub_array = (0, len(array) - 1)

    op = operator.lt
    if reverse:
        op = operator.gt

    if (sub_array[1] - sub_array[0]) > 0:
        pivot_index = random.randint(*sub_array)
        pivot_number = array[pivot_index]
        l, r = sub_array[0], sub_array[1]
        while l != r:

            while l != pivot_index:
                if op(array[l], pivot_number):
                    l += 1
                else:
                    array[l], array[pivot_index] = array[pivot_index], array[l]
                    pivot_index = l
                    break

            if l == r:
                break

            while r != pivot_index:
                if not op(array[r], pivot_number):
                    r -= 1
                else:
                    array[r], array[pivot_index] = array[pivot_index], array[r]
                    pivot_index = r
                    break

        quick_sort(array, reverse=reverse, sub_array=(sub_array[0], pivot_index - 1))
        quick_sort(array, reverse=reverse, sub_array=(pivot_index + 1, sub_array[1]))
    return array
