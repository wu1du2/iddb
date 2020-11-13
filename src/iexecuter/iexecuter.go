package iexecuter

func RunExecuter(txn_id int64) int64 {
	// time.Sleep(time.Duration(time.Second))
	println("executed successfully")
	return 0
}

func ExecuteInsertStmt(stmt string) int64 {
	println(stmt)
	return 0
}
