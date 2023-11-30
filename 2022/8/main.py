

with open("input") as f:
    content = f.readlines()
    threes = []
    counter = 0
    for c in content:
        threes.append([])
        for i in c:
            if i == "\n":
                continue
            i = int(i)
            threes[counter].append(i)
        counter += 1
        
    print("threes")
    for i in threes:
        print(i)
    
    # valid = []
    # totalvisible = (len(threes)*2) + (len(threes[0])*2) - 4
    # for i in range(1, len(threes)-1):
    #     for j in range(1, len(threes[i])-1):
    #         print(f"checking three at {i}, {j} ({threes[i][j]})")
    #         #checking for taller threes on the side
    #         # print("all threes on the right are:",end=" ")
    #         # print(threes[i][0:j])
    #         if max(threes[i][0:j]) < threes[i][j]:
    #             print("this three is visible from the left")
    #             totalvisible += 1
    #             continue
    #         # print("all threes on the left are:",end=" ")
    #         # print(threes[i][j+1:])
    #         if max(threes[i][j+1:]) < threes[i][j]:
    #             print("this three is visible from the right")
    #             totalvisible += 1
    #             continue

    #         #checking for taller threes above and below
    #         above = [k[j] for k in threes[0:i]]
    #         below = [k[j] for k in threes[i+1:]]
    #         if max(above)<threes[i][j]:
    #             print("this three is visible from above")
    #             totalvisible += 1
    #             continue
            
    #         if max(below)<threes[i][j]:
    #             print("this three is visible from below")
    #             totalvisible += 1
    #             continue
    # print("total visible threes: ", totalvisible)
    
    bestView = 0
    for i in range(1, len(threes)-1):
        for j in range(1, len(threes[i])-1):
            r,l,t,b = 0,0,0,0   
            print(f"checking three at {i}, {j} ({threes[i][j]})")
            #checking for taller threes on the side
            # print("all threes on the right are:",end=" ")
            # print(threes[i][0:j])
            left = threes[i][0:j]
            left.reverse()
            # print("left three:",end=" ")
            # print(left)
            for k in left:
                l += 1
                # print(f"left: {k} >= {threes[i][j]}? {k >= threes[i][j]}")
                if k >= threes[i][j]:
                    break
            print(f"this three can see {l} threes to the left")
            
            # print("right three:",end=" ")
            # print(threes[i][j+1:])
            for k in threes[i][j+1:]:
                r += 1
                if k >= threes[i][j]:
                    break
            print(f"this three can see {r} threes to the right")

            #checking for taller threes above and below
            above = [k[j] for k in threes[0:i]]
            below = [k[j] for k in threes[i+1:]]
            print("above three:",end=" ")
            above.reverse()
            print(above)
            for k in above:
                t += 1
                if k >= threes[i][j]:
                    break
            print(f"this three can see {t} threes above")
            
            # print("below three:",end=" ")
            # print(below)
            for k in below:
                b += 1
                if k >= threes[i][j]:
                    break
            print(f"this three can see {b} threes below")
                
            threeView = l*r*t*b
            print(f"this three has a total visibility of {threeView}")
            if bestView < threeView:
                bestView = threeView
            
    print(f"the best view is {bestView}")
            # if max(below)<threes[i][j]:
            #     print("this three is visible from below")
            #     totalvisible += 1
            #     continue