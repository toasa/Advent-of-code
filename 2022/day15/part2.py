import math


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


def euclidean_dist(p):
    return math.sqrt(p[0]**2 + p[1]**2)


def solve_part2(input):
    sensors, beacons = get_sensors_and_beacons(input)

    boundary_points = []

    for s, b in zip(sensors, beacons):
        md = manhattan_dist(s, b)
        bs = [
            (s[0] + md, s[1]),  # down
            (s[0] - md, s[1]),  # up
            (s[0], s[1] + md),  # right
            (s[0], s[1] - md),  # left
        ]
        for p in bs:
            if 0 <= p[0] <= 4000000 and 0 <= p[1] <= 4000000:
                boundary_points.append(p)

    boundary_points.sort(key=lambda x: (x[0], x[1]))

    for p in boundary_points:
        print(p)


def main():
    input = get_input()
    solve_part2(input)


main()
