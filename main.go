package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AddQuery struct {
	T     int
	Value int
}

type MaxQuery struct {
	Anode int
	Bnode int
}

type TreeNode struct {
	NodeValue int
	EdgesWith []int
	Visited   bool
}

type Tree struct {
	RootNode TreeNode
	Nodes    []TreeNode
}

func CreateNode(nodeNumber, nodeValue int) *TreeNode {
	return &TreeNode{NodeValue: nodeValue}
}

func (tree *Tree) AddNode(node *TreeNode, index int) *Tree {
	tree.Nodes = append(tree.Nodes, *node)
	if index == 1 {
		tree.RootNode = *node
	}
	return tree
}

func (tNode *TreeNode) AddEdge(edgeNode int) *TreeNode {

	tNode.EdgesWith = append(tNode.EdgesWith, edgeNode)

	return tNode
}

func QueryMax(a, b int) {
	maxValue := tNode.NodeValue

	// Recursively find the maximum value in the paths to the neighbors
	for _, neighborNumber := range tNode.EdgesWith {
		if neighbor, exists := nodes[neighborNumber]; exists {
			maxValue = max(maxValue, neighbor.QueryMax())
		}
	}

	return maxValue
}

func (tNode *TreeNode) AddValue(t *Tree, value int) *TreeNode {

	if tNode.Visited {
		return tNode
	}

	tNode.NodeValue += value
	tNode.Visited = true

	for _, edge := range tNode.EdgesWith {
		edgeNode := t.Nodes[edge-1]
		edgeNode.AddValue(t, value)
		t.Nodes[edge-1] = edgeNode
	}

	return tNode

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	s := bufio.NewScanner(os.Stdin)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	numberOfNodes, _ := strconv.Atoi(lines[0])
	tree := Tree{}

	//Build Tree
	for i := 1; i <= numberOfNodes; i++ {
		node := CreateNode(i, 0)
		tree.AddNode(node, i+1)
	}

	var xyEdges [][]int
	for i := 1; i < numberOfNodes; i++ {
		edge := strings.Split(lines[i], " ")
		aIntEdge, _ := strconv.Atoi(edge[0])
		bIntEdge, _ := strconv.Atoi(edge[1])

		intEdges := []int{aIntEdge, bIntEdge}
		xyEdges = append(xyEdges, intEdges)
	}

	for _, edge := range xyEdges {
		tree.Nodes[edge[0]-1].AddEdge(edge[1])
		tree.Nodes[edge[1]-1].AddEdge(edge[0])
	}

	fmt.Println(tree)

	numbersOfQueries, _ := strconv.Atoi(lines[numberOfNodes])
	var addQueries []AddQuery
	var maxQueries []MaxQuery

	for x := numberOfNodes + 1; x <= (numbersOfQueries + numberOfNodes); x++ {
		query := strings.Split(lines[x], " ")
		a, _ := strconv.Atoi(query[1])
		b, _ := strconv.Atoi(query[2])
		if query[0] == "add" {
			addQuery := AddQuery{T: a, Value: b}
			addQueries = append(addQueries, addQuery)
		} else {
			maxQuery := MaxQuery{Anode: a, Bnode: b}
			maxQueries = append(maxQueries, maxQuery)
		}
	}

	fmt.Println(addQueries)
}
