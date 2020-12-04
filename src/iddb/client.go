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

	iutilities.Me = getMe()
	iutilities.Me.Print()

	iutilities.Peers = getPeers()
	iutilities.Peers[0].Print()
	iutilities.Peers[1].Print()
	iutilities.Peers[2].Print()
	iutilities.Peers[3].Print()

	iutilities.Mysql = getMysqlConfig()
	println("mysql_user= ", iutilities.Mysql.Mysql_user)
	println("mysql_passwd= ", iutilities.Mysql.Mysql_passwd)
	println("mysql_db= ", iutilities.Mysql.Mysql_db)
	println("mysql_ip_port= ", iutilities.Mysql.Mysql_ip_port)
	//GET INPUT SQL STATEMENT
	var sqlstmt string
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

func getMe() iutilities.Nodes {
	return iutilities.GetMe()
}

func getPeers() []iutilities.Nodes {
	return iutilities.GetPeers()
}

func getMysqlConfig() iutilities.MysqlConfig {
	return iutilities.GetMysqlConfig()
}

// func parse(sql string)(){

// }

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
