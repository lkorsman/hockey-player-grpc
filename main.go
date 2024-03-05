package hockey_player_grpc

import (
	"context"
	"github.com/lkorsman/hockey-player-grpc/player"
	"log"
	"net"
)

type myPlayerServer struct {
	player.UnimplementedPlayerServer
}

func (s myPlayerServer) Create(ctx context.Context, req *player.PlayerRequest) (*player.PlayerResponse, error) {
	log.Println("Received message from Create")
	return &player.PlayerResponse{
		Status: "ok",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8891")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
}
