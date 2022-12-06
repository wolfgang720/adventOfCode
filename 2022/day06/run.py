def main():
    input_str: str = "nppdvjthqldpwncqszvftbrmjlhg"
    with open("input.txt", "r") as in_f:
        input_str = in_f.read()

    MARKER_LEN = 14
    marker = input_str[:MARKER_LEN]
    marker_idxs = list()
    for i, c in enumerate(input_str[MARKER_LEN:]):
        marker = marker[1:] + c
        if any(marker.count(cm) > 1 for cm in marker[1:]):
            continue
        else:
            marker_idxs.append(i + MARKER_LEN + 1)
            break

    print("Answer part 1:", marker_idxs[0])
    print("Answer part 2:")


if __name__ == "__main__":
    main()
