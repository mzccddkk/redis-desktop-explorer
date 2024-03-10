package biz

import (
	"context"
)

type Connection struct {
	Name     string `json:"name" yaml:"name"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type ConnectionRepo interface {
	Create(ctx context.Context, conn *Connection) error
	List(ctx context.Context) (*[]Connection, error)
	Delete(ctx context.Context, name string) error
}

type ConnectionUsecase struct {
	repo ConnectionRepo
}

func NewConnectionUsecase(repo ConnectionRepo) *ConnectionUsecase {
	return &ConnectionUsecase{
		repo: repo,
	}
}

func (uc *ConnectionUsecase) CreateConnection(ctx context.Context, conn *Connection) error {
	return uc.repo.Create(ctx, conn)
}

func (uc *ConnectionUsecase) ListConnection(ctx context.Context) (*[]Connection, error) {
	return uc.repo.List(ctx)
}
