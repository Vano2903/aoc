
containers = [[],[],[],[],[],[],[],[],[]]
with open("input") as f:
    content = f.readlines()
    counter = 0
    for c in content:
        if c == "\n":
            break
        print(c, end="")
        counter += 1
        for i in range(0, 36, 4):
            if c[i] == "[":
                containers[i//4].insert(0,c[i+1:i+2])
    print("end parsing")
    for c in content[counter+1:]:
        print(c, end="")
        c = c.replace("move ", "").replace("from ", "").replace("to ", "").split(" ")
        c[0] = int(c[0])
        c[1] = int(c[1])
        c[2] = int(c[2])
        # for i in range(c[0]):
        #     containers[c[2]-1].append(containers[c[1]-1].pop())
        mover = []
        for i in range(c[0]):
            mover.append(containers[c[1]-1].pop())
        mover.reverse()
        containers[c[2]-1].extend(mover)
    print(containers)

for c in containers:
    print(c[::-1][0], end="")
    

#CDLCSSHMG
#JCMHLVGMG