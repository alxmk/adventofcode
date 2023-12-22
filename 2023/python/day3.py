
def p1(nums, symbols):
    out = 0
    for p in nums.items():
        for q in symbols:
            if adjacent(p, q):
                out += int(p[1])
                break
    return out

def adjacent(npair, q):
    return npair[0][0]-1 <= q[0] <= npair[0][0]+1 and npair[0][1]-len(npair[1])-1 <= q[1] <= npair[0][1]

def p2(nums, symbols):
    sum = 0
    for q in symbols:
        adj = [int(p[1]) for p in nums.items() if adjacent(p, q)]
        if len(adj) == 2:
            sum += adj[0] * adj[1]
    return sum

def parse(lines):
    nums, symbols = {}, {}
    for i, l in enumerate(lines):
        num = ""
        for j, c in enumerate(l):
            if c.isdigit():
                num += c
                continue
            elif c != '.' and c != '\n':
                symbols[(i,j)] = c
            if num != "":
                nums[(i,j)] = num
                num = ""
    return nums, symbols

with open("2023/day3/input.txt", 'r') as f:
    lines = f.readlines()
    print(p1(*parse(lines)))
    print(p2(*parse(lines)))
