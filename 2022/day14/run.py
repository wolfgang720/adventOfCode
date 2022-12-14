from typing import List, Union


def print_field(rocks):
    max_y = max(n[1] for n in rocks)
    max_x = max(n[0] for n in rocks)
    min_x = min(n[0] for n in rocks)
    rows = [[rocks[(x, y)] if (x, y) in rocks else "." for x in range(min_x, max_x + 1)] for y in range(max_y + 1)]
    print("\n".join(("".join(row) for row in rows)))


def main():
    input_lines = ""
    with open("input.txt", "r") as in_f:
        input_lines = in_f.readlines()
    input_lines = """498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9""".split(
        "\n"
    )

    geology = dict()
    for line in input_lines:
        prev_node = None
        for node in line.split(" -> "):
            x, y = map(int, node.split(","))
            if prev_node is None:
                prev_node = (x, y)
                continue
            px, py = prev_node
            if px == x:
                rs, re = sorted((py, y))
                geology.update({(x, i): "#" for i in range(rs, re + 1)})
            elif py == y:
                rs, re = sorted((px, x))
                geology.update({(i, y): "#" for i in range(rs, re + 1)})
            prev_node = (x, y)

    def let_sand_fall(start_pos, max_y):
        sx, sy = start_pos
        while True:
            if (sx, sy + 1) not in geology:
                sy = sy + 1
            elif (sx - 1, sy + 1) not in geology:
                sx, sy = sx - 1, sy + 1
            elif (sx + 1, sy + 1) not in geology:
                sx, sy = sx + 1, sy + 1
            else:
                return (sx, sy)
            if sy + 1 >= max_y + 2:
                return (sx, sy)

    max_y = max((f[1] for f in geology))
    sand_cnt = 0
    while True:
        new_sand_pos = let_sand_fall((500, 0), max_y)
        if new_sand_pos == (500, 0):
            break
        geology[new_sand_pos] = "O"
        sand_cnt += 1

    print_field(geology)

    print("Answer part 1:", sand_cnt + 1)
    print("Answer part 2:")


if __name__ == "__main__":
    main()
