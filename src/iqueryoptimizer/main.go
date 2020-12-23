package iqueryoptimizer

import (
	"iplan"
)

//nodeid globally save the nodeid

func Optimize(oldtree iplan.PlanTree) iplan.PlanTree {
	loopmax := 1
	var newtree iplan.PlanTree
	newtree = oldtree
	for i := 0; i < loopmax; i++ {
		newtree = JointConditionPushDown(newtree)
		newtree = SelectionPushDown(newtree)
		newtree = ProjectionPushDown(newtree)
		newtree = TransmissionMinimization(newtree)
	}
	return newtree
}

func JointConditionPushDown(oldtree iplan.PlanTree) iplan.PlanTree {
	var newtree iplan.PlanTree
	newtree = oldtree
	return newtree
}

func SelectionPushDown(oldtree iplan.PlanTree) iplan.PlanTree {
	var newtree iplan.PlanTree
	newtree = oldtree
	return newtree
}

func ProjectionPushDown(oldtree iplan.PlanTree) iplan.PlanTree {
	var newtree iplan.PlanTree
	newtree = oldtree
	return newtree
}

func TransmissionMinimization(oldtree iplan.PlanTree) iplan.PlanTree {
	var newtree iplan.PlanTree
	newtree = oldtree
	return newtree
}
