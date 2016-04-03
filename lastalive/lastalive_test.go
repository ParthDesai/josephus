package lastalive

import "testing"
import "math"

// Edge test case
func TestCalculateLastAliveEdge(t *testing.T) {
	out, err := CalculateLastAlive(1)

	if err != nil {
		t.Errorf("Not Expected any error. But found:" + err.Error())
	}

	if out != 0 {
		t.Errorf("Expected: %d, Received: %d", 0, out)
	}
}

// Edge Test case 2
func TestCalculateLastAliveEdge2(t *testing.T) {
	out, err := CalculateLastAlive(math.MaxUint64)

	if err != nil {
		t.Errorf("Not Expected any error. But found:" + err.Error())
	}

	if out != 18446744073709551614 {
		t.Errorf("Expected: %d, Received: %d", 0, out)
	}
}

// Invalid input test case
func TestCalculateLastAliveInvalidInput(t *testing.T) {
	_, err := CalculateLastAlive(0)

	if err == nil {
		t.Errorf("Expected error.")
	}
}

// Normal test case
func TestCalculateLastAliveNormal(t *testing.T) {

	out, err := CalculateLastAlive(7)

	if err != nil {
		t.Errorf("Not Expected any error. But found:" + err.Error())
	}

	if out != 6 {
		t.Errorf("Expected: %d, Received: %d", 0, out)
	}

}

// Benchmark with maximum value
func BenchmarkCalculateLastAlive(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CalculateLastAlive(math.MaxUint64)
	}

}
