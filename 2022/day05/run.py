import re
from distutils.command.build_scripts import first_line_re


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    # read stacks
    first_moves_line = 0
    stacks = [list() for _ in range(9)]
    for idx, line in enumerate(input_lines):
        if line == "":
            first_moves_line = idx + 1
            break

        m = re.findall(r"\[(\w)\]\s?|(\s{4})", line)
        for col_idx, (crate, _) in enumerate(m):
            _crate = crate.strip()
            if _crate == "":
                continue
            stacks[col_idx].insert(0, _crate)
    print(stacks)

    for line in input_lines[first_moves_line:]:
        m = re.fullmatch(r"move (\d+) from (\d+) to (\d+)", line)
        n, fr, to = map(int, m.groups())

        moved = stacks[fr - 1][-n:]
        stacks[fr - 1] = stacks[fr - 1][:-n]
        stacks[to - 1].extend(reversed(moved))

    print("Answer part 1:", "".join((s[-1] for s in stacks)))
    print("Answer part 2:")


if __name__ == "__main__":
    main()
