package main

// FIFO
type Queue struct {
	items []Node
}

func (q *Queue) enqueue(n Node) {
	q.items = append(q.items, n)
}

func (q *Queue) appendToQueue(n []Node) {
	q.items = append(q.items, n...)
}

func (q *Queue) dequeue() Node {
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

type Node string

type SimpleGraph map[Node][]Node

// Where we have a specific goal (example on page 96)
func (g SimpleGraph) BreadthFirstSearchSpecificGoal(start Node, isGoal func(Node) bool) []Node {
	var q = Queue{}
	var path []Node
	q.enqueue(start)

	visited := make(map[Node]bool)
	predecessor := make(map[Node]Node)

	for !q.isEmpty() {
		item := q.dequeue()

		if isGoal(item) {
			for pred := item; pred != start; pred = predecessor[pred] {
				path = append([]Node{pred}, path...)
			}
			path = append([]Node{start}, path...)
			return path
		}

		for _, succ := range g[item] {
			if visited[succ] {
				continue
			}
			predecessor[succ] = item
			q.enqueue(succ)
			visited[succ] = true
		}
	}
	return nil

}

// return how many steps it took or the actual path
// then see which is the shortest path
