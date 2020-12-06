package main

import (
    "fmt"
    "github.com/coreos/etcd/clientv3"
    "time"
    "context"
	"strconv"
	"strings"
	"iplan"
)
 var cli *clientv3.Client

func Connect_etcd(){
	config := clientv3.Config{
        Endpoints:   []string{"localhost:2379"},
        DialTimeout: 15 * time.Second,
    }
    if cli,err := clientv3.New(config);err!=nil {
		fmt.Println(err.Error())
		return
	}	
	//fmt.Println("etcd connect succ")	
	//defer cli.Close()
}

func Disconnect_etcd({
	defer cli.Close()
})

Connect_etcd()


func Get_Tree(txnid int) (PlanTree, error){
	//Connect_etcd()
	kv := clientv3.NewKV(cli)
    var treex string
    treex = strconv.Itoa(txnid)
    treex = "/"+treex+"/"
    if getResp, err := kv.Get(context.TODO(),treex,clientv3.WithPrefix());err!=nil {
		fmt.Println("err1")
		var a PlanTree
		return a,err
	}else {
		fmt.Println("Get_Tree:")
		fmt.Println(getResp.Kvs)
		var tree_return PlanTree
		for i,mykv :=range getResp.Kvs{
			//fmt.Println(strconv.Itoa(i))
			fmt.Println(string(mykv.Key))
			fmt.Println(string(mykv.Value))

			//get node id
			if (tree_return.NodeNum < i){
				tree_return.NodeNum = i
			}
			s1 := string(mykv.Key)
			sep := "/"
			ikeys := strings.Split(s1,sep)
			// "/txnid/node id"
			ikey,err := strconv.Atoi(ikeys[2])
			if(err!=nil){
				fmt.Println("err2")				
				var b PlanTree
				return b,err 
			}

			//split
			s2 := string(mykv.Value)
			sepp := "##"
			ivalues := strings.Split(s2,sepp)
			//ivalue -> Node
			// 1##2##3##4##5##6##7##8##9##10##11##12##13
			if(len(ivalues)!=13) {
				fmt.Println("node"+string(mykv.Key)+"has wrong  attribute number:"+strconv.Itoa(len(ivalues)))
			}
			//Nodes attr
			tree_return.Nodes[ikey].Nodeid,err = strconv.Atoi(ivalues[0])
			tree_return.Nodes[ikey].Left,err = strconv.Atoi(ivalues[1])
			tree_return.Nodes[ikey].Right,err = strconv.Atoi(ivalues[2])
			tree_return.Nodes[ikey].Parent,err = strconv.Atoi(ivalues[3])
			tree_return.Nodes[ikey].Status,err = strconv.Atoi(ivalues[4])
			tree_return.Nodes[ikey].TmpTable = ivalues[5]
			tree_return.Nodes[ikey].Locate,err = strconv.Atoi(ivalues[6])
			tree_return.Nodes[ikey].TransferFlag,err = strconv.ParseBool(ivalues[7])
			tree_return.Nodes[ikey].Dest,err = strconv.Atoi(ivalues[8])
			tree_return.Nodes[ikey].NodeType,err = strconv.Atoi(ivalues[9])
			tree_return.Nodes[ikey].Where = ivalues[10]
			tree_return.Nodes[ikey].Cols = ivalues[11]
			tree_return.Nodes[ikey].Joint_cols = ivalues[12]
		}
		//fmt.Println(tree_return)
		return 	tree_return,nil
	}
}

func Get_Node(txnid int, nodeid int)(PlanTreeNode, error){
	kv := clientv3.NewKV(cli)
	var nodex string
	nodex =  "/"+strconv.Itoa(txnid)+"/"+strconv.Itoa(nodeid)
    if getResp, err := kv.Get(context.TODO(),nodex);err !=nil {
		var a PlanTreeNode
		fmt.Println(err)
		return a,err
	}else {
		fmt.Println("Get_NODE:")
		fmt.Println(getResp.Kvs)

		var Node_return PlanTreeNode

		//split
		s2 := string(getResp.Kvs[0].Value)
		sepp := "##"
		ivalues := strings.Split(s2,sepp)
		//ivalue -> Node
		// 1##2##3##4##5##6##7##8##9##10##11##12##13
		if(len(ivalues)!=13) {
			fmt.Println("node"+nodex+"has wrong  attribute number:"+strconv.Itoa(len(ivalues)))
		}
			//Nodes attr
			Node_return.Nodeid,err = strconv.Atoi(ivalues[0])
			Node_return.Left,err = strconv.Atoi(ivalues[1])
			Node_return.Right,err = strconv.Atoi(ivalues[2])
			Node_return.Parent,err = strconv.Atoi(ivalues[3])
			Node_return.Status,err = strconv.Atoi(ivalues[4])
			Node_return.TmpTable = ivalues[5]
			Node_return.Locate,err = strconv.Atoi(ivalues[6])
			Node_return.TransferFlag,err = strconv.ParseBool(ivalues[7])
			Node_return.Dest,err = strconv.Atoi(ivalues[8])
			Node_return.NodeType,err = strconv.Atoi(ivalues[9])
			Node_return.Where = ivalues[10]
			Node_return.Cols = ivalues[11]
			Node_return.Joint_cols = ivalues[12]

		return Node_return,nil
	}
}

func Set_Node(txnid int, Node_in PlanTreeNode){
	kv := clientv3.NewKV(cli)
	//key value
	ikey :=  "/"+strconv.Itoa(txnid)+"/"+strconv.Itoa(Node_in.Nodeid)
	var ivalue string
	ivalue =  strconv.Itoa(Node_in.Nodeid)+"##"+strconv.Itoa(Node_in.Left)+"##"+strconv.Itoa(Node_in.Right)+"##"
	ivalue = ivalues+strconv.Itoa(Node_in.Parent)+"##"+strconv.Itoa(Node_in.Status)+"##"+Node_in.TmpTable+"##"+strconv.Itoa(Node_in.Locate)+"##"
	ivalue = ivalues+strconv.FormatBool(Node_in.TransferFlag)+"##"+strconv.Itoa(Node_in.Dest)+"##"+strconv.Itoa(Node_in.NodeType)+"##"
	ivalue = ivalues+Node_in.Where+"##"+Node_in.Cols+"##"+Node_in.Joint_cols
	
    if putResp, err := kv.Put(context.TODO(),ikey,ivalue,clientv3.WithPrevKV());err !=nil {
		fmt.Println(err)
		return
	}
}

