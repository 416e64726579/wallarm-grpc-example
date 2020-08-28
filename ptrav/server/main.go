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

// Package main implements a server for PathTraversal service.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "grpc-wallarm/ptrav/ptrav"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement ptrav.PathTraversalServer.
type server struct {
	pb.UnimplementedPathTraversalServer
}

// ShowContent implements ptrav.PathTraversalServer
func (s *server) ShowContent(ctx context.Context, in *pb.PathRequest) (*pb.ContentReply, error) {
	log.Printf("Received: %v", in.GetPath())
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Problem with getting pwd: %v", err)
	}
	dat, err := ioutil.ReadFile(pwd + "/" + in.GetPath())
	if err != nil {
		log.Printf("Problem with opening the file: %v", err)
	}
	return &pb.ContentReply{Message: string(dat)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPathTraversalServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
