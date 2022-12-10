import os

length = 14
with open("input") as f:
    content = f.read()
    for i in range( len(content)):
        val = list(content[i:i+length])
        print(val)
        invalid = False
        s = []
        for a in val:
            s.append(a)
        s.sort()
        for j in range(length-1):
            if s[j] == s[j+1]:
                invalid = True
                break
        if not invalid:
            print("Found valid value: " + str(val))
            print(f"found {i} and is {content[i]}")
            print(f"found {i-1} and is {content[i-1]}")
            print(f"ountput {i+length}")
            break