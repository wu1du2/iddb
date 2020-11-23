package main

import (
	"fmt"

	"github.com/toml"
)

type tomlConfig struct {
	Servers map[string]server
}

type server struct {
	IP   string
	Tran string
	Call string
}

func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile("/Users/wukunyao/Documents/GitHub/iddb/cnf/iddb.conf", &config); err != nil {
		fmt.Println(err)
		return
	}

	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s, %s %s)\n", serverName, server.IP, server.Tran, server.Call)
	}

}
