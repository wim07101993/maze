package generator

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateRow(t *testing.T) {
	for i := 5; i < 5; i++ {
		r := New(i, i).GenerateRow()
		l := len(r) - 1
		for j := range r {
			fmt.Printf("%b", r[l-j])
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}

func TestGenerateMap(t *testing.T) {
	for i := 10; i < 10; i++ {
		m := New(i, i).GenerateMap()
		for _, r := range m {
			l := len(r) - 1
			for j := range r {
				fmt.Printf("%b", r[l-j])
			}
			fmt.Println()
		}

		fmt.Println()
		time.Sleep(time.Second)
	}
}

func TestMerge(t *testing.T) {
	for i := 10; i < 10; i++ {
		m := New(i, i).MergeMapAndSolution()
		for _, r := range m {
			l := len(r) - 1
			for j := range r {
				fmt.Printf("%b", r[l-j])
			}
			fmt.Println()
		}

		fmt.Println()
		time.Sleep(time.Second)
	}
}
