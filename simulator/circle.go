package simulator

import "fmt"

const (
	notstarted = iota
	executing  = iota
	completed  = iota
)

type CircleOfDeath struct {
	people      []Person
	stepPerKill int

	lastAlive uint64
}

func (c *CircleOfDeath) Init(numberOfPeople uint64, stepPerKill int) {

	if numberOfPeople < 1 {
		panic(fmt.Errorf("Number of people should be greater than or equal to 1. Found:%d", numberOfPeople))
	}

	if uint64(c.stepPerKill) >= numberOfPeople {
		panic(fmt.Errorf("Step per kill should be less then number of people."+
			" Found: NumberOfPeople:%d StepPerKill:%d", numberOfPeople, stepPerKill))
	}

	c.stepPerKill = stepPerKill

	c.people = make([]Person, numberOfPeople)

	for index, _ := range c.people {
		c.people[index].Init(index)
	}
}

func (c *CircleOfDeath) findFirstAlive() int {
	firstAlive := -1

	for index, person := range c.people {
		if person.IsAlive() {
			firstAlive = index
			break
		}
	}

	return firstAlive
}

func (c *CircleOfDeath) getNextIndex(index int) int {

	nextIndex := index + 1

	if nextIndex >= len(c.people) {
		nextIndex = 0
	}

	return nextIndex
}

func (c *CircleOfDeath) findNextAlive(index int) int {
	nextAliveIndex := index

	stepRemaining := c.stepPerKill

	for i := c.getNextIndex(index); i != index; i = c.getNextIndex(i) {
		if c.people[i].IsAlive() {
			nextAliveIndex = i
			stepRemaining -= 1
		}

		if stepRemaining == 0 {
			break
		}
	}

	return nextAliveIndex
}

func (c *CircleOfDeath) Execute() int {

	i := c.findFirstAlive()

	if i == -1 {
		panic(fmt.Errorf("Inconsistent state. No body is alive."))
	}

	for j := c.findNextAlive(i); i != j; {
		fmt.Println("Killing:", j)
		c.people[j].Kill()

		i = c.findNextAlive(i)
		j = c.findNextAlive(i)
	}

	return c.people[i].GetPosition()
}
