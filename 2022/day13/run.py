from typing import List, Union


def main():
    input_lines: str = ""
    with open("input.txt", "r") as in_f:
        input_lines = in_f.read()
    #     input_lines = """[1,1,3,1,1]
    # [1,1,5,1,1]

    # [[1],[2,3,4]]
    # [[1],4]

    # [9]
    # [[8,7,6]]

    # [[4,4],4,4]
    # [[4,4],4,4,4]

    # [7,7,7,7]
    # [7,7,7]

    # []
    # [3]

    # [[[]]]
    # [[]]

    # [1,[2,[3,[4,[5,6,7]]]],8,9]
    # [1,[2,[3,[4,[5,6,0]]]],8,9]"""

    def check_order_of_lists(l1, l2) -> Union[bool, None]:
        for i, v1 in enumerate(l1):
            try:
                v2 = l2[i]
            except IndexError:
                return False
            if type(v1) is int and type(v2) is int:
                if v1 == v2:
                    continue
                return v1 < v2
            else:
                _v1 = v1 if type(v1) is list else [v1]
                _v2 = v2 if type(v2) is list else [v2]
                sub_r = check_order_of_lists(_v1, _v2)
                if sub_r is None:
                    continue
                else:
                    return sub_r
        return True

    sum_idx = 0
    for pair_idx, pair in enumerate(input_lines.split("\n\n")):
        p1, p2 = [eval(l) for l in pair.split("\n")]
        is_in_order = check_order_of_lists(p1, p2)
        if is_in_order is None:
            print("None", p1, p2, pair_idx + 1)
        if is_in_order:
            sum_idx += pair_idx + 1

    # sort all packets
    packet_list = [[[2]], [[6]]] + [eval(l) for l in filter(None, input_lines.split("\n"))]
    sorts = 1
    while sorts > 0:
        sorts = 0
        for i in range(len(packet_list) - 1):
            if not check_order_of_lists(packet_list[i], packet_list[i + 1]):
                swp = packet_list[i] if type(packet_list[i]) is int else packet_list[i].copy()
                packet_list[i] = packet_list[i + 1] if type(packet_list[i + 1]) is int else packet_list[i + 1].copy()
                packet_list[i + 1] = swp
                sorts += 1

    print("Answer part 1:", sum_idx)
    print("Answer part 2:", (packet_list.index([[2]]) + 1) * (packet_list.index([[6]]) + 1))


if __name__ == "__main__":
    main()
