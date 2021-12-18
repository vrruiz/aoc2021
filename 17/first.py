#!/usr/bin/python3

def target_in(box, position):
    xs, xe = min(box[0]), max(box[0])
    ys, ye = min(box[1]), max(box[1])
    if position[0] >= xs and position[0] <= xe and position[1] >= ys and position[1] <= ye:
        return True
    return False

def do_step(pos, v):
    next_pos = (pos[0] + v[0], pos[1] + v[1])
    next_v = list(v)
    if next_v[0] > 0:
        next_v[0] -= 1
    elif next_v[0] < 0:
        next_v[0] += 1
    next_v[1] -= 1
    return next_pos, tuple(next_v)

def loop(target, v):
    pos = (0,0)
    max_y = 0
    reached = False
    while pos[1] >= min(target[1]):
        pos, v = do_step(pos, v)
        if pos[1] > max_y:
            max_y = pos[1]
        t_in = target_in(target,pos)
        if t_in:
            reached = True
            break
    return reached, max_y

def run(file_name):
    line = open(file_name, "r").readline().strip()
    line = line.replace("target area: ", "")
    tx, ty = line.split(", ")
    target_x = [int(i) for i in tx.replace("x=", "").split("..")]
    target_y = [int(i) for i in ty.replace("y=", "").split("..")]
    target = (target_x,target_y)

    max_y = -1
    v_max = tuple()
    for y in range(max(target_x)*2):
        for x in range(max(target_x)*2):
            reached, m_y = loop(target, (x,y))
            if reached and m_y > max_y:
                max_y = m_y
                v_max = (x,y)
    return max_y, v_max

def run_tests():
    assert target_in(((20,30),(-10,-5)), (28,-7))
    assert target_in(((20,30),(-10,-5)), (28,-4)) == False

run_tests()
y, v = run('inputest')
assert y == 45 and v == (6,9)

y, v = run('input')
print(f"Answer: y:{y} v:{v}")

