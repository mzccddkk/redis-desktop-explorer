package data

import (
	"context"

	"gopkg.in/yaml.v2"
	"redis-desktop-explorer/internal"
	"redis-desktop-explorer/internal/biz"
)

type connectionRepo struct {
	data         *Data
	localStorage *internal.Storage
}

func NewConnectionRepo(data *Data) biz.ConnectionRepo {
	return &connectionRepo{
		data:         data,
		localStorage: internal.NewStorage(internal.ConnectionsFile),
	}
}

func (r *connectionRepo) Create(ctx context.Context, conn *biz.Connection) error {
	list, err := r.List(ctx)
	if err != nil {
		return err
	}

	list = append(list, conn)
	data, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}

	return r.localStorage.Put(data)
}

func (r *connectionRepo) List(ctx context.Context) ([]*biz.Connection, error) {
	var list []*biz.Connection

	data, err := r.localStorage.Get()
	if err != nil {
		return list, err
	}

	err = yaml.Unmarshal(data, &list)
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
