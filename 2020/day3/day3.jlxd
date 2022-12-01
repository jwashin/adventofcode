using Test

function countTrees(input, right, down)
	x = 0
	y = 1
	ymax = length(input)
	xmax = length(input[1])
	treecount = 0
	while y < ymax
		x += right
		y += down
		c = input[y][(x%xmax)+1]
		# println(y, " ", x, " ", c)
		if string(c) == "#"
			treecount += 1
		end
	end
	return treecount
end

testHill = split(
	"""..##.......
	  #...#...#..
	  .#....#..#.
	  ..#.#...#.#
	  .#...##..#.
	  ..#.##.....
	  .#.#.#....#
	  .#........#
	  #.##...#...
	  #...##....#
	  .#..#...#.#""",
	"\n",
)

@test countTrees(testHill, 3, 1) == 7

function part2(input)
	product = 1
	for (r, d) in ((1, 1), (3, 1), (5, 1), (7, 1), (1, 2))
		c = countTrees(input, r, d)
		product *= c
	end
	return product

end

@test part2(testHill) == 336

input = readlines(open("input.txt"))
println("Part 1. ", countTrees(input, 3, 1))
println("Part 2. ", part2(input))