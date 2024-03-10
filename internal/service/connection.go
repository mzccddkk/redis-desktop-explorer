package service

import (
	"context"

	"redis-desktop-explorer/internal/biz"
)

type ConnectionService struct {
	ctx context.Context
	uc  *biz.ConnectionUsecase
}

func NewConnectionService(ctx context.Context, uc *biz.ConnectionUsecase) *ConnectionService {
	return &ConnectionService{
		ctx: ctx,
		uc:  uc,
	}
}

func (s *ConnectionService) CreateConnection(ctx context.Context, conn *biz.Connection) error {
	return s.uc.CreateConnection(ctx, conn)
}

func (s *ConnectionService) ListConnection(ctx context.Context) (*[]biz.Connection, error) {
	return s.uc.ListConnection(ctx)
}
