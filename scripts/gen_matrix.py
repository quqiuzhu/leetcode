#! /usr/bin/python3
#! coding: utf-8

def gen_byte(s):
    print("[][]byte{")
    for x in s:
        print("[]byte{",end='')
        for c in x:
            print("'%s'" % c,end=',')
        print("},")
    print("}")


def gen_int(s):
    print("[][]int{")
    for x in s:
        print("[]int{", end='')
        for c in x:
            print("%s" % c, end=',')
        print("},")
    print("}")

sudokus = []
sudokus.append([
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
])

sudokus.append([
    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"]
])

sudokus.append([
    ["1", "0", "1", "0", "0"],
    ["1", "0", "1", "1", "1"],
    ["1", "1", "1", "1", "1"],
    ["1", "0", "0", "1", "0"]
])

sudokus.append([
    ['A', 'B', 'C', 'E'],
    ['S', 'F', 'C', 'S'],
    ['A', 'D', 'E', 'E']
])

images = []
images.append([
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
])

images.append([
    [5, 1, 9, 11],
    [2, 4, 8, 10],
    [13, 3, 6, 7],
    [15, 14, 12, 16],
])

images.append([
    [0, 0, 0],
    [0, 1, 0],
    [0, 0, 0]
])

images.append([
    [1, 3, 1],
    [1, 5, 1],
    [4, 2, 1]
])

images.append([
    [0, 1, 2, 0],
    [3, 4, 5, 2],
    [1, 3, 1, 5]
])

images.append([
    [1,   3,  5,  7],
    [10, 11, 16, 20],
    [23, 30, 34, 50]
])

images.append([
    [2],
    [3, 4],
    [6, 5, 7],
    [4, 1, 8, 3]
])

if __name__ == "__main__":
    for s in sudokus:
        gen_byte(s)
    for i in images:
        gen_int(i)

