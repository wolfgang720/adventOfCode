from functools import reduce


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    code_map = {
        "A": "rock",
        "B": "paper",
        "C": "scissors",
        "X": "rock",
        "Y": "paper",
        "Z": "scissors",
    }

    points_map = {
        "rock": 1,
        "paper": 2,
        "scissors": 3,
        "won": 6,
        "draw": 3,
        "lost": 0,
    }

    my_points = 0
    others_points = 0
    for choices in map(str.split, input_lines):
        my_c = code_map[choices[1]]
        oth_c = code_map[choices[0]]

        # rock, paper, scissors logic -> who wins
        if my_c == oth_c:
            my_points += points_map["draw"] + points_map[my_c]
            others_points += points_map["draw"] + points_map[oth_c]
            continue

        logic_map = {
            ("rock", "rock"): "draw",
            ("paper", "paper"): "draw",
            ("scissors", "scissors"): "draw",
            ("rock", "paper"): "lost",
            ("rock", "scissors"): "won",
            ("paper", "rock"): "won",
            ("paper", "scissors"): "lost",
            ("scissors", "rock"): "lost",
            ("scissors", "paper"): "won",
        }
        my_outcome = logic_map[(my_c, oth_c)]  # type: ignore
        oth_outcome = logic_map[(oth_c, my_c)]  # type: ignore

        my_points += points_map[my_outcome] + points_map[my_c]
        others_points += points_map[oth_outcome] + points_map[oth_c]

    print("my points:", my_points)


if __name__ == "__main__":
    main()
