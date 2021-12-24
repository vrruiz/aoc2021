#!/usr/bin/python3

debug = True

def segments_in(a, b):
    find = [c in b for c in a]
    return find.count(True) == len(a)

def segments_notin(a, b):
    return ''.join([c for c in a if c not in b])

def unsolved(left, solved):
    return [e for e in left if e not in solved.values()]

def solve_in(left, solved, length, contains, number):
    not_solved = unsolved(left, solved)
    for e in not_solved:
            if len(e) == length:
                if contains != None and segments_in(solved[contains], e):
                   solved[number] = e
                elif contains == None:
                   solved[number] = e
    return solved

def sorts(s):
    return ''.join(sorted(s))

def run(file_name):
    known = ((0,1),(1,7),(2,4),(9,8)) # array position, number (1,7,4,8)
    solve = ((6,4,9),(6,7,0),(6,None,6),(5,7,3)) # length, contains, number

    result = []
    for line in open(file_name).readlines():
        left, right = [s.split(' ') for s in line.strip().split(' | ')]  # Split lines
        left = sorted(left, key=len)  # Sort words/numbers by length
        solved = {v : left[k] for k, v in known}  # Get known numbers by length
        for s in solve:
           solved = solve_in(left, solved, s[0], s[1], s[2])
        for e in unsolved(left, solved):
            if len(e) == 5 and len(segments_notin(solved[6], e)) == 1:
                solved[5] = e
            elif len(e) == 5:
                solved[2] = e
        solutions = {sorts(v) : k for k, v in solved.items()}
        digits = 0
        for r in right:
            digits = 10 * digits + solutions[sorts(r)]
        result.append(digits)
    return result

def main():
    result = run('inputest')
    assert result == [5353,8394,9781,1197,9361,4873,8418,4548,1625,8717,4315]
    result = run('input')
    total = sum(result)
    print(f"Answer: {total}")


if __name__ == '__main__':
    main()
