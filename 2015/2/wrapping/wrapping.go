package wrapping

import (
	"strconv"
	"strings"
)

type Box struct {
	length int
	width  int
	height int
}

func (dm Box) PaperTotal() int {
	s1 := dm.length * dm.width
	s2 := dm.width * dm.height
	s3 := dm.height * dm.length
	min := s1
	if s2 < min {
		min = s2
	}
	if s3 < min {
		min = s3
	}
	return 2*(s1+s2+s3) + min
}

func (dm Box) RibbonNeeded() int {
	p1 := (dm.length + dm.width) * 2
	p2 := (dm.width + dm.height) * 2
	p3 := (dm.length + dm.height) * 2

	min := p1
	if p2 < p1 {
		min = p2
	}
	if p3 < min {
		min = p3
	}

	v := dm.length * dm.width * dm.height

	return min + v
}

func MakeBox(input string) *Box {
	dims := strings.Split(input, "x")
	l, _ := strconv.Atoi(dims[0])
	w, _ := strconv.Atoi(dims[1])
	h, _ := strconv.Atoi(dims[2])

	out := Box{l, w, h}
	return &out
}

func GetPaperTotal(s string) int {

	item := MakeBox(s)
	return item.PaperTotal()

}

func GetRibbonNeeded(s string) int {
	item := MakeBox(s)
	return item.RibbonNeeded()
}
