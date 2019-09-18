package generator

import (
	"fmt"
	"testing"
)

func TestGeneratePath(t *testing.T) {
	p := GeneratePath(5, 5)
	fmt.Println(p)
}
