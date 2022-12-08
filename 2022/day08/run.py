from functools import reduce


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    def find_visible_trees_from_dir(trees):
        return {c for c, t in enumerate(trees) if all(x < t for x in trees[:c]) or all(x < t for x in trees[c + 1 :])}

    def find_scenic_score_for_tree_in_line(t_idx, trees):

        dist = 0
        for t in trees[t_idx + 1 :]:
            dist += 1
            if t >= trees[t_idx]:
                break
        score = dist
        dist = 0
        for t in reversed(trees[:t_idx]):
            dist += 1
            if t >= trees[t_idx]:
                break
        return score * dist

    visible_trees = set()
    scenic_scores = dict()
    for r, row in enumerate(input_lines):
        _row = list(map(int, row))
        # visible from left or right
        visible_trees.update({(r, c) for c in find_visible_trees_from_dir(_row)})
        for c in range(len(_row)):
            scenic_scores[(r, c)] = find_scenic_score_for_tree_in_line(c, _row)

    for c in range(len(input_lines[0])):
        # visible from top or bottom
        _col = [int(line[c]) for line in input_lines]
        visible_trees.update({(r, c) for r in find_visible_trees_from_dir(_col)})
        for r in range(len(_col)):
            scenic_scores[(r, c)] *= find_scenic_score_for_tree_in_line(r, _col)

    aw1 = len(visible_trees)
    aw2 = max(scenic_scores.values())

    print("Answer part 1:", aw1)
    print("Answer part 2:", aw2)


if __name__ == "__main__":
    main()
