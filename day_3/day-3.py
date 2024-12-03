import sys
import re
import math
import itertools
import functools
import collections
import requests

day = 3

def read_input() -> str:
    with open("./input_day_{day}.txt".format(day=day), "r") as f:
        input_data = f.read()
    return input_data

s = read_input()

# Part 1
pairs = re.findall(r"mul\((\d+),(\d+)\)", s)
print(sum(int(pair[0]) * int(pair[1]) for pair in pairs))

# Part 2
pairs2 = []
dontchunks = re.split(r"don't\(\)", s)
for chunk in dontchunks[1:]: 
    dochunks = re.split(r"do\(\)", chunk)
    if len(dochunks) > 1:
        dochunks = dochunks[1:]
        for dochunk in dochunks:
            dopairs = re.findall(r"mul\((\d+),(\d+)\)", dochunk)
            pairs2.extend(dopairs)
        

dopairs = re.findall(r"mul\((\d+),(\d+)\)", dontchunks[0])
pairs2.extend(dopairs)
print(sum(int(pair[0]) * int(pair[1]) for pair in pairs2))