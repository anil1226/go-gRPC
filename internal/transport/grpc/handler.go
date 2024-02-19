package grpc

import (
	"context"
	"net"

	"github.com/anil1226/go-gRPC/internal/rocket"
	rkt "github.com/anil1226/protos-monorepo/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(string) (rocket.Rocket, error)
	InsertRocket(rocket.Rocket) error
	DeleteRocket(string) error
}

type Handler struct {
	RocketService RocketService
}

func New(rktserv RocketService) Handler {
	return Handler{
		RocketService: rktserv,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	rck, err := h.RocketService.GetRocketByID(req.Id)
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rck.ID,
			Name: rck.Name,
			Type: rck.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	err := h.RocketService.InsertRocket(rocket.Rocket{
		ID:   req.Rocket.Id,
		Name: req.Rocket.Name,
		Type: req.Rocket.Type,
	})
	if err != nil {
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Status: "Inserted Successfully",
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	err := h.RocketService.DeleteRocket(req.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "Deleted Successfully",
	}, nil
}
