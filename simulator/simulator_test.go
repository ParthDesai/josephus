package simulator

import "testing"

// Error test case 1
func TestErrorCase1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected error. But not found any.")
		}
	}()

	circle := new(CircleOfDeath)
	circle.Init(3, 5)
}

// Error test case 2
func TestErrorCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected error. But not found any.")
		}
	}()

	circle := new(CircleOfDeath)
	circle.Init(0, 5)
}

func TestSimulator(t *testing.T) {
	circle := new(CircleOfDeath)
	circle.Init(7, 1)

	output := circle.Execute()

	if output.LastAlive != 6 {
		t.Error("Expected 7. Got:", output)
	}
}
