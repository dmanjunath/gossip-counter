package main

import (
	"context"
	"flag"
	"fmt"
	"gossip-counter/peers"
	"gossip-counter/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var nodeId int64
	var action string
	flag.StringVar(&action, "action", "increment", "action to perform")
	flag.Int64Var(&nodeId, "nodeId", 1, "node id")
	flag.Parse()

	peer := peers.GetPeerTarget(nodeId)
	conn, err := grpc.NewClient(peer, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", peer, err)
	}
	defer conn.Close()

	c := proto.NewGossipServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if action == "increment" {
		r, err := c.Increment(ctx, &proto.GossipMessage{})
		if err != nil {
			log.Fatalf("could not increment: %v", err)
		}
		fmt.Printf("Increment success: %t\n", r.GetSuccess())
	} else if action == "getCount" {
		r, err := c.GetGossipCount(ctx, &proto.GossipMessage{})
		if err != nil {
			log.Fatalf("could not get gossip count: %v", err)
		}
		fmt.Printf("Gossip count: %d\n", r.GetCount())
	} else {
		log.Fatalf("invalid action: %s", action)
	}
}
