package data

import (
	"log"

	"github.com/google/wire"
	"redis-desktop-explorer/internal"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewConnectionStorage,
	NewConnectionRepo,
)

type Data struct {
	connectionStorage *internal.Storage
}

func NewData(connectionStorage *internal.Storage) (*Data, func(), error) {
	cleanup := func() {
		log.Println("closing the data resources")
	}

	return &Data{
		connectionStorage: connectionStorage,
	}, cleanup, nil
}

func NewConnectionStorage(name string) *internal.Storage {
	return internal.NewStorage(name)
}
