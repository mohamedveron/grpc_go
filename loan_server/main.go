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

	pb "github.com/mohamedveron/grpc_go/domain"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedLoanServer
	LoansDB    map[string]*pb.NewLoan

}

func NewServer() *server {
	loansDB := make(map[string]*pb.NewLoan)

	return &server{
		UnimplementedLoanServer: pb.UnimplementedLoanServer{},
		LoansDB:                 loansDB,
	}
}


func (s *server) AddLoan(ctx context.Context, in *pb.NewLoan) (*pb.NewLoan, error) {
	log.Printf("Received: %v", in.GetName())

	loan := &pb.NewLoan{
		Id:       in.GetId(),
		Name:     in.GetName(),
		Amount:   in.GetAmount(),
		Duration: in.GetDuration(),
	}

	s.LoansDB[in.GetId()] = loan

	return loan, nil
}

func (s *server) GetLoans(ctx context.Context, in *pb.HelloRequest) (*pb.ItemResponse, error) {

	var loans pb.ItemResponse

	for _, value := range s.LoansDB {
		loans.Items = append(loans.Items, value)
	}

	return &loans, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoanServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
