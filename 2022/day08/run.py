def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    def find_visible_trees_from_dir(trees):
        return {c for c, t in enumerate(trees) if all(x < t for x in trees[:c]) or all(x < t for x in trees[c + 1 :])}

    visible_trees = set()
    for r, row in enumerate(input_lines):
        # visible from left or right
        news = find_visible_trees_from_dir(list(map(int, row)))
        visible_trees.update({(r, c) for c in news})

    for c in range(len(input_lines[0])):
        # visible from top or bottom
        visible_trees.update({(r, c) for r in find_visible_trees_from_dir([int(line[c]) for line in input_lines])})

    aw1 = len(visible_trees)
    aw2 = 0

    print("Answer part 1:", aw1)
    print("Answer part 2:", aw2)


if __name__ == "__main__":
    main()
