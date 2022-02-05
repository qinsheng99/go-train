package binaryTree

import (
	"fmt"
	"math"

	"github.com/qinsheng99/goWeb/library/dataStructure/queue"
	stack2 "github.com/qinsheng99/goWeb/library/dataStructure/stack"
)

// BinaryNode 二叉树
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

func getWidthUseMap(node *BinaryNode) {
	// 获取二叉树的最大宽度
	var m = make(map[interface{}]int, 20)
	var max, level, nodes = -1, 1, 0
	q:= queue.NewQueue()
	q.Add(node)
	m[node] = 1
	for !q.IsEmpty() {
		head := q.Pop().(*BinaryNode)
		curLevel := m[head]
		if curLevel == level {
			nodes++
		} else {
			level++
			nodes = 1
		}
		max = int(math.Max(float64(max), float64(nodes)))
		if head.left != nil {
			m[head.left] = curLevel + 1
			q.Add(head.left)
		}

		if head.right != nil {
			m[head.right] = curLevel + 1
			q.Add(head.right)
		}
	}
	fmt.Println(max)
}

func getWidthNoUseMap(node *BinaryNode) {
	// 获取二叉树的最大宽度  不使用map
	var max, curLevel, curEnd, curNextEnd = 0, 0, node, NewBinaryNoVal()
	q := queue.NewQueue()
	q.Add(node)
	for !q.IsEmpty() {
		head := q.Pop().(*BinaryNode)
		if head.left != nil {
			q.Add(head.left)
			curNextEnd = head.left
		}
		if head.right != nil {
			q.Add(head.right)
			curNextEnd = head.right
		}
		curLevel++
		if head == curEnd {
			max = int(math.Max(float64(max), float64(curLevel)))
			curLevel = 0
			curEnd = curNextEnd
		}
	}
	fmt.Println(max)
}
func souSuoBinaryTree(node *BinaryNode) bool {
	//搜索二叉树：根节点左子树不为空，那么它左子树上面的所有节点的值都小于它的根节点的值，如果它的右子树不为空，那么它右子树任意节点的值都大于他的根节点的值
	var MIN = math.MinInt
	if node == nil {
		return true
	}
	stack := stack2.CreateStack(20)
	for !stack.IsEmpty() || node != nil {
		if node != nil {
			stack.Push(node)
			node = node.left
		} else {
			node = stack.Pop().(*BinaryNode)
			if node.val <= MIN {
				return false
			} else {
				MIN = node.val
			}
			node = node.right
		}
	}
	return true
	//if node == nil {
	//	return true
	//}
	//i := souSuoBinaryTree(node.left)
	//if !i {
	//	return false
	//}
	//if node.val <= MIN {
	//	return false
	//} else {
	//	MIN = node.val
	//}
	//return souSuoBinaryTree(node.right)
}
type SouSuo struct {
	res      bool
	min, max int
}

func souSuoBinaryTree1(node *BinaryNode) *SouSuo {
	// 递归实现搜索二叉树
	// 左右都是搜索二叉树，左边最大值小于根，右边最小值大于根
	if node == nil {
		return nil
	}
	leftS := souSuoBinaryTree1(node.left)
	rightS := souSuoBinaryTree1(node.right)

	min, max := node.val, node.val
	if leftS != nil {
		min = int(math.Min(float64(min), float64(leftS.min)))
		max = int(math.Max(float64(max), float64(leftS.max)))
	}

	if rightS != nil {
		min = int(math.Min(float64(min), float64(rightS.min)))
		max = int(math.Max(float64(max), float64(rightS.max)))
	}
	var res = true

	if leftS != nil && (!leftS.res || leftS.max >= node.val) {
		res = false
	}

	if rightS != nil && (!rightS.res || rightS.min <= node.val) {
		res = false
	}

	return &SouSuo{res: res, max: max, min: min}
}
func CBT(node *BinaryNode) bool {
	/**
			1、一个节点，左孩子为空，右孩子不为空，不是满二叉树
			2、一个节点左右都空，或者左不为空右为空，剩下的节点必须都是子节点(不存在左右孩子)
	 */
	var b = false // 是否遇到第一个孩子不全的节点
	q := queue.NewQueue()
	q.Add(node)
	var l, r = NewBinaryNoVal(), NewBinaryNoVal()
	for !q.IsEmpty() {
		head := q.Pop().(*BinaryNode)
		l = head.left
		r = head.right
		if (b && (l != nil || r != nil)) || (l == nil && r != nil) {
			return false
		}
		if l != nil {
			q.Add(head.left)
		}
		if r != nil {
			q.Add(head.right)
		}
		if l == nil || r == nil {
			b = true
		}
	}
	return true
}
func balance(node *BinaryNode) (res bool, high int) {
	// 判断是否是平衡二叉树
	if node == nil {
		return true, 0
	}

	leftB, leftH := balance(node.left)
	rightB, rightH := balance(node.right)
	high = int(math.Max(float64(leftH), float64(rightH)) + 1)
	// 左数是平衡二叉树  右树是平衡二叉树  左右高度差小于2
	res = leftB && rightB && int(math.Abs(float64(leftH-rightH))) < 2
	return
}

