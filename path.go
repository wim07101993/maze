package maze

import "strings"

var (
	pathChars [3]byte
)

func init() {
	pathChars[0] = ' '
	pathChars[1] = '0'
	pathChars[2] = byte(0)
}

func IsPath(b byte) bool {
	for _, c := range pathChars {
		if b == c {
			return true
		}
	}
	return false
}

type Path []*Coordinate

func (p Path) Last() *Coordinate {
	return p[len(p)-1]
}

func (p Path) Contains(x *Coordinate) bool {
	for _, c := range p {
		if c == x {
			return true
		}
	}

	return false
}

func (p Path) String() string {
	maxRow := 0
	bss := [][]bool{}
	for _, c := range p {
		for c.Y+1 > len(bss) {
			bss = append(bss, []bool{})
		}
		for c.X+1 > len(bss[c.Y]) {
			bss[c.Y] = append(bss[c.Y], false)
		}
		for c.X+1 > maxRow {
			maxRow = c.X + 1
		}
		bss[c.Y][c.X] = true
	}

	builder := strings.Builder{}
	for _, bs := range bss {
		for i := 0; i < maxRow; i++ {
			if i >= len(bs) || !bs[i] {
				builder.WriteString(".")
			} else {
				builder.WriteString("x")
			}
		}
	}

	return builder.String()
}
