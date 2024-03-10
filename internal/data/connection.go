package data

import (
	"context"

	"gopkg.in/yaml.v2"
	"redis-desktop-explorer/internal/biz"
)

type connectionRepo struct {
	data *Data
}

func NewConnectionRepo(data *Data) biz.ConnectionRepo {
	return &connectionRepo{
		data: data,
	}
}

func (r *connectionRepo) Create(ctx context.Context, conn *biz.Connection) error {
	data, err := yaml.Marshal(conn)
	if err != nil {
		return err
	}

	return r.data.connectionStorage.Put(data)
}

func (r *connectionRepo) List(ctx context.Context) (*[]biz.Connection, error) {
	var list *[]biz.Connection

	data, err := r.data.connectionStorage.Get()
	if err != nil {
		return list, err
	}

	err = yaml.Unmarshal(data, list)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *connectionRepo) Delete(ctx context.Context, name string) error {
	// TODO implement me
	panic("implement me")
}

func (r *connectionRepo) Get(ctx context.Context, conn *biz.Connection) (*biz.Connection, error) {
	return nil, nil
}
