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


def sort(lists):
    # Bubble sort
    for i in range(1, len(lists)):
        for j in range(0, len(lists)-i):
            if comp_list(lists[j], lists[j+1]) == Order.GT:
                lists[j], lists[j+1] = lists[j+1], lists[j]


def solve_part2(input):
    lists = []
    for i in range(0, len(input), 3):
        lists.append(eval(input[i]))
        lists.append(eval(input[i+1]))

    divider1 = [[2]]
    divider2 = [[6]]
    lists.append(divider1)
    lists.append(divider2)

    sort(lists)

    for i, l in enumerate(lists, 1):
        if l == divider1:
            idx1 = i
        elif l == divider2:
            idx2 = i

    print(idx1 * idx2)


def main():
    input = get_input()
    solve_part2(input)


main()
