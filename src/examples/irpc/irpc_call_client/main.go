package main

import (
	"irpccall"
	"iutilities"
)

func main() {
	testnodeid := 1
	ip := iutilities.Peers[testnodeid].IP
	port := iutilities.Peers[testnodeid].Call
	// address := "10.77.70.161:50054"
	address := ip + ":" + port
	var txnid int64
	txnid = 1
	is_Suc := irpccall.RunCallClient(address, txnid)
	println("irpc.RunCallClient(", address, ",", txnid, "),is_Suc=", is_Suc)
}
