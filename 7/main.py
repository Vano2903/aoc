
dirs = {"dir":"/", "files":[], "subdirs":[]}
with open("input", encoding="utf-8") as f:
    content = f.readlines()
    lastdir = ""
    index = 0
    for c in content:
        if c.contains("$ cd "):
            directory =  c.replace("$ cd ", "")
            if directory == "..":
                a #idk yet
            if directory == "/":
                b #goto root
            else:
                lastdir = c.replace("$ cd ", "")
                dirs["subdirs"].apppend({"dir":lastdir, "files":[], "subdirs":[]})
        if c.contains("ls -l"):
            if c.contains("dir "):
                dirs[].
            
            
# [{"dir":a, "files":[1237,123,433], subdirs:[{"dir":b, "files":[1237,123,433]}, "subdirs":[]]}]