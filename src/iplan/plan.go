package iplan

import "fmt"

const (
	//MaxNodeNum define the max num of the plantree nodes
	MaxNodeNum = 100
	MaxFragNum = 10
)

//PlanTreeNode is the basic node for a PlanTree
type PlanTreeNode struct {
	Nodeid       int64
	Left         int64 //when there is only on child, save in Left
	Right        int64
	Parent       int64
	Status       int64  //1 for finished, 0 for waiting
	TmpTable     string //the name of tmp_table
	Locate       int64  // the sitenum,n for site_n
	TransferFlag bool   //true for transer, false for not
	Dest         int64  //the sitenum of the dest
	NodeType     int64  //1 for table, 2 for select, 3 for projection, 4 for join, 5 for union
	//detail string//according to node_type, (1)table_name for table, (2)where_condition for select, (3)col_name for projection, (4)join_type for join, (5)nil for union
	Where string
	Cols  string
	//cols_type string
	Joint_cols string //"customer_id,id"
	//union string
}

//PlanTree saves the tree nodes which are built by iparser
type PlanTree struct {
	NodeNum int64
	Root    int64 //add rootnum to find root
	Nodes   [MaxNodeNum]PlanTreeNode
}

type FragNode struct {
	FragId        int64 //1,2,3,4
	FragName      string
	SiteNum       int64 //1,2,3,4
	FragCondition string
	Ip            string //10.77.10.161
}

type FragTree struct {
	FragNum   int64
	Frags     [MaxFragNum]FragNode
	FragType  int64 //0 for Horizontal, 1 for vertical
	TableName string
}

func (pt *PlanTree) Print() {
	var i int64
	for i = 0; i < pt.NodeNum; i++ {
		fmt.Println(pt.Nodes[i+1])
	}
	return
}
