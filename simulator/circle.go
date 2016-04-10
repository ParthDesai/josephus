package simulator

import "fmt"

const (
	notstarted = iota
	executing  = iota
	completed  = iota
)

// Result is used to indicate the result of the simulation.
// LastAlive referes to the position of the person, who will be lastalive.
// KillingOrder referes to the position, in which the killing  was performed.
type Result struct {
	LastAlive    int
	KillingOrder []int
}

// CircleOfDeath represents a circle in josephus problem.
// This object is responsible to run the simulation, and return the result.
type CircleOfDeath struct {
	people      []Person
	stepPerKill int
}

// Init initializes the CircleOfDeath.
// It takes two argument, first is the number of people in the circle
// Second is steps per kill.
func (c *CircleOfDeath) Init(numberOfPeople uint64, stepPerKill int) {

	if numberOfPeople < 1 {
		panic(fmt.Errorf("Number of people should be greater than or equal to 1. Found:%d", numberOfPeople))
	}

	if uint64(stepPerKill) >= numberOfPeople {
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

// Execute executes the simulation
// It returns Result struct instance.
func (c *CircleOfDeath) Execute() Result {

	i := c.findFirstAlive()

	result := Result{}
	result.KillingOrder = make([]int, 0)

	if i == -1 {
		panic(fmt.Errorf("Inconsistent state. No body is alive."))
	}

	for j := c.findNextAlive(i); i != j; {
		c.people[j].Kill()

		result.KillingOrder = append(result.KillingOrder, j)

		i = c.findNextAlive(i)
		j = c.findNextAlive(i)
	}

	result.LastAlive = c.people[i].GetPosition()

	return result
}
