 #!/usr/bin/python3

def explore(nodes, visited, path_list, node):
    if node == 'start':
        return path_list
    if node != 'end' and node.islower():
        if visited.count(node) == 2:
            return path_list
        max_visits = max([visited.count(e) for e in visited if e.islower()])
        if visited.count(node) > 0 and max_visits > 1:
            return path_list
    visits = visited + [node]
    if node == 'end':
        path_list.append(visits)
        return path_list
    next_nodes = [n for n in nodes[node]]
    for next_node in next_nodes:
        path_list = explore(nodes, visits, path_list, next_node)
    return path_list
    
def run(file_name):
    maps = [n.split('-') for n in [l.strip() for l in open(file_name, 'r').readlines()]]
    nodes = {}
    for path in maps:
        if not path[0] in nodes:
            nodes[path[0]] = []
        nodes[path[0]].append(path[1])
        if not path[1] in nodes:
            nodes[path[1]] = []
        if path[0] != 'start':
            nodes[path[1]].append(path[0])
    path_list = []
    for node in nodes['start']:
        visited = ['start']
        path_list = explore(nodes, visited, path_list, node)
    return len(path_list)

def main():
    result = run('inputest')
    assert result == 36
    result = run('inputest-2')
    assert result == 103
    result = run('input')
    print(f"Answer: {result}")

if __name__ == '__main__':
    main()
