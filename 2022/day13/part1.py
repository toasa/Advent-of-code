from enum import Enum


class Order(Enum):
    LT = 1
    EQ = 2
    GT = 3


def get_input():
    f = open('input')
    lines = f.readlines()
    lines = list(map(lambda s: s.rstrip(), lines))
    f.close()
    return lines


def comp_list(list1, list2) -> Order:
    for e1, e2 in zip(list1, list2):
        if isinstance(e1, int) and isinstance(e2, int):
            if e1 < e2:
                return Order.LT
            elif e1 > e2:
                return Order.GT
            continue

        if isinstance(e1, list) and isinstance(e2, list):
            o = comp_list(e1, e2)
            if o != Order.EQ:
                return o
            continue

        if isinstance(e1, int):
            o = comp_list([e1], e2)
        else:
            o = comp_list(e1, [e2])
        if o != Order.EQ:
            return o

    if len(list1) < len(list2):
        return Order.LT
    elif len(list1) > len(list2):
        return Order.GT

    return Order.EQ


def is_right_order(list1, list2):
    return comp_list(list1, list2) == Order.LT


def solve_part1(input):
    count = 0
    idx = 1
    for i in range(0, len(input), 3):
        if is_right_order(eval(input[i]), eval(input[i+1])):
            count += idx
        idx += 1

    print(count)


def main():
    input = get_input()
    solve_part1(input)


main()
