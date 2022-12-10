
field = []
head = {'x': 0, 'y': 0}

class node:
    

def move_head(direction, distance):
    global head
    
    if direction == 'U':
        for i in range(distance):
            head['y'] += 1
            print(head)
            #convert head to string and hash it
            fields.append(hash(str(head)))

            
    elif direction == 'D':
        for i in range(distance):
            head['y'] -= 1
            print(head)
    elif direction == 'R':
        for i in range(distance):
            head['x'] += 1
            print(head)
            
    elif direction == 'L':
        for i in range(distance):
            head['x'] -= 1
            print(head)
            


if __name__ == "__main__":
    with open("inputtest") as f:
        content = f.readlines()
        for c in content:
            mov = c.split(' ')
            move_head(mov[0], int(mov[1]))
        print(head)
