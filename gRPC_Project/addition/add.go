package addition

import (
	"context"
	"log"
	"main/proto"
)

type Server struct {
	proto.UnimplementedAdditionServiceServer
}

func (s *Server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	log.Printf("Addition Service: Received %d + %d", req.A, req.B)

	result := req.A + req.B

	log.Printf("Result: %d", result)
	return &proto.AddResponse{
		Res: result,
	}, nil

}
