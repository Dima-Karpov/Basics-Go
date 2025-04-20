package main

import "fmt"

const MAX_VERTEX_SIZE = 5
const STACKSIZE = 1000

type Vertex struct {
	data    string
	visited bool // Have you visited
}

// Stack saves current vertices
var top = -1
var stacks [STACKSIZE]int

func push(element int) {
	top++
	stacks[top] = element
}

func pop() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	top--

	return data
}

func peek() int {
	if top == -1 {
		return -1
	}

	var data = stacks[top]
	return data
}

func isEmpty() bool {
	if top <= -1 {
		return true
	}

	return false
}

var size = 0 // Current vertex size
var vertexs [MAX_VERTEX_SIZE]Vertex

// An array of topological sort results, recording the sorted sequence
// number of each node
var topologys [MAX_VERTEX_SIZE]Vertex
var adjacencyMatrix [MAX_VERTEX_SIZE][MAX_VERTEX_SIZE]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertexs[size] = vertex
	size++
}

// Add adjacent edges
func addEdge(from int, to int) {
	adjacencyMatrix[from][to] = 1
}

func removeVertex(vertex int) {
	if vertex != size-1 {
		// If the vertex is the last element, the end
		// The vertixes are removed from the vertex array
		for i := vertex; i < size-1; i++ {
			vertexs[i] = vertexs[i+1]
		}
		for row := vertex; row < size-1; row++ {
			for col := 0; col < size-1; col++ {
				adjacencyMatrix[row][col] = adjacencyMatrix[row+1][col]
			}
		}

		for col := vertex; col < size-1; col++ {
			for row := 0; row < size-1; row++ {
				adjacencyMatrix[row][col] = adjacencyMatrix[row][col+1]
			}
		}
	}
	size-- // Decrease the number of vertices
}

func topologySort() {
	for {
		if size <= 0 {
			break
		}
		// Get o no successor node
		var noSuccessorVertex = getNoSuccessorVertex()
		if noSuccessorVertex == -1 {
			fmt.Print("There is ring in Graph \n")
			return
		}
		// Copy the deleted node to the sorted array
		topologys[size-1] = vertexs[noSuccessorVertex]
		// Delete no successor node
		removeVertex(noSuccessorVertex)
	}
}

func getNoSuccessorVertex() int {
	var existSuccessor = false
	for row := 0; row < size; row++ {
		// For each vertex
		existSuccessor = false
		// If the node has a fixed row, each column has a 1, indicating that
		// the node has a successor, terminating the loop

		for col := 0; col < size; col++ {
			if adjacencyMatrix[row][col] == 1 {
				existSuccessor = true
				break
			}
		}
		if !existSuccessor {
			// If the node has no successor, return its subscript
			return row
		}
	}

	return -1
}

func printGraph() {
	fmt.Print("Two-dimensional array traversal vertex edge and adjacente array: \n")
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
	}
	fmt.Printf("\n")

	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
		for j := 0; j < MAX_VERTEX_SIZE; j++ {
			fmt.Printf("%d ", adjacencyMatrix[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
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

	// Two-dimensional array traversal output vertex edge and adjacent array
	printGraph()

	fmt.Print("\nDepth-First Search traversal output: \n")
	fmt.Print("\nDirected Graph topological Sorting: \n")
	topologySort()
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", topologys[i].data)
	}
}
