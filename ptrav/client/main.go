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

// Package main implements a client for PathTraversal service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-wallarm/ptrav/ptrav"

	"google.golang.org/grpc"
)

const (
	defaultAddress = "localhost:50051"
	defaultName    = "world"
	defaultPath    = "default"
)

func main() {
	// Set up a connection to the server.
	address := defaultAddress
	path := defaultPath
	if len(os.Args) > 2 {
		address = os.Args[1]
		path = os.Args[2]
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPathTraversalClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ShowContent(ctx, &pb.PathRequest{Path: path})
	if err != nil {
		log.Fatalf("could not show: %v", err)
	}
	log.Printf("Content of the %s:\n%s", path, r.GetMessage())
}
