package lastalive

import "testing"

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
