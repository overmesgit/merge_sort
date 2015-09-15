from timeit import timeit
from unittest.case import TestCase
from python.bubble import bubble_sort
from python.merge import merge_sort, merge_sort_not_recursion
from python.quick import quick_sort
from random import sample


class SortTest(TestCase):
    def setUp(self):
        def test_case_gen(start=5, factor=125, number=10):
            yield []
            yield [5]
            yield from (sample(range(i), i) for i in range(start, factor * number + start, factor))

        self.case_gen = test_case_gen

    def test_merge_sort(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), merge_sort(test_case))
            self.assertEqual(sorted(test_case, reverse=True), merge_sort(test_case, reverse=True))

    def test_bubble_sort(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), bubble_sort(test_case))
            self.assertEqual(sorted(test_case, reverse=True), bubble_sort(test_case, reverse=True))

    def test_quick_sort(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), quick_sort(test_case))
            self.assertEqual(sorted(test_case, reverse=True), quick_sort(test_case, reverse=True))

    def test_merge_sort_not_recursion(self):
        for test_case in self.case_gen():
            self.assertEqual(sorted(test_case), merge_sort_not_recursion(test_case))
            self.assertEqual(sorted(test_case, reverse=True), merge_sort_not_recursion(test_case, reverse=True))

    def test_time(self):
        n = 50000
        test_input = sample(range(n), n)
        test_functions = (merge_sort, merge_sort_not_recursion, quick_sort)
        for func in test_functions:
            running_time = timeit(lambda: func(test_input.copy()), number=1)
            print('{}: {}'.format(func.__name__, running_time))
