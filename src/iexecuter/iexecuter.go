package iexecuter

import (
	"fmt"
	"iplan"
	// "database/sql"
	"strconv"
	"strings"
    _ "github.com/go-sql-driver/mysql"
)

var site int
var mysql_user string
var mysql_passwd string
var mysql_db string
var mysql_ip_port string
func RunExecuter(txn_id int64) int64 {
	// get the plan through txn_id
	//
	println("executed successfully")
	return 0
}

func ExecuteCreateStmt(stmt string) int64 {
	println(stmt)
	return 0
}

func Init() {
	// TODO:set site id
	site = 1
	mysql_user = "root"
	mysql_passwd = "123456"
	mysql_db = "test"
	mysql_ip_port = "127.0.0.1:3306"
}

func RunTree(plan_tree iplan.PlanTree) int64 {
	Init()
	for {
		// TODO:获取etcd tree

		// 检测树是否完全执行完毕
		if (TreeIsComplete(plan_tree)) {
			break
		}
		// 从tree中找可用代码
		execute_id := FindOneNode(plan_tree, 0)
		fmt.Println(execute_id)
		if (execute_id == -1) {
			continue
		}
		var pn *iplan.PlanTreeNode
		pn = &plan_tree.Nodes[execute_id]
		// 执行某个节点
		ExecuteOneNode(pn, plan_tree)
		print("executed node")
		println(pn.Nodeid)
		print("current node state ")
		var current_node iplan.PlanTreeNode
		current_node = plan_tree.Nodes[execute_id]
		fmt.Println(current_node)
		// TODO:更新etcd状态
	}
	println("executed successfully")
	return 0
}

func TreeIsComplete(plan_tree iplan.PlanTree) bool {
	var root_node iplan.PlanTreeNode
	root_node = plan_tree.Nodes[0]
	if (root_node.Status == 1) {
		return true
	} else {
		return false
	}
}

func FindOneNode(plan_tree iplan.PlanTree, node_id int) int {
	// fmt.Println("go into node")
	// fmt.Println(node_id)
	var can_execute_id int
	// 判断当前节点的状态，如果是ok，则返回-1
	var current_node iplan.PlanTreeNode
	current_node = plan_tree.Nodes[node_id]
	// fmt.Println(current_node)
	if (current_node.Status == 1) {
		can_execute_id = -1
		// fmt.Println("current node has done")
		return can_execute_id;
	}
	// 否则判断两个子节点的状态，如果都是ok，则返回当前节点id
	var left_node iplan.PlanTreeNode
	var right_node iplan.PlanTreeNode
	var can_execute bool
	can_execute = true
	if (current_node.Left != -1) {
		left_node = plan_tree.Nodes[current_node.Left]
		if (left_node.Status == 0) { can_execute = false }
	}
	if (current_node.Right != -1) {
		right_node = plan_tree.Nodes[current_node.Right]
		if (right_node.Status == 0) { can_execute = false }
	}
	if (can_execute && current_node.Locate == site) {
		// fmt.Println("current node can execute")
		can_execute_id = node_id
		return can_execute_id;
	}
	if (current_node.Left != -1) {
		can_execute_id = FindOneNode(plan_tree, current_node.Left)
	}
	if (can_execute_id != -1) { return can_execute_id }
	if (current_node.Right != -1) {
		can_execute_id = FindOneNode(plan_tree, current_node.Right)
	}
	if (can_execute_id != -1) { return can_execute_id }
	return -1
}

func ExecuteOneNode(plan_node *iplan.PlanTreeNode, plan_tree iplan.PlanTree) {
	switch {
	case plan_node.NodeType == 1 /*Table*/:
		return
	case plan_node.NodeType == 2 /*Select or Filter*/:
		ExecuteFilter(plan_node, plan_tree)
		return
	case plan_node.NodeType == 3 /*projuection*/:
		ExecuteProjection(plan_node, plan_tree)
		return
	case plan_node.NodeType == 4 /*join*/:
		ExecuteJoin(plan_node, plan_tree)
		return
	case plan_node.NodeType == 5 /*union*/:
		ExecuteUnion(plan_node, plan_tree)
		return
	}
}

func ExecuteFilter(plan_node *iplan.PlanTreeNode, plan_tree iplan.PlanTree) {
	// mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	// db, err := sql.Open("mysql", mysql)
	// TODO: assert(plan_node.Right = -1)

	tablename := plan_tree.Nodes[plan_node.Left].TmpTable
	query := "create table tmp_table_" + strconv.Itoa(plan_node.Nodeid) + " select * from " + tablename + " where " + plan_node.Where

	println(query)
	// stmt, err := db.Prepare(query)
	// res, err := stmt.Exec()
    // checkErr(err)
	// println(res)
	plan_node.TmpTable = "tmp_table_" + strconv.Itoa(plan_node.Nodeid)
	plan_node.Status = 1
}

func ExecuteProjection(plan_node *iplan.PlanTreeNode, plan_tree iplan.PlanTree) {
	// mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	// db, err := sql.Open("mysql", mysql)
	// TODO: assert(plan_node.Right = -1)

	tablename := plan_tree.Nodes[plan_node.Left].TmpTable
	query := "create table tmp_table_" + strconv.Itoa(plan_node.Nodeid) + " select " + plan_node.Cols + " from " + tablename
	println(query)

	// stmt, err := db.Prepare(query)
	// res, err := stmt.Exec()
    // checkErr(err)
	// println(res)
	plan_node.TmpTable = "tmp_table_" + strconv.Itoa(plan_node.Nodeid)
	plan_node.Status = 1
}

func ExecuteJoin(plan_node *iplan.PlanTreeNode, plan_tree iplan.PlanTree) {
	// mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	// db, err := sql.Open("mysql", mysql)
	// TODO: assert(plan_node.Right != -1)

	tablename1 := plan_tree.Nodes[plan_node.Left].TmpTable
	tablename2 := plan_tree.Nodes[plan_node.Right].TmpTable
	cols := strings.Split(plan_node.Joint_cols, ",")
	query := "create table tmp_table_" + strconv.Itoa(plan_node.Nodeid) + " select * from " + tablename1 + "," + tablename2 + " where " + tablename1 + "." + cols[0] + "==" + tablename2 + "." + cols[1]
	println(query)

	// stmt, err := db.Prepare(query)
	// res, err := stmt.Exec()
    // checkErr(err)
	// println(res)
	plan_node.TmpTable = "tmp_table_" + strconv.Itoa(plan_node.Nodeid)
	plan_node.Status = 1
}

func ExecuteUnion(plan_node *iplan.PlanTreeNode, plan_tree iplan.PlanTree) {
	// mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	// db, err := sql.Open("mysql", mysql)
	// TODO: assert(plan_node.Right != -1)

	tablename1 := plan_tree.Nodes[plan_node.Left].TmpTable
	tablename2 := plan_tree.Nodes[plan_node.Right].TmpTable
	query := "create table tmp_table_" + strconv.Itoa(plan_node.Nodeid) + " select * from " + tablename1 + "union" + "select * from " + tablename2
	println(query)

	// stmt, err := db.Prepare(query)
	// res, err := stmt.Exec()
    // checkErr(err)
	// println(res)
	plan_node.TmpTable = "tmp_table_" + strconv.Itoa(plan_node.Nodeid)
	plan_node.Status = 1
}

// func ExecuteInsert(plan_node *iplan.PlanTreeNode) {

// }

// func ExecuteTransmission(plan_node *iplan.PlanTreeNode) {

// }

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
