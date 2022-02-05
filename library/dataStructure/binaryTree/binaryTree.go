package binaryTree

import (
	"fmt"

	stack2 "github.com/qinsheng99/goWeb/library/dataStructure/stack"
)

type BinaryNode struct {
	val   int
	left  *BinaryNode
	right *BinaryNode
}

func NewBinary(val int) *BinaryNode {
	return &BinaryNode{val: val}
}
func NewBinaryNoVal() *BinaryNode {
	return nil
}

func Recursive1(node *BinaryNode) {
	//先序遍历  先打印头结点，在打印左节点，在打印右节点  递归方式
	if node == nil {
		return
	}
	fmt.Print(node.val)
	Recursive1(node.left)
	Recursive1(node.right)
}

func NoRecursive1(node *BinaryNode) {
	//先序遍历  先打印头结点，在打印左节点，在打印右节点 非递归   使用头右左压栈
	stack := stack2.CreateStack(20)
	if node != nil {
		stack.Push(node)
		for !stack.IsEmpty() {
			head := stack.Pop().(*BinaryNode)
			fmt.Print(head.val)
			if head.right != nil {
				stack.Push(head.right)
			}
			if head.left != nil {
				stack.Push(head.left)
			}
		}
	}
	defer stack.Clear()
}

func Recursive2(node *BinaryNode) {
	//中序遍历  先打印左结点，在打印头节点，在打印右节点 递归方式
	if node == nil {
		return
	}
	Recursive2(node.left)
	fmt.Print(node.val)
	Recursive2(node.right)

}

func NoRecursive2(node *BinaryNode) {
	//中序遍历  先打印左结点，在打印头节点，在打印右节点   左右头压栈
	stack := stack2.CreateStack(20)
	if node != nil {
		for !stack.IsEmpty() || node != nil {
			if node != nil {
				stack.Push(node)
				node = node.left
			} else {
				node = stack.Pop().(*BinaryNode)
				fmt.Print(node.val)
				node = node.right
			}
		}
	}
}

func Recursive3(node *BinaryNode) {
	//后序遍历:先打印左结点，在打印右节点，在打印头节点 递归方式
	if node == nil {
		return
	}
	Recursive3(node.left)
	Recursive3(node.right)
	fmt.Print(node.val)
}
func NoRecursive3(node *BinaryNode) {
	//后序遍历:先打印左结点，在打印右节点，在打印头节点 非递归方式 头右左使用头左右压栈
	stack := stack2.CreateStack(20)
	stack1 := stack2.CreateStack(20)
	if node != nil {
		stack.Push(node)
		for !stack.IsEmpty() {
			head := stack.Pop().(*BinaryNode)
			stack1.Push(head)
			if head.left != nil {
				stack.Push(head.left)
			}
			if head.right != nil {
				stack.Push(head.right)
			}
		}
		for !stack1.IsEmpty() {
			fmt.Print(stack1.Pop().(*BinaryNode).val)
		}
	}
	defer stack.Clear()
	defer stack1.Clear()
}

func BTest1() {
	var binary = NewBinary(1)
	binary.left = NewBinary(2)
	binary.right = NewBinary(3)
	binary.left.left = NewBinary(4)
	binary.left.right = NewBinary(5)
	binary.right.left = NewBinary(6)
	binary.right.right = NewBinary(7)

	Recursive1(binary)
}
