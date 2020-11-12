package main

import (
	"irpc"
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
	var table irpc.Table
	table.Operation = "INSERT INTO PUBLISHER (ID, NATION)"
	table.Rowlength = 10
	table.Record = make([]string, 10)
	table.Record[0] = "(0,'USA')"
	table.Record[1] = "(1,'USA')"
	table.Record[2] = "(2,'CHN')"
	table.Record[3] = "(3,'USA')"
	table.Record[4] = "(4,'CHN')"
	table.Record[5] = "(5,'USA')"
	table.Record[6] = "(6,'USA')"
	table.Record[7] = "(7,'USA')"
	table.Record[8] = "(8,'CHN')"
	table.Record[9] = "(9,'USA')"
	irpc.RunTranClient(address, table)
}
