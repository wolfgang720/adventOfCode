from functools import reduce


def aoc_22_1_1():
    input_lines: list[str] = list()
    with open("./input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))
    if input_lines[-1] != "":
        input_lines.append("")

    elves: list[tuple[list[int], int]] = list()
    current_elf: list[int] = list()
    for line in input_lines:

        if line == "":
            elves.append((current_elf, reduce(lambda x, y: x + y, current_elf, 0)))
            current_elf = list()
            continue
        calories = int(line, 10)
        current_elf.append(calories)

    elves.sort(key=lambda x: x[1])

    # calories of elf carrying most
    print(elves[-1][1])

    # calories of top 3 elves carrying most
    print(reduce(lambda x, y: x + y[1], elves[-3:], 0))


if __name__ == "__main__":
    aoc_22_1_1()
