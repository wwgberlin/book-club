package bfs

import (
	"testing"
)

func TestBreadthFirstSearchSpecificGoalExampleOne(t *testing.T) {
	start := Node("start")
	goal := Node("end")

	var isGoal = func(n Node) bool {
		return n == goal
	}

	A := Node("A")
	B := Node("B")
	C := Node("C")
	D := Node("D")
	E := Node("E")

	sg := SimpleGraph{
		start: []Node{A, B},
		A:     []Node{D},
		B:     []Node{C, E},
		C:     []Node{D},
		E:     []Node{D},
		D:     []Node{goal},
	}

	result := sg.BreadthFirstSearchSpecificGoal(start, isGoal)
	if !result {
		t.Errorf("This test has failed!")
	}
}

func TestBreadthFirstSearchSpecificGoalExampleTwo(t *testing.T) {
	start := Node("start")

	var isGoal = func(name Node) bool {
		// person is a mango seller if their name ends with the letter `m`
		return name[len(name)-1] == 'm'
	}

	Alice := Node("Alice")
	Claire := Node("Claire")
	Bob := Node("Bob")
	Jonny := Node("Jonny")
	Thom := Node("Thom")
	Peggy := Node("Peggy")
	Anuj := Node("Anuj")

	sg := SimpleGraph{
		start:  []Node{Alice, Bob, Claire},
		Alice:  []Node{Peggy},
		Claire: []Node{Thom, Jonny},
		Bob:    []Node{Anuj},
	}

	result := sg.BreadthFirstSearchSpecificGoal(start, isGoal)
	if !result {
		t.Errorf("This test has failed!")
	}
}
