package main

import (
	"iparser"
	"iplan"
	"iqueryanalyzer"
	"iqueryoptimizer"
)

func main() {
	//test_old()
	var pt iplan.PlanTree
	sqlstmt := `select customer.name,orders.quantity
	from customer,orders
	where customer.id=orders.customer_id`
	pt = iparser.Parse(sqlstmt, 1)
	pt = iqueryanalyzer.Analyze(pt)
	pt = iqueryoptimizer.Optimize(pt)
	return
}
