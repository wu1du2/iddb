package main

import (
	"iplan"
)

func main() {
	var pln iplan.PlanTreeNode
	pln.Left = 1
	pln.Right = 2
	println(pln.Left, pln.Right)

}
