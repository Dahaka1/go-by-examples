package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr []int
}

func (q *Queue) Push(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Shift() (int, error) {
	if len(q.arr) == 0 {
		return 0, errors.New("empty queue")
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	return val, nil
}

func isNodeExistsBFS(adj map[int][]int, node int, x int) (bool, int) {
	queue := Queue{}
	visited := make(map[int]interface{})

	queue.Push(node)

	for {
		val, err := queue.Shift()
		if err != nil {
			break
		}
		visited[node] = nil
		if val == x {
			return true, len(visited)
		}
		for i := 0; i < len(adj[node]); i++ {
			n := adj[node][i]
			if _, ok := visited[n]; !ok {
				visited[n] = nil
				queue.Push(n)
			}
		}
	}

	return false, 0
}

func isNodeExistsDFS(adj map[int][]int, visited map[int]interface{}, node int, x int) (bool, int) {
	if node == x {
		return true, len(visited)
	}
	if _, ok := visited[node]; ok {
		return false, 0
	}
	visited[node] = nil
	for i := 0; i < len(adj[node]); i++ {
		if _, ok := visited[adj[node][i]]; !ok {
			reached, steps := isNodeExistsDFS(adj, visited, adj[node][i], x)
			if reached {
				return reached, steps
			}
		}
	}
	return false, 0
}

func isNodeExists(adj map[int][]int, x int, alg string) (bool, int) {
	switch alg {
	case "BFS":
		return isNodeExistsBFS(adj, 1, x)
	case "DFS":
		return isNodeExistsDFS(adj, make(map[int]interface{}), 1, x)
	default:
		return false, 0
	}
}

func main() {
	g := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4, 5},
		4: {5},
		5: {},
	}
	x := 3
	res, smallestPath := isNodeExists(g, x, "BFS")
	switch res {
	case true:
		fmt.Printf("Node %d exists in graph\nSteps number from root: %d", x, smallestPath)
	}
}
