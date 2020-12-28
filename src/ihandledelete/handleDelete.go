package ihandledelete

import (
	"fmt"
	//"strings"
	"iexec"
	"imeta"
	"iparser"
	"iqueryanalyzer"
	"iqueryoptimizer"
	"irpccall"
	"iutilities"
	"strconv"

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
		//step1 find cid
		sqlstmt := "select cid from customer where xxxx"
		var txnID int64
		//txnID needs to be unique!
		txnID = 4433
		plantree := iparser.Parse(sqlstmt, txnID)
		plantree = iqueryanalyzer.Analyze(plantree)
		plantree = iqueryoptimizer.Optimize(plantree)

		ipaddr0 := iutilities.Peers[0].IP + ":" + iutilities.Peers[0].Call

		ipaddr1 := iutilities.Peers[1].IP + ":" + iutilities.Peers[1].Call

		iutilities.Waitgroup.Add(1)
		go irpccall.RunCallClient(ipaddr0, txnID)

		iutilities.Waitgroup.Add(1)
		go irpccall.RunCallClient(ipaddr1, txnID)

		iutilities.Waitgroup.Wait()

		plantree, err = imeta.Get_Tree(txnID)

		if err != nil {
			iutilities.CheckErr(err)
		}

		res := iexec.GetResult(plantree, txnID)

		println(res)

		//step2 delete by cid

		//delete from customer_0 where cid=xxx
		//delete from customer_1 where cid=xxx

		fmt.Println("...")
	}
	return TotalNum, outsql, siten

}
