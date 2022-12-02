def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    code_map = {
        "A": "rock",
        "B": "paper",
        "C": "scissors",
        "X": "lost",
        "Y": "draw",
        "Z": "won",
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
    for choices in map(str.split, input_lines):
        my_outcome = code_map[choices[1]]
        oth_c = code_map[choices[0]]

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

        for (mine, others), outcome in logic_map.items():
            if outcome == my_outcome and others == oth_c:
                my_points += points_map[my_outcome] + points_map[mine]
                break

    print("my points:", my_points)


if __name__ == "__main__":
    main()
