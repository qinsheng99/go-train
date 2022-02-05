package nodeList

import (
	"fmt"
	"math"
)

// Node 链表
type Node struct {
	val int
	next *Node
}

func newNode(val int) *Node {
	return &Node{val: val}
}
func newNodeNil() *Node {
	return nil
}

func NodeList() bool {
	var node = newNode(1)
	node.next = newNode(2)
	node.next.next = newNode(3)
	node.next.next.next = newNode(2)
	node.next.next.next.next = newNode(1)
	return listPartition1(node)
}

func listPartition1(node *Node) (res bool)  {
	res = true
	var n1, n2 = node, node
	for n2.next != nil && n2.next.next != nil {
		n1 = n1.next		//mid
		n2 = n2.next.next	//end
	}

	n2 = n1.next  // 2 1
	n1.next = nil
	n3 := newNodeNil()

	for n2 != nil {
		n3 = n2.next
		n2.next = n1
		n1 = n2
		n2 = n3
	}

	n3 = n1
	n2 = node

	for n1 != nil && n2 != nil { // check palindrome
		if n1.val != n2.val {
			res = false
			break
		}
		n1 = n1.next // left to mid
		n2 = n2.next // right to mid
	}

	n1 = n3.next
	n3.next = nil

	for n1 != nil {
		n2 = n1.next
		n1.next = n3
		n3 = n1
		n1 = n2
	}
	return
}

func listPartition2(node *Node, pivot int) *Node {
	// 给定一个数，将小于，等于，大于的数分开放
	var sH, sT, eH, eT, bH, bT = newNodeNil(), newNodeNil(), newNodeNil(), newNodeNil(), newNodeNil(), newNodeNil()
	for node != nil {
		Next := node.next
		node.next = nil
		if node.val < pivot {
			if sH == nil {
				sH = node
				sT = node
			} else {
				sT.next = node // sH和sT地址一样，这一步相当于是给sH的Next赋值
				sT = node      // sT赋值后，和sH内存地址发生变化
			}
		} else if node.val > pivot {
			if bH == nil {
				bH = node
				bT = node
			} else {
				bT.next = node
				bT = node
			}
		} else {
			if eH == nil {
				eH = node
				eT = node
			} else {
				eT.next = node
				eT = node
			}
		}
		node = Next
	}
	fmt.Println(sH, sT, eH, eT, bH, bT)
	// 小于区域的尾巴，连等于区域的头，等于区域的尾巴连大于区域的头
	if sT != nil {
		sT.next = eH
		if eT == nil {
			eT = sT
		} else {
		}
	}

	if eT != nil {
		eT.next = bH
	}
	if sH != nil {
		return sH
	}
	if eH != nil {
		return eH
	} else {
		return bH
	}
}

func getIntersectNode(node1, node2 *Node) *Node {
	// 给定两个点，判断是否相交
	loop1 := isLoop(node1)
	loop2 := isLoop(node2)

	if loop1 == nil && loop2 == nil {
		return noLoop(node1, node2)
	}

	if loop1 != nil && loop2 != nil {
		return bothLoop(node1, loop1, node2, loop2)
	}
	return nil
}

func noLoop(node1, node2 *Node) *Node {
	var cur1, cur2, n = node1, node2, 0

	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}

	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}
	if cur1 != cur2 {
		return nil
	}

	if n > 0 { //谁长，谁的头变成cur1
		cur1 = node1
	} else {
		cur1 = node2
	}
	// 谁短，谁的头变成cur2
	if cur1 == node1 {
		cur2 = node2
	} else {
		cur2 = node1
	}

	n = int(math.Abs(float64(n)))
	for n != 0 {
		n--
		cur1 = cur1.next
	}

	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

func bothLoop(node1, loop1, node2, loop2 *Node) *Node {
	var cur1, cur2, n = newNodeNil(), newNodeNil(), 0
	if loop1 == loop2 {
		cur1 = node1
		cur2 = node2
		for cur1.next != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2.next != loop2 {
			n--
			cur2 = cur2.next
		}
		if n > 0 { //谁长，谁的头变成cur1
			cur1 = node1
		} else {
			cur1 = node2
		}
		// 谁短，谁的头变成cur2
		if cur1 == node1 {
			cur2 = node2
		} else {
			cur2 = node1
		}
		n = int(math.Abs(float64(n)))
		for n != 0 {
			n--
			cur1 = cur1.next
		}

		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}

func isLoop(node *Node) *Node {
	//是否成环
	if node == nil || node.next == nil || node.next.next == nil {
		return nil
	}
	var s, f = node.next, node.next.next
	for s != f {
		s = s.next
		f = f.next.next
		if f.next == nil || f.next.next == nil {
			return nil
		}
	}
	f = node
	for s != f {
		s = s.next
		f = f.next
	}
	return s
}
