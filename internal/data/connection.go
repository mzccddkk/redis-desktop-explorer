package data

import (
	"context"
	"fmt"

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

func getDefaultName(host string, port int, list []*biz.Connection, idx int) string {
	connName := fmt.Sprintf("%s_%d_%d", host, port, idx)
	if idx == 0 {
		connName = fmt.Sprintf("%s_%d", host, port)
	}
	for _, c := range list {
		if c.Name == connName {
			idx++
			connName = getDefaultName(host, port, list, idx)
		}
	}
	return connName
}

func (r *connectionRepo) Create(ctx context.Context, conn *biz.Connection) error {
	list, err := r.List(ctx)
	if err != nil {
		return err
	}

	if conn.Name == "" {
		conn.Name = getDefaultName(conn.Host, conn.Port, list, 0)
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
