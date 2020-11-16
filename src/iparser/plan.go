package Plan


type PlanTreeNode struct {
	left int
	right int
	parent int 
	status int //1 for finished, 0 for waiting
	tmp_table string //the name of tmp_table
	locate int // the sitenum,n for site_n
	transfer_flag bool //1 for transer, 0 for not
	dest int //the sitenum of the dest
	node_type int //1 for table, 2 for select, 3 for projuection, 4 for join, 5 for union
	detail string//according to node_type, (1)table_name for table, (2)where_condition for select, (3)col_name for projection, (4)join_type for join, (5)nil for union
}
type PlanTree struct {
	node_num int
	nodes [10] PlanTreeNode
}


