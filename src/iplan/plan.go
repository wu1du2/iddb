package iplan

type PlanTreeNode struct {
	Left         int
	Right        int
	Parent       int
	Status       int    //1 for finished, 0 for waiting
	TmpTable     string //the name of tmp_table
	Locate       int    // the sitenum,n for site_n
	TransferFlag bool   //1 for transer, 0 for not
	Dest         int    //the sitenum of the dest
	NodeType     int    //1 for table, 2 for select, 3 for projuection, 4 for join, 5 for union
	//detail string//according to node_type, (1)table_name for table, (2)where_condition for select, (3)col_name for projection, (4)join_type for join, (5)nil for union
	Where string
	Cols  string
	//cols_type string
	//joint_tables string
	//union string
}
type PlanTree struct {
	NodeNum int
	Nodes   [10]PlanTreeNode
}
