package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 使用自定义工厂函数
// 注意返回了局部变量的地址
// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

// 为结构定义方法  显示定义和命名方法接受者
// 定义结构体的方法 （node treeNode）接受者
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

// 遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}



// 值接受者 与 指针接受者
// 要改变内容必须使用指针接受者
// 结构过大也考虑使用指针接受者
// 一致性：如有指针接受者，最好都使用指针接受者
// 值接受者是 go 语言特有
// 值/指针接受者均可接受值/指针