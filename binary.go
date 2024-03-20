package main

import "fmt"

type myTreeNode struct {
	Data  int
	Left  *myTreeNode
	Right *myTreeNode
}

func createNode(data int) *myTreeNode {
	return &myTreeNode{Data: data}
}

func (n *myTreeNode) getTreeNodNum() int {
	if n == nil {
		return 0
	}
	return 1 + n.Left.getTreeNodNum() + n.Right.getTreeNodNum()
}

func (n *myTreeNode) getTreeDepth() int {
	maxDepth := 0
	if n == nil {
		return maxDepth
	}

	leftDepth := n.Left.getTreeDepth()
	rightDepth := n.Right.getTreeDepth()

	if leftDepth > rightDepth {
		maxDepth = leftDepth
	} else {
		maxDepth = rightDepth
	}

	return 1 + maxDepth

}

func initTree() *myTreeNode {
	root := createNode(10)
	root.Left = createNode(20)
	root.Right = createNode(30)
	root.Left.Left = createNode(40)
	root.Left.Right = createNode(50)
	root.Right.Left = createNode(60)
	root.Right.Right = createNode(70)
	return root
}

func binary() {
	root := initTree()

	nodeNumber := root.getTreeNodNum()
	depth := root.getTreeDepth()

	fmt.Println("Number of nodes in the tree: ", nodeNumber)
	fmt.Println("Depth of the tree: ", depth)

}
