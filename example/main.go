package main

import (
	"fmt"
	"math/rand"

	"github.com/bt/btioc"
)

// A trivial random number generator
type RandomGenerator struct {
	min int
	max int
}

// Returns a random number between min and max.
func (r *RandomGenerator) Random() int {
	return rand.Intn(r.max - r.min) + r.min
}

func main() {
	min := 100
	max := 200

	// Instantiate a new random number generator.
	rng := RandomGenerator{
		min: min,
		max: max,
	}

	// Register the random number generator in the container.
	btioc.Register("random", rng)

	// Later on, retrieve the object for use.
	obj, _ := btioc.Retrieve("random")

	// Type assert so we can access the RandomGenerator's functions.
	random := obj.(RandomGenerator)

	// Print out the random number.
	fmt.Printf("Random number between %d and %d : %d", min, max, random.Random())
}
