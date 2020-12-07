package main

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

func main() {
	sql := "SELECT a FROM t1,t2 WHERE t1.a=t2.a"
	// sql := "select * from x_order where userId = 1 order by id desc limit 10,1;"
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
		println("ERROR with parser!\n")
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		//expr := stmt.Where.Expr
		//fmt.Printf("%v", expr.Operator)
		switch comparisonExpr := stmt.Where.Expr.(type) {
		case *sqlparser.ComparisonExpr:
			fmt.Printf("%v", comparisonExpr.Operator)
		}
	case *sqlparser.Insert:
	}

}
