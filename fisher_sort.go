package main

import (
	"fmt"
	"math/rand"
)

// Sample randomly selects k integers from the range [min, max).
// Given that Sample's algorithm is to run Fisher-Yates shuffle k times, it's complexity is O(k).
func Sample(k, min, max int) (sampled []int) {
	swapped := make(map[int]int, k)
	for i := 0; i < k; i++ {
		// generate a random number r, where i <= r < max-min
		r := rand.Intn(max-min-i) + i

		// swapped[i], swapped[r] = swapped[r], swapped[i]
		vr, ok := swapped[r]
		if ok {
			sampled = append(sampled, vr+min)
		} else {
			sampled = append(sampled, r+min)
		}
		vi, ok := swapped[i]
		if ok {
			swapped[r] = vi
		} else {
			swapped[r] = i
		}
	}
	return
}

func main() {
	for i := 0; i < 10; i++ {
		reply := Sample(3, 30, 70)
		fmt.Printf("%d, %d, %d\n", reply[0], reply[1], reply[2])
	}
}
