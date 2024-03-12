package service

import (
	"context"

	"redis-desktop-explorer/internal/biz"
)

type ConnectionService struct {
	ctx context.Context
	uc  *biz.ConnectionUsecase
}

func NewConnectionService(uc *biz.ConnectionUsecase) *ConnectionService {
	return &ConnectionService{
		uc: uc,
	}
}

func (s *ConnectionService) CreateConnection(conn *biz.Connection) error {
	return s.uc.CreateConnection(s.ctx, conn)
}

func (s *ConnectionService) ListConnection() ([]*biz.Connection, error) {
	return s.uc.ListConnection(s.ctx)
}
