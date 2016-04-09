// This is in-efficient implementation of Josephus problem.
// There is a straightforward equation to solve the problem,
// But It would be unfair, so I have implemented a algorithm
// that will compute the answer recursively.
package main

import (
	"flag"
	"fmt"

	"github.com/ParthDesai/josephus/simulator"
)

// Get N from the command line
// Arguments: <None>
// Return value: *UInt64
func getCmdN() *uint64 {
	n := flag.Uint64("N", 0, "Total number of people in the circle ")

	flag.Parse()

	return n
}

func main() {

	n := getCmdN()

	if *n == 0 {
		flag.PrintDefaults()
		return
	}

	circle := new(simulator.CircleOfDeath)

	circle.Init(*n, 1)

	fmt.Println("The number of person who will be alive is:", circle.Execute())
}
