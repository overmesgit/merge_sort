from operator import lt, gt


def bubble_sort(input_array, reverse=False):
    input_length = len(input_array)
    operator = lt if reverse else gt
    for i in range(input_length):

        value = input_array[i]
        index = i
        for j in range(i, input_length):

            if operator(value, input_array[j]):
                value = input_array[j]
                index = j

        input_array[i], input_array[index] = input_array[index], input_array[i]
    return input_array
