package main

import (
    "fmt"
    "github.com/coreos/etcd/clientv3"
    "time"
    "context"
    "strconv"
)
 var cli *clientv3.Client

type TreeNode struct {
	Nodeid		 int
	Left         int
	Right        int
	Parent       int
	Status       int    //1 for finished, 0 for waiting
	TmpTable     string //the name of tmp_table
	Locate       int    // the sitenum,n for site_n
	TransferFlag bool   //1 for transer, 0 for not
	Dest         int    //the sitenum of the dest
	NodeType     int    //1 for table, 2 for select, 3 for projuection, 4 for join, 5 for union
    //detail string
    //according to node_type, (1)table_name for table, (2)where_condition for select, (3)col_name for projection, (4)join_type for join, (5)nil for union
	Where string
	Cols  string
	//cols_type string
	Joint_cols string
	//union string
}
type PlanTree struct {
	NodeNum int
	Nodes   [10]TreeNode
}

func Get_Tree(txnid int){
	kv := clientv3.NewKV(cli)
    var treex string
    treex = strconv.Itoa(txnid)
    treex = "/txn"+treex+"/"
    if getResp, err := kv.Get(context.TODO(),treex,clientv3.WithPrefix());err!=nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println("Get_Tree:")
		fmt.Println(getResp.Kvs)
	}
}

func Get_Node(txnid int, nodeid int){
	kv := clientv3.NewKV(cli)
	var nodex string
	nodex =  "/txn"+strconv.Itoa(txnid)+"/node"+strconv.Itoa(nodeid)
    if getResp, err := kv.Get(context.TODO(),nodex);err !=nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println("Get_NODE:")
		fmt.Println(getResp.Kvs)
	}
}

func Set_Node(txnid int, nodeid int, nodevalue string){
	kv := clientv3.NewKV(cli)
	var nodex string
	nodex =  "/txn"+strconv.Itoa(txnid)+"/node"+strconv.Itoa(nodeid)
    if putResp, err := kv.Put(context.TODO(),nodex,nodevalue,clientv3.WithPrevKV());err !=nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println("Header")
		fmt.Println(putResp.Header)
		if putResp.PrevKv !=nil {
			fmt.Println("Set_NODE:")
			fmt.Println(string(putResp.PrevKv.Value))
		}
	}
}


func main() {
    /*
        DialTimeout time.Duration `json:"dial-timeout"`
        Endpoints []string `json:"endpoints"`
    */
    /*var (
		config clientv3.Config
		client *clientv3.Client
		err error
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
		leaseResp,leaseResp1 *clientv3.LeaseGrantResponse
	)*/
	var err error

    config := clientv3.Config{
        Endpoints:   []string{"localhost:2379"},
        DialTimeout: 15 * time.Second,
    }

    if cli,err = clientv3.New(config);err!=nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Println("etcd connect succ")
	
	var tag int
	for true{
		fmt.Println("1:Get_Tree; 2:Get_Node; 3:Set_Node; 0: exit")
		fmt.Scanf("%d",&tag)
		fmt.Println(tag)
		if (tag==1) {
			var nid int
			fmt.Println("input txn id")
			fmt.Scanf("%d",&nid)
			Get_Tree(nid)
		}
		if (tag==2) {
			var tid int
			var nid int
			fmt.Println("input txn id, node id")
			fmt.Scanf("%d %d",&tid, &nid)
			Get_Node(tid, nid)
		}
		if (tag==3) {
			var tid int
			var nid int
			var val string
			fmt.Println("input txn id, node id, value")
			fmt.Scanf("%d %d  %s",&tid, &nid, &val)
			Set_Node(tid, nid, val)
		}
		if (tag==0){
			fmt.Println("bye bye")
			break
		}

	}
	defer cli.Close()

}

