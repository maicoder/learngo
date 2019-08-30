package main

import (
	"fmt"
	"imooc.com/ccmouse/learngo/tree"
)

// 通过组合扩展系统或别人写的功能
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

// 不论地址还是结构本身，一律使用 . 来访问成员
func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	root.Print()
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var proot1 *tree.Node
	proot1.SetValue(200)
	proot1 = &root
	proot1.SetValue(300)
	proot1.Print()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
