package itrans

import (
	"database/sql"
	"iutilities"

	_ "github.com/go-sql-driver/mysql"
)

var site int64
var mysql_user string
var mysql_passwd string
var mysql_db string
var mysql_ip_port string

func Init() {
	// TODO:set site id
	site = iutilities.GetMe().NodeId
	// site = 1
	mysql_user = "root"
	mysql_passwd = "123456"
	mysql_db = "test"
	mysql_ip_port = "127.0.0.1:3306"
}

func ExecuteCreateStmt(stmt string) int64 {
	Init()
	mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	db, err := sql.Open("mysql", mysql)

	println(stmt)
	println(err)
	stmts, err := db.Prepare(stmt)
	res, err := stmts.Exec()
	println(res)
	println(err)
	return 0
}
