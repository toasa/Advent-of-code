def get_input():
    f = open('input')
    lines = f.readlines()
    lines = list(map(lambda s: s.rstrip(), lines))
    f.close()
    return lines


LEN = 550
SRC_OF_SAND = (0, 500)


def print_cave(cave):
    for i in range(10):
        # print("".join(cave[i]))
        print("".join(cave[i][494:504]))


def build_cave(input):
    cave = [["." for _ in range(LEN)] for _ in range(LEN)]

    rock_paths = []
    for line in input:
        paths = []

        while line.find(",") != -1:

            i = line.find(",")
            j = line.find(" ")
            if j == -1:
                j = len(line)
            x = int(line[:i])
            y = int(line[i+1:j])
            paths.append((y, x))

            line = line[j+4:]
        rock_paths.append(paths)

    for paths in rock_paths:
        for n in range(len(paths)-1):
            p1 = paths[n]
            p2 = paths[n+1]

            for i in range(p1[0], p2[0]+1):
                cave[i][p1[1]] = "#"
            for i in range(p2[0], p1[0]+1):
                cave[i][p1[1]] = "#"
            for i in range(p1[1], p2[1]+1):
                cave[p1[0]][i] = "#"
            for i in range(p2[1], p1[1]+1):
                cave[p1[0]][i] = "#"

    return cave


def get_square(cave, pos):
    return cave[pos[0]][pos[1]]


def set_square(cave, pos, tile):
    cave[pos[0]][pos[1]] = tile


def pour_one_sand(cave):
    pos = SRC_OF_SAND
    while True:
        if pos[0] + 1 >= LEN:
            return False

        # Down
        if get_square(cave, (pos[0]+1, pos[1])) == ".":
            pos = (pos[0]+1, pos[1])
            continue

        # Down-left
        if get_square(cave, (pos[0]+1, pos[1]-1)) == ".":
            pos = (pos[0]+1, pos[1]-1)
            continue
        # Down-right
        if get_square(cave, (pos[0]+1, pos[1]+1)) == ".":
            pos = (pos[0]+1, pos[1]+1)
            continue

        break

    set_square(cave, pos, "o")
    return True


def pour_sand(cave):
    count = 0
    while pour_one_sand(cave):
        count += 1

    return count


def solve_part1(input):
    cave = build_cave(input)

    count = pour_sand(cave)

    print(count)


def main():
    input = get_input()
    solve_part1(input)


main()
