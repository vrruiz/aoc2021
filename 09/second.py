#!/usr/bin/python3
from pprint import pprint
from first import minimum_board, read_board

debug = False

def explore_position(board, walked_board, pos):
    adyacents = [[0,-1],[-1,0],[1,0],[0,1]]
    points = [[pos[0] + a[0], pos[1] + a[1]] for a in adyacents if a[0] + pos[0] >= 0 and a[0] + pos[0] < len(walked_board[0]) and a[1] + pos[1] >= 0 and a[1] + pos[1] < len(walked_board)]
    for point in points:
        x, y = point
        if board[y][x] != 9 and walked_board[y][x] != 1:
            walked_board[y][x] = 1
            explore_position(board, walked_board, [x,y])

def walk_board(board, positions):
    walked_board = [[0 for x in range(0, len(board[0]))] for y in range(0, len(board))]
    for pos in positions:
        x, y = pos
        walked_board[y][x] = 1
        explore_position(board, walked_board, [x,y])
    size = 0
    for row in walked_board:
        for col in row:
            if col == 1:
                size += 1
    if debug: pprint(walked_board)
    return size

board = read_board("input")
min_board, total = minimum_board(board)
if debug: pprint(board)
if debug: pprint(min_board)

sizes = []
for y in range(0, len(board)):
    for x in range(0, len(board[y])):
        if min_board[y][x] == 1:
            sizes.append(walk_board(board, [[x,y]]))
sizes.sort()
largest = sizes[-3:]
total = 1
for s in largest:
    total *= s
if debug: print(sizes, largest)
print(f"Answer: {total}")