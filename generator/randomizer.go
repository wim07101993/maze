package generator

import (
	"math"
	"math/rand"
	"time"
)

type Randomizer struct {
	src       rand.Source
	cache     int64
	remaining int
}

func NewRandomizer() *Randomizer {
	return &Randomizer{
		src: rand.NewSource(time.Now().UnixNano()),
	}
}

func (r *Randomizer) Gen(size uint) int64 {
	if r.remaining < int(size) {
		r.cache = r.src.Int63()
		r.remaining = 63
	}

	mask := math.Pow(2, float64(size)) - 1
	result := r.cache & int64(mask)
	r.cache >>= size
	r.remaining -= int(size)

	return result
}

func (r *Randomizer) GenInt() int {
	return int(r.src.Int63())
}
