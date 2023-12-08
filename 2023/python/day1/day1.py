numchars = range(ord('1'), ord('9')+1)
numstrs = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

def val(c):
    return c - ord('0')

def p1(l):
    digits = list(val(c) for c in l.encode("utf-8") if c in numchars)
    # digits = list(map(lambda c: val(c), filter(lambda c: c in numchars, l.encode("utf-8"))))
    return (digits[0] * 10) + digits[-1]

def p2(l):
    digits = []
    for i, c in enumerate(l.encode("utf-8")):
        for j, n in enumerate(numstrs):
            if l[i:].startswith(n):
                digits.append(j+1)
        if c in numchars:
            digits.append(val(c))
    return (digits[0] * 10) + digits[-1]

with open("2023/day1/input.txt", 'r') as f:
    lines = f.readlines()
    print(sum(map(p1,lines)))
    print(sum(map(p2,lines)))
