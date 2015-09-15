import operator
import random


def quick_sort(array, reverse=False, sub_array=None):
    if not sub_array:
        sub_array = (0, len(array) - 1)

    op_l, op_r = operator.lt, operator.gt
    if reverse:
        op_l, op_r = operator.gt, operator.lt

    if (sub_array[1] - sub_array[0]) > 0:
        pivot_index = random.randint(*sub_array)
        pivot_number = array[pivot_index]
        l, r = sub_array[0], sub_array[1]
        while l < r:

            while op_l(array[l], pivot_number):
                l += 1
            while op_r(array[r], pivot_number):
                r -= 1
            if l <= r:
                array[l], array[r] = array[r], array[l]
                l += 1
                r -= 1

        quick_sort(array, reverse=reverse, sub_array=(sub_array[0], r))
        quick_sort(array, reverse=reverse, sub_array=(l, sub_array[1]))
    return array
