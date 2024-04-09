## gossip-counter

#### Overview
Build a distributed counter such that a message can be sent to any counter node and the count will be incremented across all nodes.

There will be 3 node servers as well as a client script that can send a increment request to one of the 3 node servers. The data model and logic for how the node servers replicate this data amongst themselves is what you'll modify and entirely up to you.

You can decide the constraints and trade offs of the system, favoring higher consistency or availability and the behavior of network partitions. There's no right answer here, different use cases could warrant different criteria.

Things to note:
- Right now all nodes in this system act as leaders. If you want to add a leader or leader election, that's your choice
- Assume all peers are trusted

#### Dependencies
- Go 1.21+
- protoc with go gprc and protoc modules


#### Evaluation Criteria

You'll submit two things:
1. Updated source code
2. Write up explaining the tradeoffs you made and why you made them.

The primary criteria for evaluation is your thought process about the problem. You don't have to handle all edge cases in code but document them either in comments or the write up

#### Project Structure

```
gossip-counter/
├── cmd/
│   ├── client/
│   │   └── main.go          # Client script to send gossip messages to peers
│   └── server/
│       └── main.go          # Server entry point
├── gossip/
│   └── gossip.go            # Core logic for the gossip protocol
├── peers/
│   └── peers.go             # Peers management eg get network peers, build peer string
├── proto/
│   └── gossip.proto         # protobuf definition
│   └── gossip.pb.go         # protobuf implementation
│   └── gossip_grpc.pb.go    # gRPC service definition
|
├── go.mod                   # Go module file
└── go.sum                   # Go sum file
```

#### Running the code

To start a gossip server node run `go run cmd/server/main.go --nodeId <id>` where id can be 1,2,3. To start all three gossip nodes, in three separate terminals start each of the 3 nodes
```
go run cmd/server/main.go --nodeId 1
go run cmd/server/main.go --nodeId 2
go run cmd/server/main.go --nodeId 3
```

The client script allows you to either increment the counter on a given node or get the value from a given node
```
# to increment the counter on a node
go run cmd/client/main.go --nodeId 1 --action=increment

# to get the latest value from a node
go run cmd/client/main.go --nodeId 1 --action=getCount
```

#### What to modify

You're free to modify anything and everything. Right now it's a simple counter with a count variable on the server, you're free to track the counts in any way you want both on the individual ndoes as well as how it gossips. At the moment there's no communication betweent the servers, you're free to add any data model you like.

Suggestions for places to update have `// TODO` comments, you can search for that.

If you need to update the gossip.proto file, run the following to generate new protobuf and gRPC files `protoc --go_out=. --go-grpc_out=. --proto_path=. proto/gossip.proto`. 

If you get the following error, run the below commands

```
# error
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.

# commands to run
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```