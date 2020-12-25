package iparser

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"iplan"
	"os"
	"strings"
	// "reflect"
)

//nodeid globally save the nodeid
// var nodeid int64 = 1

//Tmptableid globally save the tmptableid
var Tmptableid int64 = 0

//GetTmpTableName can get latest TmpTableName
func GetTmpTableName() (TmpTableName string) {
	TmpTableName = "Transaction_" + fmt.Sprintf("%d", TransactionID) + "_TmpTable_" + fmt.Sprintf("%d", Tmptableid)
	Tmptableid++
	return TmpTableName
}

//ResetCols reset cols to get a unique colname
func ResetCols(strin string) (strout string) {
	strout = ""
	f := func(c rune) bool {
		if c == ' ' || c == ',' {
			return true
		}
		return false
	}
	arr := strings.FieldsFunc(strin, f)
	for i, str := range arr {
		switch str {
		case "publisher.id":
			arr[i] = "pid"
		case "publisher.name":
			arr[i] = "pname"
		case "publisher.nation":
			arr[i] = "nation"
		case "book.id":
			arr[i] = "bid"
		case "book.title":
			arr[i] = "title"
		case "book.authors":
			arr[i] = "authors"
		case "book.publisher_id":
			arr[i] = "bpid"
		case "book.copies":
			arr[i] = "copies"
		case "customer.id":
			arr[i] = "cid"
		case "customer.name":
			arr[i] = "cname"
		case "customer.rank":
			arr[i] = "rank"
		case "orders.customer_id":
			arr[i] = "ocid"
		case "orders.book_id":
			arr[i] = "obid"
		case "orders.quantity":
			arr[i] = "quantity"
		}
	}
	for _, str := range arr {
		strout += str + " "
	}
	return strout
}

var logicalPlanTree iplan.PlanTree
var root int64

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
	node.Rel_cols = ""
	node.Cols = ""
	node.Joint_type = -1
	node.Joint_cols = ""
	return node
}

//InitalPlanTree init PlanTree
func InitalPlanTree() (planTree iplan.PlanTree) {
	for i := 0; i < iplan.MaxNodeNum; i++ {
		planTree.Nodes[i] = InitalPlanTreeNode()

	}
	planTree.Root = -1
	planTree.NodeNum = 0
	return planTree
}

//findEmptyNode will return the idx of first empty node
func findEmptyNode() int64 {
	for i, node := range logicalPlanTree.Nodes {
		if i != 0 && node.Nodeid == -1 {
			return int64(i)
		}
	}
	println("Error when creating node, no empty node left!")
	return -1
}

//AddTableNode add table node
func AddTableNode(newNode iplan.PlanTreeNode) {
	if logicalPlanTree.NodeNum == 0 {
		root = findEmptyNode()
		newNode.Nodeid = root
		logicalPlanTree.Nodes[root] = newNode
		logicalPlanTree.NodeNum++
	} else {
		pos := findEmptyNode()
		newNode.Nodeid = pos
		logicalPlanTree.Nodes[pos] = newNode
		logicalPlanTree.NodeNum++

		newroot := findEmptyNode()
		logicalPlanTree.Nodes[newroot] = CreateJoinNode(GetTmpTableName())
		logicalPlanTree.NodeNum++

		logicalPlanTree.Nodes[newroot].Nodeid = newroot
		logicalPlanTree.Nodes[newroot].Left = root
		logicalPlanTree.Nodes[newroot].Right = pos
		logicalPlanTree.Nodes[pos].Parent = newroot
		logicalPlanTree.Nodes[root].Parent = newroot
		root = newroot

	}

}

//AddSelectionNode add selection node
func AddSelectionNode(newNode iplan.PlanTreeNode) {
	newroot := findEmptyNode()
	newNode.Nodeid = newroot
	newNode.Left = root
	logicalPlanTree.Nodes[newroot] = newNode
	logicalPlanTree.NodeNum++
	logicalPlanTree.Nodes[root].Parent = newroot
	root = newroot
}

//AddProjectionNode add projection node
func AddProjectionNode(newNode iplan.PlanTreeNode) {
	newroot := findEmptyNode()
	newNode.Nodeid = newroot
	newNode.Left = root
	logicalPlanTree.Nodes[newroot] = newNode
	logicalPlanTree.NodeNum++
	logicalPlanTree.Nodes[root].Parent = newroot
	root = newroot
}

//CreateTableNode create table node
func CreateTableNode(tableName string) iplan.PlanTreeNode {
	node := InitalPlanTreeNode()
	node.NodeType = 1
	node.TmpTable = tableName

	return node
}

//CreateSelectionNode create selection nnode
func CreateSelectionNode(TmpTableName string, where string) iplan.PlanTreeNode {
	node := InitalPlanTreeNode()
	node.NodeType = 2
	node.TmpTable = TmpTableName
	node.Where = where
	return node
}

//CreateProjectionNode create projection node
func CreateProjectionNode(TmpTableName string, cols string) iplan.PlanTreeNode {
	node := InitalPlanTreeNode()
	node.NodeType = 3
	node.TmpTable = TmpTableName
	node.Cols = cols
	return node
}

//CreateJoinNode create join node
func CreateJoinNode(TmpTableName string) iplan.PlanTreeNode {
	node := InitalPlanTreeNode()
	node.NodeType = 4
	node.TmpTable = TmpTableName
	return node
}

//CreateUnionNode create union node
func CreateUnionNode(TmpTableName string) iplan.PlanTreeNode {
	node := InitalPlanTreeNode()
	node.NodeType = 5
	node.TmpTable = TmpTableName

	return node
}

//buildSelect for handle select statment
func buildSelect(sel *sqlparser.Select) iplan.PlanTree {
	logicalPlanTree = InitalPlanTree()
	if sel.From == nil {
		fmt.Println("cannot build plan tree without From")
		os.Exit(1)
	}
	for _, table := range sel.From {
		tableName := sqlparser.String(table)
		AddTableNode(CreateTableNode(tableName))
	}

	if sel.Where != nil {
		whereString := sqlparser.String(sel.Where.Expr)
		whereString = ResetCols(whereString)
		AddSelectionNode(CreateSelectionNode(GetTmpTableName(), whereString))
	}

	if sel.SelectExprs == nil {
		fmt.Println("cannot build plan tree without select")
		os.Exit(1)
	}
	projectionString := sqlparser.String(sel.SelectExprs)
	projectionString = ResetCols(projectionString)
	AddProjectionNode(CreateProjectionNode(GetTmpTableName(), projectionString))
	logicalPlanTree.Root = root
	return logicalPlanTree
}

func buildInsert(sel *sqlparser.Insert) {
	fmt.Println(sqlparser.String(sel.Table))
	fmt.Println(sqlparser.String(sel.Columns))
	fmt.Println(sqlparser.String(sel.Rows))
}
