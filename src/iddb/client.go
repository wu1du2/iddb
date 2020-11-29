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
	"iutilities"
	// "iparser"

	"fmt"
	"os"
	"strings"
)

var (
	me    iutilities.Nodes
	peers []iutilities.Nodes
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
		println(i, v)
	}

	me = getMe()
	me.Print()

	peers = getPeers()
	peers[0].Print()
	peers[1].Print()
	peers[2].Print()
	peers[3].Print()

	//GET INPUT SQL STATEMENT
	var sqlstmt string
	for {
		println("please enter SQL statement end with ; (q to quit)")
		sqlstmt = ScanLine()
		println(sqlstmt)
		if strings.EqualFold(sqlstmt, "q") {
			break
		}

	}
	return
}

func getMe() iutilities.Nodes {
	return iutilities.GetMe()
}

func getPeers() []iutilities.Nodes {
	return iutilities.GetPeers()
}

// func parse(sql string)(){

// }

func ScanLine() string {
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
