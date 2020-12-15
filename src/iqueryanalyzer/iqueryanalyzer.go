package iqueryanalyzer

import (
	"iparser"
	"iplan"
)

//nodeid globally save the nodeid
var nodeid int64

// var tmptableid int64 = iparser.Tmptableid

// func createMetaTable(tableName string) (metaTables []iplan.PlanTreeNode) {
// 	if tableName == "Customer" {

// 	} else if tableName == "Orders" {

// 	} else {
// 		println("ERROR when creating metatables! cannot find tableName in metadata")
// 	}
// 	return metaTables
// }

func createUnionNode(TmpTableName string) (node iplan.PlanTreeNode) {
	node = iparser.InitalPlanTreeNode()
	node.Nodeid = nodeid
	node.Status = 0
	node.TmpTable = TmpTableName
	node.NodeType = 5
	node.Locate = 1
	node.TransferFlag = false

	return node
}

//Analyze can transer a logicalPlanTree to a physicalPlanTree
func Analyze(logicalPlanTree iplan.PlanTree) (physicalPlanTree iplan.PlanTree) {
	physicalPlanTree = logicalPlanTree
	nodeid = physicalPlanTree.NodeNum
	for idx, node := range physicalPlanTree.Nodes {
		if node.NodeType == 1 { //if is table
			// metaTableNode := node
			// metaTableNode.Locate = getLocate(metaTableNode.TmpTable)
			tableName := node.TmpTable
			if tableName == "customer" {
				customerTableNode1 := node
				customerTableNode1.Nodeid = nodeid
				customerTableNode1.TmpTable = "customer_1"
				customerTableNode1.Locate = 1
				customerTableNode1.TransferFlag = false
				physicalPlanTree.Nodes[nodeid] = customerTableNode1
				physicalPlanTree.Nodes[nodeid].Parent = node.Nodeid
				physicalPlanTree.Nodes[idx].Left = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++

				customerTableNode2 := node
				customerTableNode2.Nodeid = nodeid
				customerTableNode2.TmpTable = "customer_2"
				customerTableNode2.Locate = 2
				customerTableNode2.TransferFlag = true
				customerTableNode2.Dest = 1
				physicalPlanTree.Nodes[nodeid] = customerTableNode2
				physicalPlanTree.Nodes[nodeid].Parent = node.Nodeid
				physicalPlanTree.Nodes[idx].Right = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++

				physicalPlanTree.Nodes[idx].NodeType = 5 //chang from table to union
				physicalPlanTree.Nodes[idx].Status = 0

			} else if tableName == "orders" {
				//1
				// physicalPlanTree.Nodes[nodeid] = createUnionNode(iparser.GetTmpTableName())
				physicalPlanTree.Nodes[idx].NodeType = 5 //chang from table to union
				physicalPlanTree.Nodes[idx].Status = 0
				// nodeid++
				//2
				ordersTableNode4 := node
				ordersTableNode4.Nodeid = nodeid
				ordersTableNode4.TmpTable = "orders_4"
				ordersTableNode4.Locate = 4
				ordersTableNode4.TransferFlag = true
				ordersTableNode4.Dest = 1
				physicalPlanTree.Nodes[nodeid] = ordersTableNode4
				physicalPlanTree.Nodes[nodeid].Parent = node.Nodeid
				physicalPlanTree.Nodes[idx].Right = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++
				//3
				physicalPlanTree.Nodes[nodeid] = createUnionNode(iparser.GetTmpTableName())
				physicalPlanTree.Nodes[nodeid].Parent = node.Nodeid
				physicalPlanTree.Nodes[idx].Left = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++
				//4
				ordersTableNode3 := node
				ordersTableNode3.Nodeid = nodeid
				ordersTableNode3.TmpTable = "orders_3"
				ordersTableNode3.Locate = 3
				ordersTableNode3.TransferFlag = true
				ordersTableNode3.Dest = 1
				physicalPlanTree.Nodes[nodeid] = ordersTableNode3
				physicalPlanTree.Nodes[nodeid].Parent = physicalPlanTree.Nodes[nodeid-1].Nodeid
				physicalPlanTree.Nodes[nodeid-1].Right = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++
				//5
				physicalPlanTree.Nodes[nodeid] = createUnionNode(iparser.GetTmpTableName())
				physicalPlanTree.Nodes[nodeid].Parent = physicalPlanTree.Nodes[nodeid-2].Nodeid
				physicalPlanTree.Nodes[nodeid-2].Left = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++
				//6
				ordersTableNode2 := node
				ordersTableNode2.Nodeid = nodeid
				ordersTableNode2.TmpTable = "orders_2"
				ordersTableNode2.Locate = 2
				ordersTableNode2.TransferFlag = true
				ordersTableNode2.Dest = 1
				physicalPlanTree.Nodes[nodeid] = ordersTableNode2
				physicalPlanTree.Nodes[nodeid].Parent = physicalPlanTree.Nodes[nodeid-1].Nodeid
				physicalPlanTree.Nodes[nodeid-1].Right = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++
				//7
				ordersTableNode1 := node
				ordersTableNode1.Nodeid = nodeid
				ordersTableNode1.TmpTable = "orders_1"
				ordersTableNode1.Locate = 1
				ordersTableNode1.TransferFlag = false
				physicalPlanTree.Nodes[nodeid] = ordersTableNode1
				physicalPlanTree.Nodes[nodeid].Parent = physicalPlanTree.Nodes[nodeid-2].Nodeid
				physicalPlanTree.Nodes[nodeid-2].Left = physicalPlanTree.Nodes[nodeid].Nodeid
				nodeid++

			} else {
				println("ERROR when creating metatables! cannot find tableName in metadata")
			}

		}

	}
	physicalPlanTree.NodeNum = nodeid
	return physicalPlanTree
}
