import math

ROUNDS = 20

class monkey(object):
    number:int
    starting: list[int, ...]
    passed: list[int, ...]
    operation: str
    value: int
    division: int
    monkey1num: int
    monkey1 = None
    monkey2num: int
    monkey2 = None
    
    def __init__(self, number=0,starting=[], operation="",value=0, division=0, monkey1=0, monkey2=0):
        self.number = number
        self.starting = starting
        self.operation = operation
        self.value = value
        self.division = division
        self.monkey1num = monkey1
        self.monkey2num = monkey2
        
    def link_monkeys(self, monkey1, monkey2):
        self.monkey1 = monkey1
        self.monkey2 = monkey2
    
    def fetch(self,number):
        self.passed.append(number)
    
    def check(self):
        for worry_lvl in self.starting:
            if self.operation == "+":
                worry_lvl = worry_lvl + self.value
            elif self.operation == "*":
                worry_lvl = worry_lvl * self.value
            worry_lvl = math.floor(worry_lvl/3)
            
            if worry_lvl % self.division:
                self.monkey1.fetch(worry_lvl)
            else:
                self.monkey2.fetch(worry_lvl)

class monkey_master(object):
    monkeys: list[monkey, ...]
    
    def __init__(self):
        self.monkeys = []
    
    def link_monkeys(self):
        for m in self.monkeys:
            m1 = self.get_monkey(m.monkey1num)
            m2 = self.get_monkey(m.monkey2num)
            m.link_monkeys(m1,m2)
            
    def get_monkey(self, number):
        for monkey in self.monkeys:
            if monkey.number == number:
                return monkey
        return None

if __name__ == "__main__":
    monkey_master = monkey_master()
    
    with open("inputtest") as f:
        content = f.readlines()
        m = monkey()
        for c in content:
            if "Monkey" in c:
                print(int(c.split(" ")[1][:-2]))
                m.number = int(c.split(" ")[1][:-2])
            elif "Starting items: " in c:
                c = c.replace("  Starting items: ","")
                l = c.split(", ")
                items = []
                for i in l:
                    items.append(int(i))
                m.starting = items
            elif "  Operation:" in c:
                c = c.replace("Operation: new = old","").strip()
                vals = c.split(" ")
                if vals[1] == "old":
                    m.operation = "^"
                    m.value = 2
                else:
                    m.operation = vals[0]
                    m.value = int(vals[1])
            elif "  Test: divisible by " in c:
                c = c.replace("  Test: divisible by ","")
                m.division = int(c)
            elif "    If true:" in c:
                m.monkey1 = int(c.replace("If true: throw to monkey","").strip())
            elif "    If false:" in c:
                m.monkey2 = int(c.replace("If false: throw to monkey","").strip())
            
            else:
                print(m.number)
                monkey_master.monkeys.append(m)            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            