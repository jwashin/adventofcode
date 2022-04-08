package elevator

const plus = '('
const minus = ')'

func GetFloor(instructions string) int {
	floor := 0
	for _, instruction := range instructions {
		if instruction == plus {
			floor += 1
		}
		if instruction == minus {
			floor -= 1
		}

	}
	return floor
}

func GetPositionX(instructions string, x int) int {
	floor := 0
	for idx, instruction := range instructions {
		if instruction == plus {
			floor += 1
		}
		if instruction == minus {
			floor -= 1
			if floor == -1 {
				// idx is zero-based, so add 1 for this
				return idx + 1
			}
		}

	}
	return -1
}
