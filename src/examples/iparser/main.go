package main

import (
	// "fmt"
	"iparser"
	"iqueryanalyzer"
)

func main() {
	// sql := "SELECT t1.a,t2.b FROM t1,t2"

	sql := `select customer.name,orders.quantity
	from customer,orders
	where customer.id=orders.customer_id`

	// sql := `select *
	// from customer,orders
	// where customer.id=orders.customer_id
	// and customer.id=1`

	logicalPlanTree := iparser.Parse(sql, 0)

	// iqueryanalyzer.ShowPlanTree(logicalPlanTree)
	physicalPlanTree := iqueryanalyzer.Analyze(logicalPlanTree)

	iqueryanalyzer.ShowPlanTree(physicalPlanTree)

}
