/*
tran_server
Author: Kunyao Wu
the implementation of irpc transmission server
*/
package irpctran

import (
	"context"
	"iexecuter"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	tport = ":50053"
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
	iexecuter.ExecuteCreateStmt(createstmt)
	return &IrpcStatus{IsSuc: 1}, nil
}

/*
Regitering transmission server
*/
func RunTranServer() {
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
