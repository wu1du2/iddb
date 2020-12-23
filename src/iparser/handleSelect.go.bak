package iparser

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"iplan"
	// "reflect"
)

//nodeid globally save the nodeid
var nodeid int64 = 1

//Tmptableid globally save the tmptableid
var Tmptableid int64 = 0

// var plantree iplan.PlanTree

//GetTmpTableName can get latest TmpTableName
func GetTmpTableName() (TmpTableName string) {
	TmpTableName = "Transaction_" + fmt.Sprintf("%d", TransactionID) + "_TmpTable_" + fmt.Sprintf("%d", Tmptableid)
	Tmptableid++
	return TmpTableName
}

//InitalPlanTreeNode init node
func InitalPlanTreeNode() (node iplan.PlanTreeNode) {
	node.Nodeid = -1
	node.Left = -1
	node.Right = -1
	node.Parent = -1
	node.Status = -1
	node.TmpTable = ""
	node.Locate = -1
	node.TransferFlag = false
	node.Dest = -1
	node.NodeType = -1
	node.Where = ""
	node.Cols = ""
	node.Joint_cols = ""
	return node
}

//InitalPlanTree init PlanTree
func InitalPlanTree() (planTree iplan.PlanTree) {
	for i := 0; i < iplan.MaxNodeNum; i++ {
		planTree.Nodes[i] = InitalPlanTreeNode()

	}
	planTree.NodeNum = 0
	return planTree
}

func createProjectionNode(TmpTableName string, Cols string) (node iplan.PlanTreeNode) {
	node = InitalPlanTreeNode()
	node.Nodeid = nodeid
	// nodeid++
	node.Status = 0
	node.TmpTable = TmpTableName
	node.NodeType = 3
	node.Cols = Cols

	node.Locate = 0
	node.TransferFlag = false

	return node
}

func createWhereNode(TmpTableName string, Where string) (node iplan.PlanTreeNode) {
	node = InitalPlanTreeNode()
	node.Nodeid = nodeid
	// nodeid++
	node.Status = 0
	node.TmpTable = TmpTableName
	node.NodeType = 2
	node.Where = Where

	node.Locate = 0
	node.TransferFlag = false

	return node
}

func createJoinNode(TmpTableName string) (node iplan.PlanTreeNode) {
	node = InitalPlanTreeNode()
	node.Nodeid = nodeid
	// nodeid++
	node.Status = 0
	node.TmpTable = TmpTableName
	node.NodeType = 4

	node.Locate = 0
	node.TransferFlag = false

	return node
}

func createTableNode(tablename string) (node iplan.PlanTreeNode) {
	node = InitalPlanTreeNode()
	node.Nodeid = nodeid
	// nodeid++
	node.Status = 1
	node.TmpTable = tablename
	node.NodeType = 1

	node.Locate = 0
	node.TransferFlag = false

	return node
}

//HandleSelect for handle select statment
func HandleSelect(sel *sqlparser.Select) iplan.PlanTree {
	planTree := InitalPlanTree()

	//handle projection; root node is projectionnode
	planTree.Nodes[nodeid] = createProjectionNode(GetTmpTableName(), sqlparser.String(sel.SelectExprs))
	nodeid++
	// println(planTree.Nodes[0].Cols)

	//handle where
	planTree.Nodes[nodeid] = createWhereNode(GetTmpTableName(), sqlparser.String(sel.Where))

	planTree.Nodes[nodeid].Parent = planTree.Nodes[nodeid-1].Nodeid
	planTree.Nodes[nodeid-1].Left = planTree.Nodes[nodeid].Nodeid
	nodeid++
	//handle join; only when there are more than 2 tables
	if len(sel.From) > 1 {
		planTree.Nodes[nodeid] = createJoinNode(GetTmpTableName())
		planTree.Nodes[nodeid].Parent = planTree.Nodes[nodeid-1].Nodeid
		planTree.Nodes[nodeid-1].Left = planTree.Nodes[nodeid].Nodeid
		nodeid++
	}

	//Handle from/tablenode
	switch len(sel.From) {
	case 1:
		// println(sqlparser.String(sel.From))
		createTableNode(sqlparser.String(sel.From))
	case 2:
		// println(sqlparser.String(sel.From[0]))
		planTree.Nodes[nodeid] = createTableNode(sqlparser.String(sel.From[0]))
		planTree.Nodes[nodeid].Parent = planTree.Nodes[nodeid-1].Nodeid
		planTree.Nodes[nodeid-1].Left = planTree.Nodes[nodeid].Nodeid
		nodeid++
		planTree.Nodes[nodeid] = createTableNode(sqlparser.String(sel.From[1]))
		planTree.Nodes[nodeid].Parent = planTree.Nodes[nodeid-2].Nodeid
		planTree.Nodes[nodeid-2].Right = planTree.Nodes[nodeid].Nodeid
		nodeid++
	case 3:
	case 4:
	default:
		println("The num of tables is bigger than 4, not support yet!")
	}
	// for i := 0; i < len(sel.From); i++ {
	// 	planTree.Nodes[nodeid] = createTableNode(sqlparser.String(sel.From[i]))
	// 	planTree.Nodes[nodeid].Parent = planTree.Nodes[nodeid-1].Nodeid
	// 	planTree.Nodes[nodeid-1].Left = planTree.Nodes[nodeid].Nodeid
	// 	nodeid++
	// }
	planTree.NodeNum = nodeid
	return planTree

}
