package binaryTree

type NBinaryNode struct {
	val    int
	left   *NBinaryNode
	right  *NBinaryNode
	parent *NBinaryNode
}

func NewNBinary(val int) *NBinaryNode {
	return &NBinaryNode{val: val}
}
func NewNBinaryNoVal() *NBinaryNode {
	return nil
}
