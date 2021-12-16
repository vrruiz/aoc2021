#!/Usr/bin/python3

from collections import Counter

def do_step(template, rules):
    new_template = ""
    if len(template) == 1:
        return template
    for i in range(len(template) - 1):
        pair = template[i:i+2]
        new_template += template[i]
        if pair in rules:
            new_template += rules[pair]
            # print(pair, rules[pair], new_template)
    new_template += template[-1]
    return new_template

def run(file_name, steps):
    header = True
    template = ''
    pair_rules = {}
    for line in open(file_name, 'r').readlines():
        line = line.strip()
        if header and line != '':
            template = line
        elif header and line ==  '':
            header = False
        else:
            k, v = line.split(' -> ')
            pair_rules[k] = v
    new_templates  = []
    for i in range(steps):
        template = do_step(template, pair_rules)
        new_templates.append(template)
    return new_templates

def count_elements(template):
    counter = Counter([c for c in template])
    i = 0
    for k, v in counter.items():
        if i == 0 or max_c[1] < v:
            max_c = (k,v)
        if i == 0 or min_c[1] > v:
            min_c = (k, v)
        i += 1
    return min_c, max_c

def test_run():
    result = run('inputest', 4)
    assert len(result) == 4
    assert result[0] == "NCNBCHB"
    assert result[1] == "NBCCNBBBCBHCB"
    assert result[2] == "NBBBCNCCNBBNBNBBCHBHHBCHB"
    assert result[3] == "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"

test_run()
result = run('input', 10)
min_c, max_c = count_elements(result[9])
print(result[9])
print(f"Answer: {max_c[1] - min_c[1]} {min_c} {max_c}")


