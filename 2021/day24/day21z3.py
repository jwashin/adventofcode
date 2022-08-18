from z3 import *
from common24 import extract_parameters

# https://www.keiruaprod.fr/blog/2021/12/29/a-comprehensive-guide-to-aoc-2021-day-24.html
def solve(div_check_add: list, part: int) -> int:
    solver = Optimize()
    z = 0  # this is our running z, which has to be zero at the start and end
    # We have 14 inputs, they all are integers between 1 and 9 included
    ws = [Int(f'w{i}') for i in range(14)]
    for i in range(14):
        solver.add(And(ws[i] >= 1, ws[i] <= 9))
    # The value where we concatenate our input digits
    digits_base_10 = Int(f"digits_base_10")
    solver.add(digits_base_10 == sum((10 ** i) * d for i, d in enumerate(ws[::-1])))
    # We implement the subroutine as a list of constraints, one for each of the 14 blocks:
    for (i, [div, check, add]) in enumerate(div_check_add):
        z = If(z % 26 + check == ws[i], z / div, z / div * 26 + ws[i] + add)
    # The final z value must be zero
    solver.add(z == 0)
    if part == 1:
        solver.maximize(digits_base_10)
    else:
        solver.minimize(digits_base_10)
    assert (solver.check() == sat)  # the solver must find a solution
    return solver.model().eval(digits_base_10)


input = open("input.txt").read()
div_check_add = extract_parameters(input)
print(solve(div_check_add, 1))
print(solve(div_check_add, 2))