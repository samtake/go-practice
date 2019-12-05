package tree

import (
	"fmt"
)

//定义
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//树的遍历:给结构定义方法
func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

//值不会变
// func (node treeNode) setValue(value int) {
// 	node.value = value
// }

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = value
}

//工厂函数创建
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value} //这里返回的是局部变量的地址给外部使用，在go依旧是可以的
}
