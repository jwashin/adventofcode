using Test
using Combinatorics

function part1(aList)
    x = combinations(aList, 2)
    for item in x
        if reduce(+, item) == 2020
            return reduce(*, item)
        end
    end
end

@test part1((1721, 979, 366, 299, 675, 1456)) == 514579

function part2(aList)
    x = combinations(aList, 3)
    for item in x
        if reduce(+, item) == 2020
            return reduce(*, item)
        end
    end
end

@test part2((1721, 979, 366, 299, 675, 1456)) == 241861950

input = parse.(Int, readlines(open("input.txt")))

print("Part1:", part1(input), "\n")

print("Part2:", part2(input), "\n")
