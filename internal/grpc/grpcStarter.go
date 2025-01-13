package grpc

import (
	"fmt"
	raidService "github.com/Russia9/Muskrat/internal/raid/delivery/grpc"
	"github.com/Russia9/Muskrat/pkg/domain"
	"google.golang.org/grpc"
	"net"
)

type Starter struct {
	grpcServer *grpc.Server
	port       int
}

func NewGrpcStarter(port int, uc domain.RaidUsecase) *Starter {
	grpcServer := grpc.NewServer()
	raidService.Register(grpcServer, uc)
	return &Starter{
		grpcServer: grpcServer,
		port:       port,
	}
}

func (s *Starter) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", s.port))
	if err != nil {
		panic(err)
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *Starter) Stop() {
	s.grpcServer.GracefulStop()
}
