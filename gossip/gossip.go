package gossip

import (
	"context"
	"gossip-counter/proto"
	"log"
	"sync"
)

type Server struct {
	proto.UnimplementedGossipServiceServer
	Counter int64
	Peers   []string
	mu      sync.RWMutex
}

func (s *Server) Increment(ctx context.Context, in *proto.GossipMessage) (*proto.IncrementResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Counter++
	log.Printf("Counter updated to: %v", s.Counter)
	return &proto.IncrementResponse{Success: true}, nil
}

func (s *Server) GetGossipCount(ctx context.Context, in *proto.GossipMessage) (*proto.GetCountResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return &proto.GetCountResponse{Count: s.Counter}, nil
}

// TODO implement cross-node communication here
