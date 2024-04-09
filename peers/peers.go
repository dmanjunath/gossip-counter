package peers

import (
	"fmt"
	"log"
)

var Peers = []string{":50051", ":50052", ":50053"}
var PeerIds = map[int64]string{
	1: ":50051",
	2: ":50052",
	3: ":50053",
}

func GetPeers() []string {
	return Peers
}

func VerifyPeerId(nodeId int64) {
	if _, ok := PeerIds[nodeId]; !ok {
		log.Fatalf("invalid node id: %v", nodeId)
	}
}

func GetPeerTarget(nodeId int64) string {
	VerifyPeerId(nodeId)
	peer := fmt.Sprintf("localhost%s", PeerIds[nodeId])
	return peer
}
