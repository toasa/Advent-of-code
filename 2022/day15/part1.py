def get_input():
    f = open('input')
    lines = f.readlines()
    lines = list(map(lambda s: s.rstrip(), lines))
    f.close()
    return lines


def get_sensors_and_beacons(input):
    sensors = []
    beacons = []
    for line in input:
        i = line.find("=")
        j = line.find(",")
        x = int(line[i+1:j])

        line = line[j:]

        i = line.find("=")
        j = line.find(":")
        y = int(line[i+1:j])
        sensors.append((y, x))

        line = line[j:]

        i = line.find("=")
        j = line.find(",")
        x = int(line[i+1:j])
        y = int(line[j+4:])

        beacons.append((y, x))

    return sensors, beacons


def manhattan_dist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


class Row:
    MAX_LEN = 10000000
    CHECK_ROW = 2000000

    def __init__(self):
        self.pos = [False for _ in range(self.MAX_LEN)]
        self.neg = [False for _ in range(self.MAX_LEN)]

    def set(self, x):
        if x >= 0:
            self.pos[x] = True
        else:
            self.neg[-x] = True


def solve_part1(input):
    sensors, beacons = get_sensors_and_beacons(input)
    row = Row()

    for l in zip(sensors, beacons):
        s = l[0]
        b = l[1]

        vert_p_with_check_row = (row.CHECK_ROW, s[1])

        md = manhattan_dist(s, b)
        if md < manhattan_dist(s, vert_p_with_check_row):
            continue

        for i in range(-row.MAX_LEN + 1, row.MAX_LEN):
            if manhattan_dist(s, (row.CHECK_ROW, i)) <= md:
                row.set(i)

    # Subtract 1 to remove beacon
    res = row.pos.count(True) + row.neg.count(True) - 1
    print(res)


def main():
    input = get_input()
    solve_part1(input)


main()
