package main

import (
	"irpctran"
)

/*
type Table struct {
	Operation string    //insert into TableName (Col1, Col2, Col3)
	Rowlength int64     //10
	Record    []string  //(Val1, Val2, Val3)
}
*/
func main() {
	address := "localhost:50053"
	var table irpctran.Table
	table.Createstmt = "Create Table PUBLISHER (ID int, NATION varchar(255) );"
	irpctran.RunTranClient(address, table)
}
