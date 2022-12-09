import re


def move_head(pos, direction):
    x, y = pos
    if direction == "U":
        return (x, y + 1)
    if direction == "D":
        return (x, y - 1)
    if direction == "L":
        return (x - 1, y)
    if direction == "R":
        return (x + 1, y)


def move_tail(head_pos, tail_pos):
    hx, hy = head_pos
    tx, ty = tail_pos

    if abs(hx - tx) <= 1 and abs(hy - ty) <= 1:
        return (tx, ty)

    if hx == tx:
        return (tx, ty + 1 if hy > ty else ty - 1)

    if hy == ty:
        return (tx + 1 if hx > tx else tx - 1, ty)

    return (tx + 1 if hx > tx else tx - 1, ty + 1 if hy > ty else ty - 1)


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    start_pos = (0, 0)
    head_pos = start_pos
    all_tail_pos: list[tuple[int, int]] = [start_pos]

    for line in input_lines:
        m = re.fullmatch(r"([UDLR]+)\s(\d+)", line)
        direction, distance = m.groups()
        for _ in range(int(distance)):
            head_pos = move_head(head_pos, direction)
            all_tail_pos.append(move_tail(head_pos, all_tail_pos[-1]))

    aw1 = len(set(all_tail_pos))
    aw2 = ""

    print("Answer part 1:", aw1)
    print("Answer part 2:", aw2)


if __name__ == "__main__":
    main()
