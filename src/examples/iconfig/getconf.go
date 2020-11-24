package main

import (
	"fmt"

	"github.com/toml"
)

const (
	configfile = "/Users/wukunyao/Documents/GitHub/iddb/cnf/iddb.conf"
)

type tomlConfig struct {
	NodeId  int64
	Cluster map[string]nodes
}

type nodes struct {
	NodeId int64
	IP     string
	Tran   string
	Call   string
}

func (n *nodes) print() {
	println("NodeId= ", n.NodeId)
	println("IP= ", n.IP)
	println("Tran= ", n.Tran)
	println("Call= ", n.Call)
}

func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		fmt.Println(err)
		return
	}
	printME(config)

}

func GetMe(config tomlConfig) nodes {
	for _, node := range config.Cluster {
		if node.NodeId == config.NodeId {
			return node
		}
	}
	var node nodes
	return node
}

func printME(config tomlConfig) {
	node := GetMe(config)
	node.print()
	return
}
