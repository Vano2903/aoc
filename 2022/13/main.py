import ast
import functools

ordered_pairs = []
not_ordered_pairs = []
valids = []

# def print_packet_pair(counter,p1, p2,spacing="", header=True):
#     if header:
#         print(f"{spacing}- Compare {p1} vs {p2}")
#     lenP1 = len(p1)
#     lenP2 = len(p2)
#     ordered = False
    
#     is_from_recursion_order = False
#     is_recursion_break = False
#     finished_elements = lenP1 == 0 or lenP2 == 0
#     # if not finished_elements:
#     for i in range(min(lenP1, lenP2)):
#         if (counter in ordered_pairs or counter in not_ordered_pairs):
#             return
        
#         print(f"{spacing}  - Compare {p1[i]} vs {p2[i]}")

#         if isinstance(p1[i], list):
#             if isinstance(p2[i], list):
#                 if p1[i] == [] and p2[i] == []:
#                     print("comparing two empty set, continuing")
#                     ordered = False
#                     finished_elements = True
#                     continue
#                 is_recursion_break = print_packet_pair(counter, p1[i], p2[i], spacing=spacing+"  ", header=False)
#             else:
#                 newp2 = [p2[i]]
#                 print(f"{spacing}    - Mixed types; convert right to {newp2} and retry comparison")
#                 is_from_recursion_order = True
#                 is_recursion_break = print_packet_pair(counter, p1[i], newp2, spacing=spacing+"    ", header=True)
#         elif isinstance(p2[i], list):
#             newp1 = [p1[i]]
#             print(f"{spacing}    - Mixed types; convert left to {newp1} and retry comparison")
#             is_from_recursion_order = True
#             is_recursion_break = print_packet_pair(counter, newp1, p2[i], spacing=spacing+"    ", header=True)
#         else:
#             is_from_recursion_order = False
#             if p1[i] < p2[i]:
#                 print(f"{spacing}    - Left side is smaller so in the right order")
#                 ordered = True
#                 is_recursion_break = True
#                 ordered_pairs.append(counter)
#                 break
            
#             elif p1[i] > p2[i]:
#                 print(f"{spacing}    - Right side is smaller so inputs are not in the right order")
#                 is_recursion_break = True
#                 not_ordered_pairs.append(counter)
#                 break
            
#             finished_elements = abs(lenP1 - lenP2) >=1

#     if not (counter in ordered_pairs or counter in not_ordered_pairs):
#         #or not is_from_recursion_order or not is_recursion_break:
#         if not ordered:
#             if finished_elements:
#                 if lenP1 < lenP2:
#                     print(f"{spacing}  - Left side ran out of items, so inputs are in the right order")
#                     ordered_pairs.append(counter)
#                 elif lenP1 > lenP2:
#                     print(f"{spacing}  - Right side ran out of items, so inputs are not in the right order")
#                     not_ordered_pairs.append(counter)
#     return is_recursion_break

def is_p1_before_p2(p1, p2, silent):
    lenP1 = len(p1)
    lenP2 = len(p2)
    ordered = None
    if not silent:
        print(f"- Compare h {p1} vs {p2}")
    
    is_from_recursion_order = False
    is_recursion_break = False
    finished_elements = lenP1 == 0 or lenP2 == 0
    # if not finished_elements:
    for i in range(min(lenP1, lenP2)):
        # if (counter in ordered_pairs or counter in not_ordered_pairs):
        #     return
        if not silent:
            print(f"  - Compare {p1[i]} vs {p2[i]}")
            # if (p1 == [[1], [2, 3, 4]] and p2 ==[1,[2,[3,[4,[5,6,0]]]],8,9]) or (p2 == [[1], [2, 3, 4]] and p1==[1,[2,[3,[4,[5,6,0]]]],8,9]):

        if isinstance(p1[i], list):
            if isinstance(p2[i], list):
                if p1[i] == [] and p2[i] == []:
                    ordered = False
                    finished_elements = True
                    if not silent:
                        print("    - comparing two empty set, continuing")
                    continue
                is_from_recursion_order = True
                ordered= is_p1_before_p2(p1[i], p2[i],silent)
                return ordered
            else:
                newp2 = [p2[i]]
                if not silent:
                    print(f"      - Mixed types; convert right to {newp2} and retry comparison")
                is_recursion_break = False
                is_from_recursion_order = True
                ordered=is_p1_before_p2(p1[i], newp2,silent)
                return ordered
                
        elif isinstance(p2[i], list):
            newp1 = [p1[i]]
            if not silent:
                print(f"      - Mixed types; convert left to {newp1} and retry comparison")
            is_from_recursion_order = True
            is_recursion_break = False
            ordered = is_p1_before_p2(newp1, p2[i],silent)
            return ordered
            
        else:
            is_recursion_break = True
            if p1[i] < p2[i]:
                ordered = True
                if not silent:
                    print("      - Left side is smaller so in the right order")
                return True
            
            elif p1[i] > p2[i]:
                if not silent:
                    print("      - Right side is smaller so it is not in the right order")
                ordered = False
                return False
            finished_elements = abs(lenP1 - lenP2) >=1

    #or not is_from_recursion_order or not is_recursion_break:
    if ordered is None and not is_from_recursion_order:
        if finished_elements:
            if lenP1 < lenP2:
                if not silent:
                    print("    - Left side ran out of items, so inputs are in the right order")
                return True
            elif lenP1 > lenP2:
                if not silent:
                    print("    - Right side ran out of items, so inputs are not in the right order")
                return False
    return ordered

def sort_packets(packets):
    n = len(packets)
    # optimize code, so if the array is already sorted, it doesn't need
    # to go through the entire process
    swapped = False
    # Traverse through all array elements
    for i in range(n):
        # range(n) also work but outer loop will
        # repeat one time more than needed.
        # Last i elements are already in place
        for j in range(0, n-1):

            # traverse the array from 0 to n-i-1
            # Swap if the element found is greater
            # than the next element
            if is_p1_before_p2(packets[j] , packets[j + 1], False):
                swapped = True
                packets[j], packets[j + 1] = packets[j + 1], packets[j]

    #check if the array is sorted correctly
    # print()
    # print()
    # print()
    # print("checking if the array is sorted correctly")
    # for i in range(n):
    #     for j in range(0, n-1):
    #         # print("checking", packets[i], "with",packets[j])
    #         if not is_p1_before_p2(packets[i] , packets[j],True):
    #             swapped = True
                # print("ERRORrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")

with open("input") as f:
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
    # for p in packets:
    #     print(f"== Pair {counter} ==")
    #     print_packet_pair(counter,p[0], p[1])
    #     print()
    #     counter+=1
    #     print()

    # print(ordered_pairs)
    # print("p1:",sum(ordered_pairs))
    
    
    sorted_packets = []
    for p in packets:
        sorted_packets.append(p[0])
        sorted_packets.append(p[1])
    
    copy = []
    for p in sorted_packets:
        copy.append(p)
        
    sorted_packets.append([[2]])
    sorted_packets.append([[6]])
    
    
    sort_packets(sorted_packets)
    print()
    print()
    sorted_packets.reverse()
    for p in sorted_packets:
        print(p)
    
    
    print()
    print(sorted_packets.index([[2]])+1)
    print(sorted_packets.index([[6]])+1)
    print((sorted_packets.index([[2]])+1)*(sorted_packets.index([[6]])+1))
    
    # print(is_p1_before_p2(copy[0], copy[1], False))
    
    # print(is_p1_before_p2(copy[2], copy[15]))

    