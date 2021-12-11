#!/usr/bin/python3
import math

debug = False


def syntax_scoring(file_name):
    pairs = {
        '}' : '{',
        ')' : '(',
        ']' : '[',
        '>' : '<'
    }
    points = {
        ')': 1,
        ']': 2,
        '}': 3,
        '>': 4,
    }
    reverse_pairs = {v : k for k, v in pairs.items()}

    fd = open(file_name, 'r')
    lines = [line.strip() for line in fd.readlines()]
    point_list = []
    for line in lines:
        stack = ""
        adding = ""
        subtotal = 0
        corrupted = False
        for c in line:
            if debug: print(f"syntax_scoring: {c}, {stack}")
            if c in ['{', '[', '(', '<']:
                # Opening chars
                # if debug: print(f"  Add char {c}")
                stack += c
            elif len(stack) > 0:
                # Closing chars
                # if debug: print(f"  Closing char {c} pair:{pairs[c]} stack:{stack[-1]} match:{pairs[c] == stack[-1]}")
                if pairs[c] == stack[-1]:
                    if debug: print(f"  Remove {pairs[c]} from {stack}")
                    stack = stack[:-1]
                else:
                    if debug: print(f"  Cannot remove")
                    corrupted = True
                    break
        if not corrupted and len(stack) > 0:
            completion = [reverse_pairs[c] for c in stack[::-1]]
            for char in completion:
                subtotal = subtotal * 5 + points[char]
            point_list.append(subtotal)
            if debug: print(f"  Incomplete string: line:{line} stack:{stack} sub:{subtotal}")
        elif debug: print(f"Complete {line}")
    point_list.sort()
    assert len(point_list) % 2 == 1
    total = point_list[math.floor(len(point_list) / 2)]
    if debug: print(f"Total: {point_list} {total}")
    return total

def main():
    score = syntax_scoring('inputest')
    print(f"Answer test: {score}")
    assert score == 288957 

    score = syntax_scoring('input')
    print(f"Answer: {score}")

if __name__ == '__main__':
    main()
