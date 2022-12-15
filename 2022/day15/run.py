import re
from itertools import combinations, permutations
from typing import List, Union


def main():
    input_lines: List[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = in_f.readlines()
    #     input_lines = """Sensor at x=2, y=18: closest beacon is at x=-2, y=15
    # Sensor at x=9, y=16: closest beacon is at x=10, y=16
    # Sensor at x=13, y=2: closest beacon is at x=15, y=3
    # Sensor at x=12, y=14: closest beacon is at x=10, y=16
    # Sensor at x=10, y=20: closest beacon is at x=10, y=16
    # Sensor at x=14, y=17: closest beacon is at x=10, y=16
    # Sensor at x=8, y=7: closest beacon is at x=2, y=10
    # Sensor at x=2, y=0: closest beacon is at x=2, y=10
    # Sensor at x=0, y=11: closest beacon is at x=2, y=10
    # Sensor at x=20, y=14: closest beacon is at x=25, y=17
    # Sensor at x=17, y=20: closest beacon is at x=21, y=22
    # Sensor at x=16, y=7: closest beacon is at x=15, y=3
    # Sensor at x=14, y=3: closest beacon is at x=15, y=3
    # Sensor at x=20, y=1: closest beacon is at x=15, y=3""".split(
    #         "\n"
    #     )

    def calc_distance(sx, sy, bx, by):
        return abs(sx - bx) + abs(sy - by)

    def calc_f(x, y):
        return x * 4000000 + y

    s_and_bs = list()
    beacons = set()
    for line in input_lines:
        m = re.fullmatch(r"Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)", line.strip())
        sx, sy, bx, by = map(int, m.groups())

        distance = calc_distance(sx, sy, bx, by)
        s_and_bs.append((sx, sy, bx, by, distance))
        beacons.add((bx, by))

    min_bx = min(sx - d for sx, sy, bx, by, d in s_and_bs)
    max_bx = max(sx + d for sx, sy, bx, by, d in s_and_bs)
    # print(min_bx, max_bx)
    # print(beacons)
    y = 2000000
    positions = list()
    # for x in range(0, 4000000):
    #     for y in range(0, 4000000):
    #         if not any(calc_distance(sx, sy, x, y) <= distance for sx, sy, _, _, distance in s_and_bs):
    #             positions.append((x, y, calc_f(x, y)))
    #             print(x, y, calc_f(x, y))

    # intersections = find_frame_of_sensor(s_and_bs[0][0], s_and_bs[0][1], s_and_bs[0][4])
    # intersections = set()
    # for s1, s2 in combinations(s_and_bs, 2):
    #     x1, y1, _, _, d1 = s1
    #     x2, y2, _, _, d2 = s2
    #     if calc_distance(x1, y1, x2, y2) <= d1 + d2 + 2:
    #         print("nodes:", (x1, y1, d1, x2, y2, d2))
    #         print(find_sensor_frame_intersection(x1, y1, d1, x2, y2, d2))

    found = list()
    for y in range(0, 4_000_001):
        if y % 10000 == 0:
            print(y)
        x = 0
        while x <= 4_000_000:
            for xs, ys, _, _, d in s_and_bs:
                if calc_distance(x, y, xs, ys) <= d:
                    # skip to x end of sensor range
                    xd = d - abs(y - ys)
                    x = xs + xd + 1
                    break
            else:
                found.append((x, y, calc_f(x, y)))
                x += 1
    print(found)

    print("Answer:", positions)
    # 5567804
    # 5567803 (-888612 4807592)
    # 6275922 (-1710661 5722461)


if __name__ == "__main__":
    main()
