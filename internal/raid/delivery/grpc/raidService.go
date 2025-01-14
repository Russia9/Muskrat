package grpc

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/protobuf"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type RaidService struct {
	protobuf.UnimplementedRaidServer
	raid domain.RaidUsecase
}

func Register(grpcSrv *grpc.Server, uc domain.RaidUsecase) {
	protobuf.RegisterRaidServer(grpcSrv, &RaidService{raid: uc})
}

func (r *RaidService) GetRaidInfo(ctx context.Context, info *protobuf.RaidInfo) (*protobuf.Null, error) {
	log.Info().Interface("info", info).Msg("GetRaidInfo")
	err := r.raid.UpdateOrCreate(ctx, info.BossName, info.Cell, info.Time)
	if err != nil {
		return nil, err
	}
	return &protobuf.Null{}, nil
}
