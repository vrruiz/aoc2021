#!/usr/bin/python3

def print_paper(dots, size):
    paper = []
    for i in range(size[1]):
        paper.append(['.'] * size[0])
    for dot in dots:
        paper[dot[1]][dot[0]] = '#'
    for row in range(len(paper)):
        print(''.join(paper[row]))

def fold(dots, size, fold_line, axis):
    new_dots = {}
    diff = 1
    if size[axis] % 2 == 0:
        diff = 0
    for k,v in dots.items():
        if k[axis] > fold_line:
            if axis == 0:
                new_pos = (size[axis] - diff - k[axis], k[1])
            else:
                new_pos = (k[0], size[axis] - diff - k[axis])
            if new_pos not in new_dots:
                new_dots[new_pos] = 1
        if k[axis] < fold_line and k not in new_dots:
            new_dots[k] = 1
    new_size = ()
    if axis == 0:
        new_size = (fold_line, size[1])
    else:
        new_size = (size[0], fold_line)
    return new_dots, new_size

def run(file_name, limit=0):
    fold_mode = False
    fold_lines = []
    dots = {}
    for line in open(file_name, 'r').readlines():
        line = line.strip()
        if line == '':
            fold_mode = True
            continue
        if fold_mode:
            line = line.replace('fold along ', '')
            axis, l = line.split('=')
            if axis == 'x':
                axis = 0
            else:
                axis = 1
            fold_lines.append((axis, int(l)))
        else:
            x,y = line.split(',')
            dots[(int(x), int(y))] = 1
    size_x = max([d[0] for d in dots]) + 1
    size_y = max([d[1] for d in dots]) + 1
    size = (size_x, size_y)
    for f in fold_lines:
        dots, size = fold(dots, size, f[1], f[0])
    return dots, size

def test_fold():
    dots = {(0,0) : 1, (1,0) : 1, (2,0) : 1}
    folded_dots, folded_size = fold(dots, (3, 1), 1, 0)
    assert len(folded_dots.items()) == 1
    assert (0,0) in dots

    dots = {(0,0) : 1, (2,0) : 1, (3,0) : 1}
    folded_dots, folded_size = fold(dots, (4,1), 2, 0)
    assert len(folded_dots.items()) == 2
    assert (0,0) in folded_dots
    assert (1,0) in folded_dots
    assert folded_size == (2,1)

    dots = {(0,0) : 1, (0,1) : 1, (0,2) : 1}
    folded_dots, folded_size = fold(dots, (1,3), 1, 1)
    assert len(folded_dots.items()) == 1
    assert (0,0) in dots

    dots = {(0,0) : 1, (0,2) : 1, (0,3) : 1}
    folded_dots, folded_size = fold(dots, (1,4), 2, 1)
    assert len(folded_dots.items()) == 2
    assert (0,0) in folded_dots
    assert (0,1) in folded_dots
    assert folded_size == (1,2)

test_fold()
dots, size = run('input')
print_paper(dots, size)
