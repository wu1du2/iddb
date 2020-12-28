package iparser

import (
	"fmt"
	//"strings"
	"strconv"
	//"iqueryanalyzer"
	"github.com/xwb1989/sqlparser"
)

func HandleDelete(sql string) (int64, [4]string, [4]int64) {
	var TotalNum int64
	var outsql [4]string
	var siten [4]int64

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
		println("ERROR with parser!\n")
	}
	sel := stmt.(*sqlparser.Delete)

	tablename := sqlparser.String(sel.TableExprs)
	fmt.Println(tablename)
	switch tablename {
	case "orders":
		TotalNum = 4
		i := 0
		for i < int(TotalNum) {
			siten[i] = int64(i + 1)
			outsql[i] = "delete from orders_" + strconv.Itoa(i+1) + sqlparser.String(sel.Where)
			i = i + 1
		}
	case "book":
		TotalNum = 3
		i := 0
		for i < int(TotalNum) {
			siten[i] = int64(i + 1)
			outsql[i] = "delete from book_" + strconv.Itoa(i+1) + sqlparser.String(sel.Where)
			i = i + 1
		}
	case "publisher":
		TotalNum = 4
		i := 0
		for i < int(TotalNum) {
			siten[i] = int64(i + 1)
			outsql[i] = "delete from publisher_" + strconv.Itoa(i+1) + sqlparser.String(sel.Where)
			i = i + 1
		}
	case "customer":
		fmt.Println("...")
	}
	return TotalNum, outsql, siten

}
