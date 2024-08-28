package main

import (
	"fmt"
	geomutil "github.com/denisstrizhkin/geomutil"
)

func main() {
	TestIsLeafNode()
}

func TestIsLeafNode() {
	// Our test tree
	//
	//     __1
	//    /   \
	//   2     3
	//  / \
	// 4   5
	//
	compfunc := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}
	tree := geomutil.NewBinTree(compfunc)
	tree.InsertNode(1)
	tree.InsertNode(2)
	tree.InsertNode(3)
	tree.InsertNode(4)
	tree.InsertNode(5)

	fmt.Println(tree.Search(1))

	val, err := tree.Search(3)
	val1, err1 := tree.Search(10)
	fmt.Println(val, err)
	fmt.Println(val1, err1)
}
