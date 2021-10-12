package main

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
		A:     []Node{D, start},
		B:     []Node{C, E},
		C:     []Node{D},
		E:     []Node{D},
		D:     []Node{goal},
	}

	result := sg.BreadthFirstSearchSpecificGoal(start, isGoal)
	if len(result) != 4 {
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
	if len(result) != 3 {
		t.Errorf("This test has failed!")
	}
}

func TestBreadthFirstSearchSpecificGoalNoPath(t *testing.T) {
	start := Node("start")
	goal := Node("end")

	var isGoal = func(n Node) bool {
		return false
	}

	A := Node("A")
	B := Node("B")
	C := Node("C")
	D := Node("D")
	E := Node("E")

	sg := SimpleGraph{
		start: []Node{A, B},
		A:     []Node{D, start},
		B:     []Node{C, E},
		C:     []Node{D},
		E:     []Node{D},
		D:     []Node{goal},
	}

	result := sg.BreadthFirstSearchSpecificGoal(start, isGoal)
	if result != nil {
		t.Errorf("This test has failed!")
	}
}

func TestBreadthFirstSearchSpecificGoalOnePathToRuleThemAll(t *testing.T) {
	start := Node("start")
	goal := Node("end")

	var isGoal = func(n Node) bool {
		return n == start
	}

	A := Node("A")
	B := Node("B")
	C := Node("C")
	D := Node("D")
	E := Node("E")

	sg := SimpleGraph{
		start: []Node{A, B},
		A:     []Node{D, start},
		B:     []Node{C, E},
		C:     []Node{D},
		E:     []Node{D},
		D:     []Node{goal},
	}

	result := sg.BreadthFirstSearchSpecificGoal(start, isGoal)
	if len(result) != 1 {
		t.Errorf("This test has failed!")
	}
}