func fullBinaryTree(node *BinaryNode) (high, nodes int) {
	// 判断是否是满二叉树
	//int(math.Pow(2, float64(high)))-1 == nodes
	if node == nil {
		return 0, 0
	}
	leftH, leftN := fullBinaryTree(node.left)
	rightH, rightN := fullBinaryTree(node.right)

	nodes = leftN + rightN + 1
	high = int(math.Max(float64(leftH), float64(rightH))) + 1
	return
}

func lowAncestors(node, o1, o2 *BinaryNode) *BinaryNode {
	//找到两个点的最低公共祖先
	if node == nil || node == o1 || node == o2 {
		return node
	}
	left := lowAncestors(node.left, o1, o2)
	right := lowAncestors(node.right, o1, o2)
	//o1和o2不互为祖先，返回交点
	if left != nil && right != nil {
		return node
	}

	// o1和o2，有一个是另一个的祖先，那么另一面一定返回nil
	if left != nil {
		return left
	} else {
		return right
	}
}

func subsequent(node *NBinaryNode) *NBinaryNode {
	//找后继节点：中序遍历每个节点的下一个节点
	if node == nil {
		return node
	}
	if node.right != nil { //有右子树，则右子树的左子树为后继节点
		return f(node.right)
	} else {
		parent := node.parent
		for parent.left != node && parent != nil {
			node = parent
			parent = node.parent
		}
		return parent
	}
}

func f(node *NBinaryNode) *NBinaryNode {
	if node == nil {
		return node
	}
	for node.left != nil {
		node = node.left
	}
	return node
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

func BTest2() {
	var binary = NewBinary(10)
	binary.left = NewBinary(8)
	binary.right = NewBinary(15)
	binary.left.left = NewBinary(5)
	binary.left.left.right = NewBinary(6)
	binary.left.right = NewBinary(9)
	//binary.right.left = NewBinary(1)
	binary.right.right = NewBinary(17)
	binary.right.right.left = NewBinary(16)

	NoRecursive2(binary)
	fmt.Println(souSuoBinaryTree(binary))
}

func BTest3() {
	var binary = NewBinary(1)
	binary.left = NewBinary(2)
	binary.right = NewBinary(3)
	binary.left.left = NewBinary(4)
	binary.left.left.left = NewBinary(4)
	binary.left.right = NewBinary(5)
	binary.right.left = NewBinary(6)
	//binary.right.right = NewBinary(7)
	fmt.Println(CBT(binary))
}
func BTest4() {
	var binary = NewBinary(1)
	binary.right = NewBinary(2)
	binary.right.right = NewBinary(3)
	//binary.left.left = NewBinary(4)
	//binary.left.left.left = NewBinary(4)
	//binary.left.right = NewBinary(5)
	//binary.right.left = NewBinary(6)
	//binary.right.right = NewBinary(7)
	fmt.Println(balance(binary))
}

func BTest5() {
	var binary = NewBinary(1)
	binary.left = NewBinary(2)
	binary.right = NewBinary(3)
	binary.left.left = NewBinary(4)
	binary.left.right = NewBinary(5)
	binary.right.left = NewBinary(6)
	fmt.Println(lowAncestors(binary, binary.left.left, binary.left))
}

func BTest6() {
	var binary = NewNBinary(1)
	binary.parent = nil
	binary.left = NewNBinary(2)
	binary.left.parent = binary
	binary.left.left = NewNBinary(4)
	binary.left.left.parent = binary.left
	binary.left.right = NewNBinary(5)
	binary.left.right.parent = binary.left
	binary.left.right.left = NewNBinary(10)
	binary.left.right.left.parent = binary.left.right
	binary.left.right.right = NewNBinary(100)
	binary.left.right.right.parent = binary.left.right

	binary.right = NewNBinary(3)
	binary.right.parent = binary
	binary.right.left = NewNBinary(6)
	binary.right.left.parent = binary.right

	fmt.Println(subsequent(binary.left.right.right))
}
