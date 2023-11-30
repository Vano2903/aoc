import ast
    
def flatten(lst):
    for el in lst:
        if isinstance(el, list):  # N.B. this only works for lists, not general
                                  # collections e.g. sets, tuples, dicts, etc...
            # recurse
            yield from flatten(el)
        else:
            # generate
            yield el
            
    
def sort(packets):
    for i in range(len(packet)):
        for j in range(0, len(packet)-i-1):
            #sort by first element (it can be an empty list)
            if len(packets[j]) == 0:
                packets[j], packets[j+1] = packets[j+1], packets[j]
            for k in range(0, abs(len(packets[j])-len(packets[j+1]))-1):
                if packets[j][k] > packets[j+1][k]:
                    packets[j], packets[j+1] = packets[j+1], packets[j]
                    break
                
    return packets

if __name__ == "__main__":
    with open("inputtest") as f:
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
            p1 = list(flatten(p[0]))
            p2 = list(flatten(p[1]))
            sorted_packets.append(p1)
            sorted_packets.append(p2)
            print(p1)
            print(p2)

        sorted_packets.append([2])
        sorted_packets.append([6])
        
        print(sorted_packets)
        print()
        p = sort(sorted_packets)
        print(p)
        # # # for i in range(len(sorted_p
        # # ordered = order_packages(sorted_packets)
        # sort_packets(sorted_packets)
        # print()
        # print()
        # sorted_packets.reverse()
        # for p in sorted_packets:
        #     print(p)
    
    