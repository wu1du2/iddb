package main

import (
	"fmt"
	"log"

	"github.com/marianogappa/sqlparser"
)

func main() {
	// sql := "select Customer.name ,Orders.quantity,Book.title from Customer,Orders,Book where Customer.id=Orders.customer_id and Book.id=Orders.book_id  and Customer.rank=1 and Book.copies>5000"
	sql := "SELECT * FROM Customer"
	query, err := sqlparser.Parse(sql)
	if err != nil {
		println("ERROR!")
		log.Fatal(err)
	}
	fmt.Printf("%+#v", query)

}
