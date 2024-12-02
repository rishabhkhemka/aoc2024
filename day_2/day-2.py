import sys
import re
import math
import itertools
import functools
import collections
import requests

day = 2

def read_input() -> str:
    with open("./input_day_{day}.txt".format(day=day), "r") as f:
        input_data = f.read()
    return input_data

def parse_input():
    input_data = read_input().splitlines()
    reports = []
    for line in input_data:
        level = line.split()
        reports.append(list(int(report) for report in level))
    return reports

reports = parse_input()
diff = []
safe = 0
dampsafe = 0

for level in reports:
    level = list(level)
    diffSet = set()
    for (a,b) in zip(level, level[1:]):
        diffSet.add(b-a)

    if diffSet <= {1,2,3} or diffSet <= {-1,-2,-3}:
        safe+=1
    else:
        for i in range(len(level)):
            newLevel = level[:i] + level[i+1:]
            diff = set()
            for (a,b) in zip(newLevel, newLevel[1:]):
                diff.add(b-a)
            if diff <= {1,2,3} or diff <= {-1,-2,-3}:
                dampsafe+=1
                break

print(safe)
print(safe+dampsafe)