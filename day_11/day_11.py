import sys
import re
import math
import itertools
import functools
import collections
import requests


day = 11

def read_input() -> str:
    with open("./input_day_{day}.txt".format(day=day), "r") as f:
        input_data = f.read()
    return input_data

def parse_input(input):
    return [int(num) for num in read_input().split()]


def blink(num) -> list:
    if num == 0:
        return [1]
    elif len(str(num)) % 2 == 0:
        return [num // 10**(len(str(num)) // 2), num % 10**(len(str(num)) // 2)]
    else:
        return [num * 2024]

@functools.cache
def blink_times(num, times) -> int:
    if times == 1:
        blinked = blink(num)
        return len(blinked)
    
    blinked = blink(num)
    return sum(blink_times(n, times-1) for n in blinked)

def solve(numbers, times): 
    return sum(blink_times(num, times) for num in numbers)


# print(solve_p1((125, 17), 25))
print(solve(parse_input(read_input()), 25))
print(solve(parse_input(read_input()), 75))


