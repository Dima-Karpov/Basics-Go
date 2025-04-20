package main

import "fmt"

const MAX_VERTEX_SIZE = 5

// Queue save current vertices
const QUEUESIZE = 40

type Queue struct {
	queue [QUEUESIZE]int
	head  int
	tail  int
}

var q *Queue = nil

func initQueue() {
	q = new(Queue)
	q.head = 0
	q.tail = 0
}

func isQueueEmpty() bool {
	if q.head == q.tail {
		return true
	} else {
		return false
	}
}

func enQueue(data int) bool {
	if q.tail == QUEUESIZE {
		fmt.Print("The queue was full and could not join.\n")
		return false
	}
	q.queue[q.tail] = data
	q.tail++
	return true
}

func deleteQueue() int {
	if q.head == q.tail {
		fmt.Print("The queue was empty and could not join.\n")
	}
	var data = q.queue[q.head]
	q.head++
	return data
}

// //// queue end //////////////////////////
type Vertex struct {
	data    string
	visited bool // Have you visited
}

var size = 0 //Current vertex size
var vertexs [MAX_VERTEX_SIZE]Vertex
var adjacencyMatrix [MAX_VERTEX_SIZE][MAX_VERTEX_SIZE]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertexs[size] = vertex
	size++
}

// Add adjacent adges
func addEdge(from int, to int) {
	// A -> B != B -> A
	adjacencyMatrix[from][to] = 1
}

// Clear reset
func clear() {
	for i := 0; i < size; i++ {
		vertexs[i].visited = false
	}
}

func breadthFirstSearch() {
	// Start searching from the first vertex
	vertexs[0].visited = true
	fmt.Printf("%s", vertexs[0].data)
	enQueue(0)

	var col int
	for {
		if isQueueEmpty() {
			break
		}
		var row = deleteQueue()
		// Get adjacent vertex positions that have not been visited
		col = findAdjacencyUnVisitedVertex(row)
		// Loop through all vertices connected to the current vertex
		for {
			if col == -1 {
				break
			}
			vertexs[col].visited = true
			fmt.Printf("%s", vertexs[col].data)
			enQueue(col)
			col = findAdjacencyUnVisitedVertex(row)
		}
	}
	clear()
}

// Get adjacent vertex positions that have not been visited
func findAdjacencyUnVisitedVertex(row int) int {
	for col := 0; col < size; col++ {
		if adjacencyMatrix[row][col] == 1 && !vertexs[col].visited {
			return col
		}
	}
	return -1
}

func printGraph() {
	fmt.Print("Two-dimensional array traversal vertex edge and adjacent array: \n ")
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
	}
	fmt.Print("\n")

	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
		for j := 0; j < MAX_VERTEX_SIZE; j++ {
			fmt.Printf("%d ", adjacencyMatrix[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
	initQueue()

	addVertex("A")
	addVertex("B")
	addVertex("C")
	addVertex("D")
	addVertex("E")

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(0, 3)
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(2, 3)
	addEdge(3, 4)

	//Two-dimensional array traversal output vertex edge and adjacent array
	printGraph()

	fmt.Print("\nBreadth-first search traversal output: \n")
	breadthFirstSearch()
}
