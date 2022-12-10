
counter = 0
with open("input", encoding="utf-8") as f:
    data = f.read().splitlines()
    for d in data:
        contents = d.split(",")
        v1,v2 = int(contents[0].split("-")[0]),int(contents[0].split("-")[1])
        v3,v4 = int(contents[1].split("-")[0]),int(contents[1].split("-")[1])
        # if (v1 <= v3 <= v4 <= v2) or (v3 <= v1 <= v2 <= v4):
        if (v1 <= v3 <= v2) or (v3 <= v1 <= v4):
            counter += 1

print(counter)
