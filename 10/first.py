#!/usr/bin/python3

debug = True


def syntax_scoring(file_name):
    pairs = {
        '}' : '{',
        ')' : '(',
        ']' : '[',
        '>' : '<'
    }
    points = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }

    fd = open(file_name, 'r')
    lines = [line.strip() for line in fd.readlines()]
    total = 0
    for line in lines:
        stack = ""
        subtotal = 0
        for c in line:
            if debug: print(f"syntax_scoring: {c}, {stack}")
            if c in ['{', '[', '(', '<']:
                # Opening chars
                if debug: print(f"  Add char {c}")
                stack = stack + c
            elif len(stack) > 0:
                # Closing chars
                if debug: print(f"  Closing char {c} pair:{pairs[c]} stack:{stack[-1]} match:{pairs[c] == stack[-1]}")
                if pairs[c] == stack[-1]:
                    if debug: print(f"  Remove {pairs[c]} from {stack}")
                    stack = stack[:-1]
                else:
                    if debug: print(f"  Expected {c}")
                    subtotal += points[c]
                    break
            elif len(stack) == 0:
                print(f"syntax_scoring: Empty stack")
                continue
        if subtotal == 0:
            print(f"Incomplete {line}")
        else:
            print(f"Corrupted {line} {subtotal}")
            total += subtotal
    return total

def main():
    score = syntax_scoring('inputest')
    print(f"Answer test: {score}")
    assert score == 26397

    score = syntax_scoring('input')
    print(f"Answer: {score}")

if __name__ == '__main__':
    main()
