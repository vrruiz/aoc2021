#!/usr/bin/python3

def print_paper(paper):
    for r in paper:
        print(''.join(r))

def fold(first_half, second_half):
    folded_paper = []
    assert len(first_half) == len(second_half)
    assert len(first_half[0]) == len(second_half[0])
    for y in range(len(first_half)):
        folded_paper.append(first_half[y])
        for x in range(len(first_half[0])):
            # print(y,x, len(second_half[y]), len(second_half))
            if second_half[y][x] == '#':
                folded_paper[y][x] = '#'
    return folded_paper

def fold_y(paper, line):
    first_half = paper[:line]
    second_half = paper[line+1:]
    second_half.reverse()
    return fold(first_half, second_half)

def fold_x(paper, line):
    first_half = [r[:line] for r in paper]
    second_half = [r[line+1:] for r in paper]
    for r in second_half:
        r.reverse()
    folded = fold(first_half, second_half)
    return folded

def count_dots(paper):
    return sum([e.count('#') for e in paper])

def run(file_name, limit=0):
    fold_mode = False
    fold_lines = []
    dots = []
    for line in open(file_name, 'r').readlines():
        line = line.strip()
        if line == '':
            fold_mode = True
            continue
        if fold_mode:
            line = line.replace('fold along ', '')
            axis, n = line.split('=')
            fold_lines.append([axis, int(n)])
        else:
            x,y = line.split(',')
            dots.append([int(x), int(y)])
    size_x = max([d[0] for d in dots]) + 1
    size_y = max([d[1] for d in dots]) + 1
    paper = []
    for y in range(size_y):
        paper.append(['.'] * size_x)
    assert len(paper) == size_y
    assert len(paper[0]) == size_x
    #assert size_x % 2 == 1
    #assert size_y % 2 == 1
    for d in dots:
        paper[d[1]][d[0]] = '#'
    count = []
    if limit == 0:
        limit = len(fold_lines)
    for i in range(0, limit):
        # print(f"-- {fold_lines[i]} lx:{len(paper[0])} ly:{len(paper)}")
        if fold_lines[i][0] == 'x':
            paper = fold_x(paper, fold_lines[i][1])
        else:
            paper = fold_y(paper, fold_lines[i][1])
        count.append(count_dots(paper))
    # print_paper(paper)
    return count

result = run('inputest', 1)
print(f"Answer test: {result[0]}")
assert result[0] == 17

result = run('input', 1)
print(f"Answer: {result[0]}")

