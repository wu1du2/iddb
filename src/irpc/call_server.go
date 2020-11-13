/*
call_server
Author: Kunyao Wu
the implementation of irpc executer caller server
*/
package irpc

import (
	"context"
	"iexecuter"
	"log"
	"net"

	"google.golang.org/grpc"
)

type cserver struct {
	UnimplementedExecuterCallerServer
}

const (
	cport = ":50054"
)

// ExecuterCall implementation, defined in irpc.pb.go
func (s *cserver) ExecuterCall(ctx context.Context, in *IrpcCallReq) (*IrpcStatus, error) {
	iexecuter.RunExecuter(in.Txnid)
	return &IrpcStatus{IsSuc: 1}, nil
}

func RunCallServer() {
	lis, err := net.Listen("tcp", cport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening cport %v **************", cport)
	s := grpc.NewServer()
	RegisterExecuterCallerServer(s, &cserver{})
	log.Printf("Server Registering Successfully ***********")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
