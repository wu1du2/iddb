package itransmission

import (
	"fmt"
	"iplan"
	"database/sql"
	// "irpc"
    _ "github.com/go-sql-driver/mysql"
)

type Field struct {
	fieldName string
	fieldDesc string
	dataType  string
	isNull    string
	length    int
}

func TableInfo(dbName string) map[string]string  {
	mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	db, err := sql.Open("mysql", mysql)
	println(err)
	sqlStr := `SELECT table_name tableName,TABLE_COMMENT tableDesc
			FROM INFORMATION_SCHEMA.TABLES 
			WHERE UPPER(table_type)='BASE TABLE'
			AND LOWER(table_schema) = ? 
			ORDER BY table_name asc`

	var result = make(map[string]string)

	rows, err := db.Query(sqlStr,dbName)
	checkErr(err)

	for rows.Next() {
		var tableName,tableDesc string
		err = rows.Scan(&tableName, &tableDesc)
		checkErr(err)

		if len(tableDesc) == 0 {
			tableDesc = tableName
		}
		result[tableName] = tableDesc
	}
	return result
}

func FieldInfo(dbName,tableName string) [] Field{
	mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
	db, err := sql.Open("mysql", mysql)
	println(err)

	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns 
			WHERE table_schema = ? AND table_name = ?`


	var result [] Field
	
	rows, err := db.Query(sqlStr,dbName,tableName)
	checkErr(err)

	for rows.Next() {
		var f Field
		err = rows.Scan(&f.fieldName, &f.fieldDesc, &f.dataType, &f.isNull, &f.length)
		checkErr(err)

		result = append(result, f)
	}
	return result
}

func ExecuteTransmission(plan_node *iplan.PlanTreeNode) {
	println("TransferFlag")
	if (plan_node.TransferFlag) {
		mysql := mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_ip_port + ")/" + mysql_db + "?charset=utf8"
		db, err := sql.Open("mysql", mysql)
		println(err)
		// create table
		query := "show create table " + plan_node.TmpTable
		println(query)
		rows, err := db.Query(query)
		rows.Next()
		var table_name sql.NullString
		var create_sql sql.NullString
		err = rows.Scan(&table_name, &create_sql)
		checkErr(err)
		fmt.Println(table_name)
		fmt.Println(create_sql)
		// get table desc
		tableInfo := TableInfo(mysql_db)
		fmt.Println(tableInfo)

		filedInfo := FieldInfo(mysql_db,plan_node.TmpTable)

		for _,item := range filedInfo {
			fmt.Println(item)
		}
		// insert 
		insert_query := "insert into " + plan_node.TmpTable + " values "
		query = "select * from " + plan_node.TmpTable
		println(query)
		rows, err = db.Query(query)
		for rows.Next() {
			var uid sql.NullInt32
			var username sql.NullString
			var rank sql.NullInt32
	
			err = rows.Scan(&uid, &username, &rank)
			checkErr(err)
			fmt.Println(uid)
			fmt.Println(username)
			fmt.Println(rank)
		}

		err = rows.Scan(&table_name, &create_sql)
		checkErr(err)
		fmt.Println(table_name)
		fmt.Println(create_sql)		
		// irpc.RunTranClient(address,create_sql)
		plan_node.Status = 1
	}
}
