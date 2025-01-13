package grpc

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/protobuf"
	"google.golang.org/grpc"
)

type RaidService struct {
	protobuf.UnimplementedRaidServer
	raid domain.RaidUsecase
}

func Register(grpcServer *grpc.Server, repo domain.RaidUsecase) {
	protobuf.RegisterRaidServer(grpcServer, &RaidService{raid: repo})
}

func (r *RaidService) GetRaidInfo(ctx context.Context, info *protobuf.RaidInfo) (*protobuf.Null, error) {
	err := r.raid.Create(ctx, info.BossName, info.Cell, info.Time)
	if err != nil {
		return nil, err
	}
	return &protobuf.Null{}, nil
}
