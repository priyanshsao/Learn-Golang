package helpers

import (
	"math/rand"
	"time"
)

/*
	                        ---3---
	Creating a function that generates a random number using
	"math/rand" library. It takes 'n' as an argument, here n is
	the no. that tells compiler that the number generated should
	be between 0 and n (0 is included but n is not included).

	To generate different numbers every time we run the program,
	we need to feed it with a different seed value each time.
	A common way to do this is by using the current time as the seed.
	By seeding the random number generator(rng) with the current time,
	we ensure that the numbers generated are different on each run.
	After this read the second part of section 2.
*/
func RandomNumber(n int) int {
	seed := time.Now().UnixNano() 
	rng := rand.New(rand.NewSource(seed))
	value := rng.Intn(n)
	return value
}
