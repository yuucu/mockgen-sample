package random

import "math/rand"

type IFRandom interface {
	Intn(n int) int
}

type Random struct{}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) Intn(n int) int {
	return rand.Intn(n)
}
