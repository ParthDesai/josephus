package simulator

import "fmt"

type Person struct {
	isAlive         bool
	currentPosition int
	positionHistory []int
}

func (p *Person) init(initialPosition int) {
	p.isAlive = true
	p.currentPosition = initialPosition
	p.positionHistory = make([]int, 10)
}

func (p *Person) GetCurrentPosition() int {
	return p.currentPosition
}

func (p *Person) SetCurrentPosition(newPosition int) {
	p.positionHistory = append(p.positionHistory, p.currentPosition)
	p.currentPosition = newPosition
}

func (p *Person) Kill() {
	if !p.isAlive {
		panic(fmt.Errorf("Cannot kill already dead person at position: %d", p.currentPosition))
	}
	p.isAlive = false
}

func (p *Person) IsAlive() bool {
	return p.isAlive
}
