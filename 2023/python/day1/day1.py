numstrs = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

def p1(l):
    digits = [int(c) for _, c in enumerate(l) if c.isdigit()]
    return (digits[0] * 10) + digits[-1]

def p2(l):
    digits = []
    for i, c in enumerate(l):
        for j, n in enumerate(numstrs):
            if l[i:].startswith(n):
                digits.append(j+1)
        if c.isdigit():
            digits.append(int(c))
    return (digits[0] * 10) + digits[-1]

with open("2023/day1/input.txt", 'r') as f:
    lines = f.readlines()
    print(sum(map(p1,lines)))
    print(sum(map(p2,lines)))
