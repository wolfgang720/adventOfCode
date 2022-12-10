import re


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    signal_strength = 0

    pc = -1
    cycles = 1
    instr_cycles_left = 0
    regx = 1
    regadd = 0
    while True:
        if instr_cycles_left < 1:
            regx += regadd
            pc += 1
            if pc >= len(input_lines):
                break
            instr = input_lines[pc].split(" ")
            if instr[0] == "noop":
                regadd = 0
                instr_cycles_left = 1
            else:
                regadd = int(instr[1])
                instr_cycles_left = 2

        if cycles in (20, 60, 100, 140, 180, 220):
            print(cycles * regx, cycles, regx)
            signal_strength = signal_strength + (cycles * regx)

        instr_cycles_left -= 1
        cycles += 1

    print("Answer:", signal_strength)


if __name__ == "__main__":
    main()
