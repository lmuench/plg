package server

import (
	"context"
	"log"
	"net"

	"github.com/lmuench/plg/plg"
	pb "github.com/lmuench/plg/rpc/plg"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	plg *plg.PLG
}

func (s *server) RegisterPlugin(ctx context.Context, in *pb.Plugin) (*pb.Error, error) {
	log.Printf("Received absolute object path: %s", in.AbsObjPath)
	err := s.plg.RegisterPlugin(in.AbsObjPath)
	if err != nil {
		return nil, err
	}
	return &pb.Error{Msg: "ok"}, nil
}

func Run(plg *plg.PLG) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRegistryServer(s, &server{plg: plg})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
