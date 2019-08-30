package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 使用自定义工厂函数
// 注意返回了局部变量的地址
// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 为结构定义方法  显示定义和命名方法接受者
// 定义结构体的方法 （node treeNode）接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

// 遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 不论地址还是结构本身，一律使用 . 来访问成员
func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	root.traverse()

	root.print()
	root.setValue(100)
	root.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var proot1 *treeNode
	proot1.setValue(200)
	proot1 = &root
	proot1.setValue(300)
	proot1.print()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}

// 值接受者 与 指针接受者
// 要改变内容必须使用指针接受者
// 结构过大也考虑使用指针接受者
// 一致性：如有指针接受者，最好都使用指针接受者
// 值接受者是 go 语言特有
// 值/指针接受者均可接受值/指针