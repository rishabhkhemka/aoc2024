import sys
import re
import math
import itertools
import functools
import collections


day = 1


def read_input(day) -> str:
    with open("./input_day_{day}.txt".format(day=day), "r") as f:
        input_data = f.read()
    return input_data


def parse_input():
    input_data = read_input(day).splitlines()
    left = []
    right = []
    for line in input_data:
        numbers = line.split()
        left.append(int(numbers[0]))
        right.append(int(numbers[1]))
    return left, right


left, right = parse_input()
left.sort()
right.sort()

print(sum(abs(a - b) for a, b in zip(left, right)))

freq = collections.Counter(right)
print(sum(a * freq[a] for a in left if a in freq))
