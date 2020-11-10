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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	pb "irpc3"

	"google.golang.org/grpc"
)

const (
	address = "10.77.70.161:50053"
	// address = "localhost:50053"
	//defaultName = "wky"
)

var (
	inputTable pb.Table
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPushTableClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var temtableinfo pb.TableInfo = pb.TableInfo{TableName: "testrpc3",
		TotalRow: 10,
		TotalCol: 10}
	inputTable.TableInfo = &temtableinfo
	var i int64

	REC := make([]*pb.Record, 10)

	for i = 0; i < 10; i++ {
		pub := &pb.Publisher{}
		boo := &pb.Book{}
		pub.Id, boo.Id = i, i
		REC[i] = &pb.Record{}
		//println(REC[0], "REC", i)
		REC[i].Publisher, REC[i].Book = pub, boo
	}
	REC[0].Publisher.Nation = "CHN"
	REC[1].Publisher.Nation = "USA"
	REC[2].Publisher.Nation = "CHN"
	REC[3].Publisher.Nation = "CHN"
	REC[4].Publisher.Nation = "CHN"
	REC[5].Publisher.Nation = "USA"
	REC[6].Publisher.Nation = "USA"
	REC[7].Publisher.Nation = "CHN"
	REC[8].Publisher.Nation = "USA"
	REC[9].Publisher.Nation = "USA"
	inputTable.Record = REC

	r, err := c.PushTable(ctx, &inputTable)
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	println("result is ", r.IsSuc)
}
