package simulator

import "fmt"

// Person represents individual person in the circle
type Person struct {
	isAlive  bool
	position int
}

// Init initializes the person.
func (p *Person) Init(position int) {
	p.isAlive = true
	p.position = position
}

// GetPosition gets the position of the Person
func (p *Person) GetPosition() int {
	return p.position
}

// Kill changes state of person from alive to dead
func (p *Person) Kill() {
	if !p.isAlive {
		panic(fmt.Errorf("Cannot kill already dead person at position: %d", p.position))
	}
	p.isAlive = false
}

// IsAlive indicates whether this person is alive or not
func (p *Person) IsAlive() bool {
	return p.isAlive
}
