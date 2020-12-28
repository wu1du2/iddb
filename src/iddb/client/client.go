package main

import (
	"iparser"
	"iqueryanalyzer"

	// "iqueryanalyzer"
	"iqueryoptimizer"
	// "imeta"
	// "irpc"
	"iexec"
	// "context"
	// "iexecuter"
	// "log"
	// "net"

	"imeta"
	"iplan"
	"irpccall"
	"iutilities"
	"strconv"

	// "iparser"
	"fmt"
	"runtime"
	"strings"
)

/*
iddb client设计思路
1.默认SERVER已经启动
2.循环接收用户输入，提供类似mysql的界面
3.初始化，知道自己和peers的信息
4.生成生成txnid(全局唯一，严格递增)
*/

var (
	txnID    int64
	err      error
	ipaddr   string
	plantree iplan.PlanTree
	queries  [10]string
)

func main() {
	txnID = 576
	iutilities.LoadAllConfig()
	runtime.GOMAXPROCS(8)

	test_insert()
	test_delete()

	var sqlstmt string
	queries[0] = `
	select *
	from customer`

	queries[1] = `
	select publisher.name
	from publisher`

	queries[2] = `
	select book.title
	from book
	where copies>5000`

	queries[3] = `
	select customer_id,quantity
	from orders
	where quantity<8`

	queries[4] = `
	select book.title, book.copies, publisher.name, publisher.nation
	from book, publisher
	where book.publisher_id = publisher.id
	and publisher.nation='USA'
	and book.copies>1000`

	queries[5] = `
	select customer.name, orders.quantity
	from customer,orders
	where customer.id=orders.customer_id`

	queries[6] = `
	select customer.name, customer.rank, orders.quantity
	from customer,orders
	where customer.id=orders.customer_id
	and customer.rank=1`

	queries[7] = `
	select customer.name, orders.quantity, book.title
	from customer,orders,book
	where customer.id=orders.customer_id
	and book.id=orders.book_id
	and customer.rank=1
	and book.copies>5000`

	queries[8] = `
	select customer.name, book.title, publisher.name, orders.quantity
	from customer, book, publisher, orders
	where customer.id=orders.customer_id
	and book.id=orders.book_id
	and book.publisher_id=publisher.id
	and book.id>220000
	and publisher.nation='USA'
	and orders.quantity>1`

	queries[9] = `
	select customer.name, book.title, publisher.name, orders.quantity
	from customer, book, publisher, orders
	where customer.id=orders.customer_id
	and book.id=orders.book_id
	and book.publisher_id=publisher.id
	and customer.id>308000
	and book.copies>100
	and orders.quantity>1
	and publisher.nation='PRC'`

	println("please enter TxnId: ")
	txnID, err = strconv.ParseInt(scanLine(), 10, 64)
	if err != nil {
		iutilities.CheckErr(err)
	} else {
		println("txnID=", txnID)
	}

	for qid := 0; qid < 9; qid++ {

		if qid == 4 || qid == 7 {
			txnID += 1
			continue
		}
		// println("please enter SQL statement end with ; (q to quit)")
		// sqlstmt = scanLine()

		// sqlstmt = `select customer.name,orders.quantity
		// from customer,orders
		// where customer.id=orders.customer_id`

		sqlstmt = queries[qid]

		println(sqlstmt)
		if strings.EqualFold(sqlstmt, "q") {
			break
		}

		plantree1 := iparser.Parse(sqlstmt, txnID)
		plantree = iqueryanalyzer.Analyze(plantree1)
		plantree = iqueryoptimizer.Optimize(plantree)

		// fmt.Println("plantree is ", plantree)

		plantree.Print()

		imeta.Connect_etcd()
		println("start imeta")

		err = imeta.Build_Txn(txnID)
		if err != nil {
			iutilities.CheckErr(err)
			return
		}

		println("imeta build txn ok")

		err = imeta.Set_Tree(txnID, plantree)
		if err != nil {
			iutilities.CheckErr(err)
			return
		}
		println("imeta set tree ok")

		println("end imeta")
		for _, node := range iutilities.Peers {
			ipaddr = node.IP + ":" + node.Call
			println("call node to work ", node.NodeId)
			iutilities.Waitgroup.Add(1)
			go irpccall.RunCallClient(ipaddr, txnID)
		}

		iutilities.Waitgroup.Wait()

		plantree, err = imeta.Get_Tree(txnID)

		if err != nil {
			iutilities.CheckErr(err)
			return
		}

		iexec.PrintResult(plantree, txnID)

		println("txn", txnID, "end!")

		txnID += 1

	}

	return
}

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

func test_insert() {
	var ins_stmts [5]string
	ins_stmts[0] = `
	insert into customer(id, name, rank) values(300001, 'Xiaoming', 1)`

	ins_stmts[1] = `
	insert into publisher(id, name, nation) values(104001,'High Education Press', 'PRC')`

	ins_stmts[2] = `
	insert into customer(id, name, rank) values(300002,'Xiaohong', 1)`

	ins_stmts[3] = `
	insert into book (id, title, authors, publisher_id, copies) values(205001, 'DDB', 'Oszu', 104001, 100)`

	ins_stmts[4] = `
	insert into orders (customer_id, book_id, quantity) values(300001, 205001,5)`

	for i, ins_stmt := range ins_stmts {
		println(i, ins_stmt)

		//iparser.HandleInsert()

		//func RunRemoteStmt(siteid int64, stmt string)

	}
	return
}

func test_delete() {
	var del_stmts [5]string
	del_stmts[0] = `
	delete from orders`

	del_stmts[1] = `
	delete from book where copies = 100`

	del_stmts[2] = `
	delete from publisher where nation = 'PRC'`

	del_stmts[3] = `
	delete from customer where name='Xiaohong' AND rank=1`

	del_stmts[4] = `
	delete from customer where rank = 1`

	for i, del_stmt := range del_stmts {
		println(i, del_stmt)

		//iparser.HandleInsert()

		//func RunRemoteStmt(siteid int64, stmt string)

	}
	return
}

func RunRemoteStmt(siteid int64, stmt string) {
	node := iutilities.Peers[siteid]
	ipaddr = node.IP + ":" + node.Tran
	iexec.ExecuteRemoteCreateStmt(ipaddr, stmt)
}
