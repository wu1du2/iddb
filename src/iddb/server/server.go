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

	"irpccall"
	"irpctran"
	"iutilities"

	// "iparser"

	"os"
)

/*
iddb server设计思路
1.协程 启动两个SERVER
2.允许用户指定配置文件路径
3.初始化，知道自己和peers的信息
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

	//GET INPUT SQL STATEMENT
	go irpccall.RunCallServer()

	go irpctran.RunTranServer()

	return
}
