from functools import reduce


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    in_both_comps = list()
    for rucksack in map(str.strip, input_lines):
        l = int(len(rucksack) / 2)
        comp1, comp2 = rucksack[l:], rucksack[:l]
        in_both_comps.extend(list({i for i in rucksack if i in comp2 and i in comp1}))

    def get_prio_for_item_reducer(r, i):
        if i.islower():
            return r + ord(i) - ord("a") + 1
        else:
            return r + ord(i) - ord("A") + 27

    sum_prios = reduce(get_prio_for_item_reducer, in_both_comps, 0)

    print("Answer part 1:", sum_prios)
    print("Answer part 2:")


if __name__ == "__main__":
    main()
