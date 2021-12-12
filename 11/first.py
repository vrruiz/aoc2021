#!/usr/bin/python3
from pprint import pprint

debug = False

def zeros_flashes(size_x, size_y):
    zeros = []
    for y in range(size_y):
        zeros.append([0] * size_x)
    return zeros

def do_flash(board, flashes, x, y, count):
    board[y][x] = 0
    flashes[y][x] = 1
    count += 1
    if debug: print(f"Do flash x:{x} y:{y} f:{flashes[y][x]}")
    ady = ((-1,-1),(0,-1),(1,-1),(-1,0),(1,0),(-1,1),(0,1),(1,1))
    for a in ady:
        if debug: pprint(flashes)
        x_i, y_i = x + a[0], y + a[1]
        if debug: print(f" x:{x_i} y:{y_i} lx:{len(board[0])} ly:{len(board)} f:{flashes[y_i][x_i]}")
        if x_i >= 0 and x_i < len(board[0]) and y_i >= 0 and y_i < len(board) and flashes[y_i][x_i] == 0:
            if debug: print(f"  +x:{x_i} y:{y_i}")
            board[y_i][x_i] += 1
            if board[y_i][x_i] > 9:
                if debug: print(f"  Found flash x:{x_i} y:{y_i}")
                board, flashes, count = do_flash(board, flashes, x_i, y_i, count)
    return board, flashes, count


def do_step(board):
    flashes = zeros_flashes(len(board[0]), len(board))
    assert len(flashes) == len(board)
    assert len(flashes[0]) == len(board[0])
    if debug: print("Do step")
    count = 0
    for y in range(len(board)):
        for x in range(len(board[0])):
            if flashes[y][x] == 0:
                board[y][x] += 1
                if board[y][x] > 9:
                    board, flashes, count = do_flash(board, flashes, x, y, count)
    return board, count

def run(file_name):
    fd = open(file_name, 'r')
    board = [[int(c) for c in l.strip()] for l in fd.readlines()]
    steps = 0
    total = 0
    for i in range(100):
        board, count = do_step(board)
        total += count
        if debug: pprint(board)
    print(f"Answer {file_name}: {total}")
    return total

def main():
    total = run('inputest')
    assert total == 1656
    total = run('input')

if __name__ == '__main__':
    main()
