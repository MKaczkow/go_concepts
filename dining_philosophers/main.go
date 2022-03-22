package main

import (
	"fmt"
	"sync"
	"time"
)

const eat_rounds = 10
const count = 5
var eatGroup sync.WaitGroup

type Fork struct {
	id int
	sync.Mutex
}

type Philosopher struct {
	id int
	name string
	leftFork *Fork 
	rightFork *Fork 
}

	// set up N (N==5 at first) philosophers

	// start randomly generating their "hunger"

	// create verbose description of what's happening

	// allow graceful stop

	// allow parametrization

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func (p Philosopher) eat() {
	defer eatGroup.Done()
	for j := 0; j < eat_rounds; j++ {
		p.leftFork.Lock()
		p.rightFork.Lock()

		say("eating", p.id)
		time.Sleep(time.Second)

		p.rightFork.Unlock()
		p.leftFork.Unlock()

		say("finished eating", p.id)
		time.Sleep(time.Second)
	}
}

func set_table() (frks []*Fork, philos []*Philosopher) {

	// Create forks
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = &Fork{id: i}
	}

	// Create philosophers
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
	}

	return forks, philosophers
}

func arbitrator() {

}

func resource_hierarchy() {

}

func chandy_misra() {

}

func classic() {
	fmt.Println("Table is set...")
	fmt.Println("Philosophers start to think and they are getting hungry...")

	// Create forks
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		// forks[i] = new(Fork)
		forks[i] = &Fork{id: i}
	}

	// Create philosophers, assign them 2 forks and send them to the dining table
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
		eatGroup.Add(1)
		go philosophers[i].eat()
	}
	eatGroup.Wait()
}

func main() {
	classic()
}