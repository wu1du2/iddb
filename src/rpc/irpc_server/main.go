/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "rpc/irpc"
)

const (
	port = ":50053"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedIrpcGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) IrpcGreet(ctx context.Context, in *pb.IrpcRequest) (*pb.IrpcReply, error) {
	log.Printf("Received: %v from node 1", in.GetName())
	return &pb.IrpcReply{Message: "Received, " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening port %v **************", port)
	s := grpc.NewServer()
	pb.RegisterIrpcGreeterServer(s, &server{})
	log.Printf("Server Registering Successfully ***********")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
