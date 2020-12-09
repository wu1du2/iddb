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
	var waitgroup sync.WaitGroup
	for i, v := range os.Args {
		if i == 1 {
			println(i, v)
			iutilities.Configfile = v
			println("iutilities.Configfile= ", iutilities.Configfile)
		}

	}
	iutilities.LoadAllConfig()
	runtime.GOMAXPROCS(8)
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
	txnID := 10001
	sqlstmt := "select * from Publisher"
	var plantree iplan.PlanTree
	var err error
	plantree, err = iparser.parse(sqlstmt)
	plantree, err = iqueryanalyzer.analyze(plantree)
	plantree, err = ioptimizer.optimize(plantree)
	if err != nil {
		iutilities.CheckErr(err)
		return
	}
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
	for _, node := range iutilities.Peers {
		ipaddr = node.IP + ":" + node.Call
		go irpccall.RunCallClient(ipaddr, txnID)
	}

	waitgroup.Add(1)
	waitgroup.Wait()
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
