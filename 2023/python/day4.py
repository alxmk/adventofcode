
def p1(l):
    num = len(winners(l))
    if num == 0:
        return 0
    return 1 << num - 1

def p2(lines):
    counts = [1] * len(lines)
    for i, c in enumerate(lines):
        for j in range(i+1, i+1+len(winners(c))):
            counts[j] += counts[i]
    return sum(counts)

def winners(l):
    return list(filter(lambda n: n in l.split('|')[0].split(), filter(lambda n: n.isdigit(), l.split('|')[1].split())))

with open("2023/day4/input.txt", 'r') as f:
    lines = f.readlines()
    print(sum(map(p1,lines)))
    print(p2(lines))
