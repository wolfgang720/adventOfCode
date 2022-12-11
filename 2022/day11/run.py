import sys
from dataclasses import dataclass
from typing import List


@dataclass
class Monkey:
    items: List[int]
    worry_lvl_f: str
    test_div: int
    true_monkey: int
    false_monkey: int
    inspected: List[int]


def compute_worry_lvl(old, fs):
    new = 0
    _locals = locals()
    print(_locals)


new = 0


def main():
    _input: str = ""
    with open("input.txt", "r") as in_f:
        _input = in_f.read()

    # parse monkeys
    monkeys: List[Monkey] = list()
    for m in _input.split("\n\n"):
        m_lines = m.split("\n")
        items = list(map(int, m_lines[1].split("Starting items: ")[-1].split(", ")))
        worry_lvl_f = m_lines[2].split("Operation: ")[-1]  # exec()
        test_div = int(m_lines[3].split("Test: divisible by ")[-1])
        true_monkey = int(m_lines[4].split()[-1])
        false_monkey = int(m_lines[5].split()[-1])
        monkeys.append(Monkey(items, worry_lvl_f, test_div, true_monkey, false_monkey, list()))

    for _ in range(20):
        for monkey in monkeys:
            for old in monkey.items:
                # global new
                _locals = locals()
                monkey.inspected.append(old)

                exec("global new; " + monkey.worry_lvl_f, globals(), _locals)  # new = f(old)
                new_worry_lvl = new // 3
                # print(old, monkey.worry_lvl_f, new, new_worry_lvl)

                # test
                if new_worry_lvl % monkey.test_div == 0:
                    monkeys[monkey.true_monkey].items.append(new_worry_lvl)
                else:
                    monkeys[monkey.false_monkey].items.append(new_worry_lvl)
            monkey.items = list()
        # print(monkeys)
        # sys.exit()

    monkeys.sort(key=lambda m: len(m.inspected))
    print(list(len(m.inspected) for m in monkeys))

    print("Answer 1:", len(monkeys[-1].inspected) * len(monkeys[-2].inspected))


if __name__ == "__main__":
    main()
