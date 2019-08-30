package main

import (
	"fmt"
	"imooc.com/ccmouse/learngo/tree"
)

// 不论地址还是结构本身，一律使用 . 来访问成员
func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()

	root.Print()
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var proot1 *tree.TreeNode
	proot1.SetValue(200)
	proot1 = &root
	proot1.SetValue(300)
	proot1.Print()

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
