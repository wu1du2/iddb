package main

import (
	"irpc"
)

func main() {
	address := "localhost:50054"
	is_Suc := irpc.RunCallClient(address)
	println("irpc.RunCallClient(", address, "),is_Suc=", is_Suc)
}
