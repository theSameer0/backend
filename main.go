package main

import (
	"context"
	pb "example/grpc_backend/proto"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCapitalizeServer
}

func (s *server) CapitalName(ctx context.Context, in *pb.CapitalizeRequest) (*pb.CapitalizeReply, error) {
	var name = in.GetName()
	fmt.Printf("Got the Request of : %v\n", name)
	return &pb.CapitalizeReply{Name: strings.ToUpper(name)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCapitalizeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
