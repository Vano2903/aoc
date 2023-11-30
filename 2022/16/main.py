class valve:
    def __init__(self, name, pressure_val):
        self.name = name
        self.pressure_val = pressure_val
        self.valve = []

    def add_valve(self, valve):
        self.valve.append(valve)


if __name__ == "__main__":
    valves = []
    with open("inputtest") as f:
        content = f.readlines()
        for c in content:
            c = c.strip()
            