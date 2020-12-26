package main

import (
	"fmt"
	"iparser"
	"iqueryanalyzer"
	//"github.com/xwb1989/sqlparser"
)

func main() {
	test_old()
	// test_delete()
}

func test_delete() {

	//sql := `Delete from Orders`
	//sql := `Delete from Book where copies = 100`
	sql := `delete from Publisher where nation = 'PRC'`
	//sql := `delete from Customer where name='Xiaohong' AND rank=1`
	//sql := `delete from Customer where rank = 1`

	i, str, j := iparser.HandleDelete(sql)
	fmt.Println(i)
	fmt.Println(str)
	fmt.Println(j)
}

func test_insert() {

	//sql := `insert into customer(id,name,rank) values(3000,'Able',1)`
	sql := `insert into Customer(id, name, rank) values(300001, 'Xiaoming', 1)`
	//sql := `insert into Customer(id, name, rank) values(300002,'Xiaohong', 1)\
	//sql := `insert into Publisher(id, name, nation) values(104001,'High Education Press', 'PRC')`
	//sql := `insert into Book (id, title, authors, publisher_id, copies) values(205001, 'DDB', 'Oszu', 104001, 100)`
	//sql := `insert into Orders (customer_id, book_id, quantity) values(300001, 205001,5)`

	i, str, j := iparser.HandleInsert(sql)
	fmt.Println(i)
	fmt.Println(str)
	fmt.Println(j)
}

func test_old() {
	// sql := "SELECT t1.a,t2.b FROM t1,t2"

	sql := `select customer.name,orders.quantity
	from customer,orders
	where customer.id=orders.customer_id`

	// sql := `select *
	// from customer,orders
	// where customer.id=orders.customer_id
	// and customer.id=1`

	//sql := `insert into customer(id,name,rank) values(3000,'Able',1)`
	// iparser.Parse(sql, 0)

	logicalPlanTree := iparser.Parse(sql, 0)
	//fmt.Print(logicalPlanTree)

	// iqueryanalyzer.ShowPlanTree(logicalPlanTree)

	physicalPlanTree := iqueryanalyzer.Analyze(logicalPlanTree)

	iqueryanalyzer.ShowPlanTree(physicalPlanTree)

}
