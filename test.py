from timeit import timeit
from unittest.case import TestCase
from bubble import bubble_sort
from merge import merge_sort
from random import sample


class SortTest(TestCase):
    def setUp(self):
        def test_case_gen(start=5, factor=125, number=5):
            yield []
            yield [5]
            yield from (sample(range(i), i) for i in range(start, factor*number + start, factor))

        self.case_gen = test_case_gen

    def test_merge_sort(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), merge_sort(test_case))
            self.assertEqual(sorted(test_case, reverse=True), merge_sort(test_case, reverse=True))

    def test_bubble_sort(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), bubble_sort(test_case))
            self.assertEqual(sorted(test_case, reverse=True), bubble_sort(test_case, reverse=True))

    def test_time(self):
        n = 1000
        test_input = sample(range(n), n)
        print('Native sort: {}'.format(timeit(lambda: test_input.copy().sort(), number=10)))
        print('Bubble sort: {}'.format(timeit(lambda: bubble_sort(test_input.copy()), number=10)))
        print('Merge sort: {}'.format(timeit(lambda: merge_sort(test_input.copy()), number=10)))
