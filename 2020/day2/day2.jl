using Test

function parseLine(s)
	s1 = split(s)
	j = split(s1[1], "-")
	f = j[1]
	t = j[2]
	c = replace(s1[2], ":" => "")
	tst = s1[3]
	return (f, t, c, tst)
end

@test parseLine("1-3 a: abcde") == ("1", "3", "a", "abcde")

function isValid(s)
	f, t, c, tst = parseLine(s)
	min = parse(Int, f)
	max = parse(Int, t)
	z = length(collect(eachmatch(r"" * c, tst)))
	if z >= min && z <= max
		return true
	else
		return false
	end
end

@test isValid("1-3 a: abcde") == true
@test isValid("1-3 b: cdefg") == false
@test isValid("2-9 c: ccccccccc") == true

function isValid2(s)
	f, t, c, tst = parseLine(s)
	min = parse(Int, f)
	max = parse(Int, t)
	a = string(tst[min])
	b = string(tst[max])
	if a == c && b == c
		return false
	end
	if a == c
		return true
	end
	if b == c
		return true
	end
	return false
end

@test isValid2("1-3 a: abcde") == true
@test isValid2("1-3 b: cdefg") == false
@test isValid2("2-9 c: ccccccccc") == false

input = readlines(open("input.txt"))

print("part 1: ", length(filter((x) -> isValid(x), input)))
print("part 2: ", length(filter((x) -> isValid2(x), input)))

