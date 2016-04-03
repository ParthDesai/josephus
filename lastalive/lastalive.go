package lastalive

import (
	"fmt"
	"math"
)

func calculateLastAliveInternal(n uint64) uint64 {

	if n <= 2 {
		return 1
	}

	if n%2 == 0 {
		return (2 * calculateLastAliveInternal(n/2)) - 1
	}

	return (2 * calculateLastAliveInternal(n/2)) + 1
}

// CalculateLastAlive Calculates the number of person who will be alive
// at last.
// Maximum stack value can be used: 64*16=1024 bytes
func CalculateLastAlive(n uint64) (uint64, error) {
	if n == 0 {
		return 0, fmt.Errorf("Invalid Value. Value of N should be between 1 and %d",
			uint64(math.MaxUint64))
	}

	return calculateLastAliveInternal(n) - 1, nil
}
