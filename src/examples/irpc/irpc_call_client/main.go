package main

import (
	"irpc"
)

func main() {
	address := "localhost:50054"
	var txnid int64
	txnid = 1
	is_Suc := irpc.RunCallClient(address, txnid)
	println("irpc.RunCallClient(", address, ",", txnid, "),is_Suc=", is_Suc)
}
