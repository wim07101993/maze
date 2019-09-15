package maze

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

func GeneratePath(maxX, maxY int) Path {
	// TODO implment generate path
	return nil
}
