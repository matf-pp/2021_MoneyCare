package main

import (
	"fmt"
	"math/rand"
)

func getNormDistro(x int) {
	for i := 0; i < x; i++ {
		fmt.Print("   ", int(normalInverse(25000, 100)), "   ")
	}
}

func normalInverse(mu float32, sigma float32) float32 {
	return float32(rand.NormFloat64()*float64(sigma) + float64(mu))
}

func main() {
	getNormDistro(30)
}
