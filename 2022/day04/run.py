def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    def range_from_pair(p):
        return range(*map(int, p.split("-")))

    pairs = [(range_from_pair(rs1), range_from_pair(rs2)) for rs1, rs2 in map(lambda s: s.split(","), input_lines)]

    # find pairs where one range encloses the other
    pairs_w_enclosed_rngs = list()
    for r1, r2 in pairs:
        # r1 must be longer or equal
        if len(r1) < len(r2):
            r1, r2 = r2, r1
        if r1.start <= r2.start and r1.stop >= r2.stop:
            pairs_w_enclosed_rngs.append((r1, r2))

    print("Answer part 1:", len(pairs_w_enclosed_rngs))
    print("Answer part 2:")


if __name__ == "__main__":
    main()
