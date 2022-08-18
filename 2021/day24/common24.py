import re

#https://www.keiruaprod.fr/blog/2021/12/29/a-comprehensive-guide-to-aoc-2021-day-24.html

def extract_parameters(program):
    repeated_program = r"""inp w
mul x 0
add x z
mod x 26
div z (.*)
add x (.*)
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y (.*)
mul y x
add z y"""

    div_check_add = re.findall(repeated_program, program)
    assert (len(div_check_add) == 14), len(div_check_add)
    return [list(map(int, dca)) for dca in div_check_add]