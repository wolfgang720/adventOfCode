from hashlib import new
from typing import List


def get_elevation(s: str) -> int:
    if s == "S":
        return ord("a")
    elif s == "E":
        return ord("z")
    return ord(s)


def find_ways(x, y, height_map):
    elev = get_elevation(height_map[y][x])
    ways = list()

    for xd, yd in [(x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)]:
        if not 0 <= yd < len(height_map) or not 0 <= xd < len(height_map[yd]):
            continue
        elevd = get_elevation(height_map[yd][xd])
        if (elevd - elev) <= 1:
            ways.append(((xd, yd), elev, elevd - elev))
    return ways


def main():
    input_lines: List[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = in_f.readlines()
    #     input_lines = """Sabqponm
    # abcryxxl
    # accszExk
    # acctuvwj
    # abdefghi""".split(
    #         "\n"
    #     )
    start_pos = (-1, -1)
    end_pos = (-1, -1)
    graph = dict()
    for y, row in enumerate(input_lines):
        for x, elev in enumerate(row):
            if elev == "S":
                start_pos = (x, y)
            elif elev == "E":
                end_pos = (x, y)
            graph[(x, y)] = find_ways(x, y, input_lines)

    print(graph[(0, 0)])

    # find shortest path
    visited_nodes = {graph[start_pos][0]}
    paths = [[start_pos]]

    complete_paths = list()

    # path
    # - list of nodes
    # - cost
    visited_cnt = 0
    while len(visited_nodes) != visited_cnt:
        visited_cnt = len(visited_nodes)

        new_paths = list()
        for path in paths:
            for nxt_node in graph[path[-1]]:
                nxt_pos = nxt_node[0]
                if nxt_pos == end_pos:
                    complete_paths.append(path + [nxt_pos])
                if nxt_pos not in visited_nodes:
                    new_paths.append(path + [nxt_pos])
                visited_nodes.add(nxt_pos)
        paths = new_paths

    print([len(cp) - 1 for cp in complete_paths])

    print("Answer part 1:")
    print("Answer part 2:")


if __name__ == "__main__":
    main()
