#!/usr/bin/python3
from pprint import pprint

debug = False

def minimum_board(board):
    min_board = [[0 for col in range(len(board[0]))] for row in range(len(board))]
    total = 0
    assert len(min_board) == len(board)
    assert len(min_board[0]) == len(board[0])
    for y in range(0, len(board)):
        for x in range(0, len(board[y])):
            if debug: print(f"== x:{x} y:{y}")
            value = board[y][x]
            adyacent = []
            adj = [(0,-1),(-1,0),(0,0),(1,0),(0,1)]
            for xy in adj:
                x_i, y_i = x + xy[0] , y + xy[1]
                if debug: print(f"x_i:{x_i} y_i:{y_i}")
                if x_i >= 0 and x_i < len(board[y]) and \
                   y_i >= 0 and y_i < len(board):
                    if debug: print(f"Add: x:{x_i} y:{y_i} v:{value}")
                    adyacent.append(board[y_i][x_i])
            adyacent.sort()
            if debug: print(adyacent)
            assert len(adyacent) in [3,4,5]
            if debug: print(x, y, adyacent)
            if adyacent[0] == value and adyacent.count(value) == 1:
                if debug: print(f"Low: {value} ({x},{y}) Total: {total} - {adyacent}")
                min_board[y][x] = 1
                total += value + 1
    return min_board, total

def read_board(file_name):
    board_file = open(file_name, 'r')
    board_s = board_file.readlines()
    board = [[int(char) for char in line.strip()] for line in board_s]
    return board

def main():
    board = read_board("inputest")
    min_board, total = minimum_board(board)
    if debug: pprint(board)
    if debug: pprint(min_board)
    assert total == 15

    board = read_board("input")
    min_board, total = minimum_board(board)
    if debug: pprint(board)
    if debug: pprint(min_board)
    print(f"Answer: {total}")

if __name__ == '__main__':
    main()