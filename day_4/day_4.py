import sys
import re
import math
import itertools
import functools
import collections
import requests

day = 4

def read_input() -> str:
    with open("./input_day_{day}.txt".format(day=day), "r") as f:
        input_data = f.read()
    return input_data

def parse_input(s: str):
    # split by newlines and remove empty strings
    return [line for line in s.split("\n") if line]

s = parse_input(read_input())

# Part 1
total = 0

# match any overlapping XMAS in a string forward or backwards
pattern = r"(?=(XMAS|SAMX))"

# left to right and right to left
for line in s:
    total += len(re.findall(pattern ,line))

# top to bottom and bottom to top
for i in range(len(s[0])):
    col = ''.join(line[i] for line in s)
    total += len(re.findall(pattern ,col))

fdiagonals = collections.defaultdict(str)
bdiagonals = collections.defaultdict(str)

for i in range(len(s)): 
    for j in range(len(s[i])):
        fdiagonals[i-j] += s[i][j]
        bdiagonals[i+j] += s[i][j]

for line in fdiagonals.values():
    total += len(re.findall(pattern ,line))

for line in bdiagonals.values():
    total += len(re.findall(pattern ,line))
        
print(total)

# Part 2

total2 = 0
for i in range(len(s) - 2):
    for j in range(len(s[i]) - 2):
        square = [s[i][j:j+3], s[i+1][j:j+3], s[i+2][j:j+3]]
        fd = square[0][0] + square[1][1] + square[2][2]
        bd = square[0][2] + square[1][1] + square[2][0]
        if (fd == "MAS" or fd == "SAM") and (bd == "MAS" or bd == "SAM"):
            total2 += 1

print(total2)
