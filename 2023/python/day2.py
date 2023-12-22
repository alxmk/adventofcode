import re

threshold = {"red": 12, "green": 13, "blue": 14}

def p1(l: str):
    parts = re.sub("[:;,\n]", "", l).split()
    for i, part in enumerate(parts):
        if part in threshold and int(parts[i-1]) > threshold[part]:
            return 0
    return int(parts[1])

def p2(l: str):
    minimum = {"red": 0, "green": 0, "blue": 0}
    parts = re.sub("[:;,\n]", "", l).split()
    for i, part in enumerate(parts):
        if part in minimum:
            minimum[part] = max(minimum[part], int(parts[i-1]))
    return minimum["red"] * minimum["green"] * minimum["blue"]

with open("2023/day2/input.txt", 'r') as f:
    lines = f.readlines()
    print(sum(map(p1,lines)))
    print(sum(map(p2,lines)))
