package main

import (
	"fmt"
	"learngo/tree"
)
//定义别名
type myTreeNode struct {
	node *tree.TreeNode
}

//后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()

	right := myTreeNode{myNode.node.Right}
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	fmt.Println(root)

	//创建
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	fmt.Println(root)
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)

	fmt.Println("print\n")
	root.Print()

	fmt.Println("\nsetValue\n")
	root.Right.Left.SetValue(9)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
