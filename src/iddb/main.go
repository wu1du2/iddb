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
	"fmt"
	"iutilities"
	"strings"
)

func readconf() {

}

/*
iddb main设计思路
1.CLIENT, 默认SERVER已经启动
2.循环接收用户输入，提供类似mysql的界面
3.初始化，知道自己和peers的信息

*/

func main() {
	//INIT
	var nodes iutilities.Nodes
	nodes = iutilities.GetMe()
	nodes.Print()
	//
	quit := false
	var sqlstmt string
	for quit != true {
		println("please enter SQL statement end with ; (q to quit)")
		sqlstmt = ScanLine()
		println(sqlstmt)
		if strings.EqualFold(sqlstmt, "q") {
			quit = true
		}
	}
	return

}

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
