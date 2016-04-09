package simulator

import "fmt"

type Person struct {
	isAlive  bool
	position int
}

func (p *Person) Init(position int) {
	p.isAlive = true
	p.position = position
}

func (p *Person) GetPosition() int {
	return p.position
}

func (p *Person) Kill() {
	if !p.isAlive {
		panic(fmt.Errorf("Cannot kill already dead person at position: %d", p.position))
	}
	p.isAlive = false
}

func (p *Person) IsAlive() bool {
	return p.isAlive
}
