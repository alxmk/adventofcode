import re
import z3

def parse(lines):
    vecs = []
    for line in lines:
        vecs.append([int(s) for s in re.findall(r'-?\d+', line)])
    return vecs

def p2(vecs):
    px, py, pz, vx, vy, vz = z3.Reals("px py pz vx vy vz")
    solver = z3.Solver()
    for i, v in enumerate(vecs[:3]):
        ti = z3.Real(f"t{i}")
        solver.add(ti > 0)
        solver.add(px + ti * vx == v[0] + ti * v[3])
        solver.add(py + ti * vy == v[1] + ti * v[4])
        solver.add(pz + ti * vz == v[2] + ti * v[5])
    solver.check()
    return sum(solver.model()[var].as_long() for var in [px, py, pz])

with open("2023/day24/input.txt", 'r') as f:
    lines = f.readlines()
    print(p2(parse(lines)))
