package main

import (
	"flag"
	"gossip-counter/gossip"
	"gossip-counter/peers"
	"gossip-counter/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var nodeId int64
	flag.Int64Var(&nodeId, "nodeId", 1, "node id")
	flag.Parse()

	peers.VerifyPeerId(nodeId)
	port := peers.PeerIds[nodeId]

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterGossipServiceServer(grpcServer, &gossip.Server{
		Peers: peers.GetPeers(),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
