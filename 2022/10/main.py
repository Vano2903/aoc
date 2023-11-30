


if __name__ == "__main__":
    total = 0
    x = 1
    # cycles = [40, 40, 40, 40, 40, 40] #, 40, 120, 160, 200, 240
    current_cycle = 1
    reader = ""
    reader_position = 0
    spriteCenter = 1
    cycles = 6
    with open("input") as f:
        # print("###"+("."*(40-3)))
        content = f.readlines()
        for cycle in range(cycles):
            current_cycle = 1
            reader_position = 0
            spriteCenter = 1

            for c in content:
                if reader_position >= 40:
                    reader_position = 0
                c = c.strip()
                #first read
                print(("."*(spriteCenter-1))+"###"+("."*(40-(spriteCenter+1))))
                
                if reader_position in range(spriteCenter-1, spriteCenter+2):
                    reader += "#"
                else:
                    reader += "."
                reader_position += 1
                print((" " * (reader_position-1))+"^")
                if reader_position >= 40:
                    reader_position = 0

                #noop cycle
                if c == "noop":
                    current_cycle += 1
                    continue
                # print(f"cycle {current_cycle}: {reader} {reader_position} {spriteCenter} {c}")
                if reader_position >= 40:
                    reader_position = 0

                #second read
                current_cycle += 1
                print(("."*(spriteCenter-1))+"###"+("."*(40-(spriteCenter+1))))
                if reader_position in range(spriteCenter-1, spriteCenter+2):
                    reader += "#"
                else:
                    reader += "."
                reader_position += 1
                print((" " * (reader_position-1))+"^")
                if reader_position >= 40:
                    reader_position = 0


                #end operation for addx
                v = c.replace("addx ", "")
                # print("x:",x,"v:",v)
                spriteCenter += int(v)
                current_cycle += 1
            # print("second addx:",current_cycle)

                # print(f"cycle {cycle}: {x}")

print()
print(reader)
print()

for i in range(6):
    print(reader[i*40:(i+1)*40])
