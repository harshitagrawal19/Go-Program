package main

import (
  "fmt"
  "Fibonaci/fibonaci"
)
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci.Fib(i))
	}
}
