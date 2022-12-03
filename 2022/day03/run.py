from functools import reduce


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    in_both_comps = list()
    group_badges = list()
    group_rucksacks = list()
    for n, rucksack in enumerate(input_lines):
        l = int(len(rucksack) / 2)
        comp1, comp2 = rucksack[l:], rucksack[:l]
        in_both_comps.extend(list({i for i in rucksack if i in comp2 and i in comp1}))

        group_rucksacks.append(rucksack)
        if len(group_rucksacks) % 3 == 0:
            group_badges.append(
                [i for i in group_rucksacks[0] if i in group_rucksacks[1] and i in group_rucksacks[2]][0]
            )
            group_rucksacks = list()

    def get_prio_for_item_reducer(r, i):
        if i.islower():
            return r + ord(i) - ord("a") + 1
        else:
            return r + ord(i) - ord("A") + 27

    sum_prios = reduce(get_prio_for_item_reducer, in_both_comps, 0)
    sum_prios_badges = reduce(get_prio_for_item_reducer, group_badges, 0)

    print("Answer part 1:", sum_prios)
    print("Answer part 2:", sum_prios_badges)


if __name__ == "__main__":
    main()
