/*
tran_server
Author: Kunyao Wu
the implementation of irpc transmission server
*/
package irpctran

import (
	"context"
	"itrans"
	"iutilities"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	tport = ":" + iutilities.GetMe().Tran
)

type tserver struct {
	UnimplementedPushTableServer
}

/*
the implementation of PushTable that defined in irpc.pb.go
receive a table from remote client, make it a INSERT statement
and execute
*/
func (s *tserver) PushTable(ctx context.Context, in *Table) (*IrpcStatus, error) {
	createstmt := in.Createstmt
	itrans.ExecuteCreateStmt(createstmt)
	return &IrpcStatus{IsSuc: 1}, nil
}

/*
Regitering transmission server
*/
func RunTranServer() {
	iutilities.Me = iutilities.GetMe()
	iutilities.Me.Print()

	iutilities.Peers = iutilities.GetPeers()
	iutilities.Peers[0].Print()
	iutilities.Peers[1].Print()
	iutilities.Peers[2].Print()
	iutilities.Peers[3].Print()

	iutilities.Mysql = iutilities.GetMysqlConfig()
	println("mysql_user= ", iutilities.Mysql.Mysql_user)
	println("mysql_passwd= ", iutilities.Mysql.Mysql_passwd)
	println("mysql_db= ", iutilities.Mysql.Mysql_db)
	println("mysql_ip_port= ", iutilities.Mysql.Mysql_ip_port)

	println("irpctran")
	lis, err := net.Listen("tcp", tport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening tport %v **************", tport)
	s := grpc.NewServer()
	RegisterPushTableServer(s, &tserver{})
	log.Printf("Server Registering Successfully ***********")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
