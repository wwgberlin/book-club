package bfs

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
func (g SimpleGraph) BreadthFirstSearchSpecificGoal(start Node, isGoal func(Node) bool) bool {
	var q = Queue{}
	q.enqueue(start)

	var visited = []Node{}

	for !q.isEmpty() {
		item := q.dequeue()
		// check if node has already been visited
		for _, v := range visited {
			if v == item {
				continue
			}
		}
		// add to visited
		visited = append(visited, item)
		if isGoal(item) {
			return true

		}
		q.appendToQueue(g[item])
	}
	return false

}

// return how many steps it took or the actual path
// then see which is the shortest path
