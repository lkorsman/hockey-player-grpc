package main

import (
	"context"
	"github.com/lkorsman/hockey-player-grpc/player"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myPlayerServer struct {
	player.UnimplementedPlayerServer
}

func (s myPlayerServer) Create(ctx context.Context, req *player.PlayerRequest) (*player.PlayerResponse, error) {
	log.Println("Received message from Create")
	log.Printf("Player name: %s", req.Name)
	log.Printf("Goals name: %s", req.Goal)
	return &player.PlayerResponse{
		Status: "ok",
	}, nil
}

func (s myPlayerServer) AddGoal(ctx context.Context, req *player.Goal) (*player.GoalResponse, error) {
	log.Println("Received message from AddGoal")
	return &player.GoalResponse{
		Status: "ok",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8891")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myPlayerServer{}
	player.RegisterPlayerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
