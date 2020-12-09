package main

import (
	// "iparser"
	// "iqueryanalyzer"
	// "iqueryoptimizer"
	// "imeta"
	// "irpc"
	// "iexecuter"
	// "context"
	// "iexecuter"
	// "log"
	// "net"

	"imeta"
	"iplan"
	"irpccall"
	"irpctran"
	"iutilities"
	"sync"

	// "iparser"
	"fmt"
	"os"
	"runtime"
	"strings"
)

/*
iddb client设计思路
1.默认SERVER已经启动
2.循环接收用户输入，提供类似mysql的界面
3.初始化，知道自己和peers的信息
4.生成生成txnid(全局唯一，严格递增)
*/

func main() {
	//INIT
	for i, v := range os.Args {
		if i == 1 {
			println(i, v)
			iutilities.Configfile = v
			println("iutilities.Configfile= ", iutilities.Configfile)
		}

	}
	iutilities.LoadAllConfig()

	runtime.GOMAXPROCS(8)

	imeta.Connect_etcd()

	var sqlstmt string

	test1()

	for {
		println("please enter SQL statement end with ; (q to quit)")
		sqlstmt = scanLine()
		println(sqlstmt)
		if strings.EqualFold(sqlstmt, "q") {
			break
		}

	}

	return
}

func scanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

func test1() {
	var txnID int64
	txnID = 10001
	// sqlstmt := "select * from Publisher"
	var plantree iplan.PlanTree
	var err error
	// plantree, err = iparser.parse(sqlstmt)
	// plantree, err = iqueryanalyzer.analyze(plantree)
	// plantree, err = ioptimizer.optimize(plantree)
	plantree = generatePlanTree()
	println("start imeta")
	err = imeta.Build_Txn(txnID)
	if err != nil {
		iutilities.CheckErr(err)
		return
	}
	err = imeta.Set_Tree(txnID, plantree)
	if err != nil {
		iutilities.CheckErr(err)
		return
	}
	var ipaddr string
	println("end imeta")
	for _, node := range iutilities.Peers {
		ipaddr = node.IP + ":" + node.Call
		println("call node to work ", node.NodeId)
		go irpccall.RunCallClient(ipaddr, txnID)
	}
	var waitgroup sync.WaitGroup
	println("client end!")
	waitgroup.Add(1)
	waitgroup.Wait()
}

func generatePlanTree() iplan.PlanTree {
	fmt.Println("try to create plantree")
	var plan_tree iplan.PlanTree
	plan_tree.NodeNum = 5
	/*
	       0
	   1       2
	   3       4
	*/
	// 根结点0
	pn0 := &plan_tree.Nodes[0]
	pn0.Nodeid = 0
	pn0.Left = 1
	pn0.Right = 2
	pn0.Parent = -1
	pn0.Status = 0
	pn0.Locate = 0
	pn0.NodeType = 4
	pn0.TransferFlag = true
	pn0.Dest = 1
	pn0.Joint_cols = "id,customer_id"
	// 结点1
	pn1 := &plan_tree.Nodes[1]
	pn1.Nodeid = 1
	pn1.Left = 3
	pn1.Right = -1
	pn1.Parent = 0
	pn1.Status = 0
	pn1.Locate = 0
	pn1.NodeType = 2
	pn1.Where = "id > 2"
	// 结点2
	pn2 := &plan_tree.Nodes[2]
	pn2.Nodeid = 2
	pn2.Left = 4
	pn2.Right = -1
	pn2.Parent = 0
	pn2.Status = 0
	pn2.Locate = 0
	pn2.NodeType = 3
	pn2.Cols = "customer_id,quantity"
	// 结点3 data节点
	pn3 := &plan_tree.Nodes[3]
	pn3.Nodeid = 3
	pn3.Left = -1
	pn3.Right = -1
	pn3.Parent = 1
	pn3.Status = 1
	pn3.Locate = 0
	pn3.NodeType = 1
	pn3.TmpTable = "customer"
	// 结点4 data节点
	pn4 := &plan_tree.Nodes[4]
	pn4.Nodeid = 4
	pn4.Left = -1
	pn4.Right = -1
	pn4.Parent = 2
	pn4.Status = 1
	pn4.Locate = 0
	pn4.NodeType = 1
	pn4.TmpTable = "orders"
	return plan_tree
}

func testtrans() {
	// iutilities.Peers = iutilities.GetPeers()
	testnodeid := 0
	ip := iutilities.Peers[testnodeid].IP
	port := iutilities.Peers[testnodeid].Tran
	address := ip + ":" + port
	println(address)
	// address := "localhost:50053"
	var table irpctran.Table
	// table.Createstmt = "Create Table PUBLISHER2 (ID int, NATION varchar(255) );"
	table.Createstmt = "Create Table PUBLISHER5 (ID int, NATION varchar(255) );"
	irpctran.RunTranClient(address, table)
	table.Createstmt = "Insert Into PUBLISHER5 Values (1,'US'),(2,'US'),(3,'CHN');"
	irpctran.RunTranClient(address, table)
}

func testcall() {
	testnodeid := 0
	ip := iutilities.Peers[testnodeid].IP
	port := iutilities.Peers[testnodeid].Call
	// address := "10.77.70.161:50054"
	address := ip + ":" + port
	println(address)
	var txnid int64
	txnid = 1
	is_Suc := irpccall.RunCallClient(address, txnid)
	println("irpc.RunCallClient(", address, ",", txnid, "),is_Suc=", is_Suc)
}
