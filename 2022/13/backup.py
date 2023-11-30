import ast

ordered_pairs = []
not_ordered_pairs = []
is_in_list = False


def print_packet_pair(counter, lvl, p1, p2, spacing="", header=True):
    if header:
        print(f"{spacing}- Compare header {p1} vs {p2}")
    lenP1 = len(p1)
    lenP2 = len(p2)
    global is_in_list
    
    lvl = 0
    ordered = False
    excluded_lvl = 0
    is_from_recursion_order = False
    finished_elements = lenP1 == 0 or lenP2 == 0
    if not finished_elements:
        for i in range(min(lenP1, lenP2)):
            if not (counter in ordered_pairs or counter in not_ordered_pairs):
                print(f"{spacing}  - Compare {p1[i]} vs {p2[i]}")
                if isinstance(p1[i], list):
                    if isinstance(p2[i], list):
                        excluded_lvl = lvl + 1
                        
                        is_from_recursion_order = True
                        print_packet_pair(
                            counter,
                            lvl + 1,
                            p1[i],
                            p2[i],
                            spacing=spacing + "  ",
                            header=False,
                        )
                    else:
                        newp2 = [p2[i]]
                        print(
                            f"{spacing}   - Mixed types; convert right to {newp2} and retry comparison"
                        )
                        is_from_recursion_order = True
                        excluded_lvl = lvl + 1
                        print_packet_pair(
                            counter,
                            lvl + 1,
                            p1[i],
                            newp2,
                            spacing=spacing + " ",
                            header=True,
                        )
                    
                    is_in_list = True
                elif isinstance(p2[i], list):
                    newp1 = [p1[i]]
                    print(
                        f"{spacing}   - Mixed types; convert left to {newp1} and retry comparison"
                    )
                    is_from_recursion_order = True
                    excluded_lvl = lvl + 1
                    
                    print_packet_pair(
                        counter, lvl + 1, newp1, p2[i], spacing=spacing + " ", header=True
                    )
                    is_in_list = True
                else:
                    
                    # print(
                    #     "check:",
                    #     not (counter in ordered_pairs or counter in not_ordered_pairs),
                    # )
                    # print("check2:", not is_from_recursion_order)
                    # print("list:",is_in_list)
                    # print("check 3:", is_in_list and not is_from_recursion_order)
                    # if is_in_list and not is_from_recursion_order:
                    # if not (counter in ordered_pairs or counter in not_ordered_pairs):
                    if p1[i] != p2[i]:
                        is_from_recursion_order = False
                        if p1[i] < p2[i]:
                            print(
                                f"{spacing}   - Left side is smaller so in the right order"
                            )
                            ordered_pairs.append(counter)
                            ordered = True
                            break
                        elif p1[i] > p2[i]:
                            print(
                                f"{spacing}   - Right side is smaller so inputs are not in the right order"
                            )
                            not_ordered_pairs.append(counter)
                            break
                        finished_elements = abs(lenP1 - lenP2) >= 1
                        
    # print(lenP1<lenP2)
    if not is_from_recursion_order:
        if not ordered:
            if finished_elements:
                if lenP1 < lenP2:
                    print(
                        f"{spacing}  - Left side ran out of items, so inputs are in the right order"
                    )
                    ordered_pairs.append(counter)
                else:
                    print(
                        f"{spacing}  - Right side ran out of items, so inputs are not in the right order"
                    )
                    not_ordered_pairs.append(counter)


with open("inputtest2") as f:
    content = f.readlines()
    packets = [[]]
    counter = 0
    for c in content:
        if c == "\n":
            continue

        c.strip()
        packet = ast.literal_eval(c)
        if len(packets[-1:][0]) % 2 == 1:
            packets[-1].append(packet)
        else:
            packets.append([packet])

    packets = packets[1:]

    counter = 1
    for p in packets:
        # l = min(len(p[0]), len(p[1]))
        # print("confronting:")
        # print(p[0])
        # print(p[1])
        print(f"== Pair {counter} ==")
        print_packet_pair(counter, 0, p[0], p[1])
        print()
        counter += 1
        # print("min length is", l)
        # ordered = 0
        # not_ordered = 0
        # spacing = ""
        # for i in range(l):
        #     #check if p[0] or p[1] is a list

        #     if isinstance(p[0][i], list):
        #         if isinstance(p[1][i], list):

        #     if p[0][i] > p[1][i]:
        #         not_ordered+=1
        #         print(spacing+"- right side is smaller so not in the right order")
        #     print(p[0][i], p[1][i])
        print()

    print(ordered_pairs)
    print(sum(ordered_pairs))
