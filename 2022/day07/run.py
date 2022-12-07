import re
from functools import reduce


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    cur_path = list()

    dirs = {"": (list(), list())}

    # traverse output
    for line in input_lines:
        if m := re.fullmatch(r"\$\scd\s(.+)", line):
            d = m.groups()[0]
            if d == "..":
                cur_path = cur_path[:-1]
            elif d == "/":
                cur_path = list()
            else:
                cur_path.append(d)

        # elif re.fullmatch(r"\$\sls", line):
        #     pass
        elif m := re.fullmatch(r"dir\s(.+)", line):
            d = m.groups()[0]
            path = "/".join([*cur_path, d])
            if path not in dirs:
                dirs[path] = (list(), list())
            dirs["/".join(cur_path)][1].append(path)
        elif m := re.fullmatch(r"(\d+)\s(.+)", line):
            size, name = m.groups()
            path = "/".join(cur_path)
            if path not in dirs:
                dirs[path] = (list(), list())
            dirs["/".join(cur_path)][0].append((int(size), name))

    dir_sizes = dict()
    for path, (files, ds) in dirs.items():

        this_dir_file_sizes = reduce(lambda s, x: s + x[0], files, 0)

        pparts = path.split("/")
        for x in range(len(pparts) + 1):
            ppath = "/".join(pparts[:x])
            if ppath not in dir_sizes:
                dir_sizes[ppath] = 0
            dir_sizes[ppath] += this_dir_file_sizes
        dir_sizes[""] += this_dir_file_sizes
        for d in ds:
            if d not in dir_sizes:
                dir_sizes[d] = 0

    # print([s for d, s in dir_sizes.items() if s <= 100000])

    size1 = reduce(lambda s, x: s + x, (s for s in dir_sizes.values() if s <= 100000), 0)

    print(dirs)
    print("Answer part 1:", size1)
    print("Answer part 2:")


if __name__ == "__main__":
    main()
