package main

import (
	"iparser"
)

func main() {
	var pln iparser.PlanTreeNode
	pln.Left = 1
	pln.Right = 2
	println(pln.Left, pln.Right)

}
