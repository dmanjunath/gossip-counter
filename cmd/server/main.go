package main

import (
	"flag"
	"gossip-counter/gossip"
	"gossip-counter/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var nodeId int64
	flag.Int64Var(&nodeId, "nodeId", 1, "node id")
	flag.Parse()

	var port string
	if nodeId == 1 {
		port = ":50051"
	} else if nodeId == 2 {
		port = ":50052"
	} else if nodeId == 3 {
		port = ":50053"
	} else {
		log.Fatalf("invalid node id: %v", nodeId)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterGossipServiceServer(grpcServer, &gossip.Server{
		Peers: []string{":50051", ":50052", ":50053"},
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
