

total = 0
groupLine = 0
lines = []
with open("input", mode="r", encoding="utf-8") as f:
    for line in f:
        lines.append(line.strip())
        groupLine += 1
        if groupLine % 3 == 0:
            print(lines)
            # find the duplicate character between first and second
            compare = set(lines[0]) & set(lines[1]) & set(lines[2])
            print(compare)
            lines = []
            #from a through z have point 1 through 26
            #from A through Z have point 27 through 52
            for c in compare:
                print(c)
                if c.islower():
                    print(ord(c) - ord('a') + 1)
                    total += ord(c) - ord('a') + 1
                else:
                    print(ord(c) - ord('A') + 27)
                    total += ord(c) - ord('A') + 27
print(total)
