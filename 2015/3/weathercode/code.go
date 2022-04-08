package weathercode

type Coordinate struct {
	Rank   int
	Row    int
	Column int
	Value  int
}

func NextValue(value int) int {
	return (value * 252533) % 33554393
}

func NextLocation(old Coordinate) Coordinate {
	var newColumn int
	var newRow int
	newRank := old.Rank

	if old.Row == 1 {
		newRank += 1
		newColumn = 1
		newRow = newRank

	} else {
		newRow = old.Row - 1
		newColumn = old.Column + 1
	}

	return Coordinate{newRank, newRow, newColumn, old.Value}
}

func ValueAt(start Coordinate, row int, column int) int {
	old := start
	for {
		new := NextLocation(old)
		new.Value = NextValue(old.Value)
		if new.Row == row && new.Column == column {
			return new.Value
		}
		old = new
	}

}
