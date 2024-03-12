package service

import (
	"fmt"
	"testing"

	"redis-desktop-explorer/internal/biz"
	"redis-desktop-explorer/internal/data"
)

func newConnectionService() *ConnectionService {
	dataData, _, _ := data.NewData()
	connectionRepo := data.NewConnectionRepo(dataData)
	connectionUsecase := biz.NewConnectionUsecase(connectionRepo)
	return NewConnectionService(connectionUsecase)
}

func TestCreateConnection(t *testing.T) {
	connectionService := newConnectionService()

	conn := &biz.Connection{
		Name:     "test",
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
	}
	err := connectionService.CreateConnection(conn)
	if err != nil {
		t.Errorf("create failed: %v", err)
	}
}

func TestListConnection(t *testing.T) {
	connectionService := newConnectionService()

	connection, err := connectionService.ListConnection()
	if err != nil {
		t.Errorf("list failed: %v", err)
	}

	fmt.Println(connection)
}
