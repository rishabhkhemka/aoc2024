import sys
import re
import math
import itertools
import functools
from collections import deque
import requests


day = 12

def read_input(fileName) -> str:
    with open(fileName, "r") as f:
        input_data = f.read()
    return input_data

def parse_input(input):
    return [list(line) for line in input.splitlines()]


grid = []

def solve():
    price_p1 = 0
    price_p2 = 0
    for i, row in enumerate(grid):
        for j, plant in enumerate(row):
            if plant != '.' :
                region = []
                p1_helper_dfs((i,j), plant, region)
                perimeter = getperimeter(region)
                # print(plant, region, perimeter)
                price_p1 += (len(region)*len(perimeter))
                
                for cell in region:
                    setPlant(cell, plant)
                
                # print(plant, count_sides(region, plant))
                price_p2 += (len(region)*count_sides(region, plant))

                for cell in region:
                    setPlant(cell, ".")

    print(price_p1)
    print(price_p2)


def getperimeter(region):
    perimeter = []
    for loc in region:
        for neighbor in get_neighbors(loc):
            if neighbor not in region:
                perimeter.append(neighbor)
    
    return perimeter

def get_corner_containing_cells(intersection):
    i, j = intersection
    return [(i,j), (i-1, j), (i, j-1), (i-1, j-1)]


def get_corners(loc):
    x, y = loc
    return[(x, y), (x+1, y), (x, y+1), (x+1, y+1)]

def count_sides(region, plant):
    outsides = set()
    insides = set()
    for col in region:
        for corner in get_corners(col):
            cell_count = 0
            for cell in get_corner_containing_cells(corner):
                if cell not in region:
                    cell_count+=1

            if  cell_count == 3 or cell_count == 1 :
                outsides.add(corner)
            elif cell_count == 2 and all(withinbounds(cell) for cell in get_corner_containing_cells(corner)):
                temp = [getPlant(cell) for cell in get_corner_containing_cells(corner)]
                if (temp[0] == plant and temp[3] == plant) or (temp[1] == plant and temp[2] == plant) :
                    insides.add(corner)
                          
    return len(outsides) + (2*len(insides))

def find_sides(perimeter, sides):
    for loc in perimeter:
        merge = False
        for side in sides:
            if any(neighbor in side for neighbor in get_neighbors(loc)):
                side.add(loc)
                merge = True
                break
                    
        if not merge:
            sides.append({loc})


def p1_helper_dfs(loc, plant, region):
    if not withinbounds(loc):
        return
    
    if getPlant(loc) != plant:
        return
    
    setPlant(loc, '.')
    region.append(loc)

    neighbors = get_neighbors(loc)
    for n in neighbors:
        p1_helper_dfs(n, plant, region)
    return



def getPlant(loc):
    x, y = loc
    return grid[x][y]

def setPlant(loc, plant):
    x, y = loc
    grid[x][y] = plant

def get_neighbors(loc, diagonals=False):
    x, y = loc
    if not diagonals:
        return [(x-1, y),(x+1, y),(x, y-1),(x, y+1)]

def withinbounds(loc):
    x, y = loc
    return 0 <= x < len(grid) and 0 <= y < len(grid[0])


def main():
    global grid
    grid = parse_input(read_input("./input_day_{day}.txt".format(day=day)))
    # grid = parse_input(read_input("sample.txt"))
    solve()


if __name__ == "__main__":
    main()