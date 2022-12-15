
class sensor:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.closesr_beacon = None


class beacon:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.closest_sensor = None
    

class map:
    def __init__(self, lowest_x, biggest_x, biggest_y,beacons, sensors):
        self.lowest_x = lowest_x
        self.biggest_x = biggest_x
        self.biggest_y = biggest_y
        self.beacons = beacons
        self.sensors = sensors
        self.map = []
    
    def print_map(self):
        for y in range(self.biggest_y):
            for x in range(self.lowest_x, self.biggest_x):
                for s in self.sensors:
                    if s.x == x and s.y == y:
                        print("S", end="")
                        break
                else:
                    for b in self.beacons:
                        if b.x == x and b.y == y:
                            print("B", end="")
                            break
                    else:
                        print(".", end="")
            print()

if __name__ == "__main__":
    beacons = []
    sensors = []
    biggest_x = 0
    lowest_x = 1_000_000_000_000
    biggest_y = 0
    with open("inputtest") as f:
        content = f.readlines()
        for c in content:
            c.strip()
            c = c.replace('\n',"").replace("Sensor at ","").replace(": closest beacon is at","").replace(",","").replace("x=","").replace("y=","").split(" ")
            print(c)
            s = sensor(int(c[0]),int(c[1]))
            sensors.append(s)
            
            b = beacon(int(c[2]),int(c[3]))
            #remove duplicates from beacons
            for b2 in beacons:
                if b2.x == b.x and b2.y == b.y:
                    break
            else:
                beacons.append(b)

            #adding biggest x and y
            if int(c[0]) > biggest_x:
                biggest_x = int(c[0])
            if int(c[2]) > biggest_x:
                biggest_x = int(c[2])
                
            if int(c[0]) < lowest_x:
                lowest_x = int(c[0])
            if int(c[2]) < lowest_x:
                lowest_x = int(c[2])
            if int(c[1]) > biggest_y:
                biggest_y = int(c[1])
            if int(c[3]) > biggest_y:
                biggest_y = int(c[3])
    
    
    m = map(lowest_x, biggest_x, biggest_y, beacons, sensors)
    m.print_map()
    # print()
    # for b in beacons:
    #     print(b.x, b.y)
    # print()
    # for s in sensors:
    #     print(s.x, s.y)
        
    # print("biggest bx lx y are", biggest_x,lowest_x, biggest_y)
