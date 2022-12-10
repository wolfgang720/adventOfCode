import re


def main():
    input_lines: list[str] = list()
    with open("input.txt", "r") as in_f:
        input_lines = list(map(str.strip, in_f.readlines()))

    signal_strength = 0

    BUFFER_LEN = 40 * 6
    crt_pos = 0
    crt_buffer = "." * BUFFER_LEN

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
            # print(cycles * regx, cycles, regx)
            signal_strength = signal_strength + (cycles * regx)

        # pixel visible?
        crt_pos = cycles % BUFFER_LEN - 1
        print(crt_pos, regx)
        crt_buffer = crt_buffer[:crt_pos] + ("#" if abs(regx - crt_pos % 40) <= 1 else ".") + crt_buffer[crt_pos + 1 :]

        instr_cycles_left -= 1
        cycles += 1

    print("Answer 1:", signal_strength)
    print("Answer 2:", "\n" + "\n".join(crt_buffer[x * 40 : (x + 1) * 40] for x in range(6)))


if __name__ == "__main__":
    main()
