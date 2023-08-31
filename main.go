package main

import (
	"fmt"

	"github.com/yuucu/mockgen-sample/random"
)

func main() {
	rand := random.NewRandom()
	res := exec(rand, 10)
	fmt.Println(res)
}

func exec(rand random.IFRandom, n int) int {
	return rand.Intn(n)
}
