package iutilities

import (
	"fmt"

	"github.com/toml"
)

const (
	configfile = "/Users/wukunyao/Documents/GitHub/iddb/cnf/iddb.conf"
)

type tomlConfig struct {
	NodeId  int64
	Cluster map[string]Nodes
}

type Nodes struct {
	NodeId int64
	IP     string
	Tran   string
	Call   string
}

func (n *Nodes) Print() {
	println("NodeId= ", n.NodeId)
	println("IP= ", n.IP)
	println("Tran= ", n.Tran)
	println("Call= ", n.Call)
}

func getMe(config tomlConfig) Nodes {
	for _, node := range config.Cluster {
		if node.NodeId == config.NodeId {
			return node
		}
	}
	var node Nodes
	return node
}

func getPeers(config tomlConfig) []Nodes {
	peers := make([]Nodes, 4)
	for _, node := range config.Cluster {
		i := node.NodeId
		peers[i-1] = node
	}
	return peers
}

func printMe(config tomlConfig) {
	node := getMe(config)
	node.Print()
	return
}

func PrintMe() {
	var config tomlConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		fmt.Println(err)
		return
	}
	printMe(config)

}

func GetMe() Nodes {
	var config tomlConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		fmt.Println(err)
	}
	return getMe(config)
}

func GetPeers() []Nodes {
	var config tomlConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		fmt.Println(err)
	}
	return getPeers(config)
}
