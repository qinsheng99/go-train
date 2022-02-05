package nodeList

import "fmt"

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
