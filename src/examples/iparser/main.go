package main

import (
	"fmt"
	"iparser"
	"iqueryanalyzer"
)

func main() {
	// sql := "SELECT t1.a,t2.b FROM t1,t2 WHERE t1.a=t2.a and t1.b=t2.b"
	sql := `select Customer.name,Orders.quantity
	from Customer,Orders
	where Customer.id=Orders.customer_id`

	logicalPlanTree := iparser.Parse(sql, 0)
	// fmt.Println(logicalPlanTree)
	physicalPlanTree := iqueryanalyzer.Analyze(logicalPlanTree)
	// for _, node := range physicalPlanTree.Nodes {
	// 	println(node.TmpTable)

	// }
	fmt.Println(physicalPlanTree)
}
